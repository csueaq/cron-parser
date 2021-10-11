package cronParser

import (
	"errors"
	"strings"
)

type ServiceImp struct {
	parser Parser
}

func (s ServiceImp) Parse(cronString string) (Cron, error) {
	var cron Cron
	var err error
	items := strings.Split(cronString, " ")

	if len(items) != 6 {
		err = errors.New("invalid input, missing items")
	}

	cron.Command = items[5]

	cronItems := make(map[string]Config)
	cronItems["Minute"] = Config{
		Min: 0,
		Max: 59,
	}
	cronItems["Hour"] = Config{
		Min: 0,
		Max: 23,
	}
	cronItems["DayOfMonth"] = Config{
		Min: 1,
		Max: 31,
	}
	cronItems["Month"] = Config{
		Min: 1,
		Max: 12,
	}
	cronItems["DayOfWeek"] = Config{
		Min: 0,
		Max: 6,
	}

	for i, v := range []string{"Minute", "Hour", "DayOfMonth", "Month", "DayOfWeek"} {
		var cfg Config

		cfg, err = s.parser.Parse(cronItems[v].Min, cronItems[v].Max, items[i])

		if err == nil {
			cron.Configs = append(cron.Configs, cfg)
		}

	}

	return cron, err
}

func CreateService(parser Parser) Service {

	return &ServiceImp{parser: parser}
}
