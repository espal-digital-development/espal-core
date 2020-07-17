package text

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randomStringSrc = rand.NewSource(time.Now().UnixNano())

// RandomString generates a random string based on n-length.
func RandomString(n int) string {
	buffer := make([]byte, n)
	// A randomStringSrc.Int63() generates 63 random bits, enough for letterIdxMax characters
	for i, cache, remain := n-1, randomStringSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randomStringSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			buffer[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(buffer)
}

// LowerFirst is the reverse version of strings.Title() and makes the first letter lower case.
func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// JoinCallback joins the given strings list based on callback behavior.
func JoinCallback(a []string, callback func(v string) string) string {
	builder := strings.Builder{}
	for k := range a {
		builder.WriteString(callback(a[k]))
	}
	return builder.String()
}
