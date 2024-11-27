#!/bin/sh

set -e

echo "Running database migrations..."
./migrate

echo "Starting the application..."
exec "$@"
