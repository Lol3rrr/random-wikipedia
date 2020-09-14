#!/bin/bash

POSTGRES_ID=$(docker run -d -e POSTGRES_PASSWORD=devPassword -e POSTGRES_USER=devAdmin -e POSTGRES_DB=devWikipedia --network host postgres:12);
echo "Postgres: $POSTGRES_ID";

sleep 3s;
docker cp ./schema.sql $POSTGRES_ID:/schema.sql
docker exec -it $POSTGRES_ID psql -d devWikipedia -U devAdmin -f schema.sql
docker cp ./index.sql $POSTGRES_ID:/index.sql
docker exec -it $POSTGRES_ID psql -d devWikipedia -U devAdmin -f index.sql

VAULT_ID=$(docker run --cap-add=IPC_LOCK -d --network host -e 'VAULT_LOCAL_CONFIG={"backend": {"file": {"path": "/vault/file"}}}' vault server -dev);
echo "Vault: $VAULT_ID";

sleep 3s;
VAULT_ROOT=$(docker logs $VAULT_ID 2>&1 | awk -F': ' '/Root Token: /{print $2}');
VAULT_ADDR='http://0.0.0.0:8200' vault login $VAULT_ROOT
VAULT_ADDR='http://0.0.0.0:8200' vault secrets enable database
VAULT_ADDR='http://0.0.0.0:8200' vault write database/config/postgres plugin_name="postgresql-database-plugin" connection_url="postgresql://{{username}}:{{password}}@localhost:5432/devWikipedia?sslmode=disable" allowed_roles="random-wikipedia" username="devAdmin" password="devPassword"
VAULT_ADDR='http://0.0.0.0:8200' vault write database/roles/random-wikipedia db_name="postgres" creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; GRANT ALL ON ALL TABLES IN SCHEMA randomWikipedia TO \"{{name}}\";" revocation_statements="DROP OWNED BY \"{{name}}\"; DROP user \"{{name}}\";" default_ttl="1h"
VAULT_ADDR='http://0.0.0.0:8200' vault auth enable approle
VAULT_ADDR='http://0.0.0.0:8200' vault policy write random-wikipedia ./policy.hcl
VAULT_ADDR='http://0.0.0.0:8200' vault write auth/approle/role/random-wikipedia token_ttl="20m" policies="random-wikipedia"
VAULT_ADDR='http://0.0.0.0:8200' vault read auth/approle/role/random-wikipedia/role-id
VAULT_ADDR='http://0.0.0.0:8200' vault write -force auth/approle/role/random-wikipedia/secret-id