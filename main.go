package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	path := flag.String("file", "", "file to modify")
	inc := flag.Int("inc", 0, "increment indent/nindent by this value")
	dec := flag.Int("dec", 0, "decrement indent/nindent by this value")
	startLine := flag.Int("start-line", 1, "ingnore lines before")
	endLine := flag.Int("end-line", math.MaxInt, "ingnore lines after")
	flag.Parse()

	content, err := os.ReadFile(*path)
	handleErr(err)

	res, err := processIndent(content, *inc, *dec, *startLine, *endLine)
	handleErr(err)

	handleErr(os.WriteFile(*path, res, fs.ModeAppend))
}

const indentRegex string = `(n?indent)\s+([0-9]+)`

func processIndent(content []byte, inc, dec, startLine, endLine int) ([]byte, error) {
	indentRegex := regexp.MustCompile(indentRegex)

	scanner := bufio.NewScanner(bytes.NewReader(content))
	newLines := [][]byte{}
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Bytes()
		if lineNum < startLine || lineNum > endLine {
			newLines = append(newLines, line)
			continue
		}
		match := indentRegex.FindSubmatch(line)
		if len(match) > 2 {
			indent, err := strconv.Atoi(string(match[2]))
			handleErr(err)

			indent += inc
			indent -= dec

			line = indentRegex.ReplaceAll(line, []byte(fmt.Sprintf("$1 %d", indent)))
		}
		newLines = append(newLines, line)

		lineNum++
	}

	return bytes.Join(newLines, []byte("\n")), nil
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
