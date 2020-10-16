package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/soichisumi/go-util/logger"
	"go.uber.org/zap"
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

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}
type NoTagType struct {
	A string
	B int64
}
// considering given this struct
type ExternalOrganization struct {
	Name string `json:"name"`
	Users []User `json:"users"`
	Data NoTagType
}

type Organization struct {
	ExternalOrganization
}

func (o *Organization) UnmarshalJSON(b []byte) error {
	var m struct{
		*ExternalOrganization
		Data2 struct {
			A string `json:"wowCorruptedKeyA"`
			B int64 `json:"wowwwCuriousKeyB"`
		}`json:"data"`
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	logger.Info("", zap.Any("m", m), zap.Any("o", o))
	o.Name = m.Name
	o.Users = m.Users
	o.Data = NoTagType{
		A: m.Data2.A,
		B: m.Data2.B,
	}
	return nil
}

func Test_Json_OverrideUnmarshalMethodOfParticularField(t *testing.T) {
	tests := []struct {
		name string
		in string
		exp Organization
	}{
		{
			name: "success",
			in: `{ 
					"name": "organization name",
					"users": [{
						"name": "name",
						"email":"mail@example.com"
					}],
					"data": {"wowCorruptedKeyA": "wiiii", "wowwwCuriousKeyB": 100}
				}`,
			exp: Organization{
				ExternalOrganization: ExternalOrganization{
					Name: "organization name",
					Users: []User{{Name: "name", Email: "mail@example.com"}},
					Data: NoTagType{A: "wiiii", B: 100},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u Organization
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

func Test_SQLx(t *testing.T) {
	tests := []struct {
		name string
	}{
		{ name: "suc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			query, arg, err := sqlx.In("test IN (?)  a", []string{"s", "b", "c"})
			logger.Error("", zap.Error(err))
			logger.Info("", zap.Any("q", query), zap.Any("arg", arg))
		})
	}
}

