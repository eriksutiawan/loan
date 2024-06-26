- golang version 1.18
- set env for database "postgresql"
  DB_USER="postgres"
  DB_NAME="test_backend"
  DB_SSL_MODE="disable"
  DB_PASS="asd"
- migration in folder migration
- go mod tidy
- go mod vendor

api curl 
- create loan
  curl --location 'http://localhost:8080/loans' \
  --header 'Content-Type: application/json' \
  --data '{
      "borrower_id": "2",
      "principal_amount": 10000,
      "rate": 5.5,
      "roi": 8.2
  }'

- get loan
  curl --location 'http://localhost:8080/loans/6'

- approve loan
  curl --location --request PUT 'http://localhost:8080/loans/6/approve' \
  --header 'Content-Type: application/json' \
  --data '{
      "approval_picture": "link_to_approval_picture",
      "field_validator_id": "validator123",
      "approval_date": "2024-06-26 12:00:00"
  }'

- disburse loan 
  curl --location --request PUT 'http://localhost:8080/loans/1/disburse' \
  --header 'Content-Type: application/json' \
  --data '{
      "signed_agreement_letter": "link_to_signed_agreement",
      "field_officer_id": "officer789666",
      "disbursement_date": "2024-06-26T12:00:00Z"
  }'

