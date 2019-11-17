package codec

type Encoder interface {
	Encode(msg interface{}) (data []byte, err error)
}

type Codec interface {
	Encoder
	Decode(data []byte, msg interface{}) error
}

type IDCodec interface {
	Encoder
	DecodeWithID(msgID int16, data []byte) (msg interface{}, err error)
}

type NameCodec interface {
	Encoder
	DecodeWithName(msgName int16, data []byte) (msg interface{}, err error)
}
