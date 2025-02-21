package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Funkcija koja detektuje indentaciju (razmak) za ključnu reč u YAML-u
func detectIndentation(lines []string, keyword string) string {
	for _, line := range lines {
		if strings.Contains(line, keyword) {
			return line[:strings.Index(line, keyword)]
		}
	}
	return "  " // Podrazumevana indentacija ako ne nađe ništa
}

func main() {
	ymlPath := flag.String("path", "", "Path to the .yml file")
	flag.Parse()

	if *ymlPath == "" {
		log.Fatal("Please provide the path to the .yml file using the -path flag")
	}

	content, err := os.ReadFile(*ymlPath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	ymlContent := string(content)
	lines := strings.Split(ymlContent, "\n")

	// Ako već postoji cache korak, ne radimo ništa
	if strings.Contains(ymlContent, "uses: actions/cache") {
		log.Println("Cache step already exists. No changes made.")
		fmt.Println(ymlContent)
		return
	}

	log.Println("Cache step is not present. Adding cache for Maven dependencies...")

	var newLines []string
	for i := 0; i < len(lines); i++ {
		newLines = append(newLines, lines[i])

		// Ako pronađemo "steps:", dodajemo cache korak odmah ispod
		if strings.TrimSpace(lines[i]) == "steps:" {
			stepsIndent := detectIndentation(lines, "steps:")
			cacheStep := fmt.Sprintf(`%s- name: Cache Maven dependencies
%s  uses: actions/cache@v4
%s  with:
%s    path: ~/.m2/repository
%s    key: ${{ runner.os }}-maven-${{ hashFiles('**/pom.xml') }}
%s    restore-keys: |
%s      ${{ runner.os }}-maven-`, stepsIndent, stepsIndent, stepsIndent, stepsIndent, stepsIndent, stepsIndent, stepsIndent)

			newLines = append(newLines, cacheStep)
		}
	}

	// Spajamo linije nazad u string
	ymlContent = strings.Join(newLines, "\n")

	// Upisujemo nazad u fajl
	err = os.WriteFile(*ymlPath, []byte(ymlContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	log.Println("Cache step successfully added.")
	fmt.Println(ymlContent)
}
