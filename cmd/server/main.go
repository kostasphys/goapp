package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	goapp "goapp/internal/app/server"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmsgprefix | log.Lshortfile)
}

func main() {

	/*When we open many sockets (either simultaneously or one after the other), the server starts to reserve some
	heap memory  which is to be expected. After we close all those connections this allocated memory doesn't perish
	immediately. However after some time, using the pprof tool we can see that the memory is slowly released back to
	the system if the server is idle. Probably what is  happening is that the garbage collector takes some time to clean
	up that memory for efficiency reasons.
	*/
	// Debug.
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	// Register signal handlers for exiting
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)

	// Start.
	if err := goapp.Start(exitChannel); err != nil {
		log.Fatalf("fatal: %+v\n", err)
	}
}
