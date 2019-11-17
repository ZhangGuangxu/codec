package codec

type Codec interface {
	Decode(data []byte, msg interface{}) error
	DecodeWithID(msgID int16, data []byte) (msg interface{}, err error)
	DecodeWithName(msgName int16, data []byte) (msg interface{}, err error)
	Encode(msg interface{}) (data []byte, err error)
}
