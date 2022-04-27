package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/githubengineering/email-service/v2/emailservice"
)

// run tests defined in datadir
func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(0)
	}

	port := flag.Int("p", 8080, "port")
	flag.Parse()

	realMain(*port, false)
  return
}

func realMain(port int, dryrun bool) int {
	err := emailservice.Serve("", port, dryrun)
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    return 1
  }
	return 0
}
