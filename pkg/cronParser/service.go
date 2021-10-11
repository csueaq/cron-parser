package cronParser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ServiceImp struct {
	parser Parser
}

func (s ServiceImp) Print(cron Cron) error {
	var err error

	if len(cron.Configs) != 5 {
		err = errors.New("not enough items in cron")
	} else if len(cron.Command) == 0 {
		err = errors.New("no command provided")
	} else {
		itemName := s.getItemNames()

		longestCount := s.getLongestItemNameCount()

		for i := 0; i < 5; i++ {
			numStringArr := make([]string, 0)
			for _, k := range cron.Configs[i].values {
				numStringArr = append(numStringArr, strconv.Itoa(k))
			}

			spaceStringArr := make([]string, 0)
			for j := 0; j < (longestCount - len(itemName[i])); j++ {
				spaceStringArr = append(spaceStringArr, " ")
			}

			println(fmt.Sprintf("%s%s %s", itemName[i], strings.Join(spaceStringArr, ""), strings.Join(numStringArr, " ")))

		}

		spaceStringArr := make([]string, 0)
		for i := 0; i < longestCount-len("command"); i++ {
			spaceStringArr = append(spaceStringArr, " ")
		}
		println(fmt.Sprintf("%s%s %s", "command", strings.Join(spaceStringArr, ""), cron.Command))

	}

	return err
}

func (s ServiceImp) getItemNames() []string {
	return strings.Split(CRON_ITEMS, ",")
}

func (s ServiceImp) getLongestItemNameCount() int {
	max := 0

	for _, k := range s.getItemNames() {

		if len(k) > max {
			max = len(k)
		}
	}

	return max
}
func (s ServiceImp) getItemList() map[string]Config {
	cronItems := make(map[string]Config)
	cronItemNames := s.getItemNames()
	cronItems[cronItemNames[0]] = Config{
		Min: 0,
		Max: 59,
	}
	cronItems[cronItemNames[1]] = Config{
		Min: 0,
		Max: 23,
	}
	cronItems[cronItemNames[2]] = Config{
		Min: 1,
		Max: 31,
	}
	cronItems[cronItemNames[3]] = Config{
		Min: 1,
		Max: 12,
	}
	cronItems[cronItemNames[4]] = Config{
		Min: 0,
		Max: 6,
	}

	return cronItems
}
func (s ServiceImp) ParseInput(cronString string) (Cron, error) {
	var cron Cron
	var err error
	items := strings.Split(cronString, " ")

	if len(items) != 6 {
		return cron, errors.New("invalid input, missing items")
	}

	cron.Command = items[5]

	cronItems := s.getItemList()

	for i, v := range s.getItemNames() {
		var cfg Config

		cfg, err = s.parser.Parse(cronItems[v].Min, cronItems[v].Max, items[i])

		if err == nil {
			cron.Configs = append(cron.Configs, cfg)
		} else {
			break
		}

	}

	return cron, err
}

func CreateService(parser Parser) Service {

	return &ServiceImp{parser: parser}
}
