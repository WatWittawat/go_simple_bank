package gapi

import (
	"fmt"

	db "github.com/WatWittawat/go_simple_bank/db/sqlc"
	"github.com/WatWittawat/go_simple_bank/pb"
	"github.com/WatWittawat/go_simple_bank/token"
	"github.com/WatWittawat/go_simple_bank/utils"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	pb.UnimplementedSimpleBankServer
}

// mustEmbedUnimplementedSimpleBankServer implements [pb.SimpleBankServer].
func (s *Server) mustEmbedUnimplementedSimpleBankServer() {
	panic("unimplemented")
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token meker: %w", err)
	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	return server, nil
}
