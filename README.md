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
$ ./projdump 
Usage: ./projdump <directory> [flags]

Flags:
  -config string
        Path to config file (default "$HOME/.config/projdump/config.yaml")
  -force-config
        Force-create config file with default values
  -no-warn
        Skip warning for large dumps. Use with extreme caution!
  -output string
        Output file (overrides config)
```

Examples:

```console
projdump .
projdump . -output allmychaos.txt
projdump . -force-config  # oops, reset everything
```

That’s it. It dumps your project.
