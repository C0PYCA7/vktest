-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE containers_stats
ALTER COLUMN pingTimeMS SET DEFAULT -1,
ALTER COLUMN lastSuccessDate SET DEFAULT '1999-01-01 00:00:00';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE containers_stats
ALTER COLUMN pingTimeMS DROP DEFAULT,
ALTER COLUMN lastSuccessDate DROP DEFAULT;
-- +goose StatementEnd
