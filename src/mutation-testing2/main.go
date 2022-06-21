package main

func GreetingsByLocale(locale string) string {

	greetings := resolveGreetings(locale)

	if greetings == "" {
		return "Sorry, we didn't identified you"
	}

	return greetings
}

func resolveGreetings(locale string) string {
	if locale == "en" {
		return "Hello"
	}

	if locale == "pt" {
		return "Olá"
	}

	return ""
}
