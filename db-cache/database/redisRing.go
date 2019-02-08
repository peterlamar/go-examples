package database

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	"os"
)

var (
	redisRing  *redis.Ring
	cacheCodec *cache.Codec
	codecLocal *cache.Codec

	isCacheConnected bool
)

// ConnectCache initializes the cache cluster
func ConnectCache() {

	connPort := buildRedisEnvString()

	// Connect to Cache cluster
	cacheClient := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": connPort,
		},
	})

	SetRedisRing(cacheClient)

}

// SetRedisRing sets the redis connection and sets up serialization in
// msgpack format
func SetRedisRing(s *redis.Ring) {
	redisRing = s

	cacheCodec = &cache.Codec{
		Redis: redisRing,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	_, err := redisRing.Ping().Result()

	if err != nil {
		log.Debug(err)
		// So the application doesn't crash if Redis connection is missing
		cacheCodec.Redis = nil
		isCacheConnected = false
	} else {
		isCacheConnected = true
	}
}

// IsCacheConnected returns if the cache is connected
func IsCacheConnected() bool {
	return isCacheConnected
}

// GetRedisRing returns the RedisRing connection
func GetRedisRing() *redis.Ring {
	return redisRing
}

// GetRedisCacheCodec returns the cache codec
func GetRedisCacheCodec() *cache.Codec {
	return cacheCodec
}

func buildRedisEnvString() (rtnString string) {

	if os.Getenv("REDIS_PORT") == "" {
		os.Setenv("REDIS_PORT", "6379")
	}

	if os.Getenv("REDIS_HOST") == "" {
		rtnString = ":" + os.Getenv("REDIS_PORT")
	} else {
		rtnString = os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	}

	return
}
