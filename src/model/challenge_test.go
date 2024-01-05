package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ParseInputTimeSuite struct {
	suite.Suite
}

func TestParseInputTimeSuite(t *testing.T) {
	suite.Run(t, new(ParseInputTimeSuite))
}

func (suite *ParseInputTimeSuite) TestEmpty() {
	input := ""

	_, err := parseInputTime(input)

	suite.Assert().ErrorIs(err, ErrInvalidTime)
}

func (suite *ParseInputTimeSuite) TestFirstTime() {
	input := "1:00"

	expected := time.Time{}
	expected = expected.Add(time.Hour)

	actual, err := parseInputTime(input)

	suite.Assert().NoError(err)
	suite.Assert().Equal(expected, actual)
}
