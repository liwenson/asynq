package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var (
	gLoadavg = "/proc/loadavg"
)

// 系统负载情况
type Loadavg struct {
	La1, La5, La15 float64 // 负载
	Processes      string  //
	NumCPU         float64 // CPU 核数
	IdealLoad      float64 // 理想负载值
	FullLoad       float64 //  完全负载值
	MaxLoad        float64 // 最大可接受负载值
}

func (l *Loadavg) Load2String() string {
	return fmt.Sprintf("%.2f %.2f %.2f\t%s", l.La1, l.La5, l.La15, l.Processes)
}

func (l *Loadavg) toFloat(num string) float64 {
	floatVal, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0
	}

	return floatVal
}

// Loadavg 读取/proc/loadavg
func (l *Loadavg) Loadavg() error {
	b, err := os.ReadFile(gLoadavg)
	if err != nil {
		log.Println("reading /proc/loadavg:", err.Error())
		return err
	}

	s := strings.Fields(string(b))
	l.La1 = l.toFloat(s[0])
	l.La5 = l.toFloat(s[1])
	l.La15 = l.toFloat(s[2])
	l.Processes = s[3]
	l.NumCPU = float64(runtime.NumCPU())
	l.IdealLoad = l.NumCPU * 0.7 // 理想负载值
	l.FullLoad = l.NumCPU        // 完全负载值
	l.MaxLoad = l.FullLoad * 1.5 // 最大可接受负载值

	return nil
}
