package majsoul_api

import (
	"github.com/Yesterday17/majsoul_api/lq"
	"github.com/gogo/protobuf/proto"
)

type GameClient struct {
	socketClient
}

func NewGameClient(url string) (*GameClient, error) {
	client := &GameClient{}
	return client, client.init("FastTest", url)
}

func (l *GameClient) AuthGame(in *lq.ReqAuthGame, ch chan *lq.ResAuthGame) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("authGame", req, func(msg []byte) {
		data := &lq.ResAuthGame{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) EnterGame(in *lq.ReqCommon, ch chan *lq.ResEnterGame) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("enterGame", req, func(msg []byte) {
		data := &lq.ResEnterGame{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) SyncGame(in *lq.ReqSyncGame, ch chan *lq.ResSyncGame) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("syncGame", req, func(msg []byte) {
		data := &lq.ResSyncGame{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) FinishSyncGame(in *lq.ReqCommon, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("finishSyncGame", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) TerminateGame(in *lq.ReqCommon, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("terminateGame", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) InputOperation(in *lq.ReqSelfOperation, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("inputOperation", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) InputChiPengGang(in *lq.ReqChiPengGang, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("inputChiPengGang", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) ConfirmNewRound(ch chan *lq.ResCommon) {
	_ = l.sendRPC("confirmNewRound", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) BroadcastInGame(in *lq.ReqBroadcastInGame, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("broadcastInGame", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) InputGameGMCommand(in *lq.ReqGMCommandInGaming, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("inputGameGMCommand", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) FetchGamePlayerState(ch chan *lq.ResGamePlayerState) {
	_ = l.sendRPC("fetchGamePlayerState", ReqCommon, func(msg []byte) {
		data := &lq.ResGamePlayerState{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *GameClient) CheckNetworkDelay(ch chan *lq.ResCommon) {
	_ = l.sendRPC("checkNetworkDelay", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}
