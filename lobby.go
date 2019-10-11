package majsoul_api

import (
	"github.com/Yesterday17/majsoul_api/lq"
	"github.com/gogo/protobuf/proto"
)

type LobbyClient struct {
	socketClient
}

func NewLobbyClient(url string) (*LobbyClient, error) {
	client := &LobbyClient{}
	return client, client.init("Lobby", url)
}

func (l *LobbyClient) FetchConnectionInfo(ch chan *lq.ResConnectionInfo) {
	_ = l.sendRPC("fetchConnectionInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResConnectionInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) Signup(in *lq.ReqSignupAccount, ch chan *lq.ResSignupAccount) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("signup", req, func(msg []byte) {
		data := &lq.ResSignupAccount{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) Login(in *lq.ReqLogin, ch chan *lq.ResLogin) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("login", req, func(msg []byte) {
		data := &lq.ResLogin{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) EmailLogin(in *lq.ReqEmailLogin, ch chan *lq.ResLogin) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("emailLogin", req, func(msg []byte) {
		data := &lq.ResLogin{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) Oauth2Auth(in *lq.ReqOauth2Auth, ch chan *lq.ResOauth2Auth) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("oauth2Auth", req, func(msg []byte) {
		data := &lq.ResOauth2Auth{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) Oauth2Check(in *lq.ReqOauth2Check, ch chan *lq.ResOauth2Check) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("oauth2Check", req, func(msg []byte) {
		data := &lq.ResOauth2Check{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) Oauth2Signup(in *lq.ReqOauth2Signup, ch chan *lq.ResOauth2Signup) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("oauth2Signup", req, func(msg []byte) {
		data := &lq.ResOauth2Signup{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) Oauth2Login(in *lq.ReqOauth2Login, ch chan *lq.ResLogin) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("oauth2Login", req, func(msg []byte) {
		data := &lq.ResLogin{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreatePhoneVerifyCode(in *lq.ReqCreatePhoneVerifyCode, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createPhoneVerifyCode", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateEmailVerifyCode(in *lq.ReqCreateEmailVerifyCode, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createEmailVerifyCode", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) VerfifyCodeForSecure(in *lq.ReqVerifyCodeForSecure, ch chan *lq.ResVerfiyCodeForSecure) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("verfifyCodeForSecure", req, func(msg []byte) {
		data := &lq.ResVerfiyCodeForSecure{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) BindPhoneNumber(in *lq.ReqBindPhoneNumber, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("bindPhoneNumber", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) BindEmail(in *lq.ReqBindEmail, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("bindEmail", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ModifyPassword(in *lq.ReqModifyPassword, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("modifyPassword", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) BindAccount(in *lq.ReqBindAccount, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("bindAccount", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) Logout(in *lq.ReqLogout, ch chan *lq.ResLogout) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("logout", req, func(msg []byte) {
		data := &lq.ResLogout{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) Heatbeat(in *lq.ReqHeatBeat, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("heatbeat", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) LoginBeat(in *lq.ReqLoginBeat, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("loginBeat", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateNickname(in *lq.ReqCreateNickname, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createNickname", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ModifyNickname(in *lq.ReqModifyNickname, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("modifyNickname", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ModifyBirthday(in *lq.ReqModifyBirthday, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("modifyBirthday", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchRoom(ch chan *lq.ResSelfRoom) {
	_ = l.sendRPC("fetchRoom", ReqCommon, func(msg []byte) {
		data := &lq.ResSelfRoom{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateRoom(in *lq.ReqCreateRoom, ch chan *lq.ResCreateRoom) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createRoom", req, func(msg []byte) {
		data := &lq.ResCreateRoom{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) JoinRoom(in *lq.ReqJoinRoom, ch chan *lq.ResJoinRoom) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("joinRoom", req, func(msg []byte) {
		data := &lq.ResJoinRoom{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) LeaveRoom(ch chan *lq.ResCommon) {
	_ = l.sendRPC("leaveRoom", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ReadyPlay(in *lq.ReqRoomReady, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("readyPlay", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) StartRoom(in *lq.ReqRoomStart, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("startRoom", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) KickPlayer(in *lq.ReqRoomKick, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("kickPlayer", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ModifyRoom(in *lq.ReqModifyRoom, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("modifyRoom", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) MatchGame(in *lq.ReqJoinMatchQueue, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("matchGame", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CancelMatch(in *lq.ReqCancelMatchQueue, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("cancelMatch", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchAccountInfo(in *lq.ReqAccountInfo, ch chan *lq.ResAccountInfo) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchAccountInfo", req, func(msg []byte) {
		data := &lq.ResAccountInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ChangeAvatar(in *lq.ReqChangeAvatar, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("changeAvatar", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchAccountStatisticInfo(in *lq.ReqAccountStatisticInfo, ch chan *lq.ResAccountStatisticInfo) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchAccountStatisticInfo", req, func(msg []byte) {
		data := &lq.ResAccountStatisticInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchAccountCharacterInfo(ch chan *lq.ResAccountCharacterInfo) {
	_ = l.sendRPC("fetchAccountCharacterInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResAccountCharacterInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ShopPurchase(in *lq.ReqShopPurchase, ch chan *lq.ResShopPurchase) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("shopPurchase", req, func(msg []byte) {
		data := &lq.ResShopPurchase{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchGameRecord(in *lq.ReqGameRecord, ch chan *lq.ResGameRecord) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchGameRecord", req, func(msg []byte) {
		data := &lq.ResGameRecord{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchGameRecordList(in *lq.ReqGameRecordList, ch chan *lq.ResGameRecordList) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchGameRecordList", req, func(msg []byte) {
		data := &lq.ResGameRecordList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCollectedGameRecordList(ch chan *lq.ResCollectedGameRecordList) {
	_ = l.sendRPC("fetchCollectedGameRecordList", ReqCommon, func(msg []byte) {
		data := &lq.ResCollectedGameRecordList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchGameRecordsDetail(in *lq.ReqGameRecordsDetail, ch chan *lq.ResGameRecordsDetail) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchGameRecordsDetail", req, func(msg []byte) {
		data := &lq.ResGameRecordsDetail{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) AddCollectedGameRecord(in *lq.ReqAddCollectedGameRecord, ch chan *lq.ResAddCollectedGameRecord) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("addCollectedGameRecord", req, func(msg []byte) {
		data := &lq.ResAddCollectedGameRecord{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) RemoveCollectedGameRecord(in *lq.ReqRemoveCollectedGameRecord, ch chan *lq.ResRemoveCollectedGameRecord) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("removeCollectedGameRecord", req, func(msg []byte) {
		data := &lq.ResRemoveCollectedGameRecord{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ChangeCollectedGameRecordRemarks(in *lq.ReqChangeCollectedGameRecordRemarks, ch chan *lq.ResChangeCollectedGameRecordRemarks) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("changeCollectedGameRecordRemarks", req, func(msg []byte) {
		data := &lq.ResChangeCollectedGameRecordRemarks{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchLevelLeaderboard(in *lq.ReqLevelLeaderboard, ch chan *lq.ResLevelLeaderboard) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchLevelLeaderboard", req, func(msg []byte) {
		data := &lq.ResLevelLeaderboard{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchMultiAccountBrief(in *lq.ReqMultiAccountId, ch chan *lq.ResMultiAccountBrief) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchMultiAccountBrief", req, func(msg []byte) {
		data := &lq.ResMultiAccountBrief{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchFriendList(ch chan *lq.ResFriendList) {
	_ = l.sendRPC("fetchFriendList", ReqCommon, func(msg []byte) {
		data := &lq.ResFriendList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchFriendApplyList(ch chan *lq.ResFriendApplyList) {
	_ = l.sendRPC("fetchFriendApplyList", ReqCommon, func(msg []byte) {
		data := &lq.ResFriendApplyList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ApplyFriend(in *lq.ReqApplyFriend, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("applyFriend", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) HandleFriendApply(in *lq.ReqHandleFriendApply, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("handleFriendApply", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) RemoveFriend(in *lq.ReqRemoveFriend, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("removeFriend", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) SearchAccountById(in *lq.ReqSearchAccountById, ch chan *lq.ResSearchAccountById) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("searchAccountById", req, func(msg []byte) {
		data := &lq.ResSearchAccountById{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) SearchAccountByPattern(in *lq.ReqSearchAccountByPattern, ch chan *lq.ResSearchAccountByPattern) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("searchAccountByPattern", req, func(msg []byte) {
		data := &lq.ResSearchAccountByPattern{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchAccountState(in *lq.ReqAccountList, ch chan *lq.ResAccountStates) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchAccountState", req, func(msg []byte) {
		data := &lq.ResAccountStates{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchBagInfo(ch chan *lq.ResBagInfo) {
	_ = l.sendRPC("fetchBagInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResBagInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UseBagItem(in *lq.ReqUseBagItem, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("useBagItem", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) OpenManualItem(in *lq.ReqOpenManualItem, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("openManualItem", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) OpenRandomRewardItem(in *lq.ReqOpenRandomRewardItem, ch chan *lq.ResOpenRandomRewardItem) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("openRandomRewardItem", req, func(msg []byte) {
		data := &lq.ResOpenRandomRewardItem{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ComposeShard(in *lq.ReqComposeShard, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("composeShard", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchAnnouncement(ch chan *lq.ResAnnouncement) {
	_ = l.sendRPC("fetchAnnouncement", ReqCommon, func(msg []byte) {
		data := &lq.ResAnnouncement{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ReadAnnouncement(in *lq.ReqReadAnnouncement, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("readAnnouncement", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchMailInfo(ch chan *lq.ResMailInfo) {
	_ = l.sendRPC("fetchMailInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResMailInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ReadMail(in *lq.ReqReadMail, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("readMail", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) DeleteMail(in *lq.ReqDeleteMail, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("deleteMail", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) TakeAttachmentFromMail(in *lq.ReqTakeAttachment, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("takeAttachmentFromMail", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchAchievement(ch chan *lq.ResAchievement) {
	_ = l.sendRPC("fetchAchievement", ReqCommon, func(msg []byte) {
		data := &lq.ResAchievement{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) BuyShiLian(in *lq.ReqBuyShiLian, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("buyShiLian", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) MatchShiLian(ch chan *lq.ResCommon) {
	_ = l.sendRPC("matchShiLian", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) GoNextShiLian(ch chan *lq.ResCommon) {
	_ = l.sendRPC("goNextShiLian", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UpdateClientValue(in *lq.ReqUpdateClientValue, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("updateClientValue", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchClientValue(ch chan *lq.ResClientValue) {
	_ = l.sendRPC("fetchClientValue", ReqCommon, func(msg []byte) {
		data := &lq.ResClientValue{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ClientMessage(in *lq.ReqClientMessage, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("clientMessage", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCurrentMatchInfo(in *lq.ReqCurrentMatchInfo, ch chan *lq.ResCurrentMatchInfo) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCurrentMatchInfo", req, func(msg []byte) {
		data := &lq.ResCurrentMatchInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchReviveCoinInfo(ch chan *lq.ResReviveCoinInfo) {
	_ = l.sendRPC("fetchReviveCoinInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResReviveCoinInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) GainReviveCoin(ch chan *lq.ResCommon) {
	_ = l.sendRPC("gainReviveCoin", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchDailyTask(ch chan *lq.ResDailyTask) {
	_ = l.sendRPC("fetchDailyTask", ReqCommon, func(msg []byte) {
		data := &lq.ResDailyTask{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) RefreshDailyTask(in *lq.ReqRefreshDailyTask, ch chan *lq.ResRefreshDailyTask) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("refreshDailyTask", req, func(msg []byte) {
		data := &lq.ResRefreshDailyTask{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UseGiftCode(in *lq.ReqUseGiftCode, ch chan *lq.ResUseGiftCode) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("useGiftCode", req, func(msg []byte) {
		data := &lq.ResUseGiftCode{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchTitleList(ch chan *lq.ResTitleList) {
	_ = l.sendRPC("fetchTitleList", ReqCommon, func(msg []byte) {
		data := &lq.ResTitleList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UseTitle(in *lq.ReqUseTitle, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("useTitle", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) SendClientMessage(in *lq.ReqSendClientMessage, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("sendClientMessage", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchGameLiveInfo(in *lq.ReqGameLiveInfo, ch chan *lq.ResGameLiveInfo) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchGameLiveInfo", req, func(msg []byte) {
		data := &lq.ResGameLiveInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchGameLiveLeftSegment(in *lq.ReqGameLiveLeftSegment, ch chan *lq.ResGameLiveLeftSegment) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchGameLiveLeftSegment", req, func(msg []byte) {
		data := &lq.ResGameLiveLeftSegment{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchGameLiveList(in *lq.ReqGameLiveList, ch chan *lq.ResGameLiveList) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchGameLiveList", req, func(msg []byte) {
		data := &lq.ResGameLiveList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCommentSetting(ch chan *lq.ResCommentSetting) {
	_ = l.sendRPC("fetchCommentSetting", ReqCommon, func(msg []byte) {
		data := &lq.ResCommentSetting{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UpdateCommentSetting(in *lq.ReqUpdateCommentSetting, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("updateCommentSetting", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCommentList(in *lq.ReqFetchCommentList, ch chan *lq.ResFetchCommentList) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCommentList", req, func(msg []byte) {
		data := &lq.ResFetchCommentList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCommentContent(in *lq.ReqFetchCommentContent, ch chan *lq.ResFetchCommentContent) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCommentContent", req, func(msg []byte) {
		data := &lq.ResFetchCommentContent{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) LeaveComment(in *lq.ReqLeaveComment, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("leaveComment", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) DeleteComment(in *lq.ReqDeleteComment, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("deleteComment", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UpdateReadComment(in *lq.ReqUpdateReadComment, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("updateReadComment", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchRollingNotice(ch chan *lq.ReqRollingNotice) {
	_ = l.sendRPC("fetchRollingNotice", ReqCommon, func(msg []byte) {
		data := &lq.ReqRollingNotice{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchServerTime(ch chan *lq.ResServerTime) {
	_ = l.sendRPC("fetchServerTime", ReqCommon, func(msg []byte) {
		data := &lq.ResServerTime{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchPlatformProducts(in *lq.ReqPlatformBillingProducts, ch chan *lq.ResPlatformBillingProducts) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchPlatformProducts", req, func(msg []byte) {
		data := &lq.ResPlatformBillingProducts{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateBillingOrder(in *lq.ReqCreateBillingOrder, ch chan *lq.ResCreateBillingOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createBillingOrder", req, func(msg []byte) {
		data := &lq.ResCreateBillingOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) SolveGooglePlayOrder(in *lq.ReqSolveGooglePlayOrder, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("solveGooglePlayOrder", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CancelGooglePlayOrder(in *lq.ReqCancelGooglePlayOrder, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("cancelGooglePlayOrder", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) OpenChest(in *lq.ReqOpenChest, ch chan *lq.ResOpenChest) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("openChest", req, func(msg []byte) {
		data := &lq.ResOpenChest{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) BuyFromChestShop(in *lq.ReqBuyFromChestShop, ch chan *lq.ResBuyFromChestShop) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("buyFromChestShop", req, func(msg []byte) {
		data := &lq.ResBuyFromChestShop{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchDailySignInInfo(ch chan *lq.ResDailySignInInfo) {
	_ = l.sendRPC("fetchDailySignInInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResDailySignInInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) DoDailySignIn(ch chan *lq.ResCommon) {
	_ = l.sendRPC("doDailySignIn", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) DoActivitySignIn(in *lq.ReqDoActivitySignIn, ch chan *lq.ResDoActivitySignIn) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("doActivitySignIn", req, func(msg []byte) {
		data := &lq.ResDoActivitySignIn{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCharacterInfo(ch chan *lq.ResCharacterInfo) {
	_ = l.sendRPC("fetchCharacterInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResCharacterInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ChangeMainCharacter(in *lq.ReqChangeMainCharacter, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("changeMainCharacter", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ChangeCharacterSkin(in *lq.ReqChangeCharacterSkin, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("changeCharacterSkin", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ChangeCharacterView(in *lq.ReqChangeCharacterView, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("changeCharacterView", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) SendGiftToCharacter(in *lq.ReqSendGiftToCharacter, ch chan *lq.ResSendGiftToCharacter) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("sendGiftToCharacter", req, func(msg []byte) {
		data := &lq.ResSendGiftToCharacter{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) SellItem(in *lq.ReqSellItem, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("sellItem", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCommonView(ch chan *lq.ResCommonView) {
	_ = l.sendRPC("fetchCommonView", ReqCommon, func(msg []byte) {
		data := &lq.ResCommonView{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ChangeCommonView(in *lq.ReqChangeCommonView, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("changeCommonView", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UpgradeCharacter(in *lq.ReqUpgradeCharacter, ch chan *lq.ResUpgradeCharacter) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("upgradeCharacter", req, func(msg []byte) {
		data := &lq.ResUpgradeCharacter{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) GameMasterCommand(in *lq.ReqGMCommand, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("gameMasterCommand", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchShopInfo(ch chan *lq.ResShopInfo) {
	_ = l.sendRPC("fetchShopInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResShopInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) BuyFromShop(in *lq.ReqBuyFromShop, ch chan *lq.ResBuyFromShop) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("buyFromShop", req, func(msg []byte) {
		data := &lq.ResBuyFromShop{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) BuyFromZHP(in *lq.ReqBuyFromZHP, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("buyFromZHP", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) RefreshZHPShop(ch chan *lq.ResRefreshZHPShop) {
	_ = l.sendRPC("refreshZHPShop", ReqCommon, func(msg []byte) {
		data := &lq.ResRefreshZHPShop{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchMonthTicketInfo(ch chan *lq.ResMonthTicketInfo) {
	_ = l.sendRPC("fetchMonthTicketInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResMonthTicketInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) PayMonthTicket(in *lq.ReqPayMonthTicket, ch chan *lq.ResPayMonthTicket) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("payMonthTicket", req, func(msg []byte) {
		data := &lq.ResPayMonthTicket{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ExchangeCurrency(in *lq.ReqExchangeCurrency, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("exchangeCurrency", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ExchangeChestStone(in *lq.ReqExchangeCurrency, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("exchangeChestStone", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchServerSettings(ch chan *lq.ResServerSettings) {
	_ = l.sendRPC("fetchServerSettings", ReqCommon, func(msg []byte) {
		data := &lq.ResServerSettings{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchAccountSettings(ch chan *lq.ResAccountSettings) {
	_ = l.sendRPC("fetchAccountSettings", ReqCommon, func(msg []byte) {
		data := &lq.ResAccountSettings{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UpdateAccountSettings(in *lq.ReqUpdateAccountSettings, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("updateAccountSettings", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchModNicknameTime(ch chan *lq.ResModNicknameTime) {
	_ = l.sendRPC("fetchModNicknameTime", ReqCommon, func(msg []byte) {
		data := &lq.ResModNicknameTime{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateWechatNativeOrder(in *lq.ReqCreateWechatNativeOrder, ch chan *lq.ResCreateWechatNativeOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createWechatNativeOrder", req, func(msg []byte) {
		data := &lq.ResCreateWechatNativeOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateWechatAppOrder(in *lq.ReqCreateWechatAppOrder, ch chan *lq.ResCreateWechatAppOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createWechatAppOrder", req, func(msg []byte) {
		data := &lq.ResCreateWechatAppOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateAlipayOrder(in *lq.ReqCreateAlipayOrder, ch chan *lq.ResCreateAlipayOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createAlipayOrder", req, func(msg []byte) {
		data := &lq.ResCreateAlipayOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateAlipayScanOrder(in *lq.ReqCreateAlipayScanOrder, ch chan *lq.ResCreateAlipayScanOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createAlipayScanOrder", req, func(msg []byte) {
		data := &lq.ResCreateAlipayScanOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateAlipayAppOrder(in *lq.ReqCreateAlipayAppOrder, ch chan *lq.ResCreateAlipayAppOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createAlipayAppOrder", req, func(msg []byte) {
		data := &lq.ResCreateAlipayAppOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateJPCreditCardOrder(in *lq.ReqCreateJPCreditCardOrder, ch chan *lq.ResCreateJPCreditCardOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createJPCreditCardOrder", req, func(msg []byte) {
		data := &lq.ResCreateJPCreditCardOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateJPPaypalOrder(in *lq.ReqCreateJPPaypalOrder, ch chan *lq.ResCreateJPPaypalOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createJPPaypalOrder", req, func(msg []byte) {
		data := &lq.ResCreateJPPaypalOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateJPAuOrder(in *lq.ReqCreateJPAuOrder, ch chan *lq.ResCreateJPAuOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createJPAuOrder", req, func(msg []byte) {
		data := &lq.ResCreateJPAuOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateJPDocomoOrder(in *lq.ReqCreateJPDocomoOrder, ch chan *lq.ResCreateJPDocomoOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createJPDocomoOrder", req, func(msg []byte) {
		data := &lq.ResCreateJPDocomoOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateJPWebMoneyOrder(in *lq.ReqCreateJPWebMoneyOrder, ch chan *lq.ResCreateJPWebMoneyOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createJPWebMoneyOrder", req, func(msg []byte) {
		data := &lq.ResCreateJPWebMoneyOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateJPSoftbankOrder(in *lq.ReqCreateJPSoftbankOrder, ch chan *lq.ResCreateJPSoftbankOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createJPSoftbankOrder", req, func(msg []byte) {
		data := &lq.ResCreateJPSoftbankOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateENPaypalOrder(in *lq.ReqCreateENPaypalOrder, ch chan *lq.ResCreateENPaypalOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createENPaypalOrder", req, func(msg []byte) {
		data := &lq.ResCreateENPaypalOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateENMasterCardOrder(in *lq.ReqCreateENMasterCardOrder, ch chan *lq.ResCreateENMasterCardOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createENMasterCardOrder", req, func(msg []byte) {
		data := &lq.ResCreateENMasterCardOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateENVisaOrder(in *lq.ReqCreateENVisaOrder, ch chan *lq.ResCreateENVisaOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createENVisaOrder", req, func(msg []byte) {
		data := &lq.ResCreateENVisaOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateENJCBOrder(in *lq.ReqCreateENJCBOrder, ch chan *lq.ResCreateENJCBOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createENJCBOrder", req, func(msg []byte) {
		data := &lq.ResCreateENJCBOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CreateENAlipayOrder(in *lq.ReqCreateENAlipayOrder, ch chan *lq.ResCreateENAlipayOrder) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("createENAlipayOrder", req, func(msg []byte) {
		data := &lq.ResCreateENAlipayOrder{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchMisc(ch chan *lq.ResMisc) {
	_ = l.sendRPC("fetchMisc", ReqCommon, func(msg []byte) {
		data := &lq.ResMisc{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ModifySignature(in *lq.ReqModifySignature, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("modifySignature", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchIDCardInfo(ch chan *lq.ResIDCardInfo) {
	_ = l.sendRPC("fetchIDCardInfo", ReqCommon, func(msg []byte) {
		data := &lq.ResIDCardInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UpdateIDCardInfo(in *lq.ReqUpdateIDCardInfo, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("updateIDCardInfo", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchVipReward(ch chan *lq.ResVipReward) {
	_ = l.sendRPC("fetchVipReward", ReqCommon, func(msg []byte) {
		data := &lq.ResVipReward{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) GainVipReward(in *lq.ReqGainVipReward, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("gainVipReward", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCustomizedContestList(in *lq.ReqFetchCustomizedContestList, ch chan *lq.ResFetchCustomizedContestList) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCustomizedContestList", req, func(msg []byte) {
		data := &lq.ResFetchCustomizedContestList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCustomizedContestExtendInfo(in *lq.ReqFetchCustomizedContestExtendInfo, ch chan *lq.ResFetchCustomizedContestExtendInfo) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCustomizedContestExtendInfo", req, func(msg []byte) {
		data := &lq.ResFetchCustomizedContestExtendInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) EnterCustomizedContest(in *lq.ReqEnterCustomizedContest, ch chan *lq.ResEnterCustomizedContest) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("enterCustomizedContest", req, func(msg []byte) {
		data := &lq.ResEnterCustomizedContest{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) LeaveCustomizedContest(ch chan *lq.ResCommon) {
	_ = l.sendRPC("leaveCustomizedContest", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCustomizedContestOnlineInfo(in *lq.ReqFetchCustomizedContestOnlineInfo, ch chan *lq.ResFetchCustomizedContestOnlineInfo) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCustomizedContestOnlineInfo", req, func(msg []byte) {
		data := &lq.ResFetchCustomizedContestOnlineInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCustomizedContestByContestId(in *lq.ReqFetchCustomizedContestByContestId, ch chan *lq.ResFetchCustomizedContestByContestId) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCustomizedContestByContestId", req, func(msg []byte) {
		data := &lq.ResFetchCustomizedContestByContestId{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) StartCustomizedContest(in *lq.ReqStartCustomizedContest, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("startCustomizedContest", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) StopCustomizedContest(ch chan *lq.ResCommon) {
	_ = l.sendRPC("stopCustomizedContest", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) JoinCustomizedContestChatRoom(in *lq.ReqJoinCustomizedContestChatRoom, ch chan *lq.ResJoinCustomizedContestChatRoom) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("joinCustomizedContestChatRoom", req, func(msg []byte) {
		data := &lq.ResJoinCustomizedContestChatRoom{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) LeaveCustomizedContestChatRoom(ch chan *lq.ResCommon) {
	_ = l.sendRPC("leaveCustomizedContestChatRoom", ReqCommon, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) SayChatMessage(in *lq.ReqSayChatMessage, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("sayChatMessage", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCustomizedContestGameRecords(in *lq.ReqFetchCustomizedContestGameRecords, ch chan *lq.ResFetchCustomizedContestGameRecords) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCustomizedContestGameRecords", req, func(msg []byte) {
		data := &lq.ResFetchCustomizedContestGameRecords{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchCustomizedContestGameLiveList(in *lq.ReqFetchCustomizedContestGameLiveList, ch chan *lq.ResFetchCustomizedContestGameLiveList) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchCustomizedContestGameLiveList", req, func(msg []byte) {
		data := &lq.ResFetchCustomizedContestGameLiveList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FollowCustomizedContest(in *lq.ReqTargetCustomizedContest, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("followCustomizedContest", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) UnfollowCustomizedContest(in *lq.ReqTargetCustomizedContest, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("unfollowCustomizedContest", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchActivityList(ch chan *lq.ResActivityList) {
	_ = l.sendRPC("fetchActivityList", ReqCommon, func(msg []byte) {
		data := &lq.ResActivityList{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchAccountActivityData(ch chan *lq.ResAccountActivityData) {
	_ = l.sendRPC("fetchAccountActivityData", ReqCommon, func(msg []byte) {
		data := &lq.ResAccountActivityData{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) ExchangeActivityItem(in *lq.ReqExchangeActivityItem, ch chan *lq.ResExchangeActivityItem) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("exchangeActivityItem", req, func(msg []byte) {
		data := &lq.ResExchangeActivityItem{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CompleteActivityTask(in *lq.ReqCompleteActivityTask, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("completeActivityTask", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) CompleteActivityFlipTask(in *lq.ReqCompleteActivityTask, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("completeActivityFlipTask", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) RecieveActivityFlipTask(in *lq.ReqRecieveActivityFlipTask, ch chan *lq.ResRecieveActivityFlipTask) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("recieveActivityFlipTask", req, func(msg []byte) {
		data := &lq.ResRecieveActivityFlipTask{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchActivityFlipInfo(in *lq.ReqFetchActivityFlipInfo, ch chan *lq.ResFetchActivityFlipInfo) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchActivityFlipInfo", req, func(msg []byte) {
		data := &lq.ResFetchActivityFlipInfo{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) GainAccumulatedPointActivityReward(in *lq.ReqGainAccumulatedPointActivityReward, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("gainAccumulatedPointActivityReward", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) FetchRankPointLeaderboard(in *lq.ReqFetchRankPointLeaderboard, ch chan *lq.ResFetchRankPointLeaderboard) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("fetchRankPointLeaderboard", req, func(msg []byte) {
		data := &lq.ResFetchRankPointLeaderboard{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}

func (l *LobbyClient) GainRankPointReward(in *lq.ReqGainRankPointReward, ch chan *lq.ResCommon) {
	req, _ := proto.Marshal(in)
	_ = l.sendRPC("gainRankPointReward", req, func(msg []byte) {
		data := &lq.ResCommon{}
		_ = proto.Unmarshal(msg, data)
		ch <- data
		close(ch)
	})
}
