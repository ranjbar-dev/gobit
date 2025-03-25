package apicontroller

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary       Ping (GET)
// @Description   ping
// @Tags          Server
// @Accept        json
// @Produce       json
// @Router        /server/ping [get]
// @Success       200     {object}   nil   "Success"
func (controller *Controller) ServerPingGet(c *gin.Context) {

	controller.ok(c, nil)
}

// @Summary       Ping (POST)
// @Description   ping
// @Tags          Server
// @Accept        json
// @Produce       json
// @Router        /server/ping [post]
// @Success       200     {object}   nil   "Success"
func (controller *Controller) ServerPingPost(c *gin.Context) {

	controller.ok(c, nil)
}

// @Summary       Server time
// @Description   Get server timestamp
// @Tags          Server
// @Accept        json
// @Produce       json
// @Router        /server/timestamp [get]
// @Success       200     {int64}      int64   "Success"
func (controller *Controller) ServerTimestamp(c *gin.Context) {

	controller.ok(c, time.Now().Unix())
}

// @Summary       Server memory usage
// @Description   Get server memory usage
// @Tags          Server
// @Accept        json
// @Produce       json
// @Router        /server/memory-usage [get]
// @Success       200     {string}   string   "Success"
func (controller *Controller) ServerMemoryUsage(c *gin.Context) {

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	controller.ok(c, fmt.Sprintf("Memory Usage: %v MiB, Alloc: %v MiB, Sys: %v MiB, NumGC: %v MiB", uint64(m.Alloc/1024/1024), uint64(m.TotalAlloc/1024/1024), uint64(m.Sys/1024/1024), m.NumGC))
}
