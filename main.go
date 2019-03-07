/*
Command caat prints out a picture of a cat on high-color terminals. It's
useful when you accidentally type caat instead of cat.
*/
package main

//go:generate binder -o cat.go cat.txt

// Use github.com/davecheney/godoc2md to generate README file.
//go:generate bash -c "godoc2md -ex . | sed -e 's/\\/src\\/target\\///g' -e 's/import \".\"/import \"github.com\\/ancientlore\\/caat\"/' > README.md"

import "fmt"

func main() {
	fmt.Println(string(Lookup("/cat.txt")))
}
