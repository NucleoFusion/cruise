package enums

type PageType int

const (
	Home PageType = iota
	Containers
	Images
	Vulnerability
)

type ErrorType int

const (
	Fatal ErrorType = iota
	Warning
)

type Severity int

const (
	High Severity = iota
	Medium
)
