package http

import (
	"bufio"
	"fmt"
	"github.com/botflux/go-http-server/routing"
	"io"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

type Response struct {
	StatusCode int
	Body       string
	Headers    map[string]string
}

type Server struct {
	ListenAddr string
	Router     *routing.Router
}

func (s Server) Listen() error {
	listener, err := net.Listen("tcp", s.ListenAddr)

	if err != nil {
		return fmt.Errorf("could not listen on %s: %w", s.ListenAddr, err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Printf("could not close listener on %s: %v\n", s.ListenAddr, err)
		}
	}(listener)

	fmt.Printf("listening on port %s\n", s.ListenAddr)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("failed to accept connection, err: ", err)
			continue
		}

		go func(c net.Conn) {
			defer func(conn net.Conn) {
				err := conn.Close()
				if err != nil {
					fmt.Println("failed to close connection, err: ", err)
				}
			}(conn)
			s.handleConnection(conn)
		}(conn)
	}
}

func (s Server) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)

	request, err := s.readHTTPRequest(reader)

	fmt.Printf("received a new request %+v\n", request)

	if err != nil {
		if err != io.EOF {
			fmt.Println("failed to read bytes from connection, err: ", err)
			return
		}

		fmt.Println("connection closed")
		return
	}
	//
	//if request.Method != "GET" {
	//	err := s.respond(conn, Response{
	//		StatusCode: 503,
	//		Headers:    map[string]string{"Content-Type": "text/html; charset=utf-8"},
	//		Body:       "<html><body><h1>Not implemented</h1></body></html>",
	//	})
	//
	//	if err != nil {
	//		fmt.Println("failed to write response, err: ", err)
	//		return
	//	}
	//
	//	return
	//}
	//
	//err = s.respond(conn, Response{
	//	StatusCode: 200,
	//	Headers:    map[string]string{"Content-Type": "text/html; charset=utf-8"},
	//	Body:       "<html><body><h1>Successful response!</h1></body></html>",
	//})

	response := s.Router.Handle(request)
	err = s.respond(conn, response)

	if err != nil {
		fmt.Println("failed to write response, err: ", err)
	}
}

func (s Server) respond(conn net.Conn, response Response) error {
	contentLength := len(response.Body)
	textResponse := fmt.Sprintf("HTTP/1.1 %d\r\n", response.StatusCode)

	response.Headers["Content-Length"] = strconv.Itoa(contentLength)

	for key, value := range response.Headers {
		textResponse += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	textResponse += fmt.Sprintf("\r\n%s\r\n\r\n", response.Body)

	_, err := conn.Write([]byte(textResponse))

	return err
}

func (s Server) readHTTPRequest(r *bufio.Reader) (Request, error) {
	//request := ""
	//contentLength := 0

	b, err := r.ReadBytes('\n')

	if err != nil {
		return Request{Method: "", URL: ""}, err
	}

	firstLine := string(b[:])
	parts := strings.Split(firstLine, " ")

	headers := make(map[string]string)

	for {
		b, err := r.ReadBytes('\n')

		if err != nil {
			return Request{Method: "", URL: ""}, err
		}

		line := string(b[:])

		if line == "\r\n" {
			break
		}

		parts := strings.Split(line, ":")

		headers[parts[0]] = strings.Trim(parts[1], "\r")
	}

	contentLength, hasContentLength := headers["Content-Length"]
	parsedContentLength, err := strconv.Atoi(strings.Trim(contentLength, "\r\n "))

	if hasContentLength && err != nil {
		return Request{Method: "", URL: ""}, err
	}

	if parts[0] == "GET" || !hasContentLength {
		return Request{Method: parts[0], URL: parts[1], Headers: headers}, nil
	}

	body := make([]byte, parsedContentLength)
	_, err = r.Read(body)

	if err != nil {
		return Request{Method: "", URL: ""}, err
	}

	return Request{Method: parts[0], URL: parts[1], Headers: headers, Body: string(body)}, nil
}
