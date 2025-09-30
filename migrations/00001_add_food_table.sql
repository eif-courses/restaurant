-- +goose Up
-- +goose StatementBegin

    CREATE TABLE foods
    (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        price DOUBLE PRECISION NOT NULL
    );


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE foods;
-- +goose StatementEnd
