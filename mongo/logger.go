package mongo

import (
	"log"
	"os"
)

var (
	mongodumpLog = log.New(os.Stderr, "    mongodump     | ", 0)
)
