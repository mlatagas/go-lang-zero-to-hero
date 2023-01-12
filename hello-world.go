package main

import "fmt"

func Hello(name string, lang string) string {

	prefix := "Hello, "

	if name == "" {
		name = "World"
	}

	if lang == "Spanish" {

		prefix = "Hola, "

	} else if lang == "Zulu" {
		prefix = "Sawubona, "

	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
	fmt.Println(Hello("", "English"))
	fmt.Println(Hello("Andile", "Zulu"))
}
