-- +goose Up
-- +goose StatementBegin
CREATE TABLE customer
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    NIK           VARCHAR(255)        NOT NULL,
    full_name     VARCHAR(255)        NOT NULL,
    legal_name    VARCHAR(255),
    born_place    VARCHAR(255),
    born_date     DATE,
    salary        DECIMAL(10, 2),
    is_admin      BOOLEAN DEFAULT false,
    email         VARCHAR(255) UNIQUE NOT NULL,
    password      VARCHAR(255)        NOT NULL,
    foto_selfie   VARCHAR(255),
    foto_ktp      VARCHAR(255),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customer;
-- +goose StatementEnd
