module github.com/cheolgyu/stock-read-pub-api

go 1.16

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/cheolgyu/stock-write-common v0.0.0
	github.com/google/uuid v1.2.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/joho/godotenv v1.4.0
	github.com/julienschmidt/httprouter v1.3.0
)

replace (
	github.com/cheolgyu/stock-write-common v0.0.0 => ../stock-write-common
	github.com/cheolgyu/stock-write-model v0.0.0 => ../stock-write-model
)
