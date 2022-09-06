package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	redisv8 "storage-operate/redispkg/v8"
	"strconv"
)

var ctx = context.Background()

func hashSet(rdb *redis.Client, ctx context.Context) {
	hashKey := "hset512-hashtable-test"
	for i := 0; i < 512; i++ {
		field := strconv.Itoa(i)
		rdb.HSet(ctx, hashKey, field, field)
	}
}

func main() {
	// 实例化 Redis
	conn := &redisv8.RedisSingleObj{
		Redis_host: "127.0.0.1",
		Redis_port: 6379,
		Redis_auth: "",
	}
	// 初始化连接 Single Redis 服务端
	redisClient, err := conn.InitSingleRedis()
	if err != nil {
		fmt.Printf("[Error] - %v\n", err)
		return
	}

	hashSet(redisClient, ctx)

	// 程序执行完毕释放资源
	defer redisClient.Close()
}
