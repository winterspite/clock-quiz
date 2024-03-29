package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	Time1h00m = "1:00"
	Time2h00m = "2:00"
	Time3h00m = "3:00"
)

type InternalCheckSuite struct {
	suite.Suite
	Challenge *Challenge
}

func TestInternalCheckSuite(t *testing.T) {
	suite.Run(t, new(InternalCheckSuite))
}

func (suite *InternalCheckSuite) SetupTest() {
	suite.Challenge = &Challenge{}
}

func (suite *InternalCheckSuite) TestEmpty() {
	score, err := suite.Challenge.InternalCheck()

	suite.Assert().Error(err)
	suite.Assert().Equal(ScoreInvalid, score)
}

func (suite *InternalCheckSuite) TestInvalidInput() {
	suite.Challenge.Clock1InputString = "asdf"

	score, err := suite.Challenge.InternalCheck()

	suite.Assert().ErrorIs(err, ErrInvalidTime)
	suite.Assert().Equal(ScoreInvalid, score)
}

func (suite *InternalCheckSuite) TestValidTimesNoClocks() {
	suite.Challenge.Clock1InputString = Time1h00m
	suite.Challenge.Clock2InputString = Time2h00m
	suite.Challenge.DifferenceInputString = Time2h00m

	score, err := suite.Challenge.InternalCheck()

	suite.Assert().ErrorIs(err, ErrInvalidClock1Time)
	suite.Assert().Equal(ScoreIncorrect, score)
}

func (suite *InternalCheckSuite) TestValidTimesForClocksWrongDifference() {
	suite.Challenge.Clock1Time = time.Date(1, 1, 1, 1, 00, 00, 00, time.UTC)
	suite.Challenge.Clock1InputString = Time1h00m
	suite.Challenge.Clock2Time = time.Date(1, 1, 1, 2, 00, 00, 00, time.UTC)
	suite.Challenge.Clock2InputString = Time2h00m
	suite.Challenge.DifferenceInputString = Time2h00m
	suite.Challenge.Difference = suite.Challenge.Clock2Time.Sub(suite.Challenge.Clock1Time)

	score, err := suite.Challenge.InternalCheck()

	suite.Assert().ErrorIs(err, ErrInvalidDifference)
	suite.Assert().Equal(ScoreIncorrect, score)
}

func (suite *InternalCheckSuite) TestValidTimesForClocksCorrectDifference() {
	suite.Challenge.Clock1Time = time.Date(1, 1, 1, 1, 00, 00, 00, time.UTC)
	suite.Challenge.Clock1InputString = Time1h00m
	suite.Challenge.Clock2Time = time.Date(1, 1, 1, 2, 00, 00, 00, time.UTC)
	suite.Challenge.Clock2InputString = Time2h00m
	suite.Challenge.DifferenceInputString = Time1h00m
	suite.Challenge.Difference = suite.Challenge.Clock2Time.Sub(suite.Challenge.Clock1Time)

	score, err := suite.Challenge.InternalCheck()

	suite.Assert().NoError(err)
	suite.Assert().Equal(ScoreCorrect, score)
}

func (suite *InternalCheckSuite) Test12HourDifferences() {
	suite.Challenge.Clock1Time = time.Date(1, 1, 1, 1, 00, 00, 00, time.UTC)
	suite.Challenge.Clock1InputString = Time1h00m
	suite.Challenge.Clock2Time = time.Date(1, 1, 1, 15, 00, 00, 00, time.UTC)
	suite.Challenge.Clock2InputString = Time3h00m
	suite.Challenge.DifferenceInputString = Time2h00m
	suite.Challenge.Difference = suite.Challenge.Clock2Time.Sub(suite.Challenge.Clock1Time)

	/*
		This test case helps to account for when one clock is showing a PM time, but we can't reflect that easily.
	*/

	score, err := suite.Challenge.InternalCheck()

	suite.Assert().NoError(err)
	suite.Assert().Equal(ScoreCorrect, score)
}

func (suite *InternalCheckSuite) Test12HourDifferences2() {
	suite.Challenge.Clock1Time = time.Date(1, 1, 1, 8, 30, 00, 00, time.UTC)
	suite.Challenge.Clock1InputString = "8:30"
	suite.Challenge.Clock2Time = time.Date(1, 1, 1, 20, 25, 00, 00, time.UTC)
	suite.Challenge.Clock2InputString = "8:25"
	suite.Challenge.DifferenceInputString = "11:55"
	suite.Challenge.Difference = suite.Challenge.Clock2Time.Sub(suite.Challenge.Clock1Time)

	/*
		This test case helps to account for when one clock is showing a PM time, but we can't reflect that easily.
	*/

	score, err := suite.Challenge.InternalCheck()

	suite.Assert().NoError(err)
	suite.Assert().Equal(ScoreCorrect, score)
}
