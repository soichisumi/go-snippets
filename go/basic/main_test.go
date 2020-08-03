package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"io"
	"strings"
	"testing"
)

func Test_json(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}
	tests := []struct {
		name string
		in string
		exp User
	}{
		{
			name: "success",
			in: "{\"name\": \"name\", \"email\":\"mail@example.com\"}",
			exp: User{Name: "name", Email: "mail@example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{}
			if err := json.NewDecoder(strings.NewReader(tt.in)).Decode(&u); err != nil && err != io.EOF {
				t.Errorf("err: %+v\n", err)
				return
			}
			if !cmp.Equal(tt.exp, u) {
				t.Errorf("exp: %+v, act: %+v\n", tt.exp, u)
				return
			}
			fmt.Printf("res: %+v\n", u)
		})
	}
}
