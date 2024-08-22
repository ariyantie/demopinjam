-- +goose Up
-- +goose StatementBegin
CREATE TABLE customer_loan (
   id INT AUTO_INCREMENT PRIMARY KEY,
   customer_id INT UNIQUE,
   status ENUM('Requested', 'Approved', 'Used','Expired') NOT NULL,
   loan_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   loan_amount DECIMAL(10, 2) NOT NULL,
   used_amount DECIMAL(10, 2) NOT NULL DEFAULT 0,
   approved_date TIMESTAMP DEFAULT NULL,
   tenor INT NOT NULL,
   expired_at TIMESTAMP DEFAULT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP DEFAULT NULL,
   FOREIGN KEY (customer_id) REFERENCES customer(id),
   FOREIGN KEY (tenor) REFERENCES tenor(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customer_loan;
-- +goose StatementEnd
