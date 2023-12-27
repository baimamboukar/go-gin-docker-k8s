run:
	go run main.go
build:
	set GOOS=macos && go build -o main

compose:
	docker-compose up -d

test:
	hey -z 5m -q 5 -m GET -H "x-api-key:8113" http://127.0.0.1:8000/api/v1/startups/all