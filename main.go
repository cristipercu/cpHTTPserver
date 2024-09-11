package main

import (
	"bytes"
	"fmt"
	"log"
	"net"

	httprequest "github.com/cristipercu/cpHTTPserver/httpRequest"
	httpresponse "github.com/cristipercu/cpHTTPserver/httpResponse"
)

func main() {
  listener, err := net.Listen("tcp", ":1620")
  if err != nil {
    log.Fatal(err)
  }
  defer listener.Close()

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatal(err)
    }
    
    go handleConnection(conn)
  }

}

func handleConnection(conn net.Conn) {
  defer conn.Close()

  requestBuf := make([]byte, 1024)
  n, err := conn.Read(requestBuf)
  if err != nil {
    fmt.Println("Error reading the request: ", err)
  }
  
  data := requestBuf[:n]
  request := httprequest.NewHttpRequest(data)

  response := httpresponse.NewHttpResponse(*request)
  responseStatusLine := response.HandleRequest()

  serverName := []byte("Server: CP\r\n")
  contentType := []byte("Content-Type: text/html\r\n")

  blankLine := []byte("\r\n")

  html_response := []byte(`<html>
            <body>
            <h1>Request received!</h1>
            <body>
            </html>`)

  var buffer bytes.Buffer
  buffer.Write(responseStatusLine)
  buffer.Write(serverName)
  buffer.Write(contentType)
  buffer.Write(blankLine)
  buffer.Write(html_response)
  buffer.Write([]byte("\r\n\r\n"))

  _, err = conn.Write(buffer.Bytes())

  if err != nil {
    fmt.Println(err)
  }
}
