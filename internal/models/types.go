package models

type Config struct {
	Chains map[int]ChainConfig `json:"chains"`
}

type ChainConfig struct {
	Endpoint string         `json:"endpoint"`
	Keys     map[string]int `json:"keys"`
}

type APIKey struct {
	Key       string
	RateLimit int
	LastUsed  int64
	Requests  int
}
