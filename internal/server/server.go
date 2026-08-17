package server

import (
	"flag"
	"os"
	"time"
)

var (
	addr = flag.String("addr", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
)