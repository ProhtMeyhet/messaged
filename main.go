package main

import(
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	libmessage "github.com/ProhtMeyhet/libgomessage"
)

func main() {
	flags := NewFlagConfig()
	flags.parse()

	runtime.GOMAXPROCS(flags.MaxProcs)

	var handler libmessage.SendMessageInterface
	switch(flags.Handler) {
	case "stdout":
		handler = libmessage.NewStdout()
	case "notify":
		fallthrough
	default:
		handler = libmessage.NewNotify()
	}

	switch(flags.Type) {
	case EMPTY:
		flags.usage()
		os.Exit(1)
	case "tcpssl":
		serverConfig := libmessage.NewTcpServerConfig()
		serverConfig.SetSSL(flags.CertificateFile, flags.KeyFile)
		server := libmessage.NewTcpPlainServer(serverConfig)
		runServer(flags.Threads, server, handler)
	case "tcp":
		serverConfig := libmessage.NewTcpServerConfig()
		server := libmessage.NewTcpPlainServer(serverConfig)
		runServer(flags.Threads, server, handler)
	case "androidpn":
		fallthrough
	case "xmppandroidpn":
		serverConfig := libmessage.NewTcpServerConfig()
		server := libmessage.NewXmppAndroidpnServer(serverConfig)
		runServer(flags.Threads, server, handler)
	}

	handleSignals()
}

func send(messenger libmessage.SendMessageInterface,
		message *libmessage.Message,
		to libmessage.ToInterface) {
	go messenger.Send(message, to)
}

func runServer(threads int,
		reciever libmessage.RecieveMessageInterface,
		handler libmessage.SendMessageInterface) {
	e := reciever.StartService()
	if e != nil {
		fmt.Println("Error: " + e.Error())
		os.Exit(11)
	}

	go reciever.Receive()

	for i := 0; i <= threads; i++ {
		go handleMessage(reciever.GetMessage(), handler)
	}
}

func handleMessage(messageChannel chan *libmessage.Message,
		handler libmessage.SendMessageInterface) {
	for {
		select {
			case message := <-messageChannel:
				to := &libmessage.To{ To: message.To }
				fmt.Println("Got Message!")
				send(handler, message, to)
		}
	}
}

func handleSignals() {
	signalChannel := make(chan os.Signal, 1)
        signal.Notify(signalChannel, syscall.SIGTERM)

infinite:
	for {
		select {
		case <-signalChannel:
			break infinite
		}
	}
}
