#!/bin/bash

PGHOST="localhost"
PGPORT="5432"
PGUSER="postgres"

create_migration() {
    local NAME="$1"
    migrate create -ext sql -dir migrations -seq "$NAME"
}

apply_migration() {
    local DBNAME="$1"
    local COMMAND="$2"
    migrate -path migrations -database "postgres://$PGUSER@$PGHOST:$PGPORT/$DBNAME?sslmode=disable" $COMMAND
}

function usage() {
  echo "Usage: $0 { create <migration_name> | migrate <dbname> (up|down) }"
  echo ""
  echo "Options:"
  echo "  create <migration_name>   Create a new migration with the specified name"
  echo "  migrate <dbname> (up|down)  Apply the specified migration command to the specified database"
  echo ""
  echo "Environment variables:"
  echo "  PGHOST  Database host (default: localhost)"
  echo "  PGPORT  Database port (default: 5432)"
  echo "  PGUSER  Database user (default: postgres)"
}

if [ "$1" == "--help" ]; then
    usage
    exit 0
fi

case "$1" in 
    create)
        create_migration "$2"
        ;;
    migrate)
        apply_migration "$2" "$3"
        ;;
    *)
        echo "Error. Command non found."
        exit 1
        ;;
esac