package main

import (
	"fmt"
	"github.com/sokel/gate-notification-bot/bot"
	"log"
	"time"
)

const (
	token        = "786543701:AAHatO33Yr0QLw61cR7wARZsONLbC5_m4Rw"
	channelName  = "@sonm_gate_noty"
	reportPeriod = 5 * time.Second
)

func main() {
	err := run()
	log.Println(err)
}

func run() error {
	b, err := bot.NewBot()
	if err != nil {
		return fmt.Errorf("failed to create bot instance: %v", err)
	}

	return b.Start()
}
