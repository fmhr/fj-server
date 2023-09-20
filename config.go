package main

type Config struct {
	Cmd         string `toml:"Cmd"`
	Reactive    bool   `toml:"Reactive"`
	TesterPath  string `toml:"TesterPath"`
	VisPath     string `toml:"VisPath"`
	GenPath     string `toml:"GenPath"`
	InfilePath  string `toml:"InfilePath"`
	OutfilePath string `toml:"OutfilePath"`
	Jobs        int    `toml:"Jobs"`
}
