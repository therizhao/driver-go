package conn

import (
	"bytes"
	"errors"
	"io"
	"net"
	"time"

	"github.com/bytehouse-cloud/driver-go/driver/lib/ch_encoding"
)

type fakeConn struct{}

func (f *fakeConn) Read(b []byte) (n int, err error) {
	return 0, io.EOF
}

func (f *fakeConn) Write(b []byte) (n int, err error) {
	if string(b) == "dinosaur" {
		return len(b), nil
	}

	return 0, errors.New("fakeConn: can't write anything other than dinosaur")
}

func (f *fakeConn) Close() error {
	return nil
}

func (f *fakeConn) LocalAddr() net.Addr {
	return &net.IPAddr{}
}

func (f *fakeConn) RemoteAddr() net.Addr {
	return nil
}

func (f *fakeConn) SetDeadline(t time.Time) error {
	return nil
}

func (f *fakeConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (f *fakeConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func MockConn() *GatewayConn {
	var buffer bytes.Buffer
	encoder := ch_encoding.NewEncoder(&buffer)
	decoder := ch_encoding.NewDecoder(&buffer)
	return &GatewayConn{
		encoder:     encoder,
		decoder:     decoder,
		connOptions: &ConnConfig{},
		conn: &connect{
			Conn: &fakeConn{},
		},
		authentication: NewAuthentication("123", "123", "123"),
		//	impersonation:  NewImpersonation(1, 1, "hello", false),
		logf:     func(s string, i ...interface{}) {},
		userInfo: NewUserInfo(),
	}
}
