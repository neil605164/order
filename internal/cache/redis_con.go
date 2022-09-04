package cache

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"order/app/global"
	"order/app/global/helper"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// IRedis interface
type IRedis interface {
	Ping() error
	Set(key, value string, exp time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) (bool, error)
	Expire(key string, expire time.Duration) error
	HSet(key string, field ...string) error
	HMSet(key string, field map[string]interface{}) error
	HGet(key, field string) (string, error)
	HDel(key, field string) (int64, error)
	Publish(channel string, data []byte) error
	Subscribe(channel string) *redis.PubSub
	LPush(channel string, data []byte) error
	RPush(channel string, data []byte) error
	LPop(channel string) *redis.StringCmd
	LPos(key, value string) *redis.IntCmd
	LRange(key string, start, stop int64) *redis.StringSliceCmd
	LRem(key string, count int64, value interface{}) *redis.IntCmd
	BRPop(channel string, exp time.Duration) *redis.StringSliceCmd
}

// Redis 存取值
type Redis struct{}

var singleton *Redis
var once sync.Once

// redisPool 存放redis連線池的全域變數
var redisPool *redis.Client

func Instance() IRedis {
	once.Do(func() {
		singleton = &Redis{}
	})
	return singleton
}

func PrintRedisPool(stats *redis.PoolStats) {
	fmt.Printf("Hits=%d Misses=%d Timeouts=%d TotalConns=%d IdleConns=%d StaleConns=%d\n",
		stats.Hits, stats.Misses, stats.Timeouts, stats.TotalConns, stats.IdleConns, stats.StaleConns)
}

// RedisPoolConnect 回傳連線池的 Redis 連線
func redisPoolConnect() *redis.Client {

	if redisPool != nil {
		return redisPool
	}

	redisPool = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.RedisHost + ":" + global.Config.Redis.RedisPort,
		Password: global.Config.Redis.RedisPwd, // 密码
		// 连接池容量及闲置连接数量
		PoolSize: 50, // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		// MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		// 超时
		// DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		// ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		// WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		// PoolTimeout:  5 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		// 闲置连接检查包括IdleTimeout，MaxConnAge
		// IdleCheckFrequency: 60 * time.Second, //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		// IdleTimeout:        10 * time.Second, //闲置超时，默认5分钟，-1表示取消闲置超时检查
		// MaxConnAge:         0 * time.Second,  //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

		// 命令执行失败时的重试策略
		// MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		// MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		// MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },

		// ReadOnly = true，只择 Slave Node
		// ReadOnly = true 且 RouteByLatency = true 将从 slot 对应的 Master Node 和 Slave Node， 择策略为: 选择PING延迟最低的点
		// ReadOnly = true 且 RouteRandomly = true 将从 slot 对应的 Master Node 和 Slave Node 选择，选择策略为: 随机选择

		// ReadOnly:       true,
		// RouteRandomly:  true,
		// RouteByLatency: true,
	})

	// 正式站才有 tls 設定
	if helper.IsRelease() || helper.IsStress() {
		redisPool.Options().TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return redisPool
}

// 連線檢查
func (r *Redis) Ping() error {
	pool := redisPoolConnect()
	_, err := pool.Ping(context.TODO()).Result()
	if err != nil {
		pool.Close()
		log.Fatal("🔔🔔🔔 REDIS CONNECT ERROR: 🔔🔔🔔", err.Error())
	}
	return nil
}

// Set 存值
func (r *Redis) Set(key, value string, exp time.Duration) error {
	pool := redisPoolConnect()
	if err := pool.Set(context.TODO(), key, value, exp).Err(); err != nil {
		return err
	}
	return nil
}

// Get 取出指定的值
func (r *Redis) Get(key string) (string, error) {
	pool := redisPoolConnect()
	// 切換 0
	pool.Do(context.TODO(), "select", 0)

	result, err := pool.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return "", nil
	}
	return result, nil
}

// Delete 刪除
func (r *Redis) Delete(key string) error {
	pool := redisPoolConnect()
	// 切換 0
	pool.Do(context.TODO(), "select", 0)

	err := pool.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}
	return nil
}

// Exists Key 是否存在
func (r *Redis) Exists(key string) (bool, error) {
	pool := redisPoolConnect()
	// 切換 0
	pool.Do(context.TODO(), "select", 0)

	exist, err := pool.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	if exist > 0 {
		return true, nil
	}
	return false, nil
}

// Expire Key 設定到期時間
func (r *Redis) Expire(key string, expire time.Duration) error {
	pool := redisPoolConnect()

	// 切換 0
	pool.Do(context.TODO(), "select", 0)
	if err := pool.Expire(context.Background(), key, expire).Err(); err != nil {
		return err
	}

	return nil
}

// HSet 存 hash 值
func (r *Redis) HSet(key string, field ...string) error {
	pool := redisPoolConnect()

	// 切換 0
	pool.Do(context.TODO(), "select", 0)
	err := pool.HSet(context.Background(), key, field).Err()
	if err != nil {
		return err
	}

	return nil
}

// HMSet 存 muti hash 值
func (r *Redis) HMSet(key string, field map[string]interface{}) error {
	pool := redisPoolConnect()

	// 切換 0
	pool.Do(context.TODO(), "select", 0)
	err := pool.HMSet(context.Background(), key, field).Err()
	if err != nil {
		return err
	}

	return nil
}

// HGet 取值
func (r *Redis) HGet(key, field string) (string, error) {
	pool := redisPoolConnect()

	// 切換 0
	pool.Do(context.TODO(), "select", 0)
	data, err := pool.HGet(context.Background(), key, field).Result()
	// 無 key 值
	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return data, nil
}

// HDel 删除 hash 字段
func (r *Redis) HDel(key, field string) (int64, error) {
	pool := redisPoolConnect()

	// 切換 0
	pool.Do(context.TODO(), "select", 0)
	res, err := pool.HDel(context.Background(), key, field).Result()

	if err != nil {
		return res, err
	}
	return res, nil
}

// Publish Redis Pub 事件，for queue 推送使用
func (r *Redis) Publish(channel string, data []byte) error {
	pool := redisPoolConnect()

	_, err := pool.Publish(context.Background(), channel, data).Result()
	if err != nil {
		return err
	}
	return nil
}

// Subscribe Redis sub 事件，for queue 接收使用
func (r *Redis) Subscribe(channel string) *redis.PubSub {
	pool := redisPoolConnect()
	subscriber := pool.Subscribe(context.Background(), channel)
	return subscriber
}

// LPush Redis Pub 事件，for queue 推送使用(可存留在 queue 中)
func (r *Redis) LPush(channel string, data []byte) error {
	pool := redisPoolConnect()

	_, err := pool.LPush(context.Background(), channel, data).Result()
	if err != nil {
		return err
	}

	return nil
}

// RPush
func (r *Redis) RPush(channel string, data []byte) error {
	pool := redisPoolConnect()

	_, err := pool.RPush(context.Background(), channel, data).Result()
	if err != nil {
		return err
	}

	return nil
}

// LPop
func (r *Redis) LPop(channel string) *redis.StringCmd {
	pool := redisPoolConnect()
	lpop := pool.LPop(context.Background(), channel)
	return lpop
}

// LPos 取值
func (r *Redis) LPos(key, value string) *redis.IntCmd {
	pool := redisPoolConnect()
	args := redis.LPosArgs{}

	lpos := pool.LPos(context.Background(), key, value, args)

	return lpos
}

// LRange 取值
func (r *Redis) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	pool := redisPoolConnect()
	lrange := pool.LRange(context.Background(), key, start, stop)
	return lrange
}

// LRem 刪除指定 value
func (r *Redis) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	pool := redisPoolConnect()
	lrem := pool.LRem(context.Background(), key, count, value)
	return lrem
}

// BRPop Redis sub 事件，for queue 接收使用(監聽有內容就取用，無事件時間到自動回收)
func (r *Redis) BRPop(channel string, exp time.Duration) *redis.StringSliceCmd {
	pool := redisPoolConnect()
	subscriber := pool.BRPop(context.Background(), exp, channel)
	return subscriber
}
