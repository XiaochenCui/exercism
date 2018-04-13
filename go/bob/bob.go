// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
package bob

import (
	"regexp"
)

func Hey(remark string) string {
	var chat_mechanisms = []struct {
		description string
		input_regex string
		output      string
	}{
		{
			"nothing",
			`^\s*$`,
			"Fine. Be that way!",
		},
		{
			"yell question",
			`^[A-Z\s]+\?$`,
			"Calm down, I know what I'm doing!",
		},
		{
			"question",
			`^.+\?\s*$`,
			"Sure.",
		},
		{
			"yell",
			`^[^a-z]+[A-Z][^a-z]+!?$`,
			"Whoa, chill out!",
		},
	}
	for _, chat_mechanism := range chat_mechanisms {
		var regex = regexp.MustCompile(chat_mechanism.input_regex)
		if matched := regex.MatchString(remark); matched == true {
			return chat_mechanism.output
		}
	}
	return "Whatever."
}
