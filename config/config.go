package config

type Config struct {
	Format     string `json:"format"`
	Emoji      *Emoji `json:"emoji"`
	NotConnMes string `json:"not_connect_message"`
}

type Emoji struct {
	Play       string `json:"play"`
	Pause      string `json:"pause"`
	Stop       string `json:"stop"`
	Connect    string `json:"connect"`
	NotConnect string `json:"not_connect"`
}
