//make http request
go run .\main.go -origin='google.com'

//Flush Cache
go run main.go --clear-cache

//Fresh Mode
go run .\main.go -origin='google.com' --fresh