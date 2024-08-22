-- +goose Up
-- +goose StatementBegin
INSERT INTO customer (NIK, full_name, legal_name, born_place, born_date, salary, is_admin, email, password, foto_selfie, foto_ktp)
VALUES ('WnIUt10aFKBYMg==', 'admin', 'admin', 'City', '1990-01-01', 5000.00, true, 'admin@example.com', '$2a$10$LwQaRAREL3Mon8Xq0ZWlJuX/KGGuL3HdlLryhxTmLBK4326ZCW2oa', 'CjJE6hxJQOwUcMrMYDCXFQKBhPd7ng==', 'CjJE6hxJQOwUcMrMYDCXFQKBhPd7ng==');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
