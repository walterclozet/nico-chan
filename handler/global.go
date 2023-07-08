package handler

import (
	"encoding/json"
	"fmt"
	"nico-chan/config"
	"nico-chan/encrypt"
	"nico-chan/model"
	"nico-chan/utils"
	"time"

	"xorm.io/xorm"
)

var (
	CdnUrl   string
	ErrorMsg = `{"code":20001,"message":""}`
	MainEng  *xorm.Engine
	UserEng  *xorm.Engine

	// LLAS
	sessionKey = "12345678123456781234567812345678"
)

func init() {
	if config.Conf.Cdn.CdnUrl != "" {
		CdnUrl = config.Conf.Cdn.CdnUrl
	}

	MainEng = config.MainEng
	UserEng = config.UserEng
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func IsSigned(unitId int) bool {
	exists, err := MainEng.Table("unit_sign_asset_m").Where("unit_id = ?", unitId).Exist()
	CheckErr(err)

	return exists
}

func SignResp(ep, body, key string) (resp string) {
	signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), config.MasterVersion, body)
	sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(key))
	// fmt.Println(sign)

	resp = fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
	return
}

func GetUserStatus() map[string]any {
	var r map[string]any
	if err := json.Unmarshal([]byte(utils.ReadAllText("assets/as/userStatus.json")), &r); err != nil {
		panic(err)
	}
	return r
}

func CommonUserStatus() model.AsUserStatus {
	return model.AsUserStatus{
		Name: model.Name{
			DotUnderText: "梦路2号机@果果厨",
		},
		Nickname: model.Nickname{
			DotUnderText: "梦路",
		},
		LastLoginAt: time.Now().Unix(),
		Rank:        122,
		Exp:         416263,
		Message: model.Message{
			DotUnderText: "B站主播：梦路_YumeMichi",
		},
		RecommendCardMasterID:                     200013001,
		MaxFriendNum:                              0,
		LivePointFullAt:                           1684158189,
		LivePointBroken:                           150,
		LivePointSubscriptionRecoveryDailyCount:   1,
		LivePointSubscriptionRecoveryDailyResetAt: 1656259200,
		ActivityPointCount:                        0,
		ActivityPointResetAt:                      1684123200,
		ActivityPointPaymentRecoveryDailyCount:    10,
		ActivityPointPaymentRecoveryDailyResetAt:  1683734400,
		GameMoney:                                 115665089,
		CardExp:                                   32316167,
		FreeSnsCoin:                               280,
		AppleSnsCoin:                              0,
		GoogleSnsCoin:                             0,
		Cash:                                      0,
		SubscriptionCoin:                          30,
		BirthDate:                                 nil,
		BirthMonth:                                10,
		BirthDay:                                  18,
		LatestLiveDeckID:                          1,
		MainLessonDeckID:                          1,
		FavoriteMemberID:                          1,
		LastLiveDifficultyID:                      12092302,
		LpMagnification:                           1,
		EmblemID:                                  13100608,
		DeviceToken:                               "",
		TutorialPhase:                             99,
		TutorialEndAt:                             1622217482,
		LoginDays:                                 446,
		NaviTapCount:                              2,
		NaviTapRecoverAt:                          1684771200,
		IsAutoMode:                                false,
		MaxScoreLiveDifficultyMasterID:            12037401,
		LiveMaxScore:                              40553270,
		MaxComboLiveDifficultyMasterID:            12037401,
		LiveMaxCombo:                              301,
		LessonResumeStatus:                        1,
		AccessoryBoxAdditional:                    90,
		TermsOfUseVersion:                         0,
		BootstrapSifidCheckAt:                     1683178908,
		GdprVersion:                               0,
		MemberGuildMemberMasterID:                 1,
		MemberGuildLastUpdatedAt:                  1659485328,
	}
}
