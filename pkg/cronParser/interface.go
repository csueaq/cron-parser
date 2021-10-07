package cronParser

type Handler interface {
	ProcessUserInput(input []string) error
}

type Service interface {
	Parse(cronString string) (Cron, error)
}
