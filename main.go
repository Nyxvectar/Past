package main

import (
	"fmt"
	"github.com/Nyxvectar/Past/others"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	hiMessage   = "Nyxvectar εδώ."
	initialErr  = "Η αστρονομική αρχικοποίηση απέτυχε: %v\n"
	pulTransErr = "Αποτυχία μετάδοσης pulsar: %v\n"
	photonBurst = "Επεκτόθηκε φωτονικός εκκλάδωση ▲"
	shutdSeqMsg = "\nΛήψη παγκόσμιου σήματος διακοπής..."
)

func main() {
	sl, err := others.NewSpaceLink()
	if err != nil {
		fmt.Printf(initialErr, err)
		os.Exit(1)
	}
	defer sl.Close()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-sl.Ctx.Done():
				return
			case <-ticker.C:
				if err := sl.TransmitSignal(hiMessage); err != nil {
					fmt.Printf(pulTransErr, err)
					return
				}
				fmt.Println(photonBurst)
			}
		}
	}()
	<-sigChan
	fmt.Println(shutdSeqMsg)
}
