package cronParser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSteps(t *testing.T) {

	parser := CreateParser()

	config, err := parser.Parse(1, 10, "*/1")
	assert := assert.New(t)

	assert.Equal(nil, err, "should equal")
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, config.values, "should equal")

}

func TestParseSteps_startValue(t *testing.T) {

	parser := CreateParser()

	config, err := parser.Parse(1, 10, "2/2")
	assert := assert.New(t)

	assert.Equal(nil, err, "should equal")
	assert.Equal([]int{2, 4, 6, 8, 10}, config.values, "should equal")

}

func TestParseList(t *testing.T) {

	parser := CreateParser()

	config, err := parser.Parse(1, 10, "2,3,4,5")
	assert := assert.New(t)

	assert.Equal(nil, err, "should equal")
	assert.Equal([]int{2, 3, 4, 5}, config.values, "should equal")

}

func TestParseRange(t *testing.T) {

	parser := CreateParser()

	config, err := parser.Parse(1, 10, "2-5")
	assert := assert.New(t)

	assert.Equal(nil, err, "should equal")
	assert.Equal([]int{2, 3, 4, 5}, config.values, "should equal")

}

func TestParse_InvalidInput(t *testing.T) {

	parser := CreateParser()

	_, err := parser.Parse(1, 10, "2-a")
	assert := assert.New(t)

	assert.NotEqual(nil, err, "should equal")

}

func TestParse_InvalidRange(t *testing.T) {

	parser := CreateParser()

	_, err := parser.Parse(1, 2, "5-3")
	assert := assert.New(t)

	assert.NotEqual(nil, err, "should equal")

}

func TestParseSingle_all(t *testing.T) {

	parser := CreateParser()

	config, err := parser.Parse(1, 10, "*")
	assert := assert.New(t)

	assert.Equal(nil, err, "should equal")
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, config.values, "should equal")

}

func TestParseSingle_one(t *testing.T) {

	parser := CreateParser()

	config, err := parser.Parse(1, 10, "3")
	assert := assert.New(t)

	assert.Equal(nil, err, "should equal")
	assert.Equal([]int{3}, config.values, "should equal")

}
