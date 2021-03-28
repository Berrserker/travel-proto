echo "Running db migrations"

psql $PostgresMaster "select 1"

RETRIES=10

until psql $PostgresMaster "select 1" > /dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
  echo "Waiting for postgres server, $((RETRIES--)) remaining attempts..."
  sleep 1
done

goose -version
goose -dir $(pwd)/migrations postgres $PostgresMaster down-to 0
goose -dir $(pwd)/migrations postgres $PostgresMaster up
