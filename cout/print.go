package cout

import "fmt"

const RESET string = "\u001B[0m"
const RED string = "\u001B[31m"
const GREEN string = "\u001B[32m"
const YELLOW string = "\u001B[33m"

func Info(msg string) {
	fmt.Printf("[*] %s\n", msg)
}

func Error(msg string, colored ...bool) {
	if len(colored) == 0 || colored[0] {
		fmt.Printf("%s[X] %s%s\n", RED, msg, RESET)
	} else {
		fmt.Printf("[x] %s\n", msg)
	}
}

func Warning(msg string, colored ...bool) {
	if len(colored) == 0 || colored[0] {
		fmt.Printf("%s[-] %s%s\n", YELLOW, msg, RESET)
	} else {
		fmt.Printf("[-] %s\n", msg)
	}
}

func Positive(msg string, colored ...bool) {
	if len(colored) == 0 || colored[0] {
		fmt.Printf("%s[+] %s%s\n", GREEN, msg, RESET)
	} else {
		fmt.Printf("[+] %s\n", msg)
	}
}
