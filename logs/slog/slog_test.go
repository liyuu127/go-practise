package slog

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/slog"
	"net"
	"os"
	"testing"
	"time"
)

func TestSlogTextHandler(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}

func TestSlogJsonHandler(t *testing.T) {
	opts := slog.HandlerOptions{
		AddSource: true,
	}

	slog.SetDefault(slog.New(opts.NewJSONHandler(os.Stderr)))
	slog.Info("hello", "name", "Al", "id", "er")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}

func TestSlogDynamicLevel(t *testing.T) {
	var lvl = &slog.LevelVar{}
	lvl.Set(slog.DebugLevel)
	opts := slog.HandlerOptions{
		Level: lvl,
	}
	slog.SetDefault(slog.New(opts.NewJSONHandler(os.Stderr)))

	slog.Info("before resetting log level:")

	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	slog.Info("after resetting log level to error level:")
	lvl.Set(slog.ErrorLevel)
	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}

func TestSlogHandler(t *testing.T) {
	var ch = make(chan []byte, 100)
	attrs := []slog.Attr{
		{Key: "field1", Value: slog.StringValue("value1")},
		{Key: "field2", Value: slog.StringValue("value2")},
	}
	slog.SetDefault(slog.New(NewChanHandler(ch).WithAttrs(attrs)))
	go func() { // 模拟channel的消费者，用来消费日志
		for {
			b := <-ch
			fmt.Println(string(b))
		}
	}()

	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	time.Sleep(3 * time.Second)
}

type ChanHandler struct {
	slog.Handler
	ch  chan []byte
	buf *bytes.Buffer
}

func (h *ChanHandler) Enabled(level slog.Level) bool {
	return h.Handler.Enabled(level)
}

func (h *ChanHandler) Handle(r slog.Record) error {
	err := h.Handler.Handle(r)
	if err != nil {
		return err
	}
	var nb = make([]byte, h.buf.Len())
	copy(nb, h.buf.Bytes())
	h.ch <- nb
	h.buf.Reset()
	return nil
}

func (h *ChanHandler) WithAttrs(as []slog.Attr) slog.Handler {
	return &ChanHandler{
		buf:     h.buf,
		ch:      h.ch,
		Handler: h.Handler.WithAttrs(as),
	}
}

func (h *ChanHandler) WithGroup(name string) slog.Handler {
	return &ChanHandler{
		buf:     h.buf,
		ch:      h.ch,
		Handler: h.Handler.WithGroup(name),
	}
}

func NewChanHandler(ch chan []byte) *ChanHandler {
	var b = make([]byte, 256)
	h := &ChanHandler{
		buf: bytes.NewBuffer(b),
		ch:  ch,
	}

	h.Handler = slog.NewJSONHandler(h.buf)

	return h
}
