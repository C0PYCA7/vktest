-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE containers_stats(
    id SERIAL PRIMARY KEY,
    containerIP VARCHAR(55) NOT NULL,
    pingTimeMS INTEGER NOT NULL,
    lastSuccessDate TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE containers_stats;
-- +goose StatementEnd
