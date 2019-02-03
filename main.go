package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"motd/message"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("./motd", os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message := message.Greeting(name, phrase)
	err = ioutil.WriteFile("./motd", []byte(message), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("end")

}
