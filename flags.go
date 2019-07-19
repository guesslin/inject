package main

import (
	"flag"
)

var (
	src      = flag.String("src", "", "")
	srcMac   = flag.String("srcMac", "", "")
	dst      = flag.String("dst", "", "")
	dstMac   = flag.String("dstMac", "", "")
	nic      = flag.String("nic", "eth0", "")
	filename = flag.String("filename", "", "")
)
