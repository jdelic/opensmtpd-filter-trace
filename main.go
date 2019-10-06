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
		log.Fatalf("error %v", err.Error())
	}
}

func (g *TraceFilter) Dataline(session string, params []string) {
	check(fmt.Fprintf(logfile, "dataline %v %v\n", session, strings.Join(params, "|")))
	line := strings.Join(params[1:], "|")
	fmt.Printf("filter-dataline|%s|%s|%s", params[0], session, line)
	logfile.Sync()
}

func (g *TraceFilter) Config(config []string) {
	check(fmt.Fprintf(logfile, "config %v\n", strings.Join(config, "|")))
	logfile.Sync()
}

func (g *TraceFilter) Commit(session string, params []string) {
	check(fmt.Fprintf(logfile, "commit %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxRollback(session string, params []string) {
	check(fmt.Fprintf(logfile, "tx-rollback %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxReset(session string, params []string) {
	check(fmt.Fprintf(logfile, "tx-reset %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxRcpt(session string, params []string) {
	check(fmt.Fprintf(logfile, "tx-rcpt %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxMail(session string, params []string) {
	check(fmt.Fprintf(logfile, "tx-mail %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxEnvelope(session string, params []string) {
	check(fmt.Fprintf(logfile, "tx-envelope %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxData(session string, params []string) {
	check(fmt.Fprintf(logfile, "tx-data %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxCommit(session string, params []string) {
	check(fmt.Fprintf(logfile, "tx-commit %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxBegin(session string, params []string) {
	check(fmt.Fprintf(logfile, "tx-begin %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkTLS(session string, params []string) {
	check(fmt.Fprintf(logfile, "link-tls %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkDisconnect(session string, params []string) {
	check(fmt.Fprintf(logfile, "link-disconnect %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkGreeting(session string, params []string) {
	check(fmt.Fprintf(logfile, "link-greeting %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkIdentity(session string, params []string) {
	check(fmt.Fprintf(logfile, "link-identity %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkAuth(session string, params []string) {
	check(fmt.Fprintf(logfile, "link-auth %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkConnect(session string, params []string) {
	check(fmt.Fprintf(logfile, "link-connect %v %v\n", session, strings.Join(params, "|")))
	logfile.Sync()
}

var logfile *os.File

func main() {
	logfile, _ =  os.OpenFile("/tmp/filterlogfile.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777);
	defer logfile.Close()

	_, _ = fmt.Fprintf(logfile, "Tracefilter starting\n")
	opensmtpd.Run(&TraceFilter{})
}
