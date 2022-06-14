package logger

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

type log struct {
	*logrus.Logger
}

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	CheckError(err error, i interface{}) error
	SendErrorWithFunc(err error, i interface{})
	LogError(err error, i interface{})
	Hook() echo.MiddlewareFunc
	Print(v ...interface{})
	BeforeQuery(context.Context, *pg.QueryEvent) (context.Context, error)
	AfterQuery(context.Context, *pg.QueryEvent) error
}

func New() Logger {
	logger := logrus.New()

	return log{logger}
}

func (l log) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	strBytes, _ := q.FormattedQuery()
	if q.Err != nil {
		l.Warn(string(strBytes))
	} else {
		l.Info(string(strBytes))
	}

	return nil
}

func (l log) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l log) Print(v ...interface{}) {
	if v[0] == "sql" {
		l.WithFields(logrus.Fields{"module": "gorm", "type": "sql"}).Print(v[3])
	}
	if v[0] == "log" {
		l.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

func (l log) logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path
	if p == "" {
		p = "/"
	}

	bytesIn := req.Header.Get(echo.HeaderContentLength)
	if bytesIn == "" {
		bytesIn = "0"
	}

	if res.Status == http.StatusOK {
		l.WithFields(map[string]interface{}{
			"remote_ip":  c.RealIP(),
			"method":     req.Method,
			"path":       p,
			"referer":    req.Referer(),
			"user_agent": req.UserAgent(),
			"status":     res.Status,
			"latency":    strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		}).Info("Handled request")
	}

	if res.Status != http.StatusOK {
		l.WithFields(map[string]interface{}{
			"remote_ip":     c.RealIP(),
			"host":          req.Host,
			"uri":           req.RequestURI,
			"method":        req.Method,
			"path":          p,
			"referer":       req.Referer(),
			"user_agent":    req.UserAgent(),
			"status":        res.Status,
			"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
			"latency_human": stop.Sub(start).String(),
			"bytes_in":      bytesIn,
			"bytes_out":     strconv.FormatInt(res.Size, 10),
		}).Warn("Error request")
	}

	return nil
}

func (l log) logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return l.logrusMiddlewareHandler(c, next)
	}
}

func (l log) Hook() echo.MiddlewareFunc {
	return l.logger
}

func (l log) CheckError(err error, i interface{}) error {
	if err != nil {
		l.Warn("Function name: "+getFunctionName(i)+" | Error: ", err)
		// if err != nil publish error to pubsub to make this error posted to slack channel
	}

	return err
}

// SendErrorWithFunc sends error to slack channel but do not returns in return
func (l log) SendErrorWithFunc(err error, i interface{}) {
	l.Warn("Function name: "+getFunctionName(i)+" | Error: ", err)
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (l log) LogError(err error, i interface{}) {
	if err != nil {
		l.Warn("Function name: "+getFunctionName(i)+" | Error: ", err)
	}
}
