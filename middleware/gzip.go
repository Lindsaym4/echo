package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
)

type gzipResponseWriter struct {
	Writer io.Writer
	http.ResponseWriter
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *gzipResponseWriter) Flush() {
	if gw, ok := w.Writer.(*gzip.Writer); ok {
		gw.Flush()
	}
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (w *gzipResponseWriter) Close() error {
	if gw, ok := w.Writer.(*gzip.Writer); ok {
		return gw.Close()
	}
	return nil
}