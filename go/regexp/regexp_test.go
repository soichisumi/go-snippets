package main

import (
	"regexp"
	"testing"
)

func Test_regexp(t *testing.T) {
	tests := []struct {
		name string
		regex string
		input string
		res string
	}{
		{
			name:  "success",
			regex: "^/users/([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}).*?",
			input: "/users/2fe4ca93-040e-4072-a3ab-33353553bd18/innter-resources",
			res: "2fe4ca93-040e-4072-a3ab-33353553bd18",
		},
		{
			name:  "success2",
			regex: "^/users/([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}).*?",
			input: "/users",
			res:  "",
		},
		{
			name:  "not found",
			regex: "^/users/([0-9a-f]{12}4[0-9a-f]{3}[89ab][0-9a-f]{15}).*?",
			input: "/users/0001",
			res: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg, err := regexp.Compile(tt.regex)
			if err != nil {
				t.Errorf(err.Error())
				return
			}

			res := reg.FindStringSubmatch(tt.input)
			if len(res) == 0 && tt.res != "" ||
				(len(res) >= 2 && tt.res != res[1]) {
				t.Errorf("exp: %+v, act: %+v\n", tt.res, res)
				return
			}
		})
	}
}


