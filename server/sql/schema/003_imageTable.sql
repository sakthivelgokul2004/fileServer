
-- +goose Up
CREATE TABLE files(
  userId  UUID NOT NULL REFERENCES users(id),
  fileName TEXT NOT NULL,
  fileUrl TEXT  NOT NULL,
  typeFile TEXT NOT NULL,
  id  UUID PRIMARY KEY
);



-- +goose Down
DROP TABLE files;
