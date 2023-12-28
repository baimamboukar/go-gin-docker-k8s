run:
	go run main.go
build:
	set GOOS=macos && go build -o main

compose:
	docker-compose up -d

test:
	hey -z 2m -q 20 -m GET -H "x-api-key:8113" http://localhost:8113/api/v1/startups/all