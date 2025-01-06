-- +goose Up
CREATE TABLE urls (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL
);

-- +goose Down
DROP TABLE urls;
