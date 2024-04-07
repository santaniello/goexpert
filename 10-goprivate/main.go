package main

import (
	"fmt"
	"github.com/santaniello/fc-utils-private/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
