package gomie

type Property interface {
	withID
	withName
	withAttributes
	Datatype() PayloadType
	ReceiveValue(string)
}
