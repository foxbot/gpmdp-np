package main

type Message struct {
	Channel 	string		`json:"channel"`
	Payload 	interface{}	`json:"payload"`
}	

type Track struct {
	Title		string		`json:"title"`
	Artist		string		`json:"artist"`
	Album		string		`json:"album"`
	AlbumArtURL	string		`json:"albumArt"`
}

type Time struct {
	Current		float64		`json:"current"`
	Length		float64		`json:"total"`
}