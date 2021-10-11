package cronParser

type Cron struct {
	Command string
	Configs []Config
}

type Config struct {
	Min    int
	Max    int
	values []int
}
