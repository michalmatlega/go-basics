package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://www.aws.com",
	}
	fmt.Println(websites)

	websites["Google"] = "https://www.google.pl"
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
