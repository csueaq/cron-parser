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

const CRON_ITEMS = "minute,hour,day of month,month,day of week"
