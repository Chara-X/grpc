# grpc

## Client

```go
func main() {
	var caCert, _ = os.ReadFile("resources/cert.pem")
	var caCertPool = x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	var cli = &http.Client{Transport: &http2.Transport{TLSClientConfig: &tls.Config{RootCAs: caCertPool}}}
	Unary(cli)
	// ClientStreaming(cli)
	// ServerStreaming(cli)
	// DuplexStreaming(cli)
}
func Unary(cli *http.Client) {
	var _, _, _, resBody = grpc.StreamingCall(cli, "https://127.0.0.1:2333/User/GetOne", http.Header{}, user.GetOneRequest{Id: 2})
	defer resBody.Close()
	var u *user.User
	resBody.Decode(&u)
	fmt.Println(u)
}
func ClientStreaming(cli *http.Client) {
	var _, reqBody, res, resBody = grpc.StreamingCall(cli, "https://127.0.0.1:2333/User/Count", http.Header{}, nil)
	go func() {
		for i := 1; i < 5; i++ {
			reqBody.Encode(&user.User{Id: int32(i), Name: "Chara-X"})
			time.Sleep(time.Second)
		}
		reqBody.Close()
	}()
	defer resBody.Close()
	var count *user.CountResponse
	if resBody.Decode(&count) != nil {
		fmt.Println(res.Trailer[grpc.GrpcStatus])
	}
	fmt.Println(count)
}
func ServerStreaming(cli *http.Client) {
	var _, _, res, resBody = grpc.StreamingCall(cli, "https://127.0.0.1:2333/User/GetAll", map[string][]string{}, basic.Nil{})
	defer resBody.Close()
	for {
		var u *user.User
		if resBody.Decode(&u) != nil {
			fmt.Println(res.Trailer[grpc.GrpcStatus])
			break
		}
		fmt.Println(u)
	}
}
func DuplexStreaming(cli *http.Client) {
	var _, reqBody, res, resBody = grpc.StreamingCall(cli, "https://127.0.0.1:2333/User/Filter", map[string][]string{}, nil)
	go func() {
		for i := 1; i < 5; i++ {
			reqBody.Encode(&user.User{Id: int32(i), Name: "Chara-X"})
			time.Sleep(time.Second)
		}
		reqBody.Close()
	}()
	defer resBody.Close()
	for {
		var u *user.User
		if err := resBody.Decode(&u); err != nil {
			fmt.Println(res.Trailer[grpc.GrpcStatus])
			break
		}
		fmt.Println(u)
	}
}
```

## Server

```go
func Streaming(res http.ResponseWriter, req *http.Request) {
	var resBody, reqBody = grpc.NewEncoder(res), grpc.NewDecoder(req.Body)
	grpc.WriteHeader(res, http.StatusOK)
	for {
		var req *user.User
		if err := reqBody.Decode(&req); err != nil {
			break
		}
		if req.Id == 2 {
			res.Header().Set(grpc.GrpcMessage, "Bad Request")
			res.Header().Set(grpc.GrpcStatus, "400")
			break
		}
		resBody.Encode(req)
	}
}
```

## References

[HTTP2 client-server full-duplex connection](https://github.com/posener/h2conn)

[HTTP/2 Adventure in the Go World](https://posener.github.io/http2)

[Using gRPC with JSON](https://jbrandhorst.com/post/grpc-json)
