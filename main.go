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

const version = "0.00.2"

type param struct {
	Req  string
	Type int
	Name string
}

type rid_param []param

var data = rid_param{
	{"C", 0, "Start"},
	{"R008234", 4, "MAIN_FREQ"},
	{"R009235", 4, "GENS_FREQ"},
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
	var tempBuf []string
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	d := net.Dialer{}
	d.Timeout = 500 * time.Millisecond
	conn, err := d.Dial("tcp", addr)

	if err != nil {
		fmt.Printf("{\"status\":\"error\",\"error\":\"%s\"}", err)
		os.Exit(1)
	}
	defer conn.Close()
	//fmt.Print("{")
	for i := 0; i < len(data); i++ {
		buff := make([]byte, 1024)
		conn.Write([]byte(data[i].Req))
		//fmt.Printf("\"%s\":", data[i].Name)
		conn.Read(buff)
		//fmt.Printf("\"%s\":", data[i].Name)
		tempBuf = append(tempBuf, fmt.Sprintf("%s", buff))
		// p := ParseData(buff, i)
	}
	//	fmt.Println(tempBuf)
	PrintData(tempBuf)
}

func PrintData(tempBuf []string) {

	if tempBuf == nil {
		fmt.Print("{\"status\":\"error\",\"error\":\"error1\"}")
		os.Exit(1)
	}

	fmt.Print("{")
	for i := 0; i < len(data); i++ {
		if i == 0 {

		}
		if i > 0 && i < len(data)-1 {
			fmt.Printf("\"%s\":", data[i].Name)
			p := ParseData(tempBuf[i], i)
			if i < (len(data) - 2) {
				fmt.Printf("%s,", p)
			} else {
				fmt.Printf("%s", p)
			}
		}
		if i == len(data)-1 {
			fmt.Printf(", \"version\":\"%s\"}", version)
		}
	}
}

func ParseData(buff string, num int) string {
	if data[num].Type == 0 {
		var newBuff []byte
		for l := 0; l < len(buff)-2; l++ {
			newBuff = append(newBuff, buff[l])
		}
		return string(newBuff)
	}
	if data[num].Type == 1 {
		oldString := buff
		newString := strings.Split(oldString, string(rune(4)))
		newString = strings.Split(newString[0], "D0")
		f, _ := strconv.ParseFloat(newString[1], 64)
		return fmt.Sprintf("%0.2f", f/1000)
	}
	if data[num].Type == 2 {
		oldString := buff
		newString := strings.Split(oldString, string(rune(4)))
		newString = strings.Split(newString[0], "D0")
		f, _ := strconv.ParseFloat(newString[1], 64)
		return fmt.Sprintf("%0.2f", f/10000)
	}
	if data[num].Type == 3 {
		oldString := buff
		newString := strings.Split(oldString, string(rune(4)))

		return fmt.Sprintf("\"%s\"", newString[0])
	}
	if data[num].Type == 4 {
		oldString := buff
		newString := strings.Split(oldString, string(rune(4)))
		newString = strings.Split(newString[0], "D0")
		f, _ := strconv.ParseFloat(newString[1], 64)
		return fmt.Sprintf("%0.2f", f/1000000)
	}
	return "0"
}
