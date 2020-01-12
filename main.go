package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const version = "0.00.1"

type param struct {
	Req  string
	Type int
	Name string
}

type rid_param []param

var data = rid_param{
	{"C", 0, "Start"},
	{"R008234", 3, "R008234"},
	{"R009235", 3, "R009235"},
	{"R012229", 3, "R012229"},
	{"R018235", 3, "R018235"},
	{"R020228", 3, "R020228"},
	{"R021229", 3, "R021229"},
	{"R024232", 3, "R024232"},
	{"R033232", 1, "MAIN_L1-N"},
	{"R034233", 1, "MAIN_L2-N"},
	{"R035234", 1, "MAIN_L3-N"},
	{"R036235", 1, "GENS_L1-N"},
	{"R037236", 1, "GENS_L2-N"},
	{"R038237", 1, "GENS_L3-N"},
	{"R039238", 3, "R039238"},
	{"R041231", 3, "R041231"},
	{"R042232", 3, "R042232"},
	{"R043233", 3, "R043233"},
	{"R056237", 2, "GENS_BAT_U"},
	{"R057238", 3, "R057238"},
	{"R058239", 3, "R058239"},
	{"R130230", 3, "R130230"},
	{"R132232", 3, "R132232"},
	{"R144235", 3, "R144235"},
	{"R155237", 3, "R155237"},
	{"R161234", 3, "R161234"},
	{"R166239", 3, "R166239"},
	{"R183238", 3, "R183238"},
	{"R194240", 3, "R194240"},
	{"R216235", 3, "R216235"},
	{"R217236", 3, "R217236"},
	{"R218237", 3, "R218237"},
	{"R223233", 3, "R223233"},
	{"R228238", 3, "R228238"},
	{"R229239", 3, "R229239"},
	{"R231232", 3, "R231232"},
	{"R235236", 3, "R235236"},
	{"E", 0, "End"},
}

func main() {

	addressIP := flag.String("ip", "localhost", "a string")
	tcpPort := flag.Int("port", 2001, "a string")
	flag.Parse()

	var (
		ip   = *addressIP
		port = *tcpPort
	)
	SocketClient(ip, port)
}

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	d := net.Dialer{}
	d.Timeout = 500 * time.Millisecond
	conn, err := d.Dial("tcp", addr)

	if err != nil {
		fmt.Printf("{\"status\":\"error\",\"error\":\"%s\"}", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Print("{")
	for i := 0; i < len(data); i++ {
		if i == 0 {
			buff := make([]byte, 1024)
			conn.Write([]byte(data[i].Req))
			conn.Read(buff)
		}
		if i > 0 && i < len(data)-1 {
			buff := make([]byte, 1024)
			conn.Write([]byte(data[i].Req))
			fmt.Printf("\"%s\":", data[i].Name)
			conn.Read(buff)
			p := ParseData(buff, i)
			if i < (len(data) - 2) {
				fmt.Printf("%s,", p)
			} else {
				fmt.Printf("%s", p)
			}
		}
		if i == len(data)-1 {
			buff := make([]byte, 1024)
			conn.Write([]byte(data[i].Name))
			fmt.Printf(", \"version\":%s}", version)
			conn.Read(buff)
		}
	}
}

func ParseData(buff []byte, num int) string {
	if data[num].Type == 0 {
		return string(buff)
	}
	if data[num].Type == 1 {
		var newBuff []byte
		for l := 0; l < len(buff); l++ {
			if l >= 7 && l < 13 {
				newBuff = append(newBuff, buff[l])
			}
		}
		f, _ := strconv.ParseFloat(string(newBuff), 64)
		return fmt.Sprintf("%0.2f", f/1000)
	}
	if data[num].Type == 2 {
		var newBuff []byte
		for l := 0; l < len(buff); l++ {
			if l >= 7 && l < 13 {
				newBuff = append(newBuff, buff[l])
			}
		}
		f, _ := strconv.ParseFloat(string(newBuff), 64)
		return fmt.Sprintf("%0.2f", f/10000)
	}
	if data[num].Type == 3 {
		return fmt.Sprintf("\"%s\"", string(buff))
	}
	return "0"
}
