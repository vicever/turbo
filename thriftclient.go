package turbo

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type thriftClient struct {
	thriftService interface{}
	transport     thrift.TTransport
	factory       thrift.TProtocolFactory
}

func (t *thriftClient) init(addr string, clientCreator func(trans thrift.TTransport, f thrift.TProtocolFactory) interface{}) error {
	if t.thriftService != nil {
		return nil
	}
	log.Debugf("connecting thrift addr: %s", addr)
	err := t.connect(addr)
	if err == nil {
		t.thriftService = clientCreator(t.transport, t.factory)
	}
	return err
}

func (t *thriftClient) connect(hostPort string) (err error) {
	tSocket, err := thrift.NewTSocket(hostPort)
	if err != nil {
		return err
	}
	t.transport, err = thrift.NewTTransportFactory().GetTransport(tSocket)
	if err != nil {
		return err
	}
	if err := t.transport.Open(); err != nil {
		return err
	}
	t.factory = thrift.NewTBinaryProtocolFactoryDefault()
	return nil
}

func (t *thriftClient) close() error {
	if t.transport == nil {
		return nil
	}
	return t.transport.Close()
}

// ThriftService returns a Thrift client instance,
// example: client := turbo.ThriftService().(proto.YourServiceClient)
func ThriftService(s *Server) interface{} {
	if s == nil || s.tClient == nil || s.tClient.thriftService == nil {
		log.Panic("thrift connection not initiated!")
	}
	return s.tClient.thriftService
}
