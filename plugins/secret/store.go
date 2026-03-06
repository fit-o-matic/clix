package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"

	kit "github.com/finkt/clix-kit"
	"github.com/zalando/go-keyring"
)

// Store manages secrets using the system keyring with key tracking.
type Store struct {
	serviceName string
	metaPath    string
}

// NewStore creates a Store with metadata stored in the user's config directory.
func NewStore() (*Store, error) {
	config, err := kit.LoadConfig()
	if err != nil {
		return nil, err
	}
	cacheDir := config.GetCliCacheDir() // ensure config dir exists

	if err := os.MkdirAll(cacheDir, 0700); err != nil {
		return nil, err
	}
	return &Store{
		serviceName: config.GetCliName(),
		metaPath:    filepath.Join(cacheDir, "secret-keys.json"),
	}, nil
}

// SetEnv sets an environment variable with the secret value for the given key.
func (s *Store) SetEnv(key string) error {
	value, err := keyring.Get(s.serviceName, key)
	if err != nil {
		return err
	}
	return os.Setenv(key, value)
}

// set stores a secret in the system keyring and tracks the key name.
func (s *Store) Set(key, value string) error {
	if err := keyring.Set(s.serviceName, key, value); err != nil {
		return err
	}
	return s.trackKey(key)
}

// delete removes a secret from the system keyring and untracks the key.
func (s *Store) Delete(key string) error {
	if err := keyring.Delete(s.serviceName, key); err != nil {
		return err
	}
	return s.untrackKey(key)
}

// List returns all secret key names that exist in the keyring.
// It prunes any tracked keys that no longer exist.
func (s *Store) List() ([]string, error) {
	keys, err := s.loadKeys()
	if err != nil {
		return nil, err
	}
	var valid []string
	for _, key := range keys {
		if _, err := keyring.Get(s.serviceName, key); err == nil {
			valid = append(valid, key)
		}
	}
	if len(valid) != len(keys) {
		_ = s.saveKeys(valid) // prune orphans
	}
	return valid, nil
}

func (s *Store) loadKeys() ([]string, error) {
	data, err := os.ReadFile(s.metaPath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var keys []string
	if err := json.Unmarshal(data, &keys); err != nil {
		return nil, err
	}
	return keys, nil
}

func (s *Store) saveKeys(keys []string) error {
	data, err := json.Marshal(keys)
	if err != nil {
		return err
	}
	return os.WriteFile(s.metaPath, data, 0600)
}

func (s *Store) trackKey(key string) error {
	keys, err := s.loadKeys()
	if err != nil {
		return err
	}
	if slices.Contains(keys, key) {
		return nil
	}
	keys = append(keys, key)
	slices.Sort(keys)
	return s.saveKeys(keys)
}

func (s *Store) untrackKey(key string) error {
	keys, err := s.loadKeys()
	if err != nil {
		return err
	}
	keys = slices.DeleteFunc(keys, func(k string) bool { return k == key })
	return s.saveKeys(keys)
}
