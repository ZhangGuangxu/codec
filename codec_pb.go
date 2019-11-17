package codec

import (
	"github.com/golang/protobuf/proto"
)

type PBCodec struct {
}

func (c *PBCodec) Decode(data []byte, msg interface{}) error {
	err := proto.Unmarshal(data, msg.(proto.Message))
	if err != nil {
		return err
	}
	return nil
}

func (c *PBCodec) Encode(msg interface{}) ([]byte, error) {
	data, err := proto.Marshal(msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return data, nil
}
