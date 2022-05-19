package config

type config struct {
	Database Database `yaml:"database"`
	GRPC     GRPC     `yaml:"grpc"`
	HTTP     HTTP     `yaml:"http"`
	Farmacy  Farmacy  `yaml:"farmacy"`
}

type Database struct {
	URL string `yaml:"url"`
}

type GRPC struct {
	URL string `yaml:"url"`
}

type HTTP struct {
	URL string `yaml:"url"`
}

type Farmacy struct {
	JWTKey string `yaml:"jwtkey"`
}
