package loadbalancer

import (
	"fmt"
	"sync"
)

type APIKey struct {
	Key string
}

type LoadBalancer struct {
	ChainKeys map[int][]*APIKey
	current   map[int]int
	mu        sync.Mutex
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		ChainKeys: make(map[int][]*APIKey),
		current:   make(map[int]int),
	}
}

func (lb *LoadBalancer) UpdateAPIKeys(chainID int, keys []string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	chainKeys := make([]*APIKey, 0)
	for _, key := range keys {
		chainKeys = append(chainKeys, &APIKey{Key: key})
	}
	lb.ChainKeys[chainID] = chainKeys
	lb.current[chainID] = 0
}

func (lb *LoadBalancer) RemoveAPIKey(chainID int, key string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	chainKeys, exists := lb.ChainKeys[chainID]
	if !exists {
		return
	}

	for i, k := range chainKeys {
		if k.Key == key {
			lb.ChainKeys[chainID] = append(chainKeys[:i], chainKeys[i+1:]...)
			break
		}
	}
}

func (lb *LoadBalancer) GetNextAPIKey(chainID int) (string, error) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	chainKeys, exists := lb.ChainKeys[chainID]
	if !exists {
		return "", fmt.Errorf("no API keys found for chain %d", chainID)
	}

	if len(chainKeys) == 0 {
		return "", fmt.Errorf("no API keys available for chain %d", chainID)
	}

	currentIndex := lb.current[chainID]
	key := chainKeys[currentIndex].Key
	lb.current[chainID] = (currentIndex + 1) % len(chainKeys)

	return key, nil
}
