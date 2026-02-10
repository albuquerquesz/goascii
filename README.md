# goascii

Generate ASCII art from text.

## Install (dev)

```bash
go build -o goascii ./cmd/goascii
```

## Usage

```bash
./goascii create "alexsa"
./goascii create -f slant "alexsa"
./goascii create -o ./ascii.txt "alexsa"
```

After rendering, you can press:

- `C` to copy to clipboard
- `O` to save to a file (you will be prompted for a path)
- `Enter` to exit

If running without a TTY, you must pass `--font`.
