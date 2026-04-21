module github.com/lopolopen/t-fiber-kafka-gorm

go 1.25.0

require (
	github.com/gofiber/fiber/v2 v2.52.12
	github.com/gofiber/swagger v1.1.1
	github.com/google/wire v0.7.0
	github.com/lopolopen/gap v0.1.0-beta.1
	github.com/lopolopen/gap/broker/xkafka v0.1.0-beta.1
	github.com/lopolopen/gap/storage/xgorm v0.1.0-beta.1
	github.com/lopolopen/shoot v0.7.1
	github.com/swaggo/swag v1.16.6
	go.uber.org/automaxprocs v1.6.0
	go.yaml.in/yaml/v3 v3.0.4
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/google/subcommands v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.18.5 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/pierrec/lz4/v4 v4.1.26 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/segmentio/kafka-go v0.4.50 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/swaggo/files/v2 v2.0.2 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	golang.org/x/tools v0.42.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

tool (
	github.com/google/wire/cmd/wire
	github.com/lopolopen/gap/cmd/gapc
	github.com/lopolopen/shoot/cmd/shoot
	github.com/swaggo/swag/cmd/swag
)

replace github.com/lopolopen/gap => ../../gap

replace github.com/lopolopen/gap/broker/xkafka => ../../gap/broker/xkafka

replace github.com/lopolopen/gap/storage/xgorm => ../../gap/storage/xgorm
