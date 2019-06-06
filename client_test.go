package darksky

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientMatchLang(t *testing.T) {
	for _, tc := range []struct {
		lang     string
		options  []ClientOption
		expected Lang
	}{
		{
			lang:     "cmn", // Mandarin Chinese
			expected: LangZH,
		},
		{
			lang:     "de",
			expected: LangDE,
		},
		{
			lang:     "de-CH",
			expected: LangDE,
		},
		{
			lang:     "en",
			expected: LangEN,
		},
		{
			lang:     "en-US",
			expected: LangEN,
		},
		{
			lang:     "fr",
			expected: LangFR,
		},
		{
			lang:     "fr-CH",
			expected: LangFR,
		},
		{
			lang:     "gsw", // Swiss German
			expected: LangDE,
		},
		{
			lang:     "it",
			expected: LangIT,
		},
		{
			lang:     "it-CH",
			expected: LangIT,
		},
		{
			lang:     "zh",
			expected: LangZH,
		},
		{
			lang:     "zh-TW",
			expected: LangZHTW,
		},
		{
			lang: "en",
			options: []ClientOption{
				WithLangs([]Lang{LangDE, LangFR, LangIT}),
			},
			expected: LangDE,
		},
	} {
		t.Run(tc.lang, func(t *testing.T) {
			c := NewClient(tc.options...)
			assert.Equal(t, tc.expected, c.MatchLang(tc.lang))
		})
	}
}
