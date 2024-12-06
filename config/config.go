package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/huahuayu/1scan/internal/models"
)

type Manager struct {
	config *models.Config
	mu     sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) LoadConfig(filepath string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if filepath == "" {
		return fmt.Errorf("config filepath is empty")
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// First unmarshal into a temporary map
	var rawConfig map[string]models.ChainConfig
	if err := json.Unmarshal(data, &rawConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Convert string keys to int keys
	chains := make(map[int]models.ChainConfig)
	for k, v := range rawConfig {
		chainID, err := strconv.Atoi(k)
		if err != nil {
			return fmt.Errorf("invalid chain ID %s: %w", k, err)
		}
		chains[chainID] = v
	}

	m.config = &models.Config{
		Chains: chains,
	}

	return nil
}

func (m *Manager) IsConfigLoaded() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.config != nil
}

func (m *Manager) GetChainConfig(chainID int) (models.ChainConfig, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.config == nil {
		return models.ChainConfig{}, false
	}

	config, exists := m.config.Chains[chainID]
	return config, exists
}

func (m *Manager) GetAllChains() map[int]models.ChainConfig {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.config == nil {
		return nil
	}

	return m.config.Chains
}
