package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	mnd "github.com/invidian/go-mnd/v2"
)

func main() { singlechecker.Main(mnd.Analyzer) }
