package handler

import (
	"context"
	"encoding/json"

	"github.com/uchupx/kajian-auth/internal/dto"
	"github.com/uchupx/kajian-auth/internal/service/jwt"
	"github.com/uchupx/kajian-auth/pb"
)

type AuthGRPCHandler struct {
	pb.AuthorizationServiceServer
	JWTService jwt.CryptService
}

func (h *AuthGRPCHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := pb.GetUserResponse{}
	res.IsAuthorized = false

	jwtToken := req.GetToken()
	if jwtToken == "" {
		return &res, nil
	}

	resToken, err := h.JWTService.VerifyJWTToken(jwtToken)
	if err != nil {
		return &res, nil
	}

	bytes, err := json.Marshal(resToken)
	if err != nil {
		return &res, nil
	}

	var user dto.User

	if err := json.Unmarshal(bytes, &user); err != nil {
		return &res, nil
	}

	res.IsAuthorized = true
	res.Id = int32(user.ID)
	res.Username = user.Username
	res.Email = user.Email
	res.Name = ""
	res.CreatedAt = user.Created.String()

	return &res, nil
}
