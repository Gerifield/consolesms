package main

import (
	"github.com/gerifield/go-pushbullet"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: %s <phone number>", os.Args[0])
		return
	}

	phoneNum := os.Args[1]

	pb := pushbullet.New("<TOKEN HERE>")
	devs, err := pb.Devices()
	if err != nil {
		log.Fatalln(err)
		return
	}

	if len(devs) < 1 {
		log.Fatalln("Missing device")
		return
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
		return
	}

	user, err := pb.Me()
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = pb.PushSMS(user.Iden, devs[0].Iden, phoneNum, string(b))
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("Success")
}
