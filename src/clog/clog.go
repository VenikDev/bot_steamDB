package clog

import (
	"github.com/charmbracelet/log"
	"time"
)

var (
	Logger = log.New(log.WithTimestamp(), log.WithTimeFormat(time.Kitchen),
		log.WithCaller())
)
