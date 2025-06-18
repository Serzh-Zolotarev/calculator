package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var inputFileName, outputFileName string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите имя файла для считывания примеров:")

	inputFileName, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	inputFile, err := os.OpenFile(strings.TrimSpace(inputFileName), os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	fileReader := bufio.NewReader(inputFile)

	fmt.Println("Введите имя файла для вывода результатов:")

	outputFileName, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`(?P<first>[0-9]+)(?P<math>[+\-/*])(?P<second>[0-9]+)=\?`)
	reRes := regexp.MustCompile(`\?`)

	var res string

	for {
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}

		matches := re.FindStringSubmatch(string(line))
		if len(matches) != 0 {
			first, _ := strconv.Atoi(matches[1])
			math := matches[2]
			second, _ := strconv.Atoi(matches[3])

			switch math {
			case "+":
				res += reRes.ReplaceAllString(string(line), fmt.Sprint(first+second))
			case "-":
				res += reRes.ReplaceAllString(string(line), fmt.Sprint(first-second))
			case "/":
				res += reRes.ReplaceAllString(string(line), fmt.Sprint(first/second))
			case "*":
				res += reRes.ReplaceAllString(string(line), fmt.Sprint(first*second))
			}

			res += "\n"
		}
	}

	file, err := os.OpenFile(strings.TrimSpace(outputFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte(res))
	if err != nil {
		panic(err)
	}

	if err = writer.Flush(); err != nil {
		panic(err)
	}
}
