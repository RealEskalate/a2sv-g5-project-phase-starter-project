package config_test

import (
	"blogs/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CacheTestSuite struct {
	suite.Suite
	cacheMock *mocks.Cache
}

func (suite *CacheTestSuite) SetupTest() {
	suite.cacheMock = new(mocks.Cache)
}

func (suite *CacheTestSuite) TestGetCache() {
    // Setup expected behavior with the exact cache key used in the test
    suite.cacheMock.On("GetCache", "some_key").Return("some_value", nil)

    // Call the method under test
    result, err := suite.cacheMock.GetCache("some_key")

    // Assert results
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), "some_value", result)

    // Verify the mock expectations
    suite.cacheMock.AssertExpectations(suite.T())
}


func (suite *CacheTestSuite) TestSetCache() {
	// Setup expected behavior
	suite.cacheMock.On("SetCache", "some_key", "some_value").Return (nil)

	// Call the method
	err := suite.cacheMock.SetCache("some_key", "some_value")

	// Assert results
	assert.NoError(suite.T(), err)
}

func (suite *CacheTestSuite) TestDeleteCache() {
	// Setup expected behavior
	suite.cacheMock.On("DeleteCache", "some_key").Return (nil)

	// Call the method
	err := suite.cacheMock.DeleteCache("some_key")

	// Assert results
	assert.NoError(suite.T(), err)
}

func (suite *CacheTestSuite) TearDownTest() {
	suite.cacheMock.AssertExpectations(suite.T())
}

func (suite *CacheTestSuite) TearDownSuite() {
	suite.cacheMock = nil
}


func TestCacheTestSuite(t *testing.T) {
	suite.Run(t, new(CacheTestSuite))
}