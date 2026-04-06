#!/bin/bash
MASTER_HOST="numbers-postgres-master"
MASTER_PORT=5432
USER="postgres"
DB="numbers"
SLAVE_NAME="numbers-postgres-slave"

PASSWORD=${POSTGRES_SUPER_PASSWORD:?POSTGRES_SUPER_PASSWORD must be set}


while true; do
    status=$(PGPASSWORD=$PASSWORD psql "host=$MASTER_HOST port=$MASTER_PORT user=$USER dbname=$DB" \
             -t -c "SELECT state FROM pg_stat_replication;" | xargs)

    current=$(PGPASSWORD=$PASSWORD psql "host=$MASTER_HOST port=$MASTER_PORT user=$USER dbname=$DB" -t -c "SHOW synchronous_commit;" | xargs)

    if [[ "$status" != "streaming" && "$current" != "off" ]]; then
        echo "$(date) - Slave down, switching master to async"
        PGPASSWORD=$PASSWORD psql "host=$MASTER_HOST port=$MASTER_PORT user=$USER dbname=$DB" -c "ALTER SYSTEM SET synchronous_commit = 'off';"
        PGPASSWORD=$PASSWORD psql "host=$MASTER_HOST port=$MASTER_PORT user=$USER dbname=$DB" -c "SELECT pg_reload_conf();"
    elif [[ "$status" == "streaming" && "$current" != "on" ]]; then
        echo "$(date) - Slave streaming, switching master to sync"
        PGPASSWORD=$PASSWORD psql "host=$MASTER_HOST port=$MASTER_PORT user=$USER dbname=$DB" -c "ALTER SYSTEM SET synchronous_commit = 'on';"
        PGPASSWORD=$PASSWORD psql "host=$MASTER_HOST port=$MASTER_PORT user=$USER dbname=$DB" -c "SELECT pg_reload_conf();"
    fi




    sleep 5
done