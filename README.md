# speedtest-go-prometheus

A simple utility that runs internet test using speedtest.net and reports the results with prometheus endpoint.

I used:

- [showwin/speedtest-go](https://github.com/showwin/speedtest-go)
- [prometheus/client_golang](https://github.com/prometheus/client_golang)

Inspired by 

- [geerlingguy/internet-pi](https://github.com/geerlingguy/internet-pi)
- [MiguelNdeCarvalho/speedtest-exporter](https://github.com/MiguelNdeCarvalho/speedtest-exporter)
- A gist by [tembleking/main.go](https://gist.github.com/tembleking/0b8968dbdf36dfef6227fbfdd9bb1a82)

## Usage

Consider this line:

```bash
sgp -b :8080 -i 60
```

- `sgp` is a compiled binary of this project
- `-b` is for "bind", the address which will be passed to ListenAndServe, default is `:8080`
- `-i` is for "interval", the time between measurement in seconds, default is `60`

The server will start and will report metrics over `/metrics`.

The metrics for speed test have a prefix `speedtest`. Following metrics are available: 

- `speedtest_latency` - a gauge, ms
- `speedtest_download` - a gauge, MB/s
- `speedtest_upload` - a gauge, MB/s
