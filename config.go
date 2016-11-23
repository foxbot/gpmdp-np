package main

import(
	"log"
	"io/ioutil"
	
	"github.com/BurntSushi/toml"
)

type Config struct {
	UpdatePause		bool		`toml:"UpdateTextOnPause"`
	DownloadCover	bool		`toml:"DownloadTrackCover"`
	Outputs			[]Output 	`toml:"Output"`
}

type Output struct {
	Path			string		`toml:"Path"`
	Format			string 		`toml:"Format"`
}

func loadConfig() *Config {
	var conf Config

	contents, err := ioutil.ReadFile("./config.toml")
	if err != nil {
		log.Fatal(err)	
	}
	text := string(contents)

	_, err = toml.Decode(text, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return &conf
}