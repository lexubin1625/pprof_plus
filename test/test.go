package main

import (
	"pprof_plus"
	"time"
)

func main(){
	pprof_plus.InitConfig(pprof_plus.Config{
		TimeSeconds: 10,
		FilePath: "./pprof",
	})

	pprof_plus.Gather()
	time.Sleep(time.Second * 20)
}
