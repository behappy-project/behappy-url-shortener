package model

import (
	"behappy-url-shortener/src/util"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gookit/goutil/strutil"
	"time"
)

var rdb *redis.Client

var (
	_prefix_ = "behappy:"
)

// RedisInit 初始化redis
func RedisInit() {
	addr := RunOpts.RedisHost + ":" + RunOpts.RedisPort
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: RunOpts.RedisPass, // 密码
		DB:       RunOpts.RedisDb,   // 数据库
		PoolSize: 20,                // 连接池大小
	})
}

func Redis() *redis.Client {
	return rdb
}

func Set(longUrl string, startDate, endDate time.Time, cNew bool, expired int, callback func(reply string)) {
	// todo
	findUrl(longUrl, callback)
}

func Get() {

}

func findUrl(longUrl string, callback func(reply string)) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	rdbRes := rdb.Get(ctx, kUrl(longUrl))
	callback(rdbRes.Val())
}

func findHash(shortUrl string, callback func(reply map[string]string)) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	rdbRes := rdb.HGetAll(ctx, kHash(shortUrl))
	callback(rdbRes.Val())
}

func clickLink(shortUrl string) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	rdb.HIncrBy(ctx, kHash(shortUrl), "clicks", 1)
}

// behappy:counter
func kCounter() string {
	return _prefix_ + "counter"
}

// behappy:url:<long_url> <short_url>
func kUrl(url string) string {
	return _prefix_ + "url:" + strutil.MD5(url)
}

// behappy:hash:<id> url <long_url>
// behappy:hash:<id> hash <short_url>
// behappy:hash:<id> clicks <clicks>
func kHash(hash string) string {
	return _prefix_ + "hash:" + hash
}

func uniqId(callback func(reply string)) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	incr := rdb.Incr(ctx, kCounter())
	randomInt, _ := strutil.String(util.GetRandomInt(float64(9999), float64(999999)))
	incrVal, _ := strutil.String(incr.Val())
	hash := strutil.NewBaseEncoder(58).Encode(randomInt + incrVal)
	callback(hash)
}
