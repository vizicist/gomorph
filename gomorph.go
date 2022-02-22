package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/hypebeast/go-osc/osc"
	"github.com/vizicist/gomorph/morph"
)

var Client *osc.Client
var Verbose bool
var Quiet bool

func main() {

	verbosePtr := flag.Bool("verbose", false, "Be Verbose")
	quietPtr := flag.Bool("quiet", false, "Don't print cursor messages")
	listPtr := flag.Bool("list", false, "List Morphs")
	listenPtr := flag.Bool("listen", false, "listen for OSC events")
	ipPtr := flag.String("ip", "127.0.0.1", "IP address to listen on")
	portPtr := flag.Int("port", 4444, "Port to listen on")
	serialPtr := flag.String("serial", "*", "Morph serialnum to use")

	flag.Parse()

	Verbose = *verbosePtr
	Quiet = *quietPtr

	if *listenPtr {
		doListen(*ipPtr, *portPtr)
		return
	}

	fmt.Printf("MAIN: start\n")
	morphs, err := morph.Init(*serialPtr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("MAIN: B\n")

	for _, m := range morphs {
		fmt.Printf("Opened: Morph idx=%d serial=%s firmware=%d.%d.%d\n",
			m.Idx, m.SerialNum, m.FwVersionMajor, m.FwVersionMinor, m.FwVersionBuild)
	}

	if *listPtr {
		return
	}

	Client = osc.NewClient(*ipPtr, *portPtr)
	morph.Start(morphs, handleMorph, 1.0)
}

func handleMorph(e morph.CursorDeviceEvent) {
	if !Quiet {
		fmt.Printf("Morph: cursor %s %s %f %f %f\n", e.Ddu, e.CID, e.X, e.Y, e.Z)
	}
	msg := osc.NewMessage("/cursor")
	msg.Append(e.Ddu)
	msg.Append(e.CID)
	msg.Append(float32(e.X))
	msg.Append(float32(e.Y))
	msg.Append(float32(e.Z))
	Client.Send(msg)
}

func doListen(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	server := &osc.Server{}
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		fmt.Println("Couldn't listen: ", err)
	}
	defer conn.Close()

	fmt.Printf("Listening for OSC at %s\n", addr)
	fmt.Printf("Press \"q\" to exit\n")

	go func() {
		for {
			packet, err := server.ReceivePacket(conn)
			if err != nil {
				fmt.Println("Server error: " + err.Error())
				os.Exit(1)
			}

			if packet != nil {
				switch pkt := packet.(type) {
				default:
					fmt.Println("Unknow packet type!")

				case *osc.Message:
					fmt.Printf("OSC Message: ")
					osc.PrintMessage(pkt)

				case *osc.Bundle:
					fmt.Println("OSC Bundle:")
					bundle := pkt
					for i, message := range bundle.Messages {
						fmt.Printf("OSC Bundle Message #%d: ", i+1)
						osc.PrintMessage(message)
					}
				}
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		c, err := reader.ReadByte()
		if err != nil {
			os.Exit(0)
		}

		if c == 'q' {
			os.Exit(0)
		}
	}
}
