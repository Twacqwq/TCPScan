package main

import (
	"TCPScan/cmd"
	"log"
	_ "net/http/pprof"
	"time"
)

var TIMEOUT = 1 * time.Second

func main() {
	// go func() {
	// 	http.ListenAndServe(":1234", nil)
	// }()
	// defer profile.Start().Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// 加载环境变量
	// ...
	// host := flag.String("host", "10.1.1.1", "目标主机")
	// segA := flag.String("A", "10", "A网段")
	// segB := flag.String("B", "1", "B")
	// segC := flag.String("C", "1", "C网段")
	// workers := flag.Int("worker", 100, "goroutine number.")
	// flag.Parse()
	// segments := []string{*segA, *segB, *segC}
	// fmt.Println(segments, *workers)
	// // args := flag.Args()

	// // 得到A-D网段
	// // segmentList := []string{"10", "9", "10"}
	// // 开始创建扫描任务
	// // ...
	// // address := fmt.Sprintf("%s:%%d", *host)
	// // pkg.StartPort(address)
	// pkg.Start(segments, 70, 90, *workers)

	// test cli
	// cmd.Demo()
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
