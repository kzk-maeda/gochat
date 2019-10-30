insert into
 posts (uuid, body, user_id, thread_id, created_at)
values
 ($1, $2, $3, $4, $5)
returning
 id, uuid, body, user_id, thread_id, created_at