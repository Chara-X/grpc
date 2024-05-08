package grpc

import "net/http"

func WriteHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Add("Content-Type", "application/grpc")
	w.Header().Add("Trailer", GrpcMessage)
	w.Header().Add("Trailer", GrpcStatus)
	w.WriteHeader(statusCode)
	w.(http.Flusher).Flush()
}
