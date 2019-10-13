package majsoul_api

import (
	"github.com/Yesterday17/majsoul_api/lq"
	"github.com/gogo/protobuf/proto"
)

type gameClient struct {
	socketClient
}

func NewGameClient(base, url string) (*gameClient, error) {
	client := &gameClient{}
	return client, client.init(base, "FastTest", url)
}

func (l *gameClient) AuthGame(in *lq.ReqAuthGame) chan lq.ResAuthGame {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResAuthGame)
	_ = l.sendRPC("authGame", req, func(msg []byte) {
		data := lq.ResAuthGame{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) EnterGame() chan lq.ResEnterGame {
	future := make(chan lq.ResEnterGame)
	_ = l.sendRPC("enterGame", reqCommon, func(msg []byte) {
		data := lq.ResEnterGame{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) SyncGame(in *lq.ReqSyncGame) chan lq.ResSyncGame {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResSyncGame)
	_ = l.sendRPC("syncGame", req, func(msg []byte) {
		data := lq.ResSyncGame{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) FinishSyncGame() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("finishSyncGame", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) TerminateGame() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("terminateGame", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) InputOperation(in *lq.ReqSelfOperation) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("inputOperation", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) InputChiPengGang(in *lq.ReqChiPengGang) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("inputChiPengGang", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) ConfirmNewRound() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("confirmNewRound", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) BroadcastInGame(in *lq.ReqBroadcastInGame) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("broadcastInGame", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) InputGameGMCommand(in *lq.ReqGMCommandInGaming) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("inputGameGMCommand", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) FetchGamePlayerState() chan lq.ResGamePlayerState {
	future := make(chan lq.ResGamePlayerState)
	_ = l.sendRPC("fetchGamePlayerState", reqCommon, func(msg []byte) {
		data := lq.ResGamePlayerState{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *gameClient) CheckNetworkDelay() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("checkNetworkDelay", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}
