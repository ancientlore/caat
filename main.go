/*
Command caat prints out a picture of a cat on high-color terminals. It's
useful when you accidentally type caat instead of cat.

You can also `ln -s caat gti` and it will print a GTI.
*/
package main

//go:generate binder -o cat.go cat.txt gti.txt

// Use github.com/davecheney/godoc2md to generate README file.
//go:generate bash -c "godoc2md -ex . | sed -e 's/\\/src\\/target\\///g' -e 's/import \".\"/import \"github.com\\/ancientlore\\/caat\"/' > README.md"

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var img string
	exe, err := os.Executable()
	if err == nil {
		exe = filepath.Base(exe)
		switch exe {
		case "gti", "gti.exe":
			img = "/gti.txt"
		default:
			img = "/cat.txt"
		}
		// fmt.Println(exe)
	}
	fmt.Println(string(Lookup(img)))
}
