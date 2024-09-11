package httpresponse

import (
	"fmt"

	httprequest "github.com/cristipercu/cpHTTPserver/httpRequest"
)


const (
  OK = 200
  NotFound = 404
  NotImplemented = 501
)



type HttpResponse struct {
  Request httprequest.HttpRequest
  StatusCode int
  StatusCodeText string
}


func NewHttpResponse(request httprequest.HttpRequest) *HttpResponse{
  return &HttpResponse{
    Request: request,
  }
}

func (hr *HttpResponse) HandleRequest() []byte {
  hr.handleMethod()
  return []byte(fmt.Sprintf("%s %d %s\n", hr.Request.HttpVersion, hr.StatusCode, hr.StatusCodeText))
} 

func (hr *HttpResponse) statusToString() string{
  switch hr.StatusCode {
    case 200:
      return "OK"
    case 404:
      return "Not Found"
    default:
      return "Not Implemented"

  }
}

func (hr *HttpResponse) handleMethod() {
  switch hr.Request.Method {
    case "GET":
      hr.handleGet()
    default:
      hr.handleNotImplemented()
  }
}

func (hr *HttpResponse) handleGet() {
  hr.StatusCode = 200
  hr.StatusCodeText = hr.statusToString()
  
}

func (hr *HttpResponse) handleNotImplemented() {
  hr.StatusCode = 501
  hr.StatusCodeText = hr.statusToString()
}
