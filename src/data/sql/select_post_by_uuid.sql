SELECT
 id, uuid, body, user_id, thread_id, created_at
FROM
 posts
WHERE
 uuid = $1