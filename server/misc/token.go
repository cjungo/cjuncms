package misc

import (
	"fmt"
	"time"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
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
	permitManager *mid.PermitManager[string, EmployeeToken]
}

func NewJwtClaimsManager(
	permitManager *mid.PermitManager[string, EmployeeToken],
) *JwtClaimsManager {
	return &JwtClaimsManager{
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
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	token, err := ext.MakeJwtToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (manager *JwtClaimsManager) NewOperation(
	ctx cjungo.HttpContext,
	operateType uint32,
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
		OperateType:    operateType,
		OperateSummary: operateSummary,
	}
}
