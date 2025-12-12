# projdump

Walks your folder, ignores the garbage (looking at you, `node_modules`), and stitches everything else into one big dump file. Configurable via YAML.

## Install

```console
go install github.com/ayuxsec/projdump@latest
```

## Config Example

```yaml
exclude_dirs:
  - node_modules
  - vendor
  - .git

exclude_file_exts:
  - .png
  - .jpg

exclude_file_names:
  - projdump.txt

output_file: projdump.txt
```

## Usage

```console
$ projdump -h
Usage of projdump:
  -config string
        Path to config file (default "~/.config/projdump/config.yaml")
  -force-config
        Force-create config file with default values even if it exists
  -output string
        Output file path (overrides config)
  -path string
        Root directory to scan (default ".")
```

Examples:

```console
projdump                # btw don’t run this in your home directory unless you want a 3-hour dump
projdump -output allmychaos.txt
projdump -force-config  # oops, reset everything
```

That’s it. It dumps your project.
