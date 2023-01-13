package main

import "fmt"

func Hello(name string, lang string) string {

	prefix := "Hello, "

	switch lang {
	case "spanish":
		prefix = "Hola, "
	case "zulu":
		prefix = "Sawubona, "

	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
	fmt.Println(Hello("", "English"))
	fmt.Println(Hello("Andile", "Zulu"))
}
