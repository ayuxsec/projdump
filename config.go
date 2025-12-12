package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ExcludeDirs      []string `yaml:"exclude_dirs"`
	ExcludeFileExts  []string `yaml:"exclude_file_exts"`
	ExcludeFileNames []string `yaml:"exclude_file_names"`
	OutputFile       string   `yaml:"output_file"`
}

var DefaultCfgPath = filepath.Join(os.Getenv("HOME"), ".config", "projdump", "config.yaml")

func DefaultConfig() Config {
	return Config{
		ExcludeDirs:      []string{"node_modules", "vendor", ".git"},
		ExcludeFileNames: []string{"projdump.txt"},
		ExcludeFileExts:  []string{".jpeg", ".jpg", ".png", ".gif", ".svg", ".pdf"},
		OutputFile:       "projdump.txt",
	}
}

var Cfg *Config

// LoadOrCreateConfig loads the configuration from the specified path.
// If the configuration file does not exist and force is false, it creates a default one on the specified path.
// If it exists, it reads and unmarshals the configuration into the Cfg variable.
func LoadOrCreateConfig(cfgPath string, force bool) error {
	doExist, err := doExistCfg(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to check if config file exists at %s: %v", cfgPath, err)
	} else if !doExist || force {
		if !force {
			fmt.Printf("[INF] config file does not exist. creating config at %s with default values\n", cfgPath)
		} else {
			fmt.Printf("[INF] force creating config at %s with default values\n", cfgPath)
		}
		if err := createDefaultCfg(cfgPath); err != nil {
			return fmt.Errorf("failed to save default config: %v", err)
		}
	}
	return loadCfg(cfgPath)
}

func loadCfg(cfgPath string) error {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to read config file at %s: %v", cfgPath, err)
	}
	Cfg = &Config{}
	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config file at %s: %v", cfgPath, err)
	}
	return nil
}

func createDefaultCfg(cfgPath string) error {
	data, err := yaml.Marshal(DefaultConfig())
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %v", err)
	}
	if err := ensureCfgDirExist(cfgPath); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}
	if err := os.WriteFile(cfgPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write default config to %s: %v", cfgPath, err)
	}
	return nil
}

func ensureCfgDirExist(cfgPath string) error {
	return os.MkdirAll(getCfgDirPath(cfgPath), 0755)
}

func getCfgDirPath(cfgPath string) string {
	return filepath.Dir(cfgPath)
}

func doExistCfg(cfgPath string) (bool, error) {
	_, err := os.Stat(cfgPath)
	if err != nil && !os.IsNotExist(err) {
		return false, err
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}
