package main

import (
	"fmt"
	"github.com/jdelic/opensmtpd-filters-go"
	"log"
	"os"
	"strings"
)


type TraceFilter struct {}

func check(n int, err error) {
	if err != nil {
		log.Fatalf("error %d %T %v", n, err, err)
	}
}

/* config */

func (g *TraceFilter) Config(config []string) {
	check(fmt.Fprintf(logfile, "config %v\n", strings.Join(config, "|")))
	logfile.Sync()
}


/* reporters */

func (g *TraceFilter) LinkConnect(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-connect %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkDisconnect(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-disconnect %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkGreeting(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-greeting %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkIdentify(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-identify %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkTLS(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-tls %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) LinkAuth(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "link-auth %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxReset(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-reset %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxBegin(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-begin %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxMail(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-mail %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxRcpt(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-rcpt %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxEnvelope(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-envelope %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxData(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-data %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxCommit(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-commit %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) TxRollback(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "tx-rollback %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) ProtocolClient(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "protocol-client %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) ProtocolServer(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "protocol-server %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) FilterReport(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "filter-report %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) FilterResponse(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "filter-response %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}

func (g *TraceFilter) Timeout(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "timeout %v %v\n", sessionId, strings.Join(params, "|")))
	logfile.Sync()
}


/* filters */

func (g *TraceFilter) Connect(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "connect %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Helo(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "helo %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Ehlo(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "Ehlo %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) StartTLS(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "starttls %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Auth(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "auth %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) MailFrom(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "mail-from %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) RcptTo(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "rcpt-to %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Data(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "data %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Dataline(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "dataline %v %v\n", sessionId, strings.Join(params, "|")))
	line := strings.Join(params[1:], "|")
	opensmtpd.DatalineReply(params[0], sessionId, line)
}

func (g *TraceFilter) Rset(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "rset %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Quit(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "noop %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Help(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "help %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Wiz(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "wiz %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

func (g *TraceFilter) Commit(fw opensmtpd.FilterWrapper, verb string, sh opensmtpd.SessionHolder, sessionId string, params []string) {
	check(fmt.Fprintf(logfile, "commit %v %v\n", sessionId, strings.Join(params, "|")))
	opensmtpd.Proceed(params[0], sessionId)
}

var logfile *os.File

func main() {
	log.SetOutput(os.Stderr)
	var err error
	logfile, err = os.OpenFile("/tmp/filterlog.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("Can't open /tmp/filterlog.txt %v", err)
	}
	defer logfile.Close()

	log.Println("Tracefilter starting")

	traceFilter := opensmtpd.NewFilter(&TraceFilter{})
	opensmtpd.Run(traceFilter)
}
