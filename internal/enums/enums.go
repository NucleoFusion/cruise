package enums

type PageType int

const (
	Home PageType = iota
	Containers
	Images
)

type ErrorType int

const (
	Fatal = iota
	Warning
)
