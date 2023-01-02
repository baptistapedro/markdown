package myfuzz
import "github.com/gomarkdown/markdown"

func Fuzz(data []byte) int {
	markdown.Parse(data, nil)
	return 0
}
