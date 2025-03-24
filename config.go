package main

type Config struct {
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
}

func NewConfig() *Config {
	return &Config{}
}
