package main

import "fmt"

const (
	portuguese = "Portuguese"
	french     = "French"

	englishPrefix    = "What up, "
	portuguesePrefix = "Olá, "
	frenchPrefix     = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		return "What up!"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portuguesePrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Human", "Portuguese"))
}
