package config

import (
	"bufio"
	"os"
	"strings"
)

func InitEnv() {
	filePath := ".env"

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, l := range lines {
		pair := strings.SplitN(l, "=", 2)
		if len(pair) == 2 {
			err := os.Setenv(pair[0], pair[1])
			if err != nil {
				panic(err)
			}
		}
	}
}
