
-- +goose Up
CREATE TABLE files(
  userId  UUID NOT NULL REFERENCES users(id),
  fileUrl varchar(80) NOT NULL,
  typeFile varchar(10) NOT NULL,
  id  UUID PRIMARY KEY
);



-- +goose Down
DROP TABLE files;
