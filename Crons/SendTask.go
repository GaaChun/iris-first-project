package Crons

import (
	"log"
)

type SendTask struct{}

func (this SendTask) Run() {
	log.Println("Send sth...")
}
