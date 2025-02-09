-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE UNIQUE INDEX unique_container_ip ON containers_stats(containerIP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP INDEX unique_container_ip;
-- +goose StatementEnd
