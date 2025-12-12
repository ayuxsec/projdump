// walks directory, applies exclusion rules, returns list of files.
package main

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func shouldSkipDir(name string, cfg *Config) bool {
	return slices.Contains(cfg.ExcludeDirs, name)
}

func shouldSkipExt(path string, cfg *Config) bool {
	ext := strings.ToLower(filepath.Ext(path))
	for _, e := range cfg.ExcludeFileExts {
		if ext == strings.ToLower(e) {
			return true
		}
	}
	return false
}

func shouldSkipFile(name string, cfg *Config) bool {
	for _, pattern := range cfg.ExcludeFileNames {
		if strings.EqualFold(name, pattern) {
			return true
		}
	}
	return false
}

func CollectFiles(root string, cfg *Config) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // ignore errors but continue
		}

		if info.IsDir() {
			if shouldSkipDir(info.Name(), cfg) {
				return filepath.SkipDir
			}
			return nil
		}

		if shouldSkipExt(path, cfg) {
			return nil
		}

		if shouldSkipFile(path, cfg) {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files, err
}
