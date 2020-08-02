package main

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {
	assert := assert.New(t)

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	assert.Nil(err)

	var cfg = struct {
		Key string
		Teststruct struct {
			Key string
			Number int64
		}
	}{}
	err = viper.Unmarshal(&cfg)
	assert.Nil(err)

	assert.Equal("value", cfg.Key)
	assert.Equal("value", cfg.Teststruct.Key)
	assert.Equal(int64(13), cfg.Teststruct.Number)
}