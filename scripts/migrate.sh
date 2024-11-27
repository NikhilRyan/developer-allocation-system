#!/bin/bash
migrate -path pkg/db/migrations -database "postgres://postgres:postgres@localhost:5432/developer_allocation?sslmode=disable" up
