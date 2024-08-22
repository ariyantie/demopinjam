-- +goose Up
-- +goose StatementBegin
INSERT INTO tenor (tenor, value) VALUES (1, 1.0),(2,1.5),(3,1.8),(6,2.5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
