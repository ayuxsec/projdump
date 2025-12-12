// Reads all files and writes a stitched output.
package main

import (
	"fmt"
	"os"
)

func DumpFiles(files []string, outputPath string) error {
	out, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer out.Close()
	fmt.Printf("[INF] Writting %d file(s) to %s\n", len(files), outputPath)

	for _, f := range files {
		fmt.Fprintf(out, "\n---\n %s \n---\n\n", f)

		data, err := os.ReadFile(f)
		if err != nil {
			fmt.Printf("[ERR] failed to read file %s: %v", f, err)
			fmt.Fprintf(out, "[ERROR READING FILE]\n")
			continue
		}

		out.Write(data)
		out.Write([]byte("\n"))
	}

	return nil
}
