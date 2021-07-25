package main

import (
	"fmt"
	"net"
)

func main() {
	var url string
	fmt.Println("Input website")
	fmt.Scanln(&url)
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
	fmt.Println(cname)
}

func ip(url string) {
	ips, err := net.LookupIP(url)
	if err != nil {
		panic(err)
	}
	if len(ips) == 0 {
		fmt.Printf("no record\n")
	}
	for _, ip := range ips {
		fmt.Printf("%s\n", ip.String())
	}
}

func mx(url string) {
	mxs, err := net.LookupMX(url)
	if err != nil {
		panic(err)
	}

	for _, mx := range mxs {
		fmt.Printf("%s %v\n", mx.Host, mx.Pref)
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
	for _, ns := range nss {
		fmt.Printf("%s\n", ns.Host)
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
	for _, txt := range txts {
		fmt.Printf("%s\n", txt)
	}

}
