#!bin/bash

source_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

echo "########### Creating S3 Buckets ###########"
aws s3 mb s3://mmonitoring \
	--endpoint-url=http://localhost:4566 \
	--profile=localstack

