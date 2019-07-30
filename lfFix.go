package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	var re = regexp.MustCompile(`(?m)\n\d+\|`)
	// re.FindStringIndex
	i := re.FindStringIndex(string(data))

	if len(i) > 0 {
		return i[0] + 1, data[0:i[0]], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func main() {
	file, err := os.Open("logstring.txt")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("lines")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	ew, err := os.Create("error")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ew.Close()
	log.SetOutput(ew)

	scanner := bufio.NewScanner(file)
	scanner.Split(crunchSplitFunc)
	for scanner.Scan() {
		s := strings.Replace(scanner.Text(), "\\\n", `\n`, -1)
		_, err = fmt.Fprintln(f, s)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

}
