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

const version = "0.00.4"

type param struct {
	Req  string
	Type float64
	Name string
}

type rid_param []param

var data = rid_param{
	{"C", 0, "Start"},
	{"R008234", 100, "MAIN_FREQ"},
	{"R009235", 1, "GENS_FREQ"},
	{"R012229", 1, "GENS_WORKH"},
	{"R018235", 1, "R018235"},
	{"R020228", 1, "R020228"},
	{"R021229", 1, "R021229"},
	{"R024232", 1, "R024232"},
	{"R033232", 1, "MAIN_L1-N"},
	{"R034233", 1, "MAIN_L2-N"},
	{"R035234", 1, "MAIN_L3-N"},
	{"R036235", 1, "GENS_L1-N"},
	{"R037236", 1, "GENS_L2-N"},
	{"R038237", 1, "GENS_L3-N"},
	{"R039238", 1, "GENS_FUEL1"},
	{"R041231", 1, "R041231"},
	{"R042232", 1, "R042232"},
	{"R043233", 1, "R043233"},
	{"R056237", 10, "GENS_BAT_U"},
	{"R057238", 1, "R057238"},
	{"R058239", 10, "GENS_TEN_D"},
	{"R130230", 1, "R130230"},
	{"R132232", 1, "R132232"},
	{"R144235", 1, "R144235"},
	{"R155237", 1, "R155237"},
	{"R161234", 1, "R161234"},
	{"R166239", 1, "R166239"},
	{"R183238", 1, "R183238"},
	{"R194240", 1, "R194240"},
	{"R216235", 1, "R216235"},
	{"R217236", 1, "R217236"},
	{"R218237", 1, "R218237"},
	{"R223233", 1, "R223233"},
	{"R228238", 1, "R228238"},
	{"R229239", 1, "R229239"},
	{"R231232", 1, "R231232"},
	{"R235236", 1, "R235236"},
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
		fmt.Printf("{\"status\":\"error\",\"error\":\"%s\"}\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	for i := 0; i < len(data); i++ {
		buff := make([]byte, 1024)
		conn.Write([]byte(data[i].Req))
		time.Sleep(50 * time.Millisecond)
		//fmt.Printf("\"%s\":", data[i].Name)
		conn.Read(buff)
		//fmt.Printf("\"%s\":", data[i].Name)
		tempBuf = append(tempBuf, fmt.Sprintf("%s", buff))
		// p := ParseData(buff, i)
	}
	parsData(tempBuf)
}

func parsData(data []string) {
	var newArr []float64
	for i := 0; i < len(data)-1; i++ {
		if i != 0 && i != len(data) {
			newString := strings.Split(data[i], string(rune(4)))
			newString = strings.Split(newString[0], "D")
			newTempString := newString[1]
			newTempUint, _ := strconv.ParseUint(newTempString, 10, 64)
			newStr := strconv.FormatUint(newTempUint, 10)
			if newTempUint != 0 {
				newTem2Str := string(newStr[0:(len(newStr) / 2)])
				//fmt.Println(newTem2Str)
				a, _ := strconv.ParseFloat(newTem2Str, 32)
				newArr = append(newArr, a)
			} else {
				//fmt.Println(newStr)
				b, _ := strconv.ParseFloat(newStr, 32)
				newArr = append(newArr, b)
			}
		}
	}

	PrintData(newArr)
}

func PrintData(tempBuf []float64) {
	data1 := data[1:37]
	if len(tempBuf) < len(data)-2 {
		fmt.Print("{\"status\":\"error\",\"error\":\"short packet\"}\n")
		os.Exit(1)
	}
	if tempBuf == nil {
		fmt.Print("{\"status\":\"error\",\"error\":\"no data in packet\"}\n")
		os.Exit(1)
	}
	for i := 0; i < len(data1); i++ {

		if i == 0 {
			fmt.Print("{")
		}
		fmt.Printf("\"%s\":", data1[i].Name)
		p := tempBuf[i]
		fmt.Printf("%.2f,", p/data1[i].Type)

		if i == len(data1)-1 {
			fmt.Printf(" \"version\":\"%s\"}\n", version)
		}
	}
}
