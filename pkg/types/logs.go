package types

import "time"

type Monitor struct {
	Incoming <-chan Log
	Ctx      chan<- struct{} // Anything
}

type Log struct {
	Timestamp time.Time
	Message   string
}
