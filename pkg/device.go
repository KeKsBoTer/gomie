package gomie

type Device interface {
	withID
	withName
	canUpdate
	withAttributes
	Version() Version
	State() State
	Nodes() Nodes
	Extentions() Extentions
}
