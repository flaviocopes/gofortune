package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var files []string

	fortuneCommand := exec.Command("fortune", "-f")
	pipe, err := fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}
	fortuneCommand.Start()
	outputStream := bufio.NewScanner(pipe)
	outputStream.Scan()
	line := outputStream.Text()
	root := line[strings.Index(line, "/"):]

	err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, "/off/") {
			return nil
		}
		if filepath.Ext(path) == ".dat" {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	i := randomInt(1, len(files))
	randomFile := files[i]

	file, err := os.Open(randomFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	quotes := string(b)

	quotesSlice := strings.Split(quotes, "%")
	j := randomInt(1, len(quotesSlice))

	fmt.Print(quotesSlice[j])
}
