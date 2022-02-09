package cache

import (
	"errors"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestCache_Addr(t *testing.T) {
	setEnv()
	defer func() {
		unsetEnv()
	}()
	c := NewENVConfig()
	got := c.Addr()
	assert.Equal(t, "https://testi.ng:27654", got)
}

func TestCache_Pwd(t *testing.T) {
	setEnv()
	defer func() {
		unsetEnv()
	}()
	c := NewENVConfig()
	got := c.Pwd()
	assert.Equal(t, "test_db", got)
}

func TestEngineCompliance(t *testing.T) {
	got := NewRedis(nil, nil)

	assert.Implements(t, (*Engine)(nil), got)
}

func TestEngine_GetSafeNil(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "my-custom-key"
	mock.ExpectGet(key).RedisNil()

	rds := NewRedis(nil, db)

	got := rds.GetSafe("my-custom-key")

	assert.Equal(t, "", got)
}

func TestEngine_GetSafeErr(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "my-custom-key"

	err := errors.New("connection timed out")
	mock.ExpectGet(key).SetErr(err)

	rds := NewRedis(nil, db)

	got := rds.GetSafe("my-custom-key")

	assert.Equal(t, nil, got)
}

func TestEngine_GetSafeSuccess(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "my-custom-key"
	mock.ExpectGet(key).SetVal("testing val")

	rds := NewRedis(nil, db)

	got := rds.GetSafe("my-custom-key")

	assert.Equal(t, "testing val", got)
}

func TestEngine_GetNil(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "my-custom-key"
	mock.ExpectGet(key).RedisNil()

	rds := NewRedis(nil, db)

	got, err := rds.Get("my-custom-key")

	assert.Equal(t, redis.Nil, err)
	assert.Empty(t, got)
}

func TestEngine_GetErr(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "my-custom-key"

	errex := errors.New("connection timed out")
	mock.ExpectGet(key).SetErr(errex)

	rds := NewRedis(nil, db)

	got, err := rds.Get("my-custom-key")

	assert.Equal(t, errex, err)
	assert.Empty(t, got)

	assert.Equal(t, "", got)
}

func TestEngine_GetSuccess(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "my-custom-key"
	mock.ExpectGet(key).SetVal("testing val")

	rds := NewRedis(nil, db)

	got := rds.GetSafe("my-custom-key")

	assert.Equal(t, "testing val", got)
}

func TestEngine_Set(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "my-custom-key"
	mock.ExpectSet(key, "val", redis.KeepTTL).SetVal("val")
	rds := NewRedis(nil, db)

	got := rds.Set(key, "val", redis.KeepTTL)

	assert.NoError(t, got)
}

func setEnv() {
	os.Setenv("CACHE_HOST", "https://testi.ng")
	os.Setenv("CACHE_PORT", "27654")
	os.Setenv("CACHE_PASSWORD", "test_db")
}

func unsetEnv() {
	os.Unsetenv("CACHE_HOST")
	os.Unsetenv("CACHE_PORT")
	os.Unsetenv("CACHE_PASSWORD")
}
