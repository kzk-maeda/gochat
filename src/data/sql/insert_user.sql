insert into
 users (uuid, name, email, password, created_at)
values
 ($1, $2, $3, $4, $5)
returning
 id, uuid, created_at