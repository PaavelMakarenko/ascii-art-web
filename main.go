package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Input struct {
	Text   string
	Banner string
}

type Output struct {
	Ascii string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	// http status 200
	w.WriteHeader(http.StatusOK)
	// Create a new template
	template := template.Must(template.ParseFiles("templates/index.html"))

	// 404 Error handling
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Switch case for GET and POST requests
	switch r.Method {
	case "GET":
		// Send the template page and output
		template.Execute(w, nil)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// Get input values from the form in template
		details := Input{
			Text:   r.FormValue("text"),
			Banner: r.FormValue("banner"),
		}
		// Call ascii art function with required variables and create a string to output
		response := asciiArt(details.Text, details.Banner)
		responseStr := strings.Join(response, "")

		//Generate output
		output := Output{
			Ascii: responseStr,
		}
		// Send the template page and output
		err := template.Execute(w, output)
		if err != nil {
			log.Fatalf("HTTP status 400 - Bad Request: %s", err)
		}
	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}
}

func main() {
	// HTTP endpoint GET, sends HTML response to client. HTTP endpoint POST, sends response back to client
	http.HandleFunc("/", MainPage)

	// Handle css folder
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css/"))))

	// Set up server, listen for localhost:8080, throw error if unable to
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP status 500 - Internal server error: %s", err)
	}
}

func asciiArt(textInput string, banner string) []string {
	// ascii-art font fileName
	fileName := "banners/" + banner + ".txt"

	// Divide words when encountering '\n'
	listOfWords := strings.Split(textInput, "\\n")
	var lines []string

	for _, word := range listOfWords {
		if word != "" {
			for j := 1; j <= 8; j++ {
				line := ""
				for _, l := range word {
					spaceIndexOfFileLine := 31
					breakPoint := 0

					for i := 1; i <= 847; i += 9 { // The last space before letter
						spaceIndexOfFileLine += 1

						if spaceIndexOfFileLine == int(l) {
							breakPoint = i
							break
						}
					}
					num := breakPoint + j
					line += ReadExactLine(fileName, num)
					// Delets all inviseble chars in order to have clear data
					for _, t := range []string{"\a", "\b", "\t", "\n", "\v", "\f", "\r"} {
						line = strings.TrimSuffix(line, t)
					}
				}
				line += "\n"
				lines = append(lines, line)
			}
		} else {
			lines = append(lines, "\n")
		}
	}
	return lines
}

func ReadExactLine(fileName string, lineNumber int) string {
	// Open the input file
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occurred! ", err)
	}
	// Set up a new reader
	br := bufio.NewReader(inputFile)
	// Loop through the file until finds the line needed
	for i := 1; i < lineNumber; i++ {
		_, _ = br.ReadString('\n')
	}
	// Return the line
	str, _ := br.ReadString('\n')
	return str
}
