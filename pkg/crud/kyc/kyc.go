package kyc

import (
	"context"

	"github.com/NpoolPlatform/kyc-management/pkg/db"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	npool "github.com/NpoolPlatform/message/npool/kyc"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToKyc(row *ent.Kyc) *npool.KycInfo {
	return &npool.KycInfo{
		ID:                 row.ID.String(),
		UserID:             row.UserID.String(),
		AppID:              row.AppID.String(),
		CardType:           row.CardType,
		CardID:             row.CardID,
		FrontCardImg:       row.FrontCardImg,
		BackCardImg:        row.BackCardImg,
		UserHandingCardImg: row.UserHandingCardImg,
		CreateAt:           row.CreateAt,
		UpdateAt:           row.UpdateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateKycRequest) (*npool.CreateKycResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Kyc.
		Create().
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetCardType(in.GetInfo().GetCardType()).
		SetCardID(in.GetInfo().GetCardID()).
		SetFrontCardImg(in.GetInfo().GetFrontCardImg()).
		SetBackCardImg(in.GetInfo().GetBackCardImg()).
		SetUserHandingCardImg(in.GetInfo().GetUserHandingCardImg()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateKycResponse{
		Info: dbRowToKyc(info),
	}, nil
}

func GetKycByUserIDAndAppID(ctx context.Context, appID, userID uuid.UUID) (*npool.KycInfo, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}
	info, err := cli.Kyc.Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
				kyc.UserID(userID),
			),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	return dbRowToKyc(info), nil
}

func ExistKycByUserIDAndAppID(ctx context.Context, appID, userID uuid.UUID) (bool, error) {
	cli, err := db.Client()
	if err != nil {
		return false, xerrors.Errorf("fail get db client: %v", err)
	}

	return cli.Kyc.Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
				kyc.UserID(userID),
			),
		).
		Exist(ctx)
}

func GetKycByID(ctx context.Context, kycID uuid.UUID) (*npool.KycInfo, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.Kyc.Query().
		Where(
			kyc.ID(kycID),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	return dbRowToKyc(info), nil
}

func GetAll(ctx context.Context, appID uuid.UUID, limit, offset int32) ([]*npool.KycInfo, int, error) {
	if limit == 0 {
		limit = 10
	}

	cli, err := db.Client()
	if err != nil {
		return nil, 0, xerrors.Errorf("fail get db client: %v", err)
	}

	client := cli.Kyc.Query()

	if appID != uuid.Nil {
		client.Where(kyc.AppID(appID))
	}

	total, err := client.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	infos, err := client.
		Order(ent.Desc(kyc.FieldCreateAt)).
		Offset(int(offset)).
		Limit(int(limit)).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	response := []*npool.KycInfo{}
	for _, info := range infos {
		response = append(response, dbRowToKyc(info))
	}

	return response, total, nil
}

func Update(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Kyc.
		UpdateOneID(uuid.MustParse(in.GetInfo().GetID())).
		SetCardType(in.GetInfo().GetCardType()).
		SetCardID(in.GetInfo().GetCardID()).
		SetFrontCardImg(in.GetInfo().GetFrontCardImg()).
		SetBackCardImg(in.GetInfo().GetBackCardImg()).
		SetUserHandingCardImg(in.GetInfo().GetUserHandingCardImg()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.UpdateKycResponse{
		Info: dbRowToKyc(info),
	}, nil
}

func ExistCradTypeCardIDInAppExceptUserID(ctx context.Context, cardType, cardID string, appID, userID uuid.UUID) (bool, error) {
	cli, err := db.Client()
	if err != nil {
		return false, xerrors.Errorf("fail to get db client: %v", err)
	}

	return cli.Kyc.Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
				kyc.CardType(cardType),
				kyc.CardID(cardID),
				kyc.UserIDNotIn(userID),
			),
		).
		Exist(ctx)
}

func GetKycByKycIDs(ctx context.Context, kycIDs []uuid.UUID) ([]*npool.KycInfo, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail to get db client: %v", err)
	}

	infos, err := cli.Kyc.Query().
		Where(
			kyc.IDIn(kycIDs...),
		).All(ctx)
	if err != nil {
		return nil, err
	}
	response := []*npool.KycInfo{}
	for _, info := range infos {
		response = append(response, dbRowToKyc(info))
	}
	return response, nil
}
