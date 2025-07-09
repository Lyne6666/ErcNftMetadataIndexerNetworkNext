// cmd/ercnftmetadataindexernetworknext/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"ercnftmetadataindexernetworknext/internal/ercnftmetadataindexernetworknext"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := ercnftmetadataindexernetworknext.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
