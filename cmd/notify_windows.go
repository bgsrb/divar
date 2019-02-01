package main

import (
	"fmt"
	"log"

	"github.com/bgsrb/divar"
	toast "gopkg.in/toast.v1"
)

func notify(p divar.Post) {
	notification := toast.Notification{
		AppID:   "Divent",
		Title:   p.Title,
		Message: fmt.Sprintf("Price : %d", p.Price),
		Actions: []toast.Action{
			toast.Action{
				Type:      "protocol",
				Label:     "Open browser",
				Arguments: fmt.Sprintf("https://divar.ir/v/%s", p.Token),
			},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
