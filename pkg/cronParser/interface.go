package cronParser

type Handler interface {
	ProcessUserInput(input []string) error
}

type Service interface {
	Parse(cronString string) (Cron, error)
}

type Parser interface {
	Parse(min int, max int, item string) (Config, error)
}
