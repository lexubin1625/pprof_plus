##### pprof非web服务采样(用于rpc或其它非http服务等)

###### 代码实例
```go
  
import (
	"github.com/lexubin1625/pprof_plus"
	// 其他库按需载入
)

    // 配置加载
    pprof_plus.InitConfig(pprof_plus.Config{
        TimeSeconds: 10,//采集时间(单位秒)
        FilePath: "./logs/pprof/", // 生成配置路径
    })

    // pprof采集
    pprof_plus.Gather() // 采集单次

    //可以在阻塞进程中用接收信号的方式多次采集
    var OsSignal chan os.Signal
	signal.Notify(OsSignal,syscall.SIGUSR2)

	for {
		s := <-OsSignal
		switch s {
		case syscall.SIGUSR2:
			// pprof采集
			pprof_plus.Gather()
		}
	}

```

###### 查看结果

```shell
 go tool pprof -http=:任意端口 输出路径/cpu.prof
```

###### 生成文件结构
| 文件名 | 功能   | 
| :----- | :----- |
| cpu.prof | CPU Profiling |
| mem.prof | 内存分配情况 |
| goroutine.prof | 所有运行的 goroutines 堆栈跟踪 |
|block.prof|查看导致阻塞同步的堆栈跟踪|
|mutex.prof|查看导致互斥锁的竞争持有者的堆栈跟踪|
|threadcreate.prof|查看创建新OS线程的堆栈跟踪|
  
