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

func (g TraceFilter) Dataline(session string, params []string) {
	check(fmt.Fprintf(log, "dataline %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) Config(config []string) {
	_, err := fmt.Fprintf(log, "config %s\n", strings.Join(config, "|"))
	if err != nil {
		panic("error " + err.Error())
	}
}

func (g TraceFilter) Commit(session string, params []string) {
	check(fmt.Fprintf(log, "commit %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) TxRollback(session string, params []string) {
	check(fmt.Fprintf(log, "tx-rollback %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) TxReset(session string, params []string) {
	check(fmt.Fprintf(log, "tx-reset %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) TxRcpt(session string, params []string) {
	check(fmt.Fprintf(log, "tx-rcpt %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) TxMail(session string, params []string) {
	check(fmt.Fprintf(log, "tx-mail %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) TxEnvelope(session string, params []string) {
	check(fmt.Fprintf(log, "tx-envelope %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) TxData(session string, params []string) {
	check(fmt.Fprintf(log, "tx-data %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) TxCommit(session string, params []string) {
	check(fmt.Fprintf(log, "tx-commit %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) TxBegin(session string, params []string) {
	check(fmt.Fprintf(log, "tx-begin %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) LinkTLS(session string, params []string) {
	check(fmt.Fprintf(log, "link-tls %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) LinkDisconnect(session string, params []string) {
	check(fmt.Fprintf(log, "link-disconnect %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) LinkGreeting(session string, params []string) {
	check(fmt.Fprintf(log, "link-greeting %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) LinkIdentity(session string, params []string) {
	check(fmt.Fprintf(log, "link-identity %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) LinkAuth(session string, params []string) {
	check(fmt.Fprintf(log, "link-auth %s %s\n", session, strings.Join(params, "|")))
}

func (g TraceFilter) LinkConnect(session string, params []string) {
	check(fmt.Fprintf(log, "link-connect %s %s\n", session, strings.Join(params, "|")))
}

var log *os.File

func main() {
	log, _ = os.Create("/tmp/filterlog.txt")
	defer log.Close()

	_, _ = fmt.Fprintf(log, "Tracefilter starting\n")
	opensmtpd.Run(TraceFilter{})
}
