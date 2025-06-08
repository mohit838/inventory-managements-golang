CREATE TABLE roles (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE, -- 'admin', 'cashier', etc.
    code VARCHAR(50) DEFAULT NULL, -- optional slug-like alias
    is_active BOOLEAN DEFAULT TRUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_roles_is_active ON roles (is_active);

-- INSERT INTO roles (name, code) VALUES
--   ('Admin', 'admin'),
--   ('Employee', 'employee'),
--   ('Cashier', 'cashier');