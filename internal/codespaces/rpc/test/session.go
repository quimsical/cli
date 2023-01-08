package test

import (
	"context"
	"fmt"
	"net"

	"github.com/cli/cli/v2/pkg/liveshare"
	"golang.org/x/crypto/ssh"
)

type Session struct {
	channel ssh.Channel
}

func (*Session) Close() error {
	panic("unimplemented")
}

func (*Session) GetSharedServers(context.Context) ([]*liveshare.Port, error) {
	panic("unimplemented")
}

func (*Session) RebuildContainer(context.Context, bool) error {
	panic("unimplemented")
}

func (*Session) StartSSHServer(context.Context) (int, string, error) {
	panic("unimplemented")
}

func (*Session) StartSSHServerWithOptions(context.Context, liveshare.StartSSHServerOptions) (int, string, error) {
	panic("unimplemented")
}

func (s *Session) KeepAlive(reason string) {
}

func (s *Session) StartSharing(ctx context.Context, sessionName string, port int) (liveshare.ChannelID, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", ServerPort))
	if err != nil {
		return liveshare.ChannelID{}, err
	}
	s.channel = &Channel{conn}
	return liveshare.ChannelID{}, nil
}

// Creates mock SSH channel connected to the mock gRPC server
func (s *Session) OpenStreamingChannel(ctx context.Context, id liveshare.ChannelID) (ssh.Channel, error) {
	return s.channel, nil
}
