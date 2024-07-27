package grpc

import (
	"net/http"
	"os"
)

func StreamingCall(cli *http.Client, url string, header http.Header) (req *http.Request, reqBody *Encoder, res *http.Response, resBody *Decoder) {
	var r, w, _ = os.Pipe()
	req, _ = http.NewRequest(http.MethodPost, url, r)
	req.Header = header
	req.Header.Add("Content-Type", "application/grpc+json")
	req.Header.Add("Te", "trailers")
	res, _ = cli.Do(req)
	return req, NewEncoder(w), res, NewDecoder(res.Body)
}
