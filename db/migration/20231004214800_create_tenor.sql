-- +goose Up
-- +goose StatementBegin
CREATE TABLE tenor (
   id INT AUTO_INCREMENT PRIMARY KEY,
   tenor INT NOT NULL,
   value FLOAT DEFAULT 0,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tenor;
-- +goose StatementEnd
