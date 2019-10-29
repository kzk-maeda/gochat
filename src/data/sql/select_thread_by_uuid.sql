SELECT
 id, uuid, topic, user_id, created_at
FROM
 threads
WHERE
 uuid = $1