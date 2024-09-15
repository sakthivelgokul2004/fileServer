
-- name: AddImageUrl :one
INSERT INTO files(userId,fileUrl , typeFile, id )
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetFileByUserId :many 
SELECT * FROM files WHERE userId = $1;