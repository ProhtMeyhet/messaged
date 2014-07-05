package main

import(
	"fmt"
	"flag"
	"os"
)

const EMPTY = ""

type flagConfig struct {
	Type, Handler, CertificateFile, KeyFile string
	MaxProcs, Threads int
}

func NewFlagConfig() *flagConfig {
	return &flagConfig{}
}

func (flags *flagConfig) parse() {
	flag.StringVar(&flags.Type, "type", "tcpssl", "Server type (tcpssl, tcp, udp, dbus)")
	flag.StringVar(&flags.Handler, "handler", "notify", "Message Handler (notify)")
	flag.IntVar(&flags.MaxProcs, "maxprocs", 5, "set runtime.GOMAXPROCS")
	flag.IntVar(&flags.Threads, "threads", 5, "number of threads to run")
	flag.StringVar(&flags.CertificateFile, "cert", "cert.pem", "ssl certificate file")
	flag.StringVar(&flags.KeyFile, "key", "key.pem", "ssl key file")

	//flag.BoolVar(&flags.noEncryption, "no-encryption-i-know-the-risks", false,
	//		"disable encryption DANGEROUS")

	flag.Parse()

	if flags.MaxProcs < 1 {
		fmt.Println("maxprocs may not be smaller then 1!")
		os.Exit(22)
	}

	if flags.Threads < 1 {
		fmt.Println("threads may not be smaller then 1!")
		os.Exit(23)
	}
}

func (flags *flagConfig) usage() {
	flag.Usage()
}
