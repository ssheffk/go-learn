package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("level", "CRITICAL", "log level to filter for")

	flag.Parse()

	f, err := os.Open("./logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	bufReader := bufio.NewReader(f)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			log.Println(line)
		}
	}
}
