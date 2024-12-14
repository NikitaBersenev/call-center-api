CREATE TABLE IF NOT EXISTS calls (
    id SERIAL PRIMARY KEY,
    client_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(50),
    description TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'открыта',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_calls_status ON calls(status);
