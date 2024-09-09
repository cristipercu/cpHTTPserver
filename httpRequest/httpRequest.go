package httprequest

import "bytes"

type HttpRequest struct {
  Method string
  Uri string
  HttpVersion string
}


func NewHttpRequest(data []byte) *HttpRequest {
  req := &HttpRequest{
    HttpVersion: "1.1",
  }

  req.parse(data)
  return req
} 

func (r *HttpRequest) parse(data []byte) {
  lines := bytes.Split(data, []byte("\r\n"))
  requestLine := lines[0]
  words := bytes.Split(requestLine, []byte(" "))
  
  r.Method = string(words[0])

  if len(words) > 1 {
    r.Uri = string(words[1])
  }

  if len(words) > 2 {
    r.HttpVersion = string(words[2])
  }
}
