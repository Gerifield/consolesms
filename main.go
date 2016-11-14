package main

import (
	"io/ioutil"
	"log"
	"os"

	//"github.com/gerifield/go-pushbullet"
	"gopkg.in/yaml.v2"
)

type Number struct {
	Name  string `yaml:"name"`
	Phone string `yaml:"number"`
}

type Config struct {
	Token   string   `yaml:"token"`
	Numbers []Number `yaml:"numbers"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: %s <phone number/name from config>", os.Args[0])
		return
	}

	bc, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}

	var conf Config
	err = yaml.Unmarshal(bc, &conf)
	if err != nil {
		log.Fatalln(err)
		return
	}

	phoneNum := findByNameWithFallback(os.Args[1], conf.Numbers)

	pb := pushbullet.New(conf.Token)
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

func findByNameWithFallback(name string, numbers []Number) string {
	for _, it := range numbers {
		if it.Name == name {
			return it.Phone
		}
	}
	return name
}
