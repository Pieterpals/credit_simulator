# Credit Simulator

## Overview
Credit simulator for vehicle loans supporting cars and motorcycles with dynamic interest rate calculations.

## Prerequisites
- Go 1.20+
- Docker (optional)

## Installation

### Clone Repository
```bash
git clone <repository-url>
cd credit-simulator
```

### Build
```bash
go mod download
go build -o credit_simulator
```

## Running the Application

### Interactive Mode
```bash
./credit_simulator
```

## Input Requirements
- Vehicle Type: Motor/Mobil (case-insensitive)
- Vehicle Condition: Baru/Bekas
- Vehicle Year: 4-digit year
- Total Loan Amount: ≤ 100mio
- Loan Tenor: 1-6 years
- Down Payment Rules:
  - New Vehicle: ≥ 35% of loan amount
  - Used Vehicle: ≥ 25% of loan amount

## Testing
```bash
go test ./...
```

## Docker
### Build
```bash
docker build -t credit-simulator .
```

### Run
```bash
docker run credit-simulator
```

## CI/CD
GitHub Actions configured for:
- Unit testing
- Build verification
- Docker image publishing

## License
[Specify License]
