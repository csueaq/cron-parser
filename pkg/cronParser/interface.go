package cronParser

type Handler interface {
	ProcessUserInput(input []string) error
}

type Service interface {
	ParseInput(cronString string) (Cron, error)
	Print(cron Cron) error
}

type Parser interface {
	Parse(min int, max int, item string) (Config, error)
}
