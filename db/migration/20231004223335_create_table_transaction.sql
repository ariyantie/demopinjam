-- +goose Up
-- +goose StatementBegin
CREATE TABLE transaction (
     id INT AUTO_INCREMENT PRIMARY KEY,
     customer_id INT NOT NULL,
     contract_number VARCHAR(255) UNIQUE NOT NULL,
     OTR INT NOT NULL,
     admin_fee FLOAT NOT NULL,
     total_installment INT NOT NULL,
     interest FLOAT NOT NULL,
     asset_name VARCHAR(255) NOT NULL,
     status ENUM('Success', 'Failed') NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     deleted_at TIMESTAMP DEFAULT NULL,
     FOREIGN KEY (customer_id) REFERENCES customer(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transaction;
-- +goose StatementEnd
