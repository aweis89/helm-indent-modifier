package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	path := flag.String("file", "", "file to modify")
	inc := flag.Int("inc", 0, "increment by")
	dec := flag.Int("dec", 0, "decrement by")
	flag.Parse()

	content, err := ioutil.ReadFile(*path)
	handleErr(err)
	lines := strings.Split(string(content), "\n")

	indentRegex, err := regexp.Compile(`(n?indent)\s+([0-9]+)`)
	handleErr(err)

	newLines := []string{}
	for _, line := range lines {
		match := indentRegex.FindStringSubmatch(line)
		if len(match) > 2 {
			indent, err := strconv.Atoi(match[2])
			handleErr(err)

			indent += *inc
			indent -= *dec

			line = indentRegex.ReplaceAllString(line, fmt.Sprintf("$1 %d", indent))
		}
		newLines = append(newLines, line)
	}

	newContent := strings.Join(newLines, "\n")
	ioutil.WriteFile(*path, []byte(newContent), fs.ModeAppend)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
