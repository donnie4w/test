package timtest

type TidBean struct {
	Node     string
	Domain   string
	Resource string
	Termtyp  int8
}

type TimMessageBean struct {
	MsType    int8
	OdType    int8
	Id        int64
	Mid       int64
	BnType    int32
	FromTid   *TidBean
	ToTid     *TidBean
	Body      []byte
	IsOffline bool
	Timestamp int64
	Extend    map[string]string
}
