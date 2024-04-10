-- +goose Up
CREATE TABLE IF NOT EXISTS url (
    id int GENERATED ALWAYS AS IDENTITY,
    alias text UNIQUE NOT NULL,
    url text NOT NULL,

    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS url;