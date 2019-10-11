package majsoul_api

import (
	"github.com/Yesterday17/majsoul_api/lq"
	"github.com/golang/protobuf/proto"
)

var ReqCommon, _ = proto.Marshal(&lq.ReqCommon{})

func wrap(name string, msg []byte) (ret []byte) {
	ret, _ = proto.Marshal(&lq.Wrapper{
		Name: name,
		Data: msg,
	})
	return
}

func unWrap(p []byte) []byte {
	msg := &lq.Wrapper{}
	_ = proto.Unmarshal(p, msg)
	return msg.Data
}
