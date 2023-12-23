run:
	gin run -p 8081 main.go
build:
	set GOOS=macos && go build -o main