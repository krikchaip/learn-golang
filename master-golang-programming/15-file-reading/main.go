package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// ?? READ the whole file using go embed directive (for go >= v1.16)
//
//go:embed assets/names.csv
var names string

// ?? READING multiple files
//
//go:embed assets/*
var assets embed.FS

//go:embed assets/ages.csv assets/names.csv
var csvs embed.FS

func init() {
	fmt.Println(names) // Winner,Muimui,Aim,Jane

	file, _ := assets.Open("assets/intro.txt")
	defer file.Close()

	content, _ := io.ReadAll(file)
	fmt.Printf("%s\n", content) // Hello, My name is ... I am ... years old. Nice to meet you guys

	ages, _ := csvs.ReadFile("assets/ages.csv")
	fmt.Printf("%s\n", ages) // 13,16,18,20
}

func main() {
	file, err := os.Open(public(".gitkeep"))
	handleErr(err)
	defer file.Close()

	// ?? READ EXACTLY N bytes from a file
	bytes := make([]byte, 2)
	nRead, err := io.ReadFull(file, bytes)
	handleErr(err)

	log.Printf("Number of bytes read: %d\n", nRead) // 2
	log.Printf("Data read: %s\n", bytes)            // EI

	// ?? READ the whole file into memory
	content, _ := os.ReadFile(public(".gitkeep"))
	log.Printf("%s\n", content) // EIEI ... YOU ARE PRANKED!

	// ?? READ til the end, start from current file pointer
	// ** continue reading from byte 3 and so on (because we read 2bytes before)
	content, _ = io.ReadAll(file)
	log.Printf("%s\n", content) // EI ... YOU ARE PRANKED!

	// ?? RESET file cursor position to START
	// ** by calling io.ReadAll @line29, the file cursor has already moved to EOF.
	file.Seek(0, io.SeekStart)

	// ?? READ a file LINE by LINE - using Scanner
	scanner := bufio.NewScanner(file)

	// ** the default scanner is bufio.ScanLines
	// ** and that means it will scan a file line by line.
	// scanner.Split(bufio.ScanLines)

	// ** read the first line
	if !scanner.Scan() {
		handleErr(scanner.Err())
	}

	// ?? get the data from the last scanner.Scan()
	log.Println(scanner.Text()) // EIEI ... YOU ARE PRANKED!

	// ** skip CSV header
	if !scanner.Scan() {
		handleErr(scanner.Err())
	}

	var students []Student

	// ?? read the whole remaining part of the file
	for ok := scanner.Scan(); ok; ok = scanner.Scan() {
		if !ok {
			handleErr(scanner.Err())
			break
		}

		texts := strings.Split(scanner.Text(), ",")

		name := texts[0]
		age, _ := strconv.ParseUint(texts[1], 10, 32)
		grade, _ := strconv.ParseFloat(texts[2], 32)

		students = append(students, Student{
			name,
			uint(age),
			float32(grade),
		})
	}

	log.Println(students)
}

func handleErr(err error) {
	if err != nil {
		log.Printf("err type: %T\n", err)

		// ** will add date and time prefix to every log message
		// ** eg. YYYY/MM/DD HH:MM:SS {message}
		log.Fatal(err)
	}
}

func public(path string) string {
	if rune(path[0]) != '/' {
		path = "/" + path
	}

	return "./public" + path
}

// implements Stringer interface
type Student struct {
	name  string
	age   uint
	grade float32
}

func (s Student) String() string {
	return fmt.Sprintf(
		"Student{name:%s,age:%d,grade:%.1f}",
		s.name,
		s.age,
		s.grade,
	)
}
