package pbclient

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"google.golang.org/grpc"
)

type connCofnig struct {
	addr   string
	option []grpc.DialOption

	connectTimeout  int
	connectAttempts int
	connectBackoff  int

	// logger
	l *slog.Logger

	clinet *grpc.ClientConn
}

type Conn interface {
	Connect()
	Close()
}

// ConnOption は接続のオプションを設定するための関数です。
type ConnOption func(*connCofnig)

// WithOption は接続のオプションを設定するオプションです。
func (c *connCofnig) WithOption(opts ...grpc.DialOption) ConnOption {
	return func(c *connCofnig) {
		c.option = opts
	}
}

// WithConnectTimeout は接続のタイムアウトを設定するオプションです。
// このオプションを設定しない場合、デフォルト値は 10 秒です。
// タイムアウトは秒で指定します。
func WithConnectTimeout(timeout int) ConnOption {
	return func(c *connCofnig) {
		c.connectTimeout = timeout
	}
}

// WithConnectAttempts は接続の試行回数を設定するオプションです。
// このオプションを設定しない場合、デフォルト値は 3 回です。
func WithConnectAttempts(attempts int) ConnOption {
	return func(c *connCofnig) {
		c.connectAttempts = attempts
	}
}

// WithConnectBackoff は接続のバックオフを設定するオプションです。
// このオプションを設定しない場合、デフォルト値は 3 秒です。
// バックオフは秒で指定します。
func WithConnectBackoff(backoff int) ConnOption {
	return func(c *connCofnig) {
		c.connectBackoff = backoff
	}
}

// WithLogger はロガーを設定するオプションです。
func WithLogger(l *slog.Logger) ConnOption {
	return func(c *connCofnig) {
		c.l = l
	}
}

// NewConn は接続のタイムアウトを設定するオプションです。
func NewConn(addr string, opts ...ConnOption) *connCofnig {
	c := &connCofnig{
		addr: addr,
		l:    slog.New(slog.NewTextHandler(os.Stderr, nil)).WithGroup("pb-client"),
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.connectTimeout == 0 {
		c.connectTimeout = 10
	}

	if c.connectAttempts == 0 {
		c.connectAttempts = 3
	}

	if c.connectBackoff == 0 {
		c.connectBackoff = 3
	}

	return c
}

// Connect は接続を行います。
func (c *connCofnig) Connect() {

	sleep := func() {
		time.Sleep(time.Duration(c.connectBackoff) * time.Second)
	}

	for i := 0; i < c.connectAttempts; i++ {
		c.l.Info(fmt.Sprintf("connect to %s", c.addr))

		// connect to addr
		conn, err := grpc.Dial(
			c.addr,
			c.option...,
		)

		// 成功したら終了
		if err == nil {
			c.l.Info(fmt.Sprintf("connected to %s", c.addr))
			c.clinet = conn
			return
		}

		c.l.Error(fmt.Sprintf("failed to connect to %s: %s", c.addr, err))
		sleep()
	}
}

func (c *connCofnig) Client() *grpc.ClientConn {
	return c.clinet
}

// Close は接続を閉じます。
func (c *connCofnig) Close() {
	if c.clinet == nil {
		c.l.Error(fmt.Sprintf("connection to %s is not established", c.addr))
		return
	}

	c.l.Info(fmt.Sprintf("close connection to %s", c.addr))
	if err := c.clinet.Close(); err != nil {
		c.l.Error(fmt.Sprintf("failed to close connection to %s: %s", c.addr, err))
		return
	}
	c.clinet = nil
}