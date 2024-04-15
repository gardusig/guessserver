package guess

import (
	"context"
	"fmt"
	"net"

	guessproto "github.com/gardusig/guessproto/generated/go"
	"github.com/gardusig/guessserver/database"
	"github.com/gardusig/guessserver/internal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GuessServer struct {
	guessproto.UnimplementedGuessServiceServer

	db *database.SpecialNumberDb
}

func NewGuessServer() *GuessServer {
	return &GuessServer{
		db: database.NewSpecialNumberDb(),
	}
}

func (s *GuessServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", internal.GuessServicePort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	logrus.Debug("started listener at: ", lis.Addr())
	server := grpc.NewServer()
	guessproto.RegisterGuessServiceServer(server, s)
	return server.Serve(lis)
}

func (s *GuessServer) GuessNumber(ctx context.Context, req *guessproto.GuessNumberRequest) (*guessproto.GuessNumberResponse, error) {
	logrus.Debug("received request, level: ", req.Level, ", guessNumber: ", req.Guess)
	err := validateGuessNumberRequest(req)
	if err != nil {
		return nil, err
	}
	result, encryptedMessage, err := s.db.ValidateGuess(req.Level, req.Guess)
	if err != nil {
		return nil, err
	}
	response := guessproto.GuessNumberResponse{
		Result:    result,
		LockedBox: nil,
	}
	if result == internal.Equal {
		response.LockedBox = &guessproto.LockedBox{
			EncryptedMessage: *encryptedMessage,
		}
	}
	return &response, nil
}

func (s *GuessServer) OpenBox(ctx context.Context, req *guessproto.LockedBox) (*guessproto.OpenedBox, error) {
	logrus.Debug("server: OpenBox request, encryptedMessage: ", req.EncryptedMessage)
	decryptedMessage, err := s.db.ValidateLockedBox(0, req.EncryptedMessage)
	if err != nil {
		return nil, err
	}
	response := guessproto.OpenedBox{
		Message: decryptedMessage,
	}
	return &response, nil
}

func validateGuessNumberRequest(req *guessproto.GuessNumberRequest) error {
	if req.Level < internal.LevelMinThreshold {
		return fmt.Errorf("level must be at least %v", internal.LevelMinThreshold)
	}
	if req.Level > internal.LevelMaxThreshold {
		return fmt.Errorf("level must be at most %v", internal.LevelMaxThreshold)
	}
	if req.Guess < internal.GuessMinThreshold {
		return fmt.Errorf("guess must be at least %v", internal.GuessMinThreshold)
	}
	if req.Guess > internal.GuessMaxThreshold {
		return fmt.Errorf("guess must be at most %v", internal.GuessMaxThreshold)
	}
	return nil
}
