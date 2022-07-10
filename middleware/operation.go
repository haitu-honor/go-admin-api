package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

// func OperationRecord() gin.HandlerFunc {

// }

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
