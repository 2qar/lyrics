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
	var err error
	if m, _ := regexp.MatchString(`[\w\d\s]{1,}:[\w\d\s]{1,}`, os.Args[1]); m {
		args := strings.Split(os.Args[1], ":")
		lyrics, err = genius.Lyrics(args[0], args[1])
	} else if s := regexp.MustCompile(`genius.com\/([\w\d]{1,}-){1,}lyrics`).FindString(os.Args[1]); len(s) > 0 {
		lyrics, err = genius.LyricsURL("https://www." + s)
	} else {
		printerr("bad argument: not a url or artist:song pair")
		return
	}
	if err != nil {
		printerr("%s", err.Error())
		return
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
