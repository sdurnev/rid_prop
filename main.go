package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type param struct {
	Req  string
	Type int
	Name string
}

type rid_param []param

var data = rid_param{
	{"C", 0, "Start"},
	{"R008234", 3, "--"},
	{"R009235", 3, "--"},
	{"R012229", 3, "--"},
	{"R018235", 3, "--"},
	{"R020228", 3, "--"},
	{"R021229", 3, "--"},
	{"R024232", 3, "--"},
	{"R033232", 1, "MAIN_L1-N"},
	{"R034233", 1, "MAIN_L2-N"},
	{"R035234", 1, "MAIN_L3-N"},
	{"R036235", 1, "GENS_L1-N"},
	{"R037236", 1, "GENS_L2-N"},
	{"R038237", 1, "GENS_L3-N"},
	{"R039238", 3, "--"},
	{"R041231", 3, "--"},
	{"R042232", 3, "--"},
	{"R043233", 3, "--"},
	{"R056237", 2, "GENS_BAT_U"},
	{"R057238", 3, "--"},
	{"R058239", 3, "--"},
	{"R130230", 3, "--"},
	{"R132232", 3, "--"},
	{"R144235", 3, "--"},
	{"R155237", 3, "--"},
	{"R161234", 3, "--"},
	{"R166239", 3, "--"},
	{"R183238", 3, "--"},
	{"R194240", 3, "--"},
	{"R216235", 3, "--"},
	{"R217236", 3, "--"},
	{"R218237", 3, "--"},
	{"R223233", 3, "--"},
	{"R228238", 3, "--"},
	{"R229239", 3, "--"},
	{"R231232", 3, "--"},
	{"R235236", 3, "--"},
	{"E", 0, "End"},
}

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	d := net.Dialer{}
	d.Timeout = 2 * time.Second
	conn, err := d.Dial("tcp", addr)

	if err != nil {
		fmt.Printf("{\"status\":\"error\",\"error\":\"%s\"}", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Print("{")
	for i := 0; i < len(data); i++ {
		if i > 0 {
			buff := make([]byte, 1024)
			conn.Write([]byte(data[i].Req))
			fmt.Printf("\"%s\":", data[i].Req)
			n, _ := conn.Read(buff)
			fmt.Printf("%s,", buff[:n])
		}
	}
}

func main() {

	var (
		ip   = "10.10.12.23"
		port = 2001
	)
	SocketClient(ip, port)
}
