# ECS Hello World API with Golang and Terraform

This repository demonstrates how to build a production-ready infrastructure for a Golang API using **AWS ECS (Elastic Container Service)** with **Fargate**, **Terraform**, and a fully integrated **CI/CD pipeline** powered by **GitHub Actions**. It aims to help developers facing similar challenges by providing a reference architecture for deploying APIs on AWS.

## Overview

The project includes the following components:

- **Golang API**: A RESTful API built with the [Gin Framework](https://github.com/gin-gonic/gin).
    - Endpoints:
        - `/health`: Health check endpoint.
        - `/hello`: Returns a simple "world" message.
- **Infrastructure**: Defined using **Terraform** to provision:
    - An AWS ECS Cluster running on **Fargate**.
    - An Application Load Balancer (ALB) for traffic distribution.
    - A VPC with public and private subnets.
    - Amazon ECR for container image storage.
    - CloudWatch for logging.
- **CI/CD Pipeline**: A GitHub Actions workflow that automates:
    1. Building and pushing the Docker image to Amazon ECR.
    2. Updating the ECS service with the new image.
    3. Deploying the updated service.

---

## Getting Started

### 1. Clone the Repository

Clone the project to your local environment:

```bash
git clone https://github.com/<your-username>/ecs-hello-world.git
cd ecs-hello-world
```

### 2. Run the Golang API Locally
To test the API locally:

```bash
go run cmd/main.go
```
Visit the following endpoints to verify that the API is working:

- Health Check: http://localhost:8080/health
- Hello World: http://localhost:8080/hello

### 3. Deploy the Infrastructure
   Use Terraform to provision the required AWS infrastructure.

#### Step 1: Initialize Terraform
Navigate to the Terraform directory and initialize the Terraform project:

```bash
terraform init
```

### Step 2: Preview the Deployment Plan

Review the resources that Terraform will create:

```bash
terraform plan
```

### Step 3: Apply the Configuration
Apply the Terraform configuration to deploy the infrastructure:

```bash
terraform apply
```

This step provisions:

- An ECS cluster and Fargate service.
- Networking components such as VPC, subnets, and security groups.
- Amazon ECR, IAM roles, and CloudWatch logs.

### 4. CI/CD Pipeline with GitHub Actions

This repository includes a GitHub Actions workflow (.github/workflows/ci-cd.yml) to automate the build, push, and deployment process.

#### Workflow Steps

1. When changes are pushed to the main branch, the workflow will:
   - Build the Docker image.
   - Push the image to Amazon ECR.
   - Update the ECS task definition.
   - Deploy the new version to the ECS service.

2. To enable the workflow, configure the following secrets in your GitHub repository:
    - **AWS_ACCESS_KEY_ID:** Your AWS access key ID.
    - **AWS_SECRET_ACCESS_KEY:** Your AWS secret access key.


### Project Structure
Below is the high-level structure of the repository:

```
├── cmd/main.go           # Golang API source code
├── terraform/            # Terraform configurations
│   ├── main.tf           # Main Terraform script
│   ├── variables.tf      # Variables for customization
│   └── providers.tf      # AWS provider configuration
├── .github/
│   └── workflows/
│       └── main.yml     # GitHub Actions CI/CD pipeline
└── README.md             # Project documentation
```

### Customization

#### Environment Variables
The following environment variables can be updated to customize the deployment:

- AWS_REGION: The AWS region where the infrastructure will be deployed (default: us-east-1).
- ECR_REPOSITORY: The name of the Amazon ECR repository for Docker images.
- CLUSTER_NAME: The name of the ECS cluster.
- SERVICE_NAME: The name of the ECS service.
- TASK_DEFINITION_NAME: The ECS task definition name.
- CONTAINER_NAME: The name of the container in the ECS task definition.

Update these variables in .github/workflows/ci-cd.yml and variables.tf as needed.

