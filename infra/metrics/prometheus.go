package metrics

import (
	"bytes"
	"fmt"
	"gim/handler"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"strconv"
)

const (
	contentTypeHeader   = "Content-Type"
	contentLengthHeader = "Content-Length"
)

var (
	counter = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "nim_request_data", Help: "请求类别"},
		[]string{"cmdId", "authed"})
)

func newGauge(name string, desc string, num float64) prometheus.Gauge {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{Name: name, Help: desc})
	gauge.Set(num)
	return gauge
}

func Req(cmdId uint8, authed bool) {
	counter.WithLabelValues(strconv.Itoa(int(cmdId)), strconv.FormatBool(authed)).Inc()
}

func RunMetrics(handler handler.IHandler) {
	defer func() {
		if e := recover(); e != nil {
			logrus.Errorln("metrics panic ", e)
		}
	}()
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		reg := prometheus.NewRegistry()
		reg.Register(counter)
		authLen, waitLen := handler.Metrics()
		reg.Register(newGauge("nim_auth", "已认证的连接数", float64(authLen)))
		reg.Register(newGauge("nim_waitAuth", "未认证的连接数", float64(waitLen)))
		//virtualMemory, _ := mem.VirtualMemory()
		//reg.Register(newGauge("nim_memory_total", "总内存(M)", float64(virtualMemory.Total/1024/1024)))
		//reg.Register(newGauge("nim_memory_used", "已使用内存(M)", float64(virtualMemory.Used/1024/1024)))
		//reg.Register(newGauge("nim_memory_used_percent", "占比(%)", virtualMemory.UsedPercent))
		reg.Register(newGauge("nim_goroutine", "协程数", float64(runtime.NumGoroutine())))
		//cpuInfo, _ := cpu.Percent(time.Duration(time.Second), false)
		//reg.Register(newGauge("nim_cpu", "cpu占用百分比(%)", float64(cpuInfo[0])))
		entry, err := reg.Gather()
		if err != nil {
			http.Error(w, "An error has occurred during metrics collection:\n\n"+err.Error(), http.StatusInternalServerError)
			return
		}

		buf := bytes.NewBuffer(nil)
		contentType := expfmt.Negotiate(r.Header)
		enc := expfmt.NewEncoder(buf, contentType)

		for _, met := range entry {
			if err := enc.Encode(met); err != nil {
				http.Error(w, "An error has occurred during metrics encoding:\n\n"+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if buf.Len() == 0 {
			http.Error(w, "No metrics encoded, last error:\n\n"+err.Error(), http.StatusInternalServerError)
			return
		}
		header := w.Header()
		header.Set(contentTypeHeader, string(contentType))
		header.Set(contentLengthHeader, fmt.Sprint(buf.Len()))
		w.Write(buf.Bytes())
	})

	http.ListenAndServe(":3000", nil)
}
