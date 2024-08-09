package main

import (
	"context"
	"log"
	"net"

	PokemonPb "grpc-transcoder-toy/pb/pokemon_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	PokemonPb.UnimplementedPokemonServer
}

func (s *server) GetLegendaryPokemons(ctx context.Context, params *PokemonPb.EmptyRequest) (*PokemonPb.PokemonListResponse, error) {
	resp := &PokemonPb.PokemonListResponse{
		Data: []*PokemonPb.PokemonData{
			{Name: "Lugia", Type: "Water"},
			{Name: "Moltres", Type: "Fire"},
			{Name: "Zaptos", Type: "Electic"},
		},
	}

	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	PokemonPb.RegisterPokemonServer(s, &server{})
	reflection.Register(s)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
