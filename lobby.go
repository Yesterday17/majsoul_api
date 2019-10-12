package majsoul_api

import (
	"github.com/Yesterday17/majsoul_api/lq"
	"github.com/gogo/protobuf/proto"
)

type lobbyClient struct {
	socketClient
}

func NewLobbyClient(url string) (*lobbyClient, error) {
	client := &lobbyClient{}
	return client, client.init("Lobby", url)
}

func (l *lobbyClient) FetchConnectionInfo() chan lq.ResConnectionInfo {
	future := make(chan lq.ResConnectionInfo)
	_ = l.sendRPC("fetchConnectionInfo", reqCommon, func(msg []byte) {
		data := lq.ResConnectionInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) Signup(in *lq.ReqSignupAccount) chan lq.ResSignupAccount {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResSignupAccount)
	_ = l.sendRPC("signup", req, func(msg []byte) {
		data := lq.ResSignupAccount{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) Login(in *lq.ReqLogin) chan lq.ResLogin {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResLogin)
	_ = l.sendRPC("login", req, func(msg []byte) {
		data := lq.ResLogin{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) EmailLogin(in *lq.ReqEmailLogin) chan lq.ResLogin {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResLogin)
	_ = l.sendRPC("emailLogin", req, func(msg []byte) {
		data := lq.ResLogin{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) Oauth2Auth(in *lq.ReqOauth2Auth) chan lq.ResOauth2Auth {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResOauth2Auth)
	_ = l.sendRPC("oauth2Auth", req, func(msg []byte) {
		data := lq.ResOauth2Auth{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) Oauth2Check(in *lq.ReqOauth2Check) chan lq.ResOauth2Check {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResOauth2Check)
	_ = l.sendRPC("oauth2Check", req, func(msg []byte) {
		data := lq.ResOauth2Check{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) Oauth2Signup(in *lq.ReqOauth2Signup) chan lq.ResOauth2Signup {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResOauth2Signup)
	_ = l.sendRPC("oauth2Signup", req, func(msg []byte) {
		data := lq.ResOauth2Signup{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) Oauth2Login(in *lq.ReqOauth2Login) chan lq.ResLogin {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResLogin)
	_ = l.sendRPC("oauth2Login", req, func(msg []byte) {
		data := lq.ResLogin{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreatePhoneVerifyCode(in *lq.ReqCreatePhoneVerifyCode) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("createPhoneVerifyCode", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateEmailVerifyCode(in *lq.ReqCreateEmailVerifyCode) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("createEmailVerifyCode", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) VerfifyCodeForSecure(in *lq.ReqVerifyCodeForSecure) chan lq.ResVerfiyCodeForSecure {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResVerfiyCodeForSecure)
	_ = l.sendRPC("verfifyCodeForSecure", req, func(msg []byte) {
		data := lq.ResVerfiyCodeForSecure{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) BindPhoneNumber(in *lq.ReqBindPhoneNumber) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("bindPhoneNumber", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) BindEmail(in *lq.ReqBindEmail) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("bindEmail", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ModifyPassword(in *lq.ReqModifyPassword) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("modifyPassword", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) BindAccount(in *lq.ReqBindAccount) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("bindAccount", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) Logout(in *lq.ReqLogout) chan lq.ResLogout {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResLogout)
	_ = l.sendRPC("logout", req, func(msg []byte) {
		data := lq.ResLogout{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) Heartbeat(in *lq.ReqHeatBeat) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("heatbeat", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) LoginBeat(in *lq.ReqLoginBeat) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("loginBeat", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateNickname(in *lq.ReqCreateNickname) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("createNickname", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ModifyNickname(in *lq.ReqModifyNickname) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("modifyNickname", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ModifyBirthday(in *lq.ReqModifyBirthday) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("modifyBirthday", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchRoom() chan lq.ResSelfRoom {
	future := make(chan lq.ResSelfRoom)
	_ = l.sendRPC("fetchRoom", reqCommon, func(msg []byte) {
		data := lq.ResSelfRoom{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateRoom(in *lq.ReqCreateRoom) chan lq.ResCreateRoom {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateRoom)
	_ = l.sendRPC("createRoom", req, func(msg []byte) {
		data := lq.ResCreateRoom{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) JoinRoom(in *lq.ReqJoinRoom) chan lq.ResJoinRoom {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResJoinRoom)
	_ = l.sendRPC("joinRoom", req, func(msg []byte) {
		data := lq.ResJoinRoom{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) LeaveRoom() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("leaveRoom", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ReadyPlay(in *lq.ReqRoomReady) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("readyPlay", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) StartRoom(in *lq.ReqRoomStart) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("startRoom", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) KickPlayer(in *lq.ReqRoomKick) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("kickPlayer", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ModifyRoom(in *lq.ReqModifyRoom) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("modifyRoom", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) MatchGame(in *lq.ReqJoinMatchQueue) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("matchGame", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CancelMatch(in *lq.ReqCancelMatchQueue) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("cancelMatch", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchAccountInfo(in *lq.ReqAccountInfo) chan lq.ResAccountInfo {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResAccountInfo)
	_ = l.sendRPC("fetchAccountInfo", req, func(msg []byte) {
		data := lq.ResAccountInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ChangeAvatar(in *lq.ReqChangeAvatar) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("changeAvatar", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchAccountStatisticInfo(in *lq.ReqAccountStatisticInfo) chan lq.ResAccountStatisticInfo {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResAccountStatisticInfo)
	_ = l.sendRPC("fetchAccountStatisticInfo", req, func(msg []byte) {
		data := lq.ResAccountStatisticInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchAccountCharacterInfo() chan lq.ResAccountCharacterInfo {
	future := make(chan lq.ResAccountCharacterInfo)
	_ = l.sendRPC("fetchAccountCharacterInfo", reqCommon, func(msg []byte) {
		data := lq.ResAccountCharacterInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ShopPurchase(in *lq.ReqShopPurchase) chan lq.ResShopPurchase {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResShopPurchase)
	_ = l.sendRPC("shopPurchase", req, func(msg []byte) {
		data := lq.ResShopPurchase{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchGameRecord(in *lq.ReqGameRecord) chan lq.ResGameRecord {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResGameRecord)
	_ = l.sendRPC("fetchGameRecord", req, func(msg []byte) {
		data := lq.ResGameRecord{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchGameRecordList(in *lq.ReqGameRecordList) chan lq.ResGameRecordList {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResGameRecordList)
	_ = l.sendRPC("fetchGameRecordList", req, func(msg []byte) {
		data := lq.ResGameRecordList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCollectedGameRecordList() chan lq.ResCollectedGameRecordList {
	future := make(chan lq.ResCollectedGameRecordList)
	_ = l.sendRPC("fetchCollectedGameRecordList", reqCommon, func(msg []byte) {
		data := lq.ResCollectedGameRecordList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchGameRecordsDetail(in *lq.ReqGameRecordsDetail) chan lq.ResGameRecordsDetail {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResGameRecordsDetail)
	_ = l.sendRPC("fetchGameRecordsDetail", req, func(msg []byte) {
		data := lq.ResGameRecordsDetail{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) AddCollectedGameRecord(in *lq.ReqAddCollectedGameRecord) chan lq.ResAddCollectedGameRecord {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResAddCollectedGameRecord)
	_ = l.sendRPC("addCollectedGameRecord", req, func(msg []byte) {
		data := lq.ResAddCollectedGameRecord{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) RemoveCollectedGameRecord(in *lq.ReqRemoveCollectedGameRecord) chan lq.ResRemoveCollectedGameRecord {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResRemoveCollectedGameRecord)
	_ = l.sendRPC("removeCollectedGameRecord", req, func(msg []byte) {
		data := lq.ResRemoveCollectedGameRecord{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ChangeCollectedGameRecordRemarks(in *lq.ReqChangeCollectedGameRecordRemarks) chan lq.ResChangeCollectedGameRecordRemarks {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResChangeCollectedGameRecordRemarks)
	_ = l.sendRPC("changeCollectedGameRecordRemarks", req, func(msg []byte) {
		data := lq.ResChangeCollectedGameRecordRemarks{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchLevelLeaderboard(in *lq.ReqLevelLeaderboard) chan lq.ResLevelLeaderboard {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResLevelLeaderboard)
	_ = l.sendRPC("fetchLevelLeaderboard", req, func(msg []byte) {
		data := lq.ResLevelLeaderboard{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchMultiAccountBrief(in *lq.ReqMultiAccountId) chan lq.ResMultiAccountBrief {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResMultiAccountBrief)
	_ = l.sendRPC("fetchMultiAccountBrief", req, func(msg []byte) {
		data := lq.ResMultiAccountBrief{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchFriendList() chan lq.ResFriendList {
	future := make(chan lq.ResFriendList)
	_ = l.sendRPC("fetchFriendList", reqCommon, func(msg []byte) {
		data := lq.ResFriendList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchFriendApplyList() chan lq.ResFriendApplyList {
	future := make(chan lq.ResFriendApplyList)
	_ = l.sendRPC("fetchFriendApplyList", reqCommon, func(msg []byte) {
		data := lq.ResFriendApplyList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ApplyFriend(in *lq.ReqApplyFriend) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("applyFriend", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) HandleFriendApply(in *lq.ReqHandleFriendApply) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("handleFriendApply", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) RemoveFriend(in *lq.ReqRemoveFriend) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("removeFriend", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) SearchAccountById(in *lq.ReqSearchAccountById) chan lq.ResSearchAccountById {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResSearchAccountById)
	_ = l.sendRPC("searchAccountById", req, func(msg []byte) {
		data := lq.ResSearchAccountById{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) SearchAccountByPattern(in *lq.ReqSearchAccountByPattern) chan lq.ResSearchAccountByPattern {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResSearchAccountByPattern)
	_ = l.sendRPC("searchAccountByPattern", req, func(msg []byte) {
		data := lq.ResSearchAccountByPattern{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchAccountState(in *lq.ReqAccountList) chan lq.ResAccountStates {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResAccountStates)
	_ = l.sendRPC("fetchAccountState", req, func(msg []byte) {
		data := lq.ResAccountStates{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchBagInfo() chan lq.ResBagInfo {
	future := make(chan lq.ResBagInfo)
	_ = l.sendRPC("fetchBagInfo", reqCommon, func(msg []byte) {
		data := lq.ResBagInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UseBagItem(in *lq.ReqUseBagItem) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("useBagItem", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) OpenManualItem(in *lq.ReqOpenManualItem) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("openManualItem", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) OpenRandomRewardItem(in *lq.ReqOpenRandomRewardItem) chan lq.ResOpenRandomRewardItem {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResOpenRandomRewardItem)
	_ = l.sendRPC("openRandomRewardItem", req, func(msg []byte) {
		data := lq.ResOpenRandomRewardItem{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ComposeShard(in *lq.ReqComposeShard) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("composeShard", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchAnnouncement() chan lq.ResAnnouncement {
	future := make(chan lq.ResAnnouncement)
	_ = l.sendRPC("fetchAnnouncement", reqCommon, func(msg []byte) {
		data := lq.ResAnnouncement{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ReadAnnouncement(in *lq.ReqReadAnnouncement) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("readAnnouncement", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchMailInfo() chan lq.ResMailInfo {
	future := make(chan lq.ResMailInfo)
	_ = l.sendRPC("fetchMailInfo", reqCommon, func(msg []byte) {
		data := lq.ResMailInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ReadMail(in *lq.ReqReadMail) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("readMail", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) DeleteMail(in *lq.ReqDeleteMail) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("deleteMail", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) TakeAttachmentFromMail(in *lq.ReqTakeAttachment) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("takeAttachmentFromMail", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchAchievement() chan lq.ResAchievement {
	future := make(chan lq.ResAchievement)
	_ = l.sendRPC("fetchAchievement", reqCommon, func(msg []byte) {
		data := lq.ResAchievement{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) BuyShiLian(in *lq.ReqBuyShiLian) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("buyShiLian", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) MatchShiLian() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("matchShiLian", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) GoNextShiLian() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("goNextShiLian", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UpdateClientValue(in *lq.ReqUpdateClientValue) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("updateClientValue", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchClientValue() chan lq.ResClientValue {
	future := make(chan lq.ResClientValue)
	_ = l.sendRPC("fetchClientValue", reqCommon, func(msg []byte) {
		data := lq.ResClientValue{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ClientMessage(in *lq.ReqClientMessage) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("clientMessage", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCurrentMatchInfo(in *lq.ReqCurrentMatchInfo) chan lq.ResCurrentMatchInfo {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCurrentMatchInfo)
	_ = l.sendRPC("fetchCurrentMatchInfo", req, func(msg []byte) {
		data := lq.ResCurrentMatchInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchReviveCoinInfo() chan lq.ResReviveCoinInfo {
	future := make(chan lq.ResReviveCoinInfo)
	_ = l.sendRPC("fetchReviveCoinInfo", reqCommon, func(msg []byte) {
		data := lq.ResReviveCoinInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) GainReviveCoin() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("gainReviveCoin", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchDailyTask() chan lq.ResDailyTask {
	future := make(chan lq.ResDailyTask)
	_ = l.sendRPC("fetchDailyTask", reqCommon, func(msg []byte) {
		data := lq.ResDailyTask{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) RefreshDailyTask(in *lq.ReqRefreshDailyTask) chan lq.ResRefreshDailyTask {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResRefreshDailyTask)
	_ = l.sendRPC("refreshDailyTask", req, func(msg []byte) {
		data := lq.ResRefreshDailyTask{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UseGiftCode(in *lq.ReqUseGiftCode) chan lq.ResUseGiftCode {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResUseGiftCode)
	_ = l.sendRPC("useGiftCode", req, func(msg []byte) {
		data := lq.ResUseGiftCode{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchTitleList() chan lq.ResTitleList {
	future := make(chan lq.ResTitleList)
	_ = l.sendRPC("fetchTitleList", reqCommon, func(msg []byte) {
		data := lq.ResTitleList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UseTitle(in *lq.ReqUseTitle) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("useTitle", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) SendClientMessage(in *lq.ReqSendClientMessage) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("sendClientMessage", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchGameLiveInfo(in *lq.ReqGameLiveInfo) chan lq.ResGameLiveInfo {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResGameLiveInfo)
	_ = l.sendRPC("fetchGameLiveInfo", req, func(msg []byte) {
		data := lq.ResGameLiveInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchGameLiveLeftSegment(in *lq.ReqGameLiveLeftSegment) chan lq.ResGameLiveLeftSegment {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResGameLiveLeftSegment)
	_ = l.sendRPC("fetchGameLiveLeftSegment", req, func(msg []byte) {
		data := lq.ResGameLiveLeftSegment{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchGameLiveList(in *lq.ReqGameLiveList) chan lq.ResGameLiveList {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResGameLiveList)
	_ = l.sendRPC("fetchGameLiveList", req, func(msg []byte) {
		data := lq.ResGameLiveList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCommentSetting() chan lq.ResCommentSetting {
	future := make(chan lq.ResCommentSetting)
	_ = l.sendRPC("fetchCommentSetting", reqCommon, func(msg []byte) {
		data := lq.ResCommentSetting{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UpdateCommentSetting(in *lq.ReqUpdateCommentSetting) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("updateCommentSetting", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCommentList(in *lq.ReqFetchCommentList) chan lq.ResFetchCommentList {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchCommentList)
	_ = l.sendRPC("fetchCommentList", req, func(msg []byte) {
		data := lq.ResFetchCommentList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCommentContent(in *lq.ReqFetchCommentContent) chan lq.ResFetchCommentContent {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchCommentContent)
	_ = l.sendRPC("fetchCommentContent", req, func(msg []byte) {
		data := lq.ResFetchCommentContent{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) LeaveComment(in *lq.ReqLeaveComment) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("leaveComment", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) DeleteComment(in *lq.ReqDeleteComment) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("deleteComment", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UpdateReadComment(in *lq.ReqUpdateReadComment) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("updateReadComment", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchRollingNotice() chan lq.ReqRollingNotice {
	future := make(chan lq.ReqRollingNotice)
	_ = l.sendRPC("fetchRollingNotice", reqCommon, func(msg []byte) {
		data := lq.ReqRollingNotice{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchServerTime() chan lq.ResServerTime {
	future := make(chan lq.ResServerTime)
	_ = l.sendRPC("fetchServerTime", reqCommon, func(msg []byte) {
		data := lq.ResServerTime{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchPlatformProducts(in *lq.ReqPlatformBillingProducts) chan lq.ResPlatformBillingProducts {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResPlatformBillingProducts)
	_ = l.sendRPC("fetchPlatformProducts", req, func(msg []byte) {
		data := lq.ResPlatformBillingProducts{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateBillingOrder(in *lq.ReqCreateBillingOrder) chan lq.ResCreateBillingOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateBillingOrder)
	_ = l.sendRPC("createBillingOrder", req, func(msg []byte) {
		data := lq.ResCreateBillingOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) SolveGooglePlayOrder(in *lq.ReqSolveGooglePlayOrder) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("solveGooglePlayOrder", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CancelGooglePlayOrder(in *lq.ReqCancelGooglePlayOrder) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("cancelGooglePlayOrder", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) OpenChest(in *lq.ReqOpenChest) chan lq.ResOpenChest {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResOpenChest)
	_ = l.sendRPC("openChest", req, func(msg []byte) {
		data := lq.ResOpenChest{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) BuyFromChestShop(in *lq.ReqBuyFromChestShop) chan lq.ResBuyFromChestShop {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResBuyFromChestShop)
	_ = l.sendRPC("buyFromChestShop", req, func(msg []byte) {
		data := lq.ResBuyFromChestShop{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchDailySignInInfo() chan lq.ResDailySignInInfo {
	future := make(chan lq.ResDailySignInInfo)
	_ = l.sendRPC("fetchDailySignInInfo", reqCommon, func(msg []byte) {
		data := lq.ResDailySignInInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) DoDailySignIn() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("doDailySignIn", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) DoActivitySignIn(in *lq.ReqDoActivitySignIn) chan lq.ResDoActivitySignIn {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResDoActivitySignIn)
	_ = l.sendRPC("doActivitySignIn", req, func(msg []byte) {
		data := lq.ResDoActivitySignIn{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCharacterInfo() chan lq.ResCharacterInfo {
	future := make(chan lq.ResCharacterInfo)
	_ = l.sendRPC("fetchCharacterInfo", reqCommon, func(msg []byte) {
		data := lq.ResCharacterInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ChangeMainCharacter(in *lq.ReqChangeMainCharacter) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("changeMainCharacter", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ChangeCharacterSkin(in *lq.ReqChangeCharacterSkin) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("changeCharacterSkin", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ChangeCharacterView(in *lq.ReqChangeCharacterView) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("changeCharacterView", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) SendGiftToCharacter(in *lq.ReqSendGiftToCharacter) chan lq.ResSendGiftToCharacter {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResSendGiftToCharacter)
	_ = l.sendRPC("sendGiftToCharacter", req, func(msg []byte) {
		data := lq.ResSendGiftToCharacter{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) SellItem(in *lq.ReqSellItem) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("sellItem", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCommonView() chan lq.ResCommonView {
	future := make(chan lq.ResCommonView)
	_ = l.sendRPC("fetchCommonView", reqCommon, func(msg []byte) {
		data := lq.ResCommonView{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ChangeCommonView(in *lq.ReqChangeCommonView) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("changeCommonView", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UpgradeCharacter(in *lq.ReqUpgradeCharacter) chan lq.ResUpgradeCharacter {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResUpgradeCharacter)
	_ = l.sendRPC("upgradeCharacter", req, func(msg []byte) {
		data := lq.ResUpgradeCharacter{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) GameMasterCommand(in *lq.ReqGMCommand) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("gameMasterCommand", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchShopInfo() chan lq.ResShopInfo {
	future := make(chan lq.ResShopInfo)
	_ = l.sendRPC("fetchShopInfo", reqCommon, func(msg []byte) {
		data := lq.ResShopInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) BuyFromShop(in *lq.ReqBuyFromShop) chan lq.ResBuyFromShop {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResBuyFromShop)
	_ = l.sendRPC("buyFromShop", req, func(msg []byte) {
		data := lq.ResBuyFromShop{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) BuyFromZHP(in *lq.ReqBuyFromZHP) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("buyFromZHP", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) RefreshZHPShop() chan lq.ResRefreshZHPShop {
	future := make(chan lq.ResRefreshZHPShop)
	_ = l.sendRPC("refreshZHPShop", reqCommon, func(msg []byte) {
		data := lq.ResRefreshZHPShop{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchMonthTicketInfo() chan lq.ResMonthTicketInfo {
	future := make(chan lq.ResMonthTicketInfo)
	_ = l.sendRPC("fetchMonthTicketInfo", reqCommon, func(msg []byte) {
		data := lq.ResMonthTicketInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) PayMonthTicket(in *lq.ReqPayMonthTicket) chan lq.ResPayMonthTicket {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResPayMonthTicket)
	_ = l.sendRPC("payMonthTicket", req, func(msg []byte) {
		data := lq.ResPayMonthTicket{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ExchangeCurrency(in *lq.ReqExchangeCurrency) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("exchangeCurrency", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ExchangeChestStone(in *lq.ReqExchangeCurrency) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("exchangeChestStone", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchServerSettings() chan lq.ResServerSettings {
	future := make(chan lq.ResServerSettings)
	_ = l.sendRPC("fetchServerSettings", reqCommon, func(msg []byte) {
		data := lq.ResServerSettings{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchAccountSettings() chan lq.ResAccountSettings {
	future := make(chan lq.ResAccountSettings)
	_ = l.sendRPC("fetchAccountSettings", reqCommon, func(msg []byte) {
		data := lq.ResAccountSettings{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UpdateAccountSettings(in *lq.ReqUpdateAccountSettings) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("updateAccountSettings", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchModNicknameTime() chan lq.ResModNicknameTime {
	future := make(chan lq.ResModNicknameTime)
	_ = l.sendRPC("fetchModNicknameTime", reqCommon, func(msg []byte) {
		data := lq.ResModNicknameTime{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateWechatNativeOrder(in *lq.ReqCreateWechatNativeOrder) chan lq.ResCreateWechatNativeOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateWechatNativeOrder)
	_ = l.sendRPC("createWechatNativeOrder", req, func(msg []byte) {
		data := lq.ResCreateWechatNativeOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateWechatAppOrder(in *lq.ReqCreateWechatAppOrder) chan lq.ResCreateWechatAppOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateWechatAppOrder)
	_ = l.sendRPC("createWechatAppOrder", req, func(msg []byte) {
		data := lq.ResCreateWechatAppOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateAlipayOrder(in *lq.ReqCreateAlipayOrder) chan lq.ResCreateAlipayOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateAlipayOrder)
	_ = l.sendRPC("createAlipayOrder", req, func(msg []byte) {
		data := lq.ResCreateAlipayOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateAlipayScanOrder(in *lq.ReqCreateAlipayScanOrder) chan lq.ResCreateAlipayScanOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateAlipayScanOrder)
	_ = l.sendRPC("createAlipayScanOrder", req, func(msg []byte) {
		data := lq.ResCreateAlipayScanOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateAlipayAppOrder(in *lq.ReqCreateAlipayAppOrder) chan lq.ResCreateAlipayAppOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateAlipayAppOrder)
	_ = l.sendRPC("createAlipayAppOrder", req, func(msg []byte) {
		data := lq.ResCreateAlipayAppOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateJPCreditCardOrder(in *lq.ReqCreateJPCreditCardOrder) chan lq.ResCreateJPCreditCardOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateJPCreditCardOrder)
	_ = l.sendRPC("createJPCreditCardOrder", req, func(msg []byte) {
		data := lq.ResCreateJPCreditCardOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateJPPaypalOrder(in *lq.ReqCreateJPPaypalOrder) chan lq.ResCreateJPPaypalOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateJPPaypalOrder)
	_ = l.sendRPC("createJPPaypalOrder", req, func(msg []byte) {
		data := lq.ResCreateJPPaypalOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateJPAuOrder(in *lq.ReqCreateJPAuOrder) chan lq.ResCreateJPAuOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateJPAuOrder)
	_ = l.sendRPC("createJPAuOrder", req, func(msg []byte) {
		data := lq.ResCreateJPAuOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateJPDocomoOrder(in *lq.ReqCreateJPDocomoOrder) chan lq.ResCreateJPDocomoOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateJPDocomoOrder)
	_ = l.sendRPC("createJPDocomoOrder", req, func(msg []byte) {
		data := lq.ResCreateJPDocomoOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateJPWebMoneyOrder(in *lq.ReqCreateJPWebMoneyOrder) chan lq.ResCreateJPWebMoneyOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateJPWebMoneyOrder)
	_ = l.sendRPC("createJPWebMoneyOrder", req, func(msg []byte) {
		data := lq.ResCreateJPWebMoneyOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateJPSoftbankOrder(in *lq.ReqCreateJPSoftbankOrder) chan lq.ResCreateJPSoftbankOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateJPSoftbankOrder)
	_ = l.sendRPC("createJPSoftbankOrder", req, func(msg []byte) {
		data := lq.ResCreateJPSoftbankOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateENPaypalOrder(in *lq.ReqCreateENPaypalOrder) chan lq.ResCreateENPaypalOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateENPaypalOrder)
	_ = l.sendRPC("createENPaypalOrder", req, func(msg []byte) {
		data := lq.ResCreateENPaypalOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateENMasterCardOrder(in *lq.ReqCreateENMasterCardOrder) chan lq.ResCreateENMasterCardOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateENMasterCardOrder)
	_ = l.sendRPC("createENMasterCardOrder", req, func(msg []byte) {
		data := lq.ResCreateENMasterCardOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateENVisaOrder(in *lq.ReqCreateENVisaOrder) chan lq.ResCreateENVisaOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateENVisaOrder)
	_ = l.sendRPC("createENVisaOrder", req, func(msg []byte) {
		data := lq.ResCreateENVisaOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateENJCBOrder(in *lq.ReqCreateENJCBOrder) chan lq.ResCreateENJCBOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateENJCBOrder)
	_ = l.sendRPC("createENJCBOrder", req, func(msg []byte) {
		data := lq.ResCreateENJCBOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CreateENAlipayOrder(in *lq.ReqCreateENAlipayOrder) chan lq.ResCreateENAlipayOrder {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCreateENAlipayOrder)
	_ = l.sendRPC("createENAlipayOrder", req, func(msg []byte) {
		data := lq.ResCreateENAlipayOrder{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchMisc() chan lq.ResMisc {
	future := make(chan lq.ResMisc)
	_ = l.sendRPC("fetchMisc", reqCommon, func(msg []byte) {
		data := lq.ResMisc{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ModifySignature(in *lq.ReqModifySignature) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("modifySignature", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchIDCardInfo() chan lq.ResIDCardInfo {
	future := make(chan lq.ResIDCardInfo)
	_ = l.sendRPC("fetchIDCardInfo", reqCommon, func(msg []byte) {
		data := lq.ResIDCardInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UpdateIDCardInfo(in *lq.ReqUpdateIDCardInfo) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("updateIDCardInfo", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchVipReward() chan lq.ResVipReward {
	future := make(chan lq.ResVipReward)
	_ = l.sendRPC("fetchVipReward", reqCommon, func(msg []byte) {
		data := lq.ResVipReward{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) GainVipReward(in *lq.ReqGainVipReward) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("gainVipReward", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCustomizedContestList(in *lq.ReqFetchCustomizedContestList) chan lq.ResFetchCustomizedContestList {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchCustomizedContestList)
	_ = l.sendRPC("fetchCustomizedContestList", req, func(msg []byte) {
		data := lq.ResFetchCustomizedContestList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCustomizedContestExtendInfo(in *lq.ReqFetchCustomizedContestExtendInfo) chan lq.ResFetchCustomizedContestExtendInfo {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchCustomizedContestExtendInfo)
	_ = l.sendRPC("fetchCustomizedContestExtendInfo", req, func(msg []byte) {
		data := lq.ResFetchCustomizedContestExtendInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) EnterCustomizedContest(in *lq.ReqEnterCustomizedContest) chan lq.ResEnterCustomizedContest {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResEnterCustomizedContest)
	_ = l.sendRPC("enterCustomizedContest", req, func(msg []byte) {
		data := lq.ResEnterCustomizedContest{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) LeaveCustomizedContest() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("leaveCustomizedContest", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCustomizedContestOnlineInfo(in *lq.ReqFetchCustomizedContestOnlineInfo) chan lq.ResFetchCustomizedContestOnlineInfo {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchCustomizedContestOnlineInfo)
	_ = l.sendRPC("fetchCustomizedContestOnlineInfo", req, func(msg []byte) {
		data := lq.ResFetchCustomizedContestOnlineInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCustomizedContestByContestId(in *lq.ReqFetchCustomizedContestByContestId) chan lq.ResFetchCustomizedContestByContestId {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchCustomizedContestByContestId)
	_ = l.sendRPC("fetchCustomizedContestByContestId", req, func(msg []byte) {
		data := lq.ResFetchCustomizedContestByContestId{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) StartCustomizedContest(in *lq.ReqStartCustomizedContest) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("startCustomizedContest", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) StopCustomizedContest() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("stopCustomizedContest", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) JoinCustomizedContestChatRoom(in *lq.ReqJoinCustomizedContestChatRoom) chan lq.ResJoinCustomizedContestChatRoom {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResJoinCustomizedContestChatRoom)
	_ = l.sendRPC("joinCustomizedContestChatRoom", req, func(msg []byte) {
		data := lq.ResJoinCustomizedContestChatRoom{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) LeaveCustomizedContestChatRoom() chan lq.ResCommon {
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("leaveCustomizedContestChatRoom", reqCommon, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) SayChatMessage(in *lq.ReqSayChatMessage) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("sayChatMessage", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCustomizedContestGameRecords(in *lq.ReqFetchCustomizedContestGameRecords) chan lq.ResFetchCustomizedContestGameRecords {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchCustomizedContestGameRecords)
	_ = l.sendRPC("fetchCustomizedContestGameRecords", req, func(msg []byte) {
		data := lq.ResFetchCustomizedContestGameRecords{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchCustomizedContestGameLiveList(in *lq.ReqFetchCustomizedContestGameLiveList) chan lq.ResFetchCustomizedContestGameLiveList {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchCustomizedContestGameLiveList)
	_ = l.sendRPC("fetchCustomizedContestGameLiveList", req, func(msg []byte) {
		data := lq.ResFetchCustomizedContestGameLiveList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FollowCustomizedContest(in *lq.ReqTargetCustomizedContest) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("followCustomizedContest", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) UnfollowCustomizedContest(in *lq.ReqTargetCustomizedContest) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("unfollowCustomizedContest", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchActivityList() chan lq.ResActivityList {
	future := make(chan lq.ResActivityList)
	_ = l.sendRPC("fetchActivityList", reqCommon, func(msg []byte) {
		data := lq.ResActivityList{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchAccountActivityData() chan lq.ResAccountActivityData {
	future := make(chan lq.ResAccountActivityData)
	_ = l.sendRPC("fetchAccountActivityData", reqCommon, func(msg []byte) {
		data := lq.ResAccountActivityData{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) ExchangeActivityItem(in *lq.ReqExchangeActivityItem) chan lq.ResExchangeActivityItem {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResExchangeActivityItem)
	_ = l.sendRPC("exchangeActivityItem", req, func(msg []byte) {
		data := lq.ResExchangeActivityItem{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CompleteActivityTask(in *lq.ReqCompleteActivityTask) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("completeActivityTask", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) CompleteActivityFlipTask(in *lq.ReqCompleteActivityTask) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("completeActivityFlipTask", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) RecieveActivityFlipTask(in *lq.ReqRecieveActivityFlipTask) chan lq.ResRecieveActivityFlipTask {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResRecieveActivityFlipTask)
	_ = l.sendRPC("recieveActivityFlipTask", req, func(msg []byte) {
		data := lq.ResRecieveActivityFlipTask{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchActivityFlipInfo(in *lq.ReqFetchActivityFlipInfo) chan lq.ResFetchActivityFlipInfo {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchActivityFlipInfo)
	_ = l.sendRPC("fetchActivityFlipInfo", req, func(msg []byte) {
		data := lq.ResFetchActivityFlipInfo{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) GainAccumulatedPointActivityReward(in *lq.ReqGainAccumulatedPointActivityReward) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("gainAccumulatedPointActivityReward", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) FetchRankPointLeaderboard(in *lq.ReqFetchRankPointLeaderboard) chan lq.ResFetchRankPointLeaderboard {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResFetchRankPointLeaderboard)
	_ = l.sendRPC("fetchRankPointLeaderboard", req, func(msg []byte) {
		data := lq.ResFetchRankPointLeaderboard{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}

func (l *lobbyClient) GainRankPointReward(in *lq.ReqGainRankPointReward) chan lq.ResCommon {
	req, _ := proto.Marshal(in)
	future := make(chan lq.ResCommon)
	_ = l.sendRPC("gainRankPointReward", req, func(msg []byte) {
		data := lq.ResCommon{}
		_ = proto.Unmarshal(msg, &data)
		future <- data
	})
	return future
}
