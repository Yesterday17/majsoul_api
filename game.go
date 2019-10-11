package majsoul_api

import (
	"github.com/Yesterday17/majsoul_api/lq"
	"github.com/gogo/protobuf/proto"
)

type Game struct {
	socket
}

func (l *Game) Init() error {
	return l.socket.init("FastTest")
}

func (l *Game) AuthGame(in *lq.ReqAuthGame, ch chan *lq.ResAuthGame) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("authGame", req, func(msg []byte) {
		data := &lq.ResAuthGame{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) EnterGame(in *lq.ReqCommon, ch chan *lq.ResEnterGame) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("enterGame", req, func(msg []byte) {
		data := &lq.ResEnterGame{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) SyncGame(in *lq.ReqSyncGame, ch chan *lq.ResSyncGame) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("syncGame", req, func(msg []byte) {
		data := &lq.ResSyncGame{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) FinishSyncGame(in *lq.ReqCommon, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("finishSyncGame", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) TerminateGame(in *lq.ReqCommon, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("terminateGame", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) InputOperation(in *lq.ReqSelfOperation, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("inputOperation", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) InputChiPengGang(in *lq.ReqChiPengGang, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("inputChiPengGang", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) ConfirmNewRound(ch chan *lq.ResCommon) {
	_ = l.sendRPC("confirmNewRound", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) BroadcastInGame(in *lq.ReqBroadcastInGame, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("broadcastInGame", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) InputGameGMCommand(in *lq.ReqGMCommandInGaming, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("inputGameGMCommand", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) FetchGamePlayerState(ch chan *lq.ResGamePlayerState) {
	_ = l.sendRPC("fetchGamePlayerState", ReqCommon, func(msg []byte) {
		data := &lq.ResGamePlayerState{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *Game) CheckNetworkDelay(ch chan *lq.ResCommon) {
	_ = l.sendRPC("checkNetworkDelay", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}
