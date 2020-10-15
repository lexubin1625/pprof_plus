package pprof_plus

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

type Config struct {
	TimeSeconds time.Duration
	FilePath    string
}

var (
	TimeSecond time.Duration
	FilePath   string
)

func InitConfig(c Config) {
	TimeSecond = c.TimeSeconds
	if c.FilePath == "" {
		FilePath = "./"
	} else {
		FilePath = c.FilePath
	}
	FilePath = strings.TrimRight(FilePath,"/") + "/"
}

// cpu采集开始
func StartCpuProf() error {
	f, err := os.Create(FilePath + "cpu.prof")
	if err != nil {
		return err
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		f.Close()
		return err
	}

	return nil
}

func StopCpuProf() {
	pprof.StopCPUProfile()
}

func SaveHeapProf() error {
	f, err := os.Create(FilePath + "mem.prof")
	if err != nil {
		return err
	}

	if err := pprof.WriteHeapProfile(f); err != nil {
		return err
	}
	f.Close()

	return nil
}

// goroutine block
func SaveBlockProf() error {
	_, err := captureProfile("block")
	return err
}

func SaveMutexProf() error {
	_, err := captureProfile("mutex")
	return err
}

func SaveGoroutineProf() error {
	_, err := captureProfile("goroutine")
	return err
}

func SaveThreadcreateProf() error {
	_, err := captureProfile("threadcreate")
	return err
}

func captureProfile(name string) (string, error) {
	f, err := os.Create(FilePath + name + ".prof")
	if err != nil {
		return "", err
	}
	if err := pprof.Lookup(name).WriteTo(f, 0); err != nil {
		return "", nil
	}
	if err := f.Close(); err != nil {
		return "", nil
	}
	return f.Name(), nil
}

func Gather() {
	go func(t time.Duration) {
		fmt.Println("pprof gather info start")
		err := StartCpuProf()
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second * t)
		StopCpuProf()
		SaveHeapProf()
		SaveBlockProf()
		SaveMutexProf()
		SaveGoroutineProf()
		SaveThreadcreateProf()
		fmt.Println("pprof gather info end")
	}(TimeSecond)

}
