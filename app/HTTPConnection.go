package main

type HTTPConnection struct {
	conn TCPConnection
}

func connectHTTP(protocol string, address string) HTTPConnection {
	return HTTPConnection{connectTCP(protocol, address)}
}

func nextRequest() {
	
}

func (HTTPConn *HTTPConnection) Close() {
	HTTPConn.conn.Close()
}
