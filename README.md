# real-time-tcp-connection-monitor

A Go CLI that repeatedly reads /proc/net/tcp and prints the count of sockets currently in the ESTABLISHED state.

## Key features

- Continuous output every two seconds
- Uses only the standard library

## Project structure

- `main.go` — Reads and parses /proc/net/tcp
- `go.mod` — Go module definition

## Requirements

- Go 1.22+
- Linux, because it depends on /proc/net/tcp

## Setup

```bash
git clone https://github.com/biprajit007/real-time-tcp-connection-monitor.git
cd real-time-tcp-connection-monitor
go run .
```

## Usage

### Start monitoring

```bash
go run .
```

## Sample output

```text
established=42
```

## Limitations / next improvements

- Linux-only
- Counts all ESTABLISHED sockets and does not group by port/process
