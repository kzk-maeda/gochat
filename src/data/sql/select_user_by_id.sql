SELECT
 id, uuid, name, email, created_at
FROM
 users 
WHERE
 id = $1