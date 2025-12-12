package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	configPath := flag.String("config", DefaultCfgPath, "Path to config file")
	forceCreateCfg := flag.Bool("force-config", false, "Force create config file with default values even if it exists")
	root := flag.String("path", ".", "Root directory to scan")
	outputFile := flag.String("output", "", "Output file path (overrides config)")
	flag.Parse()

	if err := LoadOrCreateConfig(*configPath, *forceCreateCfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// safeguard against nil config
	if Cfg == nil {
		log.Fatalf("Config is nil after loading")
	}

	files, err := CollectFiles(*root, Cfg)
	if err != nil {
		log.Fatalf("Failed collecting files: %v", err)
	}

	if *outputFile == "" {
		*outputFile = Cfg.OutputFile
	}

	if err := DumpFiles(files, *outputFile); err != nil {
		log.Fatalf("Failed dumping files: %v", err)
	}

	fmt.Printf("Dump created: %s (%d files)\n", *outputFile, len(files))
}
