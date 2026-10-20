package main

import (
	"fmt"
	"runtime"

	"github.com/blss-tico/go-engineering-journey/modules/module-00/hello-go/greeting"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println(greeting.Message())
	fmt.Printf("target=%s/%s\n", runtime.GOOS, runtime.GOARCH)

	title := cases.Title(language.BrazilianPortuguese)
	fmt.Println(title.String("engenharia de software com go"))
}
