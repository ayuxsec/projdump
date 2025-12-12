package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	configPath := flag.String("config", DefaultCfgPath, "Path to config file")
	forceCreateCfg := flag.Bool("force-config", false, "Force-create config file with default values")
	outputFile := flag.String("output", "", "Output file (overrides config)")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <directory> [flags]\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("No directory provided.\nUsage: projdump <directory>")
	}

	root := flag.Arg(0)

	if err := LoadOrCreateConfig(*configPath, *forceCreateCfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// safeguard against nil config
	if Cfg == nil {
		log.Fatalf("Config is nil after loading")
	}

	files, err := CollectFiles(root, Cfg)
	if err != nil {
		log.Fatalf("Failed collecting files: %v", err)
	}

	if *outputFile == "" {
		*outputFile = Cfg.OutputFile
	}

	if err := DumpFiles(files, *outputFile); err != nil {
		log.Fatalf("Failed dumping files: %v", err)
	}

	fmt.Printf("Dump created of directory %s: %s (%d files)\n", root, *outputFile, len(files))
}
