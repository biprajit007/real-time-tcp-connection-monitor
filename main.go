package main
import (
  "bufio"; "fmt"; "os"; "strings"; "time"
)
func countEstablished(path string) int { f,err:=os.Open(path); if err!=nil { return 0 }; defer f.Close(); s:=bufio.NewScanner(f); c:=0; first:=true; for s.Scan(){ if first { first=false; continue }; fields:=strings.Fields(s.Text()); if len(fields)>3 && fields[3]=="01" { c++ } }; return c }
func main(){ for { fmt.Printf("established=%d\n", countEstablished("/proc/net/tcp")); time.Sleep(2*time.Second) } }
