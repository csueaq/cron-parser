package cronParser

import (
	"errors"
)

type HandlerImp struct {
	service Service
}

func (h HandlerImp) ProcessUserInput(input []string) error {

	err := h.checkUserInput(input)

	if err == nil {
		h.service.Parse(input[0])
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
