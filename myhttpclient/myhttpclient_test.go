package myhttpclient

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	wantWords := []string{
		"Responsive is better than fast.",
		"It’s not fully shipped until it’s fast.",
		"Anything added dilutes everything else.",
		"Practicality beats purity.",
		"Approachable is better than simple.",
		"Mind your words, they are important.",
		"Speak like a human.",
		"Half measures are as bad as nothing at all.",
		"Encourage flow.",
		"Non-blocking is better than blocking.",
		"Favor focus over features.",
		"Avoid administrative distraction.",
		"Design for failure.",
		"Keep it logically awesome.",
	}
	for i := 0; i < 100; i++ {
		code, body, err := Get("https://api.github.com/octocat")
		assert.NoError(t, err)
		switch code {
		case 200:
			assert.Contains(t, body, "MMMMMMMMMMMMMMMMMMMMMMMM")
			r := regexp.MustCompile(`MMMMMMMMMMMMMMMMMMMMMMM   \| ([^|]+) \|`)
			match := r.FindStringSubmatch(body)
			assert.Contains(t, wantWords, match[1])
		case 403:
			assert.Contains(t, body, "API rate limit exceeded")
		default:
			t.Fail()
		}
	}
}
