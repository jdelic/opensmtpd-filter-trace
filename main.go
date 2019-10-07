package main

import (
	"fmt"
	"github.com/jdelic/opensmtpd-filters-go"
	"os"
	"strings"
    "log"
)


type TraceFilter struct {
}

func check(n int, err error) {
	if err != nil {
		log.Fatalf("error %d %T %v", n, err, err)
	}
}

func (g *TraceFilter) Dataline(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "dataline %v %v\n", sessionId, strings.Join(params, "|")))
	line := strings.Join(params[1:], "|")
	opensmtpd.DatalineReply(params[0], sessionId, line)
	logfile.Sync()
}

func (g *TraceFilter) Config(config []string) {
	check(fmt.Fprintf(logfile, "config %v\n", strings.Join(config, "|")))
	logfile.Sync()
}

func (g *TraceFilter) Commit(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "commit %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
	logfile.Sync()
}

func (g *TraceFilter) TxRollback(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-rollback %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxReset(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-reset %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxRcpt(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-rcpt %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxMail(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-mail %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxEnvelope(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-envelope %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxData(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-data %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxCommit(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-commit %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxBegin(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-begin %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkTLS(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-tls %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkDisconnect(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-disconnect %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkGreeting(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-greeting %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkIdentity(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-identity %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkAuth(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-auth %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkConnect(sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-connect %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

var logfile *os.File

func main() {
	var err error
	logfile, err = os.OpenFile("/tmp/filterlog.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("Can't open /tmp/filterlog.txt %v", err)
	}
	defer logfile.Close()

	_, _ = fmt.Fprintf(logfile, "Tracefilter starting\n")
	opensmtpd.Run(&TraceFilter{})
}
