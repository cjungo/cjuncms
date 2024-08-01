package misc

import (
	"fmt"
	"time"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"
	"github.com/golang-jwt/jwt/v5"
)

type EmployeeToken struct {
	EmployeeId          uint32   `json:"eid,omitempty"`
	EmployeeNickname    string   `json:"nickname,omitempty"`
	EmployeePermissions []string `json:"permissions,omitempty"`
}

type JwtClaims struct {
	EmployeeToken
	jwt.RegisteredClaims
}

func (claims *JwtClaims) GetPermissions() []string {
	return claims.EmployeePermissions
}

func (claims *JwtClaims) GetStore() EmployeeToken {
	return claims.EmployeeToken
}

type JwtClaimsManager struct {
	mysql         *db.MySql
	permitManager *mid.PermitManager[string, EmployeeToken]
}

func NewJwtClaimsManager(
	mysql *db.MySql,
	permitManager *mid.PermitManager[string, EmployeeToken],
) *JwtClaimsManager {
	return &JwtClaimsManager{
		mysql:         mysql,
		permitManager: permitManager,
	}
}

func (manager *JwtClaimsManager) Renewal(ctx cjungo.HttpContext) (string, error) {
	proof, ok := manager.permitManager.GetProof(ctx)
	if !ok {
		return "", fmt.Errorf("proof 无效")
	}
	claims, ok := proof.(*JwtClaims)
	if !ok {
		return "", fmt.Errorf("proof 不是 JwtClaims 类型")
	}
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(2 * 24 * time.Hour))
	token, err := ext.MakeJwtToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

type EmployeeProfile struct {
	model.CjEmployee
}

func (manager *JwtClaimsManager) Profile(ctx cjungo.HttpContext) (*EmployeeProfile, error) {
	proof, ok := manager.permitManager.GetProof(ctx)
	if !ok {
		return nil, fmt.Errorf("proof 无效")
	}

	profile := EmployeeProfile{}
	if err := manager.mysql.Find(&profile.CjEmployee, proof.GetStore().EmployeeId).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}

type OperateType uint32

const (
	OPT_LOGIN OperateType = 0
	OPT_ADD   OperateType = 1
	OPT_DROP  OperateType = 2
	OPT_EDIT  OperateType = 3
	OPT_QUERY OperateType = 4
)

func (manager *JwtClaimsManager) NewOperation(
	ctx cjungo.HttpContext,
	operateType OperateType,
	operateSummary string,
) *model.CjOperation {
	proof, ok := manager.permitManager.GetProof(ctx)
	if !ok {
		panic("PERMIT PROOF IS EMPTY.")
	}
	return &model.CjOperation{
		OperatorID:     proof.GetStore().EmployeeId,
		OperatorType:   0,
		OperateAt:      ctx.GetReqAt(),
		OperateType:    uint32(operateType),
		OperateSummary: operateSummary,
	}
}
