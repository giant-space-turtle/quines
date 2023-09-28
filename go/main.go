package main

import "fmt"

var text = `package main

import "fmt"

var text = 

func main() {
	fmt.Print(text[:39] + "\u0060" + text + "\u0060" + text[39:])
}
`

func main() {
	fmt.Print(text[:39] + "\u0060" + text + "\u0060" + text[39:])
}
