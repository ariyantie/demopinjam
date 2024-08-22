package repository

// USERS QUERY
const (
	insertNewCostumer  = `INSERT INTO customer(NIK, full_name, legal_name, born_place, born_date, salary, is_admin, email, password, foto_selfie, foto_ktp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	getCostumerByEmail = `
        SELECT
            id,
            NIK,
            full_name,
            legal_name,
            born_place,
            born_date,
            salary,
            is_admin,
            email,
            password,
            foto_selfie,
            foto_ktp,
            created_at,
            updated_at,
            deleted_at
        FROM
            customer
        WHERE
            email = ?
    `

	queryRequestLoan = `
		INSERT INTO customer_loan (customer_id, tenor, loan_date, loan_amount, status)
		VALUES (?, ?, ?, ?, ?)
	`
	queryGetUserLimit = `  SELECT
            id,
            customer_id,
            status,
            used_amount,
            loan_amount,
            tenor,
            expired_at
        FROM
            customer_loan WHERE customer_id=?`
	queryGetListCustomerRequest = `
        SELECT
            id,
            customer_id,
            status,
            loan_amount,
            used_amount,
            tenor,
            created_at,
            updated_at,
            loan_date
        FROM
            customer_loan
        WHERE
            status = ? ORDER BY created_at desc 
    `
	queryGetListCustomerRequestByIds = `
        SELECT
            id,
            customer_id,
            status,
            loan_amount,
            used_amount,
            tenor,
            created_at,
            updated_at,
            loan_date
        FROM
            customer_loan
        WHERE
            status = ? AND id IN (%s)
    `
	queryApproveLoanRequest = `
			UPDATE customer_loan
			SET
			  status = ?,
			  approved_date = ?,
			  expired_at = ?
			WHERE
			  id = ?
`
	queryUseLimitRequest = `
			UPDATE customer_loan
			SET
			  status = ?,
			  used_amount = ?
			WHERE
			  id = ?`
)

const (
	getTenorList = `SELECT id,tenor,value FROM tenor`
	getTenorByID = `SELECT id,tenor,value FROM tenor WHERE id = ?`

	queryCreateTransaction = `INSERT INTO transaction (customer_id, contract_number, OTR, admin_fee, total_installment, interest, asset_name, status)
							VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	queryCreateSchedulePayment = `INSERT INTO schedule_payment (transaction_id, amount, status, due_date)
								VALUES(?,?,?,?)`
	queryGetListTransaction = `
						SELECT
							DATE_FORMAT(due_date , '%Y-%m') AS month,
							JSON_ARRAYAGG(
								JSON_OBJECT(
									'id', id,
									'transaction_id', transaction_id,
									'payment_date', payment_date,
									'amount', amount,
									'status', status,
									'due_date', due_date,
									'late_fee', late_fee
								)
							) AS payments
						FROM schedule_payment
						WHERE transaction_id IN (
							SELECT id
							FROM transaction
							WHERE customer_id = ?
						)
						GROUP BY month
						ORDER BY month;
`
	queryGetSchedulePayment = `
			SELECT sp.*
			FROM schedule_payment sp
			JOIN transaction t ON sp.transaction_id = t.id
			WHERE t.customer_id = ? 
				AND DATE_FORMAT(sp.due_date, '%Y-%m') = ? AND sp.status != 'Paid';`

	queryUpdateSchedulePayment = `
			UPDATE schedule_payment
			SET status = ?, payment_date = ?
			WHERE id = ?;`

	queryUpdateIdentity = `UPDATE customer
							SET
								foto_ktp = ?,
								foto_selfie = ?
							WHERE
								id = ?;`
)
