package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// SayHello returns a greeting string that indicates the language and greets the name passed in
func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

// Italian struct implements Greeter interface
type Italian struct{}

// LanguageName method implementation for the Italian struct
func (i Italian) LanguageName() string {
	return "Italian"
}

// Greet method implementation for the Italian struct
func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

// LanguageName method implementation for the Portuguese struct
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

// Greet method implementation for the Portuguese struct
func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
