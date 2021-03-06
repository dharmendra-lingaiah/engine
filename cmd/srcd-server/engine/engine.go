package engine

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	api "github.com/src-d/engine/api"
)

var _ api.EngineServer = new(Server)

type Server struct {
	version     string
	workdir     string
	datadir     string
	workdirHash string
}

func NewServer(version, workdir, datadir string) *Server {
	h := sha1.Sum([]byte(workdir))
	return &Server{
		version:     version,
		workdir:     workdir,
		datadir:     datadir,
		workdirHash: hex.EncodeToString(h[:]),
	}
}

func (s *Server) Version(ctx context.Context, req *api.VersionRequest) (*api.VersionResponse, error) {
	return &api.VersionResponse{Version: s.version}, nil
}
