package main

import (
	"context"
	"log"
	"os"

	"github.com/robert-nix/ansihtml"
)

func main() {
	config, err := NewConfig("config.toml")
	if err != nil {
		log.Fatalf("Error with config file: %v", err)
	}
	ansiText := generateAnsi(config)
	log.Println(ansiText)

	err = os.MkdirAll("web/dist/", 0755) // 0755 is standard permissions
	if err != nil {
		log.Fatalln("Error creating directory:", err)
	}
	err = os.WriteFile("web/dist/index.txt", []byte(ansiText+"\n"), 0644)
	if err != nil {
		log.Fatalf("failed to write output file in index.txt: %v", err)
	}

	ansiHtml := string(ansihtml.ConvertToHTML([]byte(ansiText)))

	htmlFile, err := os.Create("web/dist/index.html")
	if err != nil {
		log.Fatalf("failed to create output file index.html: %v", err)
	}

	log.Println("Generating HTML file")

	err = index(ansiHtml, config).Render(context.Background(), htmlFile)
	if err != nil {
		log.Fatalf("failed to write output file in index.html: %v", err)
	}
}
