package grpc

import (
	"net/http"
	"os"
)

func StreamingCall(cli *http.Client, url string, header http.Header, body any) (req *http.Request, reqBody *Encoder, res *http.Response, resBody *Decoder) {
	var r, w, _ = os.Pipe()
	if reqBody = NewEncoder(w); body != nil {
		reqBody.Encode(body)
	}
	req, _ = http.NewRequest(http.MethodPost, url, r)
	req.Header = header
	req.Header.Add("Content-Type", "application/grpc+json")
	req.Header.Add("Te", "trailers")
	res, _ = cli.Do(req)
	resBody = NewDecoder(res.Body)
	return
}
