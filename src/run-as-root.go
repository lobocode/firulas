package src

import (
	"log"
	"os"
)

func runAsRoot() {

	if os.Geteuid() == 0 {
		log.Println("You have root permission!")
	} else {
		log.Fatal("You don't have root permission!")
	}

}
