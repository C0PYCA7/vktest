-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE containers_stats RENAME COLUMN pingTimeMS TO pingTimeMKS;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE containers_stats RENAME COLUMN pingTimeMKS TO pingTimeMS;
-- +goose StatementEnd
