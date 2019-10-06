package main

import (
	"fmt"
	"github.com/jdelic/opensmtpd-filters-go"
	"os"
	"strings"
)


type TraceFilter struct {
}

func check(n int, err error) {
	if err != nil {
		panic("error " + err.Error())
	}
}

func (g *TraceFilter) Dataline(session string, params []string) {
	check(fmt.Fprintf(log, "dataline %v %v\n", session, strings.Join(params, "|")))
	line := strings.Join(params[1:], "|")
	fmt.Printf("filter-dataline|%s|%s|%s", params[0], session, line)
	log.Sync()
}

func (g *TraceFilter) Config(config []string) {
	check(fmt.Fprintf(log, "config %v\n", strings.Join(config, "|")))
	log.Sync()
}

func (g *TraceFilter) Commit(session string, params []string) {
	check(fmt.Fprintf(log, "commit %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) TxRollback(session string, params []string) {
	check(fmt.Fprintf(log, "tx-rollback %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) TxReset(session string, params []string) {
	check(fmt.Fprintf(log, "tx-reset %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) TxRcpt(session string, params []string) {
	check(fmt.Fprintf(log, "tx-rcpt %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) TxMail(session string, params []string) {
	check(fmt.Fprintf(log, "tx-mail %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) TxEnvelope(session string, params []string) {
	check(fmt.Fprintf(log, "tx-envelope %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) TxData(session string, params []string) {
	check(fmt.Fprintf(log, "tx-data %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) TxCommit(session string, params []string) {
	check(fmt.Fprintf(log, "tx-commit %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) TxBegin(session string, params []string) {
	check(fmt.Fprintf(log, "tx-begin %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) LinkTLS(session string, params []string) {
	check(fmt.Fprintf(log, "link-tls %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) LinkDisconnect(session string, params []string) {
	check(fmt.Fprintf(log, "link-disconnect %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) LinkGreeting(session string, params []string) {
	check(fmt.Fprintf(log, "link-greeting %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) LinkIdentity(session string, params []string) {
	check(fmt.Fprintf(log, "link-identity %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) LinkAuth(session string, params []string) {
	check(fmt.Fprintf(log, "link-auth %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

func (g *TraceFilter) LinkConnect(session string, params []string) {
	check(fmt.Fprintf(log, "link-connect %v %v\n", session, strings.Join(params, "|")))
	log.Sync()
}

var log *os.File

func main() {
	log, _ =  os.OpenFile("/tmp/filterlog.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777);
	defer log.Close()

	_, _ = fmt.Fprintf(log, "Tracefilter starting\n")
	opensmtpd.Run(&TraceFilter{})
}
