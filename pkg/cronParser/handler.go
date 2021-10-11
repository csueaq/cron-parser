package cronParser

import (
	"errors"
)

type HandlerImp struct {
	service Service
}

func (h HandlerImp) ProcessUserInput(input []string) error {
	var err error
	var cron Cron
	err = h.checkUserInput(input)

	if err == nil {
		cron, err = h.service.ParseInput(input[0])

		if err == nil {
			err = h.service.Print(cron)
		}
	}

	return err
}

func (h HandlerImp) checkUserInput(input []string) error {
	if len(input) < 1 {
		return errors.New("no input found")
	}

	if len(input) > 1 {
		return errors.New("too many arguments")
	}

	return nil
}
func CreateHandler(service Service) Handler {

	return &HandlerImp{
		service: service,
	}
}
