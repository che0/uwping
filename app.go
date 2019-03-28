package main

import (
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

func main() {
    
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] uwsgi://host:port/path\n\nParameters:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	http_host := flag.String("host", "", "HTTP_HOST")
	remote_addr := flag.String("remote", "127.0.0.1", "remote addr")
	modifier1 := flag.Int("modifier1", 0, "modifier1")
	expected_status := flag.Int("expected_status", 200, "expected_status")
	flag.Parse()
	arg := flag.Arg(0)
	if arg == "" {
		flag.Usage()
	}
	url, err := url.Parse(arg)
	if err != nil {
		flag.Usage()
	}
	host, _, _ := net.SplitHostPort(url.Host)
	if *http_host == "" {
		http_host = &host
	}
	response := get(url, *http_host, *remote_addr, *modifier1)
	
	response_re := regexp.MustCompile(`^HTTP/1.1 ([0-9]{3}) ([^\r\n]+)`)
	m := response_re.FindStringSubmatch(string(response))
	if len(m) < 3 {
		fmt.Printf("Unable to parse response\n")
		os.Exit(1)
	}
	
	status_code, _ := strconv.Atoi(m[1])
	fmt.Printf("%d %s\n", status_code, m[2])
	if int(status_code) != *expected_status {
		os.Exit(1)
	}
}
