package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)



func main() {
	//使用redis封装的Dial进行tcp连接
	c,err := redis.Dial("tcp","localhost:6379")
	errCheck(err)

	defer c.Close()

	//对本次连接进行set操作
	//_,setErr := c.Do("set","url","xxbandy.github.io")
	//errCheck(setErr)

	for true  {
		r,getErr := redis.String(c.Do("get","I0T"))
		errCheck(getErr)
		fmt.Println(r)
		time.Sleep(1*1.0E9)
	}
	//使用redis的string类型获取set的k/v信息


}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:",err)
		os.Exit(-1)
	}
}
