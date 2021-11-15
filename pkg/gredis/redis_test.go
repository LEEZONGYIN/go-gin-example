package gredis

import (
	"fmt"
	"testing"
)

func TestSetUp(t *testing.T) {
	err := SetUp()
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func TestSet(t *testing.T) {
	SetUp()
	conn := RedisConn.Get()
	defer conn.Close()
	res, err := conn.Do("SET", "testKey", "testValue")
	if err != nil {
		fmt.Println("set err", err)
	} else {
		fmt.Println("set res ", res)
	}
}
