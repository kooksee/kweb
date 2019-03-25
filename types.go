package kweb

type DbConfig struct {
	Schema       string
	Name         string
	DbUrl        string
	MaxIdleConns int
}
