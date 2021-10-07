package cronParser

import "fmt"

type ServiceImp struct {
}

func (s ServiceImp) Parse(cronString string) (Cron, error) {
	var cron Cron
	var err error
	fmt.Println(cronString)

	return cron, err
}

func CreateService() Service {

	return &ServiceImp{}
}
