package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

func FormatMessage(str string) string {
	lenStr := len(str)
	str = str[1 : lenStr-2]
	str = strings.Replace(str, ",", ", ", 1)
	str = strings.Replace(str, ":", ": ", 2)
	str = strings.ReplaceAll(str, "\"", "")
	return str
}

func startClients(idx int, conn *websocket.Conn, wg *sync.WaitGroup, countTest int) {

	defer wg.Done()
	defer conn.Close()

	count := 0
	strz := "[conn #" + strconv.Itoa(idx) + "]"
	for {
		_, mess, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error when reading message!!!")
			return
		}

		strMess := string(mess[:])
		strMess = FormatMessage(strMess)
		fmt.Printf("%s %s \n", strz, strMess)

		if countTest > 0 {
			count++
			if count >= countTest {
				fmt.Printf(" client %d  ENDED \n", idx)
				return
			}
		}

	}
}

func extractArgs() (int, int) {
	num := flag.Int("n", -1, "Number of parallel connections")
	test := flag.Int("t", 0, "Flag ")
	flag.Parse()
	if *num <= 0 {
		fmt.Println("-n cannot be non positive")
		os.Exit(1)
	}
	if *test < 0 {
		fmt.Println("-t cannot be negative")
		os.Exit(1)
	}

	return *num, *test
}

func main() {

	numConn, test := extractArgs()

	clientConn := []*websocket.Conn{}
	wg := sync.WaitGroup{}
	urlTemp := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/goapp/client"}

	for i := 0; i < numConn; i++ {

		conn, _, err := websocket.DefaultDialer.Dial(urlTemp.String(), nil)

		if err != nil {
			fmt.Printf("Error when opening socket, %s", err)
			os.Exit(1)
		}

		clientConn = append(clientConn, conn)

		wg.Add(1)

		var flag int
		if test > 0 {
			flag = i*test + 10
		} else {
			flag = 0
		}

		go startClients(i, clientConn[i], &wg, flag)

	}

	wg.Wait()

	fmt.Printf("ALL CONNECTIONS ENDED SUCCESSFULLY\n")

}
