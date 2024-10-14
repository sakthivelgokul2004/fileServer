
-- name: AddFileUrl :one
INSERT INTO files(userId,fileUrl , typeFile, fileName, id )
VALUES ($1, $2, $3, $4,$5)
RETURNING *;

-- name: GetFileByUserId :many 
SELECT * FROM files WHERE userId = $1;

-- name: GetFileByFileId :one
SELECT * FROM files WHERE id = $1;
-- name: DeleteByFileID :exec
DELETE  FROM files WHERE id = $1;