# GoNext

GoNext is a modern full-stack starter kit designed to help you launch your idea in production as quickly as possible. It leverages Golang for the backend, ReactJS for the frontend, and PostgreSQL for the database. The project is optimized for both rapid local development (using Docker) and scalable production deployment (using AWS infrastructure).

## Technologies Used

- **Backend:** Golang (with Ent ORM)
- **Frontend:** ReactJS (Vite)
- **Database:** PostgreSQL
- **Dev Environment:** Docker Compose
- **Production Environment:** AWS (EC2, RDS, S3, CloudFront, ALB, Terraform)

## Project Structure

```
GoNext/
├── backend/         # Golang API server
├── frontend/        # ReactJS app (Vite)
├── infra/           # Terraform for AWS infrastructure
├── docker-compose.yml
```

## Getting Started

### Development (Docker)

1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd GoNext
   ```
2. **Configure environment variables:**
   - Copy `.env.example` to `.env` and fill in required values.
   - In `frontend/go-next` copy `.env.development.example` to `.env.development` and fill in required values.
3. **Start the stack:**
   ```sh
   docker-compose up --build
   ```
4. **Access the apps:**
   - Frontend: [http://localhost:3000](http://localhost:3000)
   - Backend: [http://localhost:8080](http://localhost:8080)

### Production (AWS)

1. **Configure AWS credentials** on your machine.
    - Ensure you have AWS CLI installed and configured with your credentials.
    - Set up your AWS environment variables or use a credentials file.
2. **Edit Terraform variables:**
   - Update `infra/terraform/variables.tfvars` with your desired values.
3. **Provision infrastructure:**
   ```sh
   cd infra/terraform
   terraform init
   terraform apply -var-file=variables.tfvars
   ```
4. **Build and deploy the backend and frontend:**
   - Certificate: Ensure you have an SSL certificate set up in AWS Certificate Manager (For external registrar domain).
   - Backend: Build Go binary and upload to EC2 and run the service.
   - Frontend: Build React app and upload to S3 bucket for CloudFront distribution.

## Features
- JWT authentication
- User management
- Secure password validation
- Modern React UI with protected routes
- Infrastructure as code (Terraform)
- Ready for CI/CD integration

## License
MIT
