#!/bin/bash

echo """
AWS_REGION=
AWS_SQS_ARN=
AWS_SQS_URL=
AWS_S3_BUCKET=

DB_USER=
DB_DATABASE=
DB_HOST=
DB_PORT=
DB_PASSWORD=

KEYCLOAK_REALM=
KEYCLOAK_CLIENT=
KEYCLOAK_SECRET=
""" >> "config.env"

echo "Sample env created in config.env. Please fill the blank variables"