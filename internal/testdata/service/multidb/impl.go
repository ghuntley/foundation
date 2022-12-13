// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package multidb

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"namespacelabs.dev/foundation/internal/testdata/service/proto"
	"namespacelabs.dev/foundation/std/go/server"
	"namespacelabs.dev/foundation/universe/db/postgres"
)

type Service struct {
	postgres *postgres.DB
	rds      *postgres.DB
}

const timeout = 2 * time.Second

func addPostgres(ctx context.Context, db *postgres.DB, item string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := db.Exec(ctx, "INSERT INTO list (Item) VALUES ($1);", item)
	return err
}

func (svc *Service) AddRds(ctx context.Context, req *proto.AddRequest) (*emptypb.Empty, error) {
	log.Printf("new AddRds request: %+v\n", req)

	if err := addPostgres(ctx, svc.rds, req.Item); err != nil {
		log.Fatalf("failed to add list item: %v", err)
	}

	response := &emptypb.Empty{}
	return response, nil
}

func (svc *Service) AddPostgres(ctx context.Context, req *proto.AddRequest) (*emptypb.Empty, error) {
	log.Printf("new AddPostgres request: %+v\n", req)

	if err := addPostgres(ctx, svc.postgres, req.Item); err != nil {
		log.Fatalf("failed to add list item: %v", err)
	}

	response := &emptypb.Empty{}
	return response, nil
}

func listPostgres(ctx context.Context, db *postgres.DB) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := db.Query(ctx, "SELECT Item FROM list;")
	if err != nil {
		return nil, fmt.Errorf("failed read list: %w", err)
	}
	defer rows.Close()

	var res []string
	for rows.Next() {
		var item string
		err = rows.Scan(&item)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (svc *Service) List(ctx context.Context, _ *emptypb.Empty) (*proto.ListResponse, error) {
	log.Print("new List request\n")

	var list []string

	rdslist, err := listPostgres(ctx, svc.rds)
	if err != nil {
		log.Fatalf("failed to read list: %v", err)
	}
	list = append(list, rdslist...)

	pglist, err := listPostgres(ctx, svc.postgres)
	if err != nil {
		log.Fatalf("failed to read list: %v", err)
	}
	list = append(list, pglist...)

	response := &proto.ListResponse{Item: list}
	return response, nil
}

func WireService(ctx context.Context, srv server.Registrar, deps ServiceDeps) {
	svc := &Service{
		postgres: deps.Postgres,
		rds:      deps.Rds,
	}
	var err error
	svc.postgres, err = wirePostgres()
	if err != nil {
		log.Fatalf("failed to wire postgres: %v", err)
	}
	proto.RegisterMultiDbListServiceServer(srv, svc)
}
