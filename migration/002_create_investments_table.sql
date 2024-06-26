CREATE TABLE investments (
    id SERIAL PRIMARY KEY,
    loan_id INT NOT NULL REFERENCES loans(id),
    investor_id VARCHAR(50) NOT NULL,
    investment_amount NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
