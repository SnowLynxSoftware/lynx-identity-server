# Create Migration
```bash
./migrate create -ext sql -dir ./migrations -format unix initial
```

# Run Migrations
```bash
./migrate -database "mysql://root:p%40ssw0rd@tcp(10.0.0.24:3306)/identity" -path migrations/ up
```