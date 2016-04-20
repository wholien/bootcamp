package main

import "fmt"

func main() {
	fmt.Println("What's your name?\n")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("What's your email?\n")
	var email string
	fmt.Scanf("%s", &email)
	fmt.Printf("<header>profile</header>\n")
	fmt.Printf("<body>\n")
	fmt.Printf("Hi, my name is %s\n", name)
	fmt.Printf("My email is %s\n", email)
	fmt.Printf("</body>\n");
}
	
