update
 posts
set 
 body = $2,
 created_at = $3
where
 id = $1