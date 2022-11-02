package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"regexp"
	"strings"
	"time"
)

// Return the path pattern a s3 key path to store the lighthouse result
func GetPathFromUrl(url string) string {
	r := regexp.MustCompile(`https?:\/\/(?P<Domain>[a-zA-Z0-9.]+)(?P<Path>\/[a-zA-Z0-9\/]+)?`)
	result := r.FindStringSubmatch(url)

	hasher := sha1.New()
	hasher.Write([]byte(url))
	id := hex.EncodeToString(hasher.Sum(nil))

	if len(result) == 3 {
		domain := result[1]
		pathing := []string{
			"reports",
			domain,
			nowAsPath(),
			id + ".json",
		}

		return strings.Join(pathing, "/")
	}

	return ""
}

func nowAsPath() string {
	// YYYY/MM/DD
	datePath := time.Now().Format("2006/01/02")
	timePath := time.Now().Format("15")
	return datePath + "/" + timePath
}
