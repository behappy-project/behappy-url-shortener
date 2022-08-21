package model

import (
	"behappy-url-shortener/src/common"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang-module/carbon/v2"
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

func Set(longUrl string, startDate, endDate carbon.DateTime, expired time.Duration, cNew bool, callback func(err error, reply map[string]string)) {
	findUrl(longUrl, func(err error, reply string) {
		if err != nil {
			callback(err, nil)
		} else if reply != "" && !cNew {
			_m := make(map[string]string)
			_m["hash"] = reply
			_m["long_url"] = longUrl
			callback(nil, _m)
		} else {
			uniqId(func(err error, hash string) {
				if err != nil {
					callback(err, nil)
				} else {
					_m := make(map[string]string)
					_m["hash"] = hash
					_m["long_url"] = longUrl
					keyHash := kHash(hash)
					keyUrl := kUrl(longUrl)
					ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
					defer cancel()
					pipe := rdb.TxPipeline()
					// Execute
					//     MULTI
					//     xxxx
					//     EXEC
					// using one rdb-server roundtrip.
					pipe.Set(ctx, keyUrl, hash, redis.KeepTTL)
					pipe.HMSet(ctx, keyHash, map[string]interface{}{
						"url":        longUrl,
						"hash":       hash,
						"start_date": startDate.String(),
						"end_date":   endDate.String()})
					if expired != 0 {
						pipe.Expire(ctx, keyUrl, expired)
						pipe.Expire(ctx, keyHash, expired)
					}
					_, err := pipe.Exec(ctx)
					callback(err, _m)
				}
			})
		}

	})
}

func Get(shortUrl string, callback func(err error, reply map[string]string)) {
	findHash(shortUrl, func(err error, reply map[string]string) {
		if err != nil {
			callback(err, nil)
		} else if reply != nil && reply["url"] != "" {
			callback(nil, reply)
		} else {
			callback(common.NotFoundWithMsg("url not exists!"), nil)
		}
	})
}

func findUrl(longUrl string, callback func(err error, reply string)) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	rdbRes := rdb.Get(ctx, kUrl(longUrl))
	callback(nil, rdbRes.Val())
}

func findHash(shortUrl string, callback func(err error, reply map[string]string)) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	rdbRes := rdb.HGetAll(ctx, kHash(shortUrl))
	callback(nil, rdbRes.Val())
}

// Deprecated: Use
// 点击量
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

// 对counter进行自增,再取base64
func uniqId(callback func(err error, hash string)) {
	//randomInt, _ := strutil.String(util.GetRandomInt(float64(9999), float64(999999)))
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	incr := rdb.Incr(ctx, kCounter())
	incrVal, _ := strutil.String(incr.Val())
	hash := strutil.Base64(incrVal)
	callback(nil, hash)
}
