path "/kv/data/wikipedia/email" {
  capabilities = ["read"]
}

path "/database/creds/random-wikipedia" {
  capabilities = ["read"]
}

path "/kv/data/wikipedia/vapidKeys" {
  capabilities = ["read", "update"]
}