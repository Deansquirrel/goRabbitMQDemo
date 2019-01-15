package object

type SysConfig struct {
	Total total `toml:"total"`
}

type total struct {
	IsDebug bool `toml:"isDebug"`
}
