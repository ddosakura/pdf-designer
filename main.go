//go:generate sblocker -s ./assets -f
//go:generate go fmt statik/statik.go
//go:generate rm -rf ./example

package main

import (
	"github.com/ddosakura/gklang"
	"github.com/ddosakura/pdf-designer/cmd"
)

func main() {
	gklang.Init("PDF DESIGNER")

	cmd.Execute()
}
