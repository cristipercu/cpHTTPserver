package main

import (
	"bytes"
	"fmt"
	"log"
	"net"

	httprequest "github.com/cristipercu/cpHTTPserver/httpRequest"
)

const (
  OK = 200
  NotFound = 404
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

func statusToString(code int) string{
  switch code {
    case 200:
      return "OK"
    case 404:
      return "Not Found"
    default:
      return ""
  }
}

func handleConnection(conn net.Conn) {
  defer conn.Close()

  responseLine := createResponseLine(200)
  serverName := []byte("Server: CP\r\n")
  contentType := []byte("Content-Type: text/html\r\n")

  blankLine := []byte("\r\n")

  response := []byte(`<html>
            <body>
            <h1>Request received!</h1>
            <body>
            </html>`)

  requestBuf := make([]byte, 1024)
  n, err := conn.Read(requestBuf)
  if err != nil {
    fmt.Println("Error reading the request: ", err)
  }
  
  data := requestBuf[:n]

  hr := httprequest.NewHttpRequest(data)

  var buffer bytes.Buffer
  buffer.Write(responseLine)
  buffer.Write(serverName)
  buffer.Write(contentType)
  buffer.Write(blankLine)
  buffer.Write(response)
  buffer.Write([]byte("\r\n\r\n"))

  fmt.Println("New req", hr)


  _, err = conn.Write(buffer.Bytes())

  if err != nil {
    fmt.Println(err)
  }
}

func createResponseLine(statusCode int) []byte{
  statusCodeMessage := statusToString(statusCode)
  responseLine := fmt.Sprintf("HTTP/1.1 %v %v", statusCode, statusCodeMessage)
  return []byte(responseLine)
}

