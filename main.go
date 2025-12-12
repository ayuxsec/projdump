package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	configPath := flag.String("config", DefaultCfgPath, "Path to config file")
	forceCreateCfg := flag.Bool("force-config", false, "Force-create config file with default values")
	outputFile := flag.String("output", "", "Output file (overrides config)")
	skipWarn := flag.Bool("no-warn", false, "Skip warning for large dumps. Use with extreme caution!")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <directory> [flags]\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	root := flag.Arg(0)

	if err := LoadOrCreateConfig(*configPath, *forceCreateCfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if Cfg == nil {
		log.Fatalf("Config is nil after loading")
	}

	files, err := CollectFiles(root, Cfg)
	if err != nil {
		log.Fatalf("Failed collecting files: %v", err)
	}

	// safety: warn for huge directories
	const warnThreshold = 5000

	if len(files) > warnThreshold && !*skipWarn {
		fmt.Printf("[CRITICAL] This will dump %d files. This could result in self DoS!\n", len(files))
		if !askConfirm("Do you want to continue?") {
			fmt.Println("Aborted.")
			os.Exit(1)
		}
	}

	if *outputFile == "" {
		*outputFile = Cfg.OutputFile
	}

	if err := DumpFiles(files, *outputFile); err != nil {
		log.Fatalf("Failed dumping files: %v", err)
	}

	fmt.Printf("Dump created for directory %s: %s (%d files)\n", root, *outputFile, len(files))
}

func askConfirm(prompt string) bool {
	fmt.Print(prompt + " [y/N]: ")

	var resp string
	_, _ = fmt.Scanln(&resp) // ignore error; empty input = NO

	resp = strings.ToLower(strings.TrimSpace(resp))
	return resp == "y" || resp == "yes"
}
