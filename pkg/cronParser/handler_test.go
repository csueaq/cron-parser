package cronParser

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockService struct {
	mock.Mock
}

func (mock *mockService) ParseInput(cronString string) (Cron, error) {
	args := mock.Called(cronString)
	return args.Get(0).(Cron), args.Error(1)
}

func (mock *mockService) Print(cron Cron) error {
	args := mock.Called(cron)
	return args.Error(0)
}

func TestProcessUserInput(t *testing.T) {
	//given
	svc := &mockService{}
	handler := CreateHandler(svc)
	svc.On("ParseInput", mock.Anything).Return(Cron{}, nil)
	svc.On("Print", mock.Anything).Return(nil)

	// when
	err := handler.ProcessUserInput([]string{"*/15 0 1,15 * 1-5 /usr/bin/find"})

	// then
	assert := assert.New(t)
	assert.Equal(nil, err, "should equal")

}

func TestProcessUserInput_Fail_Parse_error(t *testing.T) {
	//given
	svc := &mockService{}
	handler := CreateHandler(svc)
	svc.On("ParseInput", mock.Anything).Return(Cron{}, errors.New("oops"))
	svc.On("Print", mock.Anything).Return(nil)

	// when
	err := handler.ProcessUserInput([]string{"*/15 0 1,15 * 1-5 /usr/bin/find"})

	// then
	assert := assert.New(t)
	assert.NotEqual(nil, err, "should equal")

}

func TestProcessUserInput_Fail_Print_error(t *testing.T) {
	//given
	svc := &mockService{}
	handler := CreateHandler(svc)
	svc.On("ParseInput", mock.Anything).Return(Cron{}, nil)
	svc.On("Print", mock.Anything).Return(errors.New("oops"))

	// when
	err := handler.ProcessUserInput([]string{"*/15 0 1,15 * 1-5 /usr/bin/find"})

	// then
	assert := assert.New(t)
	assert.NotEqual(nil, err, "should equal")

}

func TestProcessUserInput_Fail_No_Arg(t *testing.T) {
	//given
	svc := &mockService{}
	handler := CreateHandler(svc)

	// when
	err := handler.ProcessUserInput([]string{})

	// then
	assert := assert.New(t)
	assert.NotEqual(nil, err, "should equal")

}

func TestProcessUserInput_Fail_Too_Many_Arg(t *testing.T) {
	//given
	svc := &mockService{}
	handler := CreateHandler(svc)

	// when
	err := handler.ProcessUserInput([]string{"a", "b"})

	// then
	assert := assert.New(t)
	assert.NotEqual(nil, err, "should equal")

}
