package object

type SysConfig struct {
	Total    total    `toml:"total"`
	RabbitMQ rabbitMQ `toml:"rabbitMQ"`
}

type total struct {
	IsDebug bool `toml:"isDebug"`
}

type rabbitMQ struct {
	Server      string `toml:"server"`
	Port        int    `toml:"port"`
	User        string `toml:"user"`
	Pwd         string `toml:"pwd"`
	VirtualHost string `toml:"virtualHost"`
}
