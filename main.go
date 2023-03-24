package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	path := flag.String("file", "", "file to modify")
	inc := flag.Int("inc", 0, "increment indent/nindent by this value")
	dec := flag.Int("dec", 0, "decrement indent/nindent by this value")
	startLine := flag.Int("start-line", 1, "ingnore lines before")
	endLine := flag.Int("end-line", math.MaxInt, "ingnore lines after")
	flag.Parse()

	content, err := ioutil.ReadFile(*path)
	handleErr(err)

	res, err := processIndent(content, *inc, *dec, *startLine, *endLine)
	handleErr(err)

	ioutil.WriteFile(*path, res, fs.ModeAppend)
}

func processIndent(content []byte, inc, dec, startLine, endLine int) ([]byte, error) {
	indentRegex, err := regexp.Compile(`(n?indent)\s+([0-9]+)`)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	newLines := []string{}
	for ind, line := range lines {
		lineNum := ind + 1
		if lineNum < startLine || lineNum > endLine {
			newLines = append(newLines, line)
			continue
		}
		match := indentRegex.FindStringSubmatch(line)
		if len(match) > 2 {
			indent, err := strconv.Atoi(match[2])
			handleErr(err)

			indent += inc
			indent -= dec

			line = indentRegex.ReplaceAllString(line, fmt.Sprintf("$1 %d", indent))
		}
		newLines = append(newLines, line)
	}

	return []byte(strings.Join(newLines, "\n")), nil
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
