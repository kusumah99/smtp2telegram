package SmtpServer

import (
	"errors"
	"flag"
	"io"
	DataHandler "ksa-smtp-telegram/data-handler"
	"net"
	"os"
	"sync"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

// type DataHandler interface {
// 	OnMailCreated(data []byte, from string, to []string)
// }

var dataHandler DataHandler.DataHandlerInterface

func SetDataMailHandler(handler DataHandler.DataHandlerInterface) {
	dataHandler = handler
}

type message struct {
	From     string
	To       []string
	RcptOpts []*smtp.RcptOptions
	Data     []byte
	Opts     *smtp.MailOptions
}

type backend struct {
	authDisabled bool

	// mailmessage *message
	// messages []*message

	implementLMTPData bool
	lmtpStatus        []struct {
		addr string
		err  error
	}
	lmtpStatusSync chan struct{}

	// Errors returned by Data method.
	dataErrors chan error

	// Error that will be returned by Data method.
	dataErr error

	// Read N bytes of message before returning dataErr.
	dataErrOffset int64

	panicOnMail bool
	userErr     error
}

func (be *backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	if be.implementLMTPData {
		return &lmtpSession{&session{backend: be, anonymous: true}}, nil
	}

	return &session{backend: be, anonymous: true}, nil
}

type lmtpSession struct {
	*session
}

type session struct {
	backend   *backend
	anonymous bool

	msg *message
}

var _ smtp.AuthSession = (*session)(nil)

func (s *session) AuthMechanisms() []string {
	if s.backend.authDisabled {
		return nil
	}
	return []string{sasl.Plain, sasl.Login}
}

func (s *session) Auth(mech string) (sasl.Server, error) {
	if s.backend.authDisabled {
		return nil, smtp.ErrAuthUnsupported
	}

	switch mech {
	case "PLAIN":
		return sasl.NewPlainServer(func(identity, username, password string) error {
			// if identity != "" && identity != username {
			// 	return errors.New("invalid identity")
			// }
			// if username != "username" || password != "password" {
			// 	return errors.New("invalid username or password")
			// }
			s.anonymous = false
			return nil
		}), nil
	case "LOGIN":
		return sasl.NewLoginServer(func(username, password string) error {
			// if username != "username" || password != "password" {
			// 	return errors.New("invalid username or password")
			// }
			s.anonymous = false
			return nil
		}), nil
	default:
		s.anonymous = true
		return nil, errors.New("unsupported auth mechanism")
	}

}

func (s *session) Reset() {
	s.msg = &message{}
}

func (s *session) Logout() error {
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	if s.backend.userErr != nil {
		return s.backend.userErr
	}
	if s.backend.panicOnMail {
		panic("Everything is on fire!")
	}
	s.Reset()
	s.msg.From = from
	s.msg.Opts = opts
	return nil
}

func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	s.msg.To = append(s.msg.To, to)
	s.msg.RcptOpts = append(s.msg.RcptOpts, opts)
	return nil
}

func (s *session) Data(r io.Reader) error {
	if s.backend.dataErr != nil {

		if s.backend.dataErrOffset != 0 {
			io.CopyN(io.Discard, r, s.backend.dataErrOffset)
		}

		err := s.backend.dataErr
		if s.backend.dataErrors != nil {
			s.backend.dataErrors <- err
		}
		return err
	}

	if b, err := io.ReadAll(r); err != nil {
		if s.backend.dataErrors != nil {
			s.backend.dataErrors <- err
		}
		return err
	} else {
		s.msg.Data = b
		// s.backend.mailmessage.Data = b
		// s.backend.messages = append(s.backend.messages, s.msg)
		if s.backend.dataErrors != nil {
			s.backend.dataErrors <- nil
		}
	}
	// ========================== Disini kirim ke telegram ====================
	dataHandler.OnMailCreated(s.msg.Data, s.msg.From, s.msg.To)
	return nil
}

func (s *session) LMTPData(r io.Reader, collector smtp.StatusCollector) error {
	if err := s.Data(r); err != nil {
		return err
	}

	for _, val := range s.backend.lmtpStatus {
		collector.SetStatus(val.addr, val.err)

		if s.backend.lmtpStatusSync != nil {
			s.backend.lmtpStatusSync <- struct{}{}
		}
	}

	return nil
}

type failingListener struct {
	c      chan error
	closed bool
	mu     sync.Mutex
}

func newFailingListener() *failingListener {
	return &failingListener{c: make(chan error)}
}

func (l *failingListener) Send(err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.closed {
		l.c <- err
	}
}

func (l *failingListener) Accept() (net.Conn, error) {
	return nil, <-l.c
}

func (l *failingListener) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.closed {
		close(l.c)
		l.closed = true
	}
	return nil
}

func (l *failingListener) Addr() net.Addr {
	return &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 12345,
	}
}

type mockError struct {
	msg       string
	temporary bool
}

func newMockError(msg string, temporary bool) *mockError {
	return &mockError{
		msg:       msg,
		temporary: temporary,
	}
}

func (m *mockError) Error() string   { return m.msg }
func (m *mockError) String() string  { return m.msg }
func (m *mockError) Timeout() bool   { return false }
func (m *mockError) Temporary() bool { return m.temporary }

type serverConfigureFunc func(*smtp.Server)

var (
	authDisabled = func(s *smtp.Server) {
		s.Backend.(*backend).authDisabled = true
	}
)

var smtp_svr *smtp.Server

// func SetDataMailHandler(dataHandlerfunc *onDataMail.OnMailCreated) {
// 	mailHandler = dataHandler
// }

func SetConfig(serverAddress, domain string, allowInsecureAuth bool) {
	flag.Parse()

	smtp_svr = smtp.NewServer(&backend{})

	smtp_svr.Addr = serverAddress
	smtp_svr.Domain = domain
	smtp_svr.AllowInsecureAuth = allowInsecureAuth
	smtp_svr.Debug = os.Stdout

}

func ListenAndServe() error {
	return smtp_svr.ListenAndServe()
}
