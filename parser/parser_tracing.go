package parser

import (
	"fmt"
	"strings"
)

var (
	debug      bool = false
	traceLevel int  = 0
)

const traceIndentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIndentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	if !debug {
		return
	}

	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel++ }
func decIdent() { traceLevel-- }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
