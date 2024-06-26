CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    borrower_id VARCHAR(50) NOT NULL,
    principal_amount NUMERIC(10, 2) NOT NULL,
    rate NUMERIC(5, 2) NOT NULL,
    roi NUMERIC(5, 2) NOT NULL,
    state VARCHAR(20) NOT NULL DEFAULT 'proposed',
    approval_picture TEXT,
    field_validator_id VARCHAR(50),
    approval_date DATE,
    signed_agreement_letter TEXT,
    field_officer_id VARCHAR(50),
    disbursement_date DATE,
    agreement_letter_link TEXT
);