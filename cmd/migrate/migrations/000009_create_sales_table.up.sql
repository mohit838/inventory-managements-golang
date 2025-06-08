CREATE TABLE sales (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  invoice_number VARCHAR(50) NOT NULL UNIQUE,
  customer_id BIGINT DEFAULT NULL,              -- link to customers table
  employee_id BIGINT NOT NULL,                  -- who made the sale (FK to users)
  otp_code VARCHAR(10) DEFAULT NULL,            -- for OTP confirmation
  otp_verified_at DATETIME DEFAULT NULL,        -- when OTP was verified
  total_amount DECIMAL(10, 2) NOT NULL,         -- total after discount
  paid_amount DECIMAL(10, 2) NOT NULL,          -- amount paid by customer
  payment_method VARCHAR(50) DEFAULT 'cash',    -- later: card, upi, etc.
  is_active BOOLEAN DEFAULT TRUE,               -- soft disable of sale
  is_deleted BOOLEAN DEFAULT FALSE,             -- soft delete logic
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  
  FOREIGN KEY (employee_id) REFERENCES users(id),
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);
