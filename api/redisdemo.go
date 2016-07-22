package api

import (
	"github.com/labstack/echo"
	"../db"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func SetR() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		rc := db.RedisClient.Get()
		key := "zhangdong"
		value := "cool"
		// 操作redis时调用Do方法，第一个参数传入操作名称（字符串），然后根据不同操作传入key、value、数字等
		// 返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息
		n, err := rc.Do("SETNX", key, value)
		// 若操作失败则返回
		if err != nil {
			fmt.Println(err)
			return
		}
		// 返回的n的类型是int64的，所以得将1或0转换成为int64类型的再比较
		if n == int64(1) {
			// 设置过期时间为24小时
			n, _ := rc.Do("EXPIRE", key, 24*3600)
			if n == int64(1) {
				fmt.Println("success")
			}
		} else if n == int64(0) {
			fmt.Println("the key has already existed")
		}
		defer rc.Close()
		return c.String(1,"va")
	}
}

func GetR() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		rc := db.RedisClient.Get()

		key := "zhangdong"
		// 调用Do后，还得调用相应的方法才能取得数据
		// 由于之前存的value是string类型，所以用redis.String将数据转换成string类型
		value, err := redis.String(rc.Do("GET", key))
		if err != nil {
			fmt.Println("fail")
		}
		defer rc.Close()
		return c.String(1,value)
	}
}