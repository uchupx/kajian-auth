package handler

import (
	"context"

	"github.com/uchupx/kajian-auth/internal/service"
	"github.com/uchupx/kajian-auth/pb"
)

type AuthGRPCHandler struct {
	pb.AuthorizationServiceServer
	UserService *service.UserService
}

// GetUser is a function to get user data from jwt token (GRPC)
func (h *AuthGRPCHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := pb.GetUserResponse{}
	res.IsAuthorized = false

	jwtToken := req.GetToken()
	if jwtToken == "" {
		return &res, nil
	}

	user, err := h.UserService.RetrieveUser(ctx, jwtToken)
	if err != nil {
		return &res, nil
	}

	res.IsAuthorized = true
	res.Id = user.ID
	res.Username = user.Username
	res.Email = user.Email
	res.Name = ""
	res.CreatedAt = user.Created.String()

	return &res, nil
}
