# Open Graph Social Media Image Generator

This project generates Open Graph images for social media sharing. It uses a Go-based Lambda function to create images dynamically based on provided title and description parameters. The infrastructure is managed using AWS CDK.

## Architecture

The project consists of the following main components:

-   **Lambda Function (Go):**  Generates the Open Graph image. It uses the `gg` and `imaging` Go libraries to create and manipulate images. It downloads assets from S3 on cold start.
-   **Infrastructure as Code (AWS CDK):** Defines and deploys the AWS infrastructure required for the project, including the Lambda function, API Gateway, S3 bucket, and CloudFront distribution.
-   **S3 Bucket:** Stores assets such as background images and fonts used by the Lambda function.
-   **API Gateway:**  Provides an HTTP endpoint to trigger the Lambda function.
-   **CloudFront:**  A content delivery network (CDN) that caches the generated images for faster delivery and provides a custom domain with SSL certificate.

## Prerequisites

Before you can deploy this project, you need to have the following tools installed:

-   [Node.js](https://nodejs.org/en/) (v18 or later)
-   [AWS CDK Toolkit](https://docs.aws.amazon.com/cdk/v2/guide/cli.html)
-   [Go](https://go.dev/dl/) (v1.21 or later)
-   [AWS CLI](https://aws.amazon.com/cli/)
-   An AWS account

## Setup

1.  **Configure AWS Credentials:** Configure your AWS credentials using the AWS CLI.
    ```bash
    aws configure
    ```
2.  **Install Dependencies:** Navigate to the `infra` directory and install the necessary Node.js dependencies.
    ```bash
    cd infra
    npm install
    ```
3.  **Set Environment Variables:** Configure the following environment variables, either in your shell or in your CI/CD environment:
    -   `CDK_DEFAULT_ACCOUNT`: Your AWS account ID.
    -   `AWS_DEFAULT_REGION`: The AWS region you want to deploy to (e.g., `us-east-1`).
    -   `AWS_ACCESS_KEY_ID`: Your AWS access key ID.
    -   `AWS_SECRET_ACCESS_KEY`: Your AWS secret access key.
    You can set these variables in your `.bashrc` or `.zshrc` file, or pass them directly to the `make` command.

## Deployment

The project includes a `Makefile` to simplify the build and deployment process. To deploy the project, run the following command:

```bash
make deploy
```

This command will:

1. Build the Go Lambda function.
2. Package the Lambda function into a ZIP file.
3. Deploy the AWS infrastructure using CDK.

## Usage
After successful deployment, you can access the Open Graph image generator through the CloudFront distribution's domain name.

To generate an image, send a GET request to the API endpoint with the title and description query parameters:
```https://your-cloudfront-domain/prod?title=Your%20Page%20Title&description=Your%20Page%20Description```

Replace your-cloudfront-domain with the actual domain name of your CloudFront distribution.

## Customization
- Lambda Function: You can modify the Go code in the functions directory to customize the image generation process.
- Infrastructure: You can modify the CDK code in the open-graph-infra-stack.ts file to adjust the AWS infrastructure configuration.
- Assets: Replace the files in the S3 bucket to change the background, fonts, logo, etc.

## Makefile commands
- make build: Builds the Go Lambda function and packages it into a ZIP file.
- make deploy: Builds the Lambda function, installs the CDK dependencies, and deploys the infrastructure.

This file defines a GitHub Actions workflow that automates the build and deployment process. It is triggered on workflow dispatch.