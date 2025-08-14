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
	Critical Severity = iota
	High
	Medium
	Low
	Unknown
)
