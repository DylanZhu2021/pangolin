package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func Connect() {
	conn, err := redis.Dial("tcp", "81.69.173.253:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()

	//redis set操作
	_, err = conn.Do("SET", "youmen", "18")
	if err != nil {
		fmt.Println("redis set error:", err)
	}

	//redis get操作
	name, err := redis.String(conn.Do("GET", "youmen"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Get name: %s \n", name)
	}
}
