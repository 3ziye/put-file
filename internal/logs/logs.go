package logs

import (
	"net/http"
	"time"
)

// responseWriter 是 http.ResponseWriter 的包装器，用于捕获状态码
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader 重写 WriteHeader 方法以捕获状态码
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Write 重写 Write 方法以确保状态码被设置
func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}
	return rw.ResponseWriter.Write(b)
}

// RequestLogger 中间件，用于记录HTTP请求
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		ip := r.RemoteAddr
		method := r.Method
		path := r.URL.Path

		// 创建响应包装器以捕获状态码
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// 调用下一个处理器
		next.ServeHTTP(rw, r)

		// 记录请求信息
		duration := time.Since(startTime)
		Info("%s %s %s %d %s", ip, method, path, rw.statusCode, duration)
	})
}



// responseWriter 是 http.ResponseWriter 的包装器，用于捕获状态码
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader 重写 WriteHeader 方法以捕获状态码
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}