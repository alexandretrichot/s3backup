package common

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	AppLog    = log.New(os.Stderr, "[info]  s3backup  | ", 0)
	AppErrLog = log.New(os.Stderr, color.RedString("[error] s3backup  | "), 0)
)
