package model

import (
	"github.com/stretchr/testify/suite"
	"testing"
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
