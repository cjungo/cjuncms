package misc

import (
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
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
