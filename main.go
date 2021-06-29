package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/showwin/speedtest-go/speedtest"
)

var (
	addr = flag.String("b", ":8080", "b is for bind. The address to listen on for HTTP requests.")
	testInterval = flag.Int("i", 60, "i is for interval. The time in seconds between speedtest measurements.")
	latencyGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "speedtest",
			Name:      "latency",
			Help:      "Latency gauge, measured in ms.",
		})
	downloadGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "speedtest",
			Name:      "download",
			Help:      "Download speed gauge, measured in MB/s.",
		})
	uploadGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "speedtest",
			Name:      "upload",
			Help:      "Upload speed gauge, measured in MB/s.",
		})
)

func main() {

	prometheus.MustRegister(latencyGauge)
	prometheus.MustRegister(downloadGauge)
	prometheus.MustRegister(uploadGauge)

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {
			log.Print("Starting speed test...")
			user, _ := speedtest.FetchUserInfo()
			serverList, _ := speedtest.FetchServerList(user)
			targets, _ := serverList.FindServer([]int{})
			for _, s := range targets {
				s.PingTest()
				s.DownloadTest(false)
				s.UploadTest(false)

				log.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
			
				latencyGauge.Set(float64(s.Latency))
				downloadGauge.Set(s.DLSpeed)
				uploadGauge.Set(s.ULSpeed)

			}
			log.Print("Speed test completed, waiting...")
			time.Sleep(time.Duration(*testInterval) * time.Second)
		}
	}()
	log.Printf("Server started at %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
