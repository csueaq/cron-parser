package cronParser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {

	svc := CreateService(CreateParser())

	cron, err := svc.Parse("*/15 0 1,15 * 1-5 /usr/bin/find")

	assert := assert.New(t)

	expectedCron := Cron{
		Command: "/usr/bin/find",
		Configs: []Config{
			{
				Min:    0,
				Max:    59,
				values: []int{0, 15, 30, 45},
			},
			{
				Min:    0,
				Max:    23,
				values: []int{0},
			},
			{
				Min:    1,
				Max:    31,
				values: []int{1, 15},
			},
			{
				Min:    1,
				Max:    12,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			},
			{
				Min:    0,
				Max:    6,
				values: []int{1, 2, 3, 4, 5},
			},
		},
	}

	assert.Equal(nil, err, "should equal")
	assert.Equal(expectedCron, cron, "should equal")

}
