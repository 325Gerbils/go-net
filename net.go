package gonet

import (
	"bufio"
	"net"
)

// Server ...
type Server struct {
	ln   net.Listener
	conn net.Conn
}

// Start ...
func (s *Server) Start(port string) (err error) {
	s.ln, err = net.Listen("tcp", port)
	if err != nil {
		return err
	}
	s.conn, err = s.ln.Accept()
	if err != nil {
		return err
	}
	return nil
}

// OnMessage ...
func (s Server) OnMessage(on func(msg string)) {
	go func() {
		for {
			message, _ := bufio.NewReader(s.conn).ReadString('\n')
			go on(message)
		}
	}()
}

// Send ...
func (s Server) Send(msg string) {
	s.conn.Write([]byte(msg + "\n"))
}

// Client ...
type Client struct {
	ip   string
	conn net.Conn
}

// Connect ...
func (c *Client) Connect(ip string) (err error) {
	c.ip = ip
	c.conn, err = net.Dial("tcp", ip)
	if err != nil {
		return err
	}
	return nil
}

// OnMessage ...
func (c Client) OnMessage(on func(msg string)) {
	go func() {
		for {
			message, err := bufio.NewReader(c.conn).ReadString('\n')
			if err != nil {
				break
			}
			go on(message)
		}
	}()
}

// Send ...
func (c Client) Send(msg string) {
	c.conn.Write([]byte(msg + "\n"))
}

// Disconnect ...
func (c Client) Disconnect() {
	c.conn.Close()
}
