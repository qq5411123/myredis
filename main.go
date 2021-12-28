package main

import (
	"github.com/go-cmd/cmd"
	"log"
	"math"
	"strconv"
	"strings"
)

var dataSizeArr []string
var redis_n = 100000

func main()  {
	dataSizeArr = []string{"10", "20", "50", "100", "200", "1000", "5000",}
	var used_memory_dataset int
	used_memory_dataset = infoMemory("0")

	for _, dataSize := range dataSizeArr{
		redisBenchmark(dataSize)
		new_used_memory_dataset := infoMemory(dataSize)
		diff := new_used_memory_dataset - used_memory_dataset

		key_avg_memory := math.Round(float64(diff / redis_n))

		log.Println("new: ", new_used_memory_dataset, "old: ", used_memory_dataset, "diff: ", diff)
		log.Printf("when %s key avg memory is %f", dataSize, key_avg_memory)

		used_memory_dataset = new_used_memory_dataset
	}
}

func redisBenchmark(dataSize string)  {
	c := cmd.NewCmd("/usr/local/redis/bin/redis-benchmark","-n", strconv.Itoa(redis_n), "-t", "get,set", "-r", strconv.Itoa(redis_n), "-d", dataSize)
	log.Println("/usr/local/redis/bin/redis-benchmark","-n", strconv.Itoa(redis_n), "-t", "get,set", "-d", dataSize)

	_ = <-c.Start()//阻塞到命令执行完
	//输出utf8文本; 字符串数组，一行一元素
	//log.Println("get cmd Stdout:", strings.Join(s.Stdout,"\n"))
	//log.Println("get cmd Stderr:", strings.Join(s.Stderr,"\n"))
}

func infoMemory(dataSize string) int {
	c := cmd.NewCmd("/usr/local/redis/bin/redis-cli","info", "memory")
	s := <-c.Start()//阻塞到命令执行完
	if len(s.Stderr) > 0 {
		log.Fatalf("%s info fail", dataSize)
	}

	//log.Println(s.Stdout)

	var used_memory_dataset int
	for _, v := range s.Stdout{
		if strings.Contains(v, "used_memory_dataset")  {
			arr := strings.Split(v, ":")
			if len(arr) > 1 && arr[0] == "used_memory_dataset" {
				used_memory_dataset, _ = strconv.Atoi(arr[1])
			}
		}
	}
	return used_memory_dataset
	
}