package loader

type Error string

// Error is for implementing error interface
func (e Error) Error() string { return string(e) }

const (
	ErrStatusNotOK     = Error("got not-OK status code")
	ErrNilQueryPointer = Error("got nil query pointer")
	ErrNilNodesArray   = Error("got nil nodes array")
	ErrEmptyNodesArray = Error("got empty nodes array")
	ErrNilNodePointer  = Error("got nil node pointer")
	ErrNilAttrArray    = Error("got nil attr array")
	ErrEmptyAttrArray  = Error("got empty attr array")
	ErrSrcAttrNotFound = Error("attr 'src' not found")
)
