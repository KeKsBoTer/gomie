package gomie

type Node interface {
	withID
	withName
	withAttributes
	Type() string
	Properties() Properties
}
