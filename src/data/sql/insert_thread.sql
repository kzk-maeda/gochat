insert into
 threads (uuid, topic, user_id, created_at) 
values
 ($1, $2, $3, $4) 
returning
 id, uuid, topic, user_id, created_at