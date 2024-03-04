package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type utilSuite struct {
	suite.Suite
}

func (suite *utilSuite) TestIsValidCEP() {
	cep := "13330350"
	isValid := IsValidCEP(cep)
	assert.True(suite.T(), isValid)
}

func (suite *utilSuite) TestIsInValidCEP() {
	cep := "1333035a"
	isValid := IsValidCEP(cep)
	assert.False(suite.T(), isValid)

	cep = "1333035"
	isValid = IsValidCEP(cep)
	assert.False(suite.T(), isValid)

	cep = ""
	isValid = IsValidCEP(cep)
	assert.False(suite.T(), isValid)
}

func (suite *utilSuite) TestIsDigit() {
	cep := "13330350"
	isValid := IsDigit(cep)
	assert.True(suite.T(), isValid)
}

func (suite *utilSuite) TestIsNotDigit() {
	cep := "133aa350"
	isValid := IsDigit(cep)
	assert.False(suite.T(), isValid)
}

func (suite *utilSuite) TestGetEnvVar() {
	varName := "PORT"
	value := GetEnvVariable(varName)
	assert.NotEmpty(suite.T(), value)

	varName = "XPTO"
	value = GetEnvVariable(varName)
	assert.Empty(suite.T(), value)
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, new(utilSuite))
}
