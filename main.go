package main

import (
	"fmt"
	"github.com/bigheadgeorge/lyrics/genius"
	"os"
	"regexp"
	"strings"
)

func printerr(s string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "lyrics: "+s+"\n", v...)
}

func main() {
	if len(os.Args) > 2 {
		printerr("too many args")
		return
	} else if len(os.Args) == 1 {
		printerr("missing args")
		return
	}

	var lyrics []string
	if m, _ := regexp.MatchString(`[\w\d\s]{1,}:[\w\d\s]{1,}`, os.Args[1]); m {
		args := strings.Split(os.Args[1], ":")
		var err error
		lyrics, err = genius.Lyrics(args[0], args[1])
		if err != nil {
			printerr("error grabbing song lyrics: %s", err.Error())
		}
	}

	lines := len(lyrics)
	for i, line := range lyrics {
		fmt.Println(line)
		if i != lines-1 {
			if lyrics[i+1][0] == '[' {
				fmt.Println()
			}
		}
	}
}
