package main

import (
	_ "git.chunyu.me/feiwang/appserver/routers"
	"github.com/astaxie/beego"
	"git.chunyu.me/feiwang/appserver/backends"
	"github.com/docopt/docopt-go"
	"fmt"

	log "git.chunyu.me/golang/cyutils/utils/rolling_log"
	"strconv"
	"os"
)

var usage = `Usage:
  %s [-L <log_file>] [--log-level=<loglevel>] [--log-keep-days=<maxdays>] [--nodb]
  %s -V | --version

options:
   -c <config_file>
   -L	set output log file, default is stdout
   --log-level=<loglevel>	set log level: info, warn, error, debug [default: info]
   --log-keep-days=<maxdays>  set max log file keep days, default is 3 days
   --profile-addr=<profile-addr>
   --work-dir=<work-dir>
   --code-url-version=<code-url-version>
`

func main() {
	version := "20160629"
	args, err := docopt.Parse(usage, nil, true, version, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if s, ok := args["-V"].(bool); ok && s {
		fmt.Println(backends.GreenF(version))
		os.Exit(1)
	}


	// 2. 解析Log相关的配置
	log.SetLevel(log.LEVEL_INFO)

	var maxKeepDays int = 3
	if s, ok := args["--log-keep-days"].(string); ok && s != "" {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			log.PanicErrorf(err, "invalid max log file keep days = %s", s)
		}
		maxKeepDays = int(v)
	}


	// set output log file
	if s, ok := args["-L"].(string); ok && s != "" {
		f, err := log.NewRollingFile(s, maxKeepDays)
		if err != nil {
			log.PanicErrorf(err, "open rolling log file failed: %s", s)
		} else {
			defer f.Close()
			log.StdLog = log.New(f, "")
		}
	}
	log.SetLevel(log.LEVEL_INFO)
	log.SetFlags(log.Flags() | log.Lshortfile)

	appsRoot := beego.AppConfig.String("apps_root")
	backends.ListAppDir(appsRoot)

	// 添加Watch
	done := backends.NewWatcher(appsRoot, func() {
		backends.ScanAppRootDir(appsRoot)
	})
	beego.Run()
	done <- true
}

