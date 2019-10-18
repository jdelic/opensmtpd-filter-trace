package main

import (
	"bufio"
	"fmt"
	"github.com/jdelic/opensmtpd-filters-go"
	"os"
	"strings"
    "log"
)


type TraceFilter struct {
}

func (g *TraceFilter) GetCapabilities() opensmtpd.FilterDispatchMap {
	return opensmtpd.GetCapabilities(g)
}

func (g *TraceFilter) Register() {
	opensmtpd.Register(g)
}

func (g *TraceFilter) Dispatch(params []string) {
	check(fmt.Fprintf(logfile, "INPUT %v\n", strings.Join(params, "|")))
	opensmtpd.Dispatch(g, params)
}

func (g *TraceFilter) ProcessConfig(scanner *bufio.Scanner) {
	opensmtpd.ProcessConfig(g, scanner)
}

func check(n int, err error) {
	if err != nil {
		log.Fatalf("error %d %T %v", n, err, err)
	}
}

func (g *TraceFilter) Dataline(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "dataline %v %v\n", sessionId, strings.Join(params, "|")))
	line := strings.Join(params[1:], "|")
	opensmtpd.DatalineReply(params[0], sessionId, line)
	logfile.Sync()
}

func (g *TraceFilter) Config(config []string) {
	check(fmt.Fprintf(logfile, "config %v\n", strings.Join(config, "|")))
	logfile.Sync()
}

func (g *TraceFilter) Commit(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "commit %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
	logfile.Sync()
}

func (g *TraceFilter) TxRollback(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-rollback %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxReset(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-reset %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxRcpt(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-rcpt %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxMail(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-mail %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxEnvelope(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-envelope %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxData(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-data %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxCommit(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-commit %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxBegin(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-begin %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkTLS(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-tls %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkDisconnect(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-disconnect %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkGreeting(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-greeting %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkIdentity(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-identity %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkAuth(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-auth %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkConnect(verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
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
