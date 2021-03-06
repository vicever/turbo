package turbo

import (
	"google.golang.org/grpc"
)

type grpcClient struct {
	grpcService interface{}
	conn        *grpc.ClientConn
}

func (g *grpcClient) init(addr string, clientCreator func(conn *grpc.ClientConn) interface{}) error {
	// ???? support multiple grpc clients
	// ???? support grpcservice discovery
	if g.grpcService != nil {
		return nil
	}
	log.Info("[grpc]connecting addr:", addr)
	err := g.dial(addr)
	if err == nil {
		g.grpcService = clientCreator(g.conn)
	}
	return err
}

func (g *grpcClient) dial(address string) (err error) {
	g.conn, err = grpc.Dial(address, grpc.WithInsecure())
	return err
}

func (g *grpcClient) close() error {
	if g.conn == nil {
		return nil
	}
	return g.conn.Close()
}

// GrpcService returns a grpc client instance,
// example: client := turbo.GrpcService().(proto.YourServiceClient)
func GrpcService(s *Server) interface{} {
	if s == nil || s.gClient == nil || s.gClient.grpcService == nil {
		log.Panic("grpc connection not initiated!")
	}
	return s.gClient.grpcService
}
