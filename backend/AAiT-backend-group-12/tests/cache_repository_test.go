package tests

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	"blog_api/repository"
	redis_service "blog_api/infrastructure/redis"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
)

type CacheRepositorySuite struct {
	suite.Suite
	cacheRepo domain.CacheRepositoryInterface
	redisClient *redis.Client
}

// SetupSuite initializes the Redis client and repository
func (suite *CacheRepositorySuite) SetupSuite() {
	// load environment variables
	err := env.LoadEnvironmentVariables("../.env")
	suite.Require().NoError(err)

	redisClient, err := redis_service.ConnectStore(env.ENV.REDIS_URL)
	suite.Require().NoError(err)

	_, err = redisClient.Ping().Result()
	suite.Require().NoError(err)

	suite.redisClient = redisClient
	suite.cacheRepo = repository.NewCacheRepository(redisClient)
}

// TeardownSuite disconnects the Redis client after tests
func (suite *CacheRepositorySuite) TearDownSuite() {
	redis_service.DisconnectStore(suite.redisClient)
}

// SetupTest clears the Redis database before each test
func (suite *CacheRepositorySuite) SetupTest() {
	suite.redisClient.FlushDB().Err()
}

// TestCacheData_Success tests successful caching of data
func (suite *CacheRepositorySuite) TestCacheData_Success() {
	key := "test_key"
	value := "test_value"
	expiration := time.Minute * 1

	err := suite.cacheRepo.CacheData(key, value, expiration)
	suite.Require().NoError(err)

	cachedValue, err := suite.cacheRepo.GetCacheData(key)
	suite.Require().NoError(err)
	suite.Equal(value, cachedValue)
}

// TestCacheData_Overwrite tests caching where the key already exists
func (suite *CacheRepositorySuite) TestCacheData_Overwrite() {
	key := "test_key"
	value1 := "first_value"
	value2 := "second_value"
	expiration := time.Minute * 1

	// Cache the first value
	err := suite.cacheRepo.CacheData(key, value1, expiration)
	suite.Require().NoError(err)

	// Overwrite with a second value
	err = suite.cacheRepo.CacheData(key, value2, expiration)
	suite.Require().NoError(err)

	cachedValue, err := suite.cacheRepo.GetCacheData(key)
	suite.Require().NoError(err)
	suite.Equal(value2, cachedValue)
}

// TestCacheData_Expiration tests caching with expiration
func (suite *CacheRepositorySuite) TestCacheData_Expiration() {
	key := "test_key"
	value := "test_value"
	expiration := time.Millisecond * 100 // Short expiration time

	err := suite.cacheRepo.CacheData(key, value, expiration)
	suite.Require().NoError(err)

	time.Sleep(expiration + time.Millisecond*50) // Wait for expiration

	cachedValue, err := suite.cacheRepo.GetCacheData(key)
	suite.Error(err)
	suite.Equal(domain.NewError(string(redis.Nil), domain.ERR_INTERNAL_SERVER), err) // Expect no value (expired)
	suite.Equal("", cachedValue)
}

// TestIsCached_Success tests IsCached with an existing key
func (suite *CacheRepositorySuite) TestIsCached_Success() {
	key := "test_key"
	suite.redisClient.Set(key, "test_value", time.Minute)

	isCached := suite.cacheRepo.IsCached(key)
	suite.True(isCached)
}

// TestIsCached_NotCached tests IsCached with a non-existing key
func (suite *CacheRepositorySuite) TestIsCached_NotCached() {
	key := "non_existent_key"

	isCached := suite.cacheRepo.IsCached(key)
	suite.False(isCached)
}

// TestGetCacheData_Success tests successful retrieval of cached data
func (suite *CacheRepositorySuite) TestGetCacheData_Success() {
	key := "test_key"
	value := "test_value"
	suite.redisClient.Set(key, value, time.Minute)

	cachedValue, err := suite.cacheRepo.GetCacheData(key)
	suite.Require().NoError(err)
	suite.Equal(value, cachedValue)
}

// TestGetCacheData_NonExistentKey tests retrieval of non-existent key
func (suite *CacheRepositorySuite) TestGetCacheData_NonExistentKey() {
	key := "non_existent_key"

	_, err := suite.cacheRepo.GetCacheData(key)
	suite.Error(err)
	suite.Equal(domain.ERR_INTERNAL_SERVER, err.GetCode())
}

// TestCacheDataAndRetrieve tests a full cycle: caching, checking, and retrieving
func (suite *CacheRepositorySuite) TestCacheDataAndRetrieve() {
	key := "test_key"
	value := "test_value"
	expiration := time.Minute * 1

	// Cache the value
	err := suite.cacheRepo.CacheData(key, value, expiration)
	suite.Require().NoError(err)

	// Check if it's cached
	isCached := suite.cacheRepo.IsCached(key)
	suite.True(isCached)

	// Retrieve the cached value
	cachedValue, err := suite.cacheRepo.GetCacheData(key)
	suite.Require().NoError(err)
	suite.Equal(value, cachedValue)
}

// Run the test suite
func TestCacheRepositorySuite(t *testing.T) {
	suite.Run(t, new(CacheRepositorySuite))
}
