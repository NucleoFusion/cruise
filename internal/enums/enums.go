package enums

type PageType int

const (
	Home PageType = iota
	Containers
)

type ErrorType int

const (
	Fatal = iota
	Warning
)
