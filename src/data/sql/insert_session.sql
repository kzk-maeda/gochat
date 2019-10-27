insert into
 sessions (uuid, email, user_id, created_at)
values
 ($1, $2, $3, $4)
returning
 id, uuid, email, user_id, created_at