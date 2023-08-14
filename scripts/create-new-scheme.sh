#!/bin/bash
cd ..
go run -mod=mod entgo.io/ent/cmd/ent new $1
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
cd scripts