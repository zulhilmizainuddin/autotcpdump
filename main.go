package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"strings"
	"autotcpdump/parser"
	"autotcpdump/executer"
	"autotcpdump/checker"
)

func main() {
	cmdlineArgs := os.Args[1:]

	config := parser.ConfigParser{}
	if err := config.Parse("config/config.json"); err != nil {
		log.Fatal(err)
	}

	if err := checker.CheckIfPathWritable(config.PcapLocation); err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("tcpdump_%v.pcap", time.Now().Format("20060102_150405"))
	fmt.Println("directory:", config.PcapLocation, "filename:", filename)

	commandOptions := config.CommandOptions + " " + strings.Join(cmdlineArgs, " ")

	tcpdump := executer.TcpdumpExecuter{}
	if err := tcpdump.RunTcpdump(config.PcapLocation, filename, commandOptions); err != nil {
		log.Fatal(err)
	}

	if err := tcpdump.TerminateTcpdump(); err != nil {
		log.Fatal(err)
	}

	if err := tcpdump.AdbPullPcapFile(config.PcapLocation, filename); err != nil {
		log.Fatal(err)
	}

	if err := tcpdump.OpenWithWireshark(config.WiresharkLocation, filename); err != nil {
		log.Fatal(err)
	}
}
