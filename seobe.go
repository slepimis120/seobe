package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func detectIndentation(lines []string, stepsIndex int) string {
	for i := stepsIndex + 1; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "- name:") {
			return line[:strings.Index(line, "- name:")]
		} else if trimmed != "" {
			break
		}
	}
	return lines[stepsIndex][:strings.Index(lines[stepsIndex], "steps:")+1]
}

func jobUsesMvnw(lines []string, stepsIndex int) bool {
	for i := stepsIndex + 1; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])

		if strings.HasSuffix(trimmed, ":") && !strings.HasPrefix(trimmed, "- name:") {
			break
		}

		if strings.Contains(trimmed, "./mvnw") {
			return true
		}
	}
	return false
}

func removeDuplicateCheckout(lines []string) []string {
	var newLines []string
	seenCheckout := false

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "uses: actions/checkout") {
			if seenCheckout {
				if len(newLines) > 0 && strings.Contains(newLines[len(newLines)-1], "- name: checkout") {
					newLines = newLines[:len(newLines)-1]
				}
				continue
			}
			seenCheckout = true
		}
		newLines = append(newLines, line)
	}

	return newLines
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

	if strings.Contains(ymlContent, "uses: actions/cache") {
		log.Println("Cache step already exists. No changes made.")
		fmt.Println(ymlContent)
		return
	}

	log.Println("Scanning for jobs that use mvnw...")

	var newLines []string
	for i := 0; i < len(lines); i++ {
		newLines = append(newLines, lines[i])

		if strings.TrimSpace(lines[i]) == "steps:" {
			if jobUsesMvnw(lines, i) {
				stepsIndent := detectIndentation(lines, i)

				jdkStep := fmt.Sprintf(`%s- name: Set up JDK
%s  uses: actions/setup-java@v3
%s  with:
%s    distribution: 'temurin'
%s    java-version: '21'`, stepsIndent, stepsIndent, stepsIndent, stepsIndent, stepsIndent)

				newLines = append(newLines, jdkStep)
				log.Println("Added JDK setup step for a job using mvnw.")

				cacheStep := fmt.Sprintf(`%s- name: Cache Maven dependencies
%s  uses: actions/cache@v4
%s  with:
%s    path: ~/.m2/repository
%s    key: ${{ runner.os }}-maven-${{ hashFiles('**/pom.xml') }}
%s    restore-keys: |
%s      ${{ runner.os }}-maven-`, stepsIndent, stepsIndent, stepsIndent, stepsIndent, stepsIndent, stepsIndent, stepsIndent)

				newLines = append(newLines, cacheStep)
				log.Println("Added cache step for a job using mvnw.")
			}
		}
	}

	newLines = removeDuplicateCheckout(newLines)

	ymlContent = strings.Join(newLines, "\n")

	err = os.WriteFile(*ymlPath, []byte(ymlContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	log.Println("JDK setup and cache step successfully added where necessary.")
	fmt.Println(ymlContent)
}
