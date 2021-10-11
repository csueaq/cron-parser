package cronParser

import (
	"errors"
	"strconv"
	"strings"
)

type ParserImp struct {
}

func (p ParserImp) Parse(min int, max int, item string) (Config, error) {
	config := Config{
		Min: min,
		Max: max,
	}
	var err error
	var hasConfigDone bool
	hasConfigDone, err = p.checkStep(item, &config)

	if !hasConfigDone && err == nil {
		hasConfigDone, err = p.checkValueList(item, &config)

		if !hasConfigDone && err == nil {
			hasConfigDone, err = p.checkValueRange(item, &config)

			if !hasConfigDone && err == nil {
				hasConfigDone, err = p.checkValueSingle(item, &config)
			}
		}
	}

	if !hasConfigDone {
		err = errors.New("unable to parse item")
	}
	return config, err
}

func (p ParserImp) checkStep(item string, config *Config) (bool, error) {

	var err error
	var isStepConfigured bool
	steps := strings.Split(item, "/")
	listOfValues := make([]int, 0)

	if len(steps) == 2 {

		var stepInterval int
		var stepStart int
		stepInterval, err = strconv.Atoi(steps[1])
		if steps[0] == "*" {
			stepStart = config.Min
		} else {
			stepStart, err = strconv.Atoi(steps[0])
		}
		if err == nil {

			for i := stepStart; i <= config.Max; i = i + stepInterval {
				listOfValues = append(listOfValues, i)
			}

		}
	}

	if len(listOfValues) > 0 {
		isStepConfigured = true
		config.values = listOfValues
	}

	return isStepConfigured, err
}

func (p ParserImp) checkValueList(item string, config *Config) (bool, error) {
	var err error
	var isListConfigured bool
	listOfValues := make([]int, 0)
	list := strings.Split(item, ",")

	if len(list) > 1 {
		for _, k := range list {
			var value int
			value, err = strconv.Atoi(k)
			if value >= config.Min && value <= config.Max {
				listOfValues = append(listOfValues, value)
			}
		}
	}

	if len(listOfValues) > 0 {
		config.values = listOfValues
		isListConfigured = true
	}

	return isListConfigured, err
}

func (p ParserImp) checkValueRange(item string, config *Config) (bool, error) {
	var err error
	var isRangeConfigured bool
	listOfValues := make([]int, 0)
	list := strings.Split(item, "-")

	if len(list) == 2 {
		lowerBound, errLower := strconv.Atoi(list[0])
		upperBound, errUpper := strconv.Atoi(list[1])

		if errLower == nil && errUpper == nil {
			if lowerBound <= upperBound && lowerBound >= config.Min && upperBound <= config.Max {
				for i := lowerBound; i <= upperBound; i++ {
					listOfValues = append(listOfValues, i)
				}

			} else {
				err = errors.New("invalid range value")
			}
		} else {
			err = errors.New("invalid range value input")
		}
	}

	if len(listOfValues) > 0 {
		isRangeConfigured = true
		config.values = listOfValues
	}
	return isRangeConfigured, err
}

func (p ParserImp) checkValueSingle(item string, config *Config) (bool, error) {
	var err error
	var isSingleConfigured bool
	listOfValues := make([]int, 0)

	// check *
	if item == "*" {
		for i := config.Min; i <= config.Max; i++ {
			listOfValues = append(listOfValues, i)
		}
	} else {
		var value int
		value, err = strconv.Atoi(item)
		if err == nil {
			listOfValues = append(listOfValues, value)
		}

	}

	if len(listOfValues) > 0 {
		isSingleConfigured = true
		config.values = listOfValues
	}

	return isSingleConfigured, err
}

func CreateParser() Parser {

	return &ParserImp{}
}
