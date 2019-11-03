package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"os"
	"runtime"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}


func LoggerHandler(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	// Start timer
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	fmt.Println(path, raw)

	//todo 记录前置日志

	// Process request
	c.Next()

	defer func() {
		if err := recover(); err != nil {
			//日志记录异常不处理
		}
	}()
	//todo 记录后置日志
}

func RecoverHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			pnc := fmt.Sprintf("[Recovery] %s panic recovered:\n%s\n%s\n%s", time.Now().Format("2006-01-02 15:04:05"), string(httprequest), err, buf)
			fmt.Fprintf(os.Stderr, pnc)
			c.AbortWithStatus(500)
		}
	}()
	c.Next()
}
