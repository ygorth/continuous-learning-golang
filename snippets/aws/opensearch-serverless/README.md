# Golang AWS OpenSearch Index Management

Golang snippets that provide examples of how to connect to AWS OpenSearch (formerly known as Amazon Elasticsearch) using IAM-based role authentication. 

openserverless_serverless_conn_create_index.go: how to create and delete an index with non-default settings.
openserverless_serverless_conn_get_index.go: how to get index metadata.

## Prerequisites

Before running the code, ensure you have the following:

1. An AWS account with appropriate IAM roles and permissions to access AWS OpenSearch.
2. Golang installed on your system.

## Setup

1. Set environment variables:

   Make sure to set the following environment variables before running the code:

   - `AWS_OPENSEARCH_ENDPOINT`: The endpoint URL of your AWS OpenSearch cluster.
   - `AWS_REGION`: The AWS region where your AWS OpenSearch cluster is located.

**Important**: Before running the code, ensure you understand the implications of creating and deleting indexes in your AWS OpenSearch cluster. These operations can have a significant impact on your data and should be executed with caution.
