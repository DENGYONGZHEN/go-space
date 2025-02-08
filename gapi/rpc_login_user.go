package gapi

import (
	"context"
	"errors"
	db "simple-bank/db/sqlc"
	"simple-bank/pb"
	"simple-bank/util"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to find user: %s", err)
	}

	err = util.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password : %s", err)
	}

	accessToken, payload, err := server.tokenMaker.CreateToken(req.GetUsername(), server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err)
	}
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(req.GetUsername(),
		server.config.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refreshToken: %s", err)
	}
	var pgUUID pgtype.UUID
	err = pgUUID.Scan(refreshPayload.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "scan uuid err: %s", err)
	}

	medt := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           pgUUID,
		Username:     refreshPayload.Username,
		RefreshToken: refreshToken,
		UserAgent:    medt.UserAgent,
		ClientIp:     medt.ClientIP,
		IsBlocked:    false,
		ExpiresAt: pgtype.Timestamptz{
			Time:  refreshPayload.ExpiresAt.Time,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create session error: %s", err)
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(payload.ExpiresAt.Time),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt.Time),
	}

	return rsp, nil
}
