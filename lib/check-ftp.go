package checkftp

import (
	"os"
	"fmt"
	"time"
	"crypto/tls"

	"github.com/mackerelio/checkers"
	"github.com/jessevdk/go-flags"
	"github.com/BenefitsDataTrust/ftp"
)

type options struct {
	Host               string  `short:"H" long:"host" default:"localhost" description:"Hostname"`
	Port               int     `short:"P" long:"port" default:"21" description:"Port"`
	User               string  `short:"u" long:"user" default:"anonymous" description:"FTP username"`
	Password           string  `short:"p" long:"password" default:"anonymous" description:"FTP password"`
	Warning            float64 `short:"w" long:"warning" description:"Warning threshold (sec)"`
	Critical           float64 `short:"c" long:"critical" description:"Critical threshold (sec)"`
	Timeout            int     `short:"t" long:"timeout" default:"10" description:"Timeout (sec)"`
	FTPS               bool    `short:"s" long:"ftps" description:"Use FTPS"`
	ImplicitMode       bool    `short:"i" long:"implicit-mode" description:"Connects directly using TLS"`
	NoCheckCertificate bool    `long:"no-check-certificate" description:"Do not check certificate"`
}

// Do the plugin
func Do() {
	ckr := run(os.Args[1:])
	ckr.Name = "FTP"
	ckr.Exit()
}

func run(args []string) *checkers.Checker {
	var opts options
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		os.Exit(1)
	}

	if opts.Warning == 0 && opts.Critical == 0 {
		return checkers.Unknown("require threshold option (warning or critical)")
	}

	addr := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	dialOpts := []ftp.DialOption{
		ftp.DialWithTimeout(time.Duration(opts.Timeout) * time.Second),
	}

	if opts.FTPS {
		tlsOpt := ftp.DialWithTLS(
			&tls.Config{InsecureSkipVerify: opts.NoCheckCertificate, ServerName: opts.Host},
			opts.ImplicitMode,
		)
		dialOpts = append(dialOpts, tlsOpt)
	}

	stTime := time.Now()

	c, err := ftp.Dial(addr, dialOpts...)
	if err != nil {
		return checkers.Unknown(err.Error())
	}
	defer c.Quit()

	err = c.Login(opts.User, opts.Password)
	if err != nil {
		return checkers.Unknown(err.Error())
	}

	elapsed := time.Since(stTime)
	msg := fmt.Sprintf("%.3f seconds response time", elapsed.Seconds())

	if opts.Critical != 0 && elapsed.Seconds() > opts.Critical {
		return checkers.Critical(msg)
	} else if opts.Warning != 0 && elapsed.Seconds() > opts.Warning {
		return checkers.Warning(msg)
	}

	return checkers.Ok("Get,Set OK")

} 
