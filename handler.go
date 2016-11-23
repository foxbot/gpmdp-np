package main

import(
	"log"
	"io/ioutil"
	"strings"

	"github.com/mitchellh/mapstructure"
)

var(
	track 		Track
	timeInfo	Time
	playState	bool = 	false

	config 		*Config
)

func handleMessage(msg *Message) {
	switch msg.Channel {
		case "API_VERSION":
			version := msg.Payload
			log.Println("API Version", version)
			if version != ApiVersion {
				log.Println("WARNING -> This version of the GPMDP API may not be supported.")
			}
		case "track":
			err := mapstructure.Decode(msg.Payload, &track)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Track Changed:", track)
			notify()
		case "time":
			err := mapstructure.Decode(msg.Payload, &timeInfo)
			if err != nil {
				log.Fatal(err)
			}
		case "playState":
			state, ok := msg.Payload.(bool)
			if !ok {
				log.Println("playState->", msg.Payload, "failed to parse as bool")
			}
			playState = state
			log.Println("Playing:", playState)
			notify()
		default:
			return
	}
}

func notify() {
	if config == nil {
		config = loadConfig()
	}

	if config.DownloadCover && len(track.AlbumArtURL) > 0 {
		downloadFromUrl(track.AlbumArtURL, "cover.png")
	} else if config.DownloadCover {
		log.Println("No album cover for track (len:", len(track.AlbumArtURL), ")")
	}

	for _, output := range config.Outputs {
		write := output.Format
		if config.UpdatePause && !playState {
			write = ""
		} else if !playState {
			return
		} else {
			write = strings.Replace(write, "{{track}}", track.Title, 10)
			write = strings.Replace(write, "{{artist}}", track.Artist, 10)
			write = strings.Replace(write, "{{album}}", track.Album, 10)
		}

		err := ioutil.WriteFile(output.Path, []byte(write), 0666)
		if err != nil {
			log.Println("Error when writing file:",err)
		}
	}
}