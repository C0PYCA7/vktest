-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE containers_stats
ALTER COLUMN pingTimeMS DROP NOT NULL,
ALTER COLUMN lastSuccessDate DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE containers_stats
ALTER COLUMN pingTimeMS SET NOT NULL,
ALTER COLUMN lastSuccessDate SET NOT NULL;
-- +goose StatementEnd
