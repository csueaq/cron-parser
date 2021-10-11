package cronParser

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockParser struct {
	mock.Mock
}

func (mock *mockParser) Parse(min int, max int, item string) (Config, error) {
	args := mock.Called(min, max, item)
	return args.Get(0).(Config), args.Error(1)
}

func TestParse_Success(t *testing.T) {

	//given
	parser := &mockParser{}
	svc := CreateService(parser)
	parser.On("Parse", mock.Anything, mock.Anything, mock.Anything).Return(Config{}, nil)

	//when
	_, err := svc.ParseInput("*/15 0 1,15 * 1-5 /usr/bin/find")

	//then
	assert := assert.New(t)
	assert.Equal(nil, err, "should equal")

}

func TestParse_ParserFail(t *testing.T) {

	//given
	parser := &mockParser{}
	svc := CreateService(parser)
	parser.On("Parse", mock.Anything, mock.Anything, mock.Anything).Return(Config{}, errors.New("oops"))

	//when
	_, err := svc.ParseInput("*/15 0 1,15 * 1-5 /usr/bin/find")

	//then
	assert := assert.New(t)
	assert.NotEqual(nil, err, "should not equal")

}

func TestParse_ParserInput(t *testing.T) {

	//given
	parser := &mockParser{}
	svc := CreateService(parser)

	//when
	_, err := svc.ParseInput("*/15 0 1,15 * 1-5")

	//then
	assert := assert.New(t)
	assert.NotEqual(nil, err, "should not equal")

}

func TestPrint_Success(t *testing.T) {
	//given
	parser := &mockParser{}
	svc := CreateService(parser)
	cronToPrint := Cron{
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

	//when
	err := svc.Print(cronToPrint)

	//then
	assert := assert.New(t)
	assert.Equal(nil, err, "should equal")

}

func TestPrint_Fail_Not_Enough_Item(t *testing.T) {
	//given
	parser := &mockParser{}
	svc := CreateService(parser)
	cronToPrint := Cron{
		Command: "/usr/bin/find",
		Configs: []Config{
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

	//when
	err := svc.Print(cronToPrint)

	//then
	assert := assert.New(t)
	assert.NotEqual(nil, err, "should equal")

}

func TestPrint_Fail_No_Commans(t *testing.T) {
	//given
	parser := &mockParser{}
	svc := CreateService(parser)
	cronToPrint := Cron{
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

	//when
	err := svc.Print(cronToPrint)

	//then
	assert := assert.New(t)
	assert.NotEqual(nil, err, "should equal")

}
func TestParse(t *testing.T) {

	//given
	svc := CreateService(CreateParser())

	//when
	cron, err := svc.ParseInput("*/15 0 1,15 * 1-5 /usr/bin/find")

	//then
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
