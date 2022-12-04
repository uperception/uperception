#!/bin/bash

echo "########### Creating S3 Bucket ###########"
awslocal s3 mb s3://uperception-storage

echo "########### Creating SQS Queue ###########"
awslocal sqs create-queue --queue-name uperception-queue