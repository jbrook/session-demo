package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"time"
	"fmt"
	"github.com/jbrook/sessiondemo/ttlcache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cookieNameForSessionID = "sess"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
	ttl                    = time.Second * 30
	cache                  *ttlcache.Cache
	metric                 prometheus.Gauge
)

func init() {
	cache = ttlcache.NewCache(ttl)
	metric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_session_count",
			Help: "Count of active sessions",
		},
	)
	prometheus.MustRegister(metric)
	metric.Set(0)
}

func hello(ctx iris.Context) {
	session := sess.Start(ctx)
	activeSessions := 0

	if _, ok := cache.Get(session.ID()); !ok {
		cache.Set(session.ID(), "1")
	}
	activeSessions = cache.Count()
	metric.Set(float64(activeSessions))

	session.Set("foo", "bar")
	ctx.HTML("<h1>Hello, World!</h1>")
	ctx.HTML(fmt.Sprintf("Active sessions: %d ", activeSessions))
}

func main() {
	app := iris.New()

	app.Get("/hello", hello)

	app.Get("/metrics", iris.FromStd(promhttp.Handler()))

	app.Run(iris.Addr(":8080"))
}
