package handler

import (
	"context"

	"github.com/uchupx/kajian-auth/pb"
)

type AuthGRPCHandler struct {
	pb.AuthorizationServiceServer
}

func (h *AuthGRPCHandler) CredentialCheck(ctx context.Context, req *pb.CredentialCheckRequest) (*pb.CredentialCheckResponse, error) {
	return nil, nil
}
