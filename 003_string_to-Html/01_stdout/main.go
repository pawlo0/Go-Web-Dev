package main

import "fmt"

func main() {
	name := "Paulo Santos"
	subtitle := "Golang learner"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	<h3>` + subtitle + `</h3>
	</body>
	</html>
	`
	fmt.Println(tpl)
}
