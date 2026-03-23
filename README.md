# Real-Time TCP Connection Monitor

Prints the number of established TCP connections by reading `/proc/net/tcp`.

## Features
- Real-time counter loop
- No external packages

## Usage
```bash
go run .
```

## Disclaimer
Currently Linux-focused because it reads `/proc/net/tcp`.
