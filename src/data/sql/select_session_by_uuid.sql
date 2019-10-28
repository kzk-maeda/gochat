SELECT
 id, uuid, email, user_id, created_at
FROM
 sessions
WHERE
 uuid = $1