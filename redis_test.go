package main

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestRedis(t *testing.T) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connet redis err:", err)
		return
	}

	defer conn.Close()
}
