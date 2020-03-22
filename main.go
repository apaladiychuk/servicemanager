package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("./start.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(stdout)
	go readStdOut(scanner)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

func readStdOut(scanner *bufio.Scanner) {

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf(">>>%s\n", line)
	}

}
