package main

import "fmt"

func main() {
	name := "Devesh Kakad"

	tpl := `
	<!DOCTYPE html>
	<html>
	<body>
	<h1>My Name is ` + name + `</h1>
	<p>My first paragraph.</p>

	</body>
	</html>`

	fmt.Println(tpl)
}
