package main

import (
	"github.com/go-cmd/cmd"
	"log"
	"strings"
)

func main()  {
	dataSizeArr := map[string]string{"10", /*"20", "50", "100", "200", "1k", "5k",*/}
	for , v := range dataSizeArr{
		redisCmd(v)
	}


}

func redisCmd(dataSize string)  {
	c := cmd.NewCmd("/usr/local/redis/bin/redis-benchmark","-n", "100000", "-t", "get,set", "-d", dataSize)

	s := <-c.Start()//阻塞到命令执行完

	//输出utf8文本; 字符串数组，一行一元素
	log.Println("get cmd Stdout:", strings.Join(s.Stdout,"\n"))
	log.Println("get cmd Stderr:", strings.Join(s.Stderr,"\n"))
}
