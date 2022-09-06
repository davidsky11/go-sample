package redispkg

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	//redis "github.com/go-redis/redis/v8"
)

// 定义一个 RedisSingleConnection 结构体
type RedisSingleConnection struct {
	Redis_host   string
	Redis_port   uint16
	Redis_auth   string
	database     int
	singleClient *redis.Client
}

// 结构体 InitSingleRedis 方法：用于初始化 redis 数据库
func (r *RedisSingleConnection) InitSingleRedis() (err error) {
	// Redis 连接格式拼接
	redisAddr := fmt.Sprintf("%s:%d", r.Redis_host, r.Redis_port)
	// Redis 连接对象：NewClient 将客户端返回到由选项指定的 Redis 服务器
	r.singleClient = redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		Password:    r.Redis_auth,
		DB:          r.database, // 连接的database库
		IdleTimeout: 300,        // 默认Idle超时时间
		PoolSize:    10,         // 连接池
	})

	fmt.Printf("Connecting Redis : %v\n", redisAddr)

	// 验证是否连接到 redis 服务端
	res, err := r.singleClient.Ping().Result()
	if err != nil {
		fmt.Printf("Connect Failed! Err: %v\n", err)
		return err
	} else {
		fmt.Printf("Connect Successful! Ping => %v\n", res)
		return nil
	}

}

// 定义一个RedisClusterObj结构体
type RedisSentinelConnection struct {
	Redis_master   string
	Redis_addr     []string
	Redis_auth     string
	sentinelClient *redis.Client
}

// 结构体方法
func (r *RedisSentinelConnection) initSentinelClient() (err error) {
	r.sentinelClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = r.sentinelClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// 定义一个RedisClusterConnection结构体
type RedisClusterConnection struct {
	Redis_addr    []string
	Redis_auth    string
	clusterClient *redis.ClusterClient
}

// 结构体方法
func (r *RedisClusterConnection) initClusterClient() (err error) {
	r.clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = r.clusterClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
