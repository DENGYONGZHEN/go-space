package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type MetaData struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *MetaData {

	medt := &MetaData{}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			medt.UserAgent = userAgents[0]
		}
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			medt.UserAgent = userAgents[0]
		}

		if cliendIPs := md.Get(xForwardedForHeader); len(cliendIPs) > 0 {
			medt.ClientIP = cliendIPs[0]
		}
	}
	if p, ok := peer.FromContext(ctx); ok {
		medt.ClientIP = p.Addr.String()
	}

	return medt
}
