// Package genius provides a simple API for getting lyrics off of Genius.
package genius

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// parseLyrics returns the lyrics from a Genius page, or an empty []string if it couldn't find them.
func parseLyrics(r io.Reader) (lyrics []string) {
	z := html.NewTokenizer(r)

	var inLyrics bool
	var start bool
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		name, hasAttr := z.TagName()
		if string(name) == "div" && hasAttr {
			k, v, _ := z.TagAttr()
			if string(k) == "class" && string(v) == "lyrics" {
				inLyrics = true
				continue
			}
		}

		if inLyrics {
			text := string(z.Text())
			if text == "sse" {
				start = true
				continue
			}
			if text == "/sse" {
				break
			}
			if start {
				text = strings.TrimSpace(text)
				if len(text) > 0 {
					lyrics = append(lyrics, text)
				}
			}
		}
	}
	return
}

// toSlug converts a string with lots of special characters to a slug usable in Genius urls.
// ex. "Pink + White" returns "pink-white".
func toSlug(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	return strings.Replace(s, " ", "-", -1)
}

// Lyrics returns the lyrics to a song by a given artist.
func Lyrics(artist, song string) (l []string, err error) {
	url := fmt.Sprintf("https://www.genius.com/%s-%s-lyrics", toSlug(artist), toSlug(song))
	return LyricsURL(url)
}

// LyricsURL returns the lyrics from a specific page.
func LyricsURL(url string) (l []string, err error) {
	r, err := http.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()

	l = parseLyrics(r.Body)
	if len(l) == 0 {
		err = fmt.Errorf("no lyrics parsed")
	}
	return
}
