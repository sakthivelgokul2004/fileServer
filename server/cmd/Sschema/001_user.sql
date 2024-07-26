-- +goose up

CREATE TABLE USER{
id UUID PRIMARY KEY,
name TEXT NOT null
}
-- +goose down

DROP TABLE user;
