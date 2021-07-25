package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	cname(url)
	ip(url)
	mx(url)
	ns(url)
	txt(url)
	txt("_dmarc." + url)
}

func cname(url string) {
	cname, err := net.LookupCNAME(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n[CNAME] %s\n", cname)
}

func ip(url string) {
	ips, err := net.LookupIP(url)
	if err != nil {
		panic(err)
	}
	if len(ips) == 0 {
		fmt.Printf("no record\n")
	}
	fmt.Println()
	for _, ip := range ips {
		fmt.Printf("[IP] %s\n\n", ip)
	}
}

func mx(url string) {
	mxs, err := net.LookupMX(url)
	if err != nil {
		panic(err)
	}

	for _, mx := range mxs {
		fmt.Printf("[MX] %s %v\n", mx.Host, mx.Pref)
	}
}

func ns(url string) {
	nss, err := net.LookupNS(url)
	if err != nil {
		panic(err)
	}
	if len(nss) == 0 {
		fmt.Printf("no record\n")
	}
	fmt.Println()
	for _, ns := range nss {
		fmt.Printf("[NS] %s\n", ns.Host)
	}
}

func txt(url string) {
	txts, err := net.LookupTXT(url)
	if err != nil {
		panic(err)
	}
	if len(txts) == 0 {
		fmt.Printf("no record\n")
	}
	fmt.Println()
	for _, txt := range txts {
		if strings.Contains(txt, "v=DMARC1") {
			fmt.Printf("[DMARC] %s\n", txt)
			break
		}
		fmt.Printf("[TXT] %s\n", txt)
	}
}
