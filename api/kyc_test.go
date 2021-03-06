// nolint
package api

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	myconst "github.com/NpoolPlatform/kyc-management/pkg/message/const"
	testinit "github.com/NpoolPlatform/kyc-management/pkg/test-init"
	npoolcommon "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/kyc"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		log.Fatal(err)
	}
}

func TestKycAPI(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	kycInfo := &npool.KycInfo{
		UserID:             uuid.New().String(),
		AppID:              uuid.New().String(),
		CardType:           uuid.New().String(),
		CardID:             uuid.New().String()[:16],
		FrontCardImg:       "frontCardImage",
		BackCardImg:        "backCardImage",
		UserHandingCardImg: "userHandingCardImage",
	}

	createKycRequest := &npool.CreateKycRequest{
		Info: kycInfo,
	}

	ctx, cancel := context.WithTimeout(context.Background(), myconst.GrpcTimeout)
	defer cancel()

	createResp, err := kyc.Create(ctx, createKycRequest)
	if assert.Nil(t, err) {
		assert.NotNil(t, createResp.GetInfo().GetID())
		kycInfo.ID = createResp.GetInfo().GetID()
		assert.Equal(t, kycInfo.GetUserID(), createResp.GetInfo().GetUserID())
		assert.Equal(t, kycInfo.GetAppID(), createResp.GetInfo().GetAppID())
		assert.Equal(t, kycInfo.GetCardType(), createResp.GetInfo().GetCardType())
		assert.Equal(t, kycInfo.GetCardID(), createResp.GetInfo().GetCardID())
		assert.Equal(t, kycInfo.GetFrontCardImg(), createResp.GetInfo().GetFrontCardImg())
		assert.Equal(t, kycInfo.GetBackCardImg(), createResp.GetInfo().GetBackCardImg())
		assert.Equal(t, kycInfo.GetUserHandingCardImg(), createResp.GetInfo().GetUserHandingCardImg())
	}

	cli := resty.New()

	getKycByUserIDResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetKycByUserIDRequest{
			AppID:  kycInfo.GetAppID(),
			UserID: kycInfo.GetUserID(),
		}).Post("http://localhost:50120/v1/get/kyc/by/userid")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, getKycByUserIDResp.StatusCode())
		getKycByUserIDResponse := &npool.GetKycByUserIDResponse{}
		err := json.Unmarshal(getKycByUserIDResp.Body(), getKycByUserIDResponse)
		if assert.Nil(t, err) {
			assert.Equal(t, kycInfo.GetID(), getKycByUserIDResponse.GetInfo().GetID())
			assert.Equal(t, kycInfo.GetAppID(), getKycByUserIDResponse.GetInfo().GetAppID())
			assert.Equal(t, kycInfo.GetUserID(), getKycByUserIDResponse.GetInfo().GetUserID())
			assert.Equal(t, kycInfo.GetCardType(), getKycByUserIDResponse.GetInfo().GetCardType())
			assert.Equal(t, kycInfo.GetCardID(), getKycByUserIDResponse.GetInfo().GetCardID())
			assert.Equal(t, kycInfo.GetFrontCardImg(), getKycByUserIDResponse.GetInfo().GetFrontCardImg())
			assert.Equal(t, kycInfo.GetBackCardImg(), getKycByUserIDResponse.GetInfo().GetBackCardImg())
			assert.Equal(t, kycInfo.GetUserHandingCardImg(), getKycByUserIDResponse.GetInfo().GetUserHandingCardImg())
		}
	}

	getKycByAppIDResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetKycByAppIDRequest{
			TargetAppID: kycInfo.GetAppID(),
			PageInfo: &npoolcommon.PageInfo{
				Limit:  5,
				Offset: 0,
			},
		}).Post("http://localhost:50120/v1/get/kyc/by/appid")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, getKycByAppIDResp.StatusCode())
		getKycByAppIDResponse := &npool.GetKycByAppIDResponse{}
		err := json.Unmarshal(getKycByAppIDResp.Body(), getKycByAppIDResponse)
		if assert.Nil(t, err) {
			assert.NotNil(t, getKycByAppIDResponse.Infos)
			assert.LessOrEqual(t, len(getKycByAppIDResponse.Infos), 5)
		}
	}

	getAllKycResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetAllKycRequest{
			PageInfo: &npoolcommon.PageInfo{
				Limit:  5,
				Offset: 0,
			},
		}).Post("http://localhost:50120/v1/get/all/kyc")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, getAllKycResp.StatusCode())
		getAllKycResponse := &npool.GetAllKycResponse{}
		err := json.Unmarshal(getAllKycResp.Body(), getAllKycResponse)
		if assert.Nil(t, err) {
			assert.NotNil(t, getAllKycResponse.Infos)
			assert.LessOrEqual(t, len(getAllKycResponse.Infos), 5)
		}
	}

	updateKycResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.UpdateKycRequest{
			Info: &npool.KycInfo{
				ID:                 kycInfo.GetID(),
				UserID:             kycInfo.GetUserID(),
				AppID:              kycInfo.GetAppID(),
				CardType:           kycInfo.GetCardType(),
				CardID:             kycInfo.GetCardID(),
				FrontCardImg:       kycInfo.GetFrontCardImg(),
				BackCardImg:        kycInfo.GetBackCardImg(),
				UserHandingCardImg: kycInfo.GetUserHandingCardImg(),
			},
		}).Post("http://localhost:50120/v1/update/kyc")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, updateKycResp.StatusCode())
		updateKycResponse := &npool.UpdateKycResponse{}
		err := json.Unmarshal(updateKycResp.Body(), updateKycResponse)
		if assert.Nil(t, err) {
			assert.Equal(t, kycInfo.GetID(), updateKycResponse.GetInfo().GetID())
			assert.Equal(t, kycInfo.GetAppID(), updateKycResponse.GetInfo().GetAppID())
			assert.Equal(t, kycInfo.GetUserID(), updateKycResponse.GetInfo().GetUserID())
			assert.Equal(t, kycInfo.GetCardType(), updateKycResponse.GetInfo().GetCardType())
			assert.Equal(t, kycInfo.GetCardID(), updateKycResponse.GetInfo().GetCardID())
			assert.Equal(t, kycInfo.GetFrontCardImg(), updateKycResponse.GetInfo().GetFrontCardImg())
			assert.Equal(t, kycInfo.GetBackCardImg(), updateKycResponse.GetInfo().GetBackCardImg())
			assert.Equal(t, kycInfo.GetUserHandingCardImg(), updateKycResponse.GetInfo().GetUserHandingCardImg())
		}
	}

	// wrong
	newUserID := uuid.New().String()
	updateKycRespW, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.UpdateKycRequest{
			Info: &npool.KycInfo{
				ID:                 kycInfo.GetID(),
				UserID:             newUserID,
				AppID:              kycInfo.GetAppID(),
				CardType:           kycInfo.GetCardType(),
				CardID:             kycInfo.GetCardID(),
				FrontCardImg:       kycInfo.GetFrontCardImg(),
				BackCardImg:        kycInfo.GetBackCardImg(),
				UserHandingCardImg: kycInfo.GetUserHandingCardImg(),
			},
		}).Post("http://localhost:50120/v1/update/kyc")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, updateKycRespW.StatusCode())
	}

	kycIDs := []string{kycInfo.ID}
	getKycByKycIDsResp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetKycByKycIDsRequest{
			KycIDs: kycIDs,
		}).Post("http://localhost:50120/v1/get/kyc/by/kyc/ids")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, getKycByKycIDsResp.StatusCode())
	}
}
