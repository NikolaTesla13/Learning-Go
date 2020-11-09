@ECHO off

echo Building routes/routes.go
go build src/routes/routes.go

echo Installing modules
go install

echo Running main.go
go run src/main.go