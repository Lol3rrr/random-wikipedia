#!/bin/bash

POSTGRES_ID=$(docker run -d -e POSTGRES_PASSWORD=devPassword -e POSTGRES_USER=devAdmin -e POSTGRES_DB=devWikipedia --network host postgres:12);
echo "Postgres: $POSTGRES_ID";

sleep 3s;
docker exec -it $POSTGRES_ID psql -d devWikipedia -U devAdmin -c 'CREATE TABLE IF NOT EXISTS Users (ID TEXT NOT NULL PRIMARY KEY, SessionID TEXT NOT NULL, Email TEXT NOT NULL);';
docker exec -it $POSTGRES_ID psql -d devWikipedia -U devAdmin -c 'CREATE TABLE IF NOT EXISTS Passwords (ID TEXT NOT NULL PRIMARY KEY, Password TEXT NOT NULL, Expiration INTEGER);';
docker exec -it $POSTGRES_ID psql -d devWikipedia -U devAdmin -c 'CREATE TABLE IF NOT EXISTS Notifications (ID TEXT NOT NULL PRIMARY KEY, Subscription TEXT NOT NULL);';
docker exec -it $POSTGRES_ID psql -d devWikipedia -U devAdmin -c 'CREATE TABLE IF NOT EXISTS Settings (ID TEXT NOT NULL PRIMARY KEY, Notifytime INTEGER);';
docker exec -it $POSTGRES_ID psql -d devWikipedia -U devAdmin -c 'CREATE TABLE IF NOT EXISTS Userlists (ID TEXT NOT NULL PRIMARY KEY, ListID INTEGER);';
docker exec -it $POSTGRES_ID psql -d devWikipedia -U devAdmin -c 'CREATE TABLE IF NOT EXISTS Lists (ListID INTEGER PRIMARY KEY, Title TEXT NOT NULL);';

VAULT_ID=$(docker run --cap-add=IPC_LOCK -d --network host -e 'VAULT_LOCAL_CONFIG={"backend": {"file": {"path": "/vault/file"}}}' vault server -dev);
echo "Vault: $VAULT_ID";

sleep 3s;
VAULT_ROOT=$(docker logs $VAULT_ID 2>&1 | awk -F': ' '/Root Token: /{print $2}');
VAULT_ADDR='http://0.0.0.0:8200' vault login $VAULT_ROOT
VAULT_ADDR='http://0.0.0.0:8200' vault secrets enable database
VAULT_ADDR='http://0.0.0.0:8200' vault write database/config/postgres plugin_name="postgresql-database-plugin" connection_url="postgresql://{{username}}:{{password}}@localhost:5432/devWikipedia?sslmode=disable" allowed_roles="random-wikipedia" username="devAdmin" password="devPassword"
VAULT_ADDR='http://0.0.0.0:8200' vault write database/roles/random-wikipedia db_name="postgres" creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; GRANT ALL ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" revocation_statements="DROP OWNED BY \"{{name}}\"; DROP user \"{{name}}\";" default_ttl="1h"
VAULT_ADDR='http://0.0.0.0:8200' vault auth enable approle
VAULT_ADDR='http://0.0.0.0:8200' vault policy write random-wikipedia ./policy.hcl
VAULT_ADDR='http://0.0.0.0:8200' vault write auth/approle/role/random-wikipedia token_ttl="20m" policies="random-wikipedia"
VAULT_ADDR='http://0.0.0.0:8200' vault read auth/approle/role/random-wikipedia/role-id
VAULT_ADDR='http://0.0.0.0:8200' vault write -force auth/approle/role/random-wikipedia/secret-id