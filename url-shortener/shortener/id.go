package shortener

import (
	"crypto/sha1"
	"fmt"
)

func generateID(url string) string {
	sum := sha1.Sum([]byte(url))
	return fmt.Sprintf("%x", sum)[:12]
}
