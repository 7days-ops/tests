package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func Hello(name string , language string) (string , error) {
	
	if name == "" && language == "" {
		log.Println("Name and language not provided, reading from stdin...")
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your name: ")
		nameInput , _ := r.ReadString('\n')
		name = strings.TrimSpace(nameInput)
	  if name == "" {
            return "", errors.New("name cannot be empty")
    }
		fmt.Print("Enter you language:")
		langInput , _ := r.ReadString('\n')
		language = strings.TrimSpace(langInput)
		if language == "" {
						return "",errors.New("language cannot be empty")
		}
	}
	var prefix string
	switch  language {
	case "deustch":
		prefix = "Hallo"
	case "english":
		prefix = "Hello"
	case "spanish":
		prefix = "Hola"
	case "norway":
		prefix = "Hei"
	case "italian":
		prefix = "Ciao"
	default:
		return "",errors.New("need to prodive a supported language")
	}
	answer := prefix + " " + name 
	return answer,nil
}

func main() {
	answer , _ := Hello("","")
	fmt.Println(answer)
}