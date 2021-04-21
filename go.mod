module pikachu

go 1.13

replace google.golang.org/grpc v1.37.0 => google.golang.org/grpc v1.29.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/kutty-kumar/charminder v0.0.0-20210417173905-a130572e8976
	github.com/kutty-kumar/ho_oh v0.0.0-20210409062057-e437d5e35ca2
	github.com/prometheus/client_golang v1.8.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.4.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.37.0
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.8
)
