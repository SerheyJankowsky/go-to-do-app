package config

var Config = struct {
	DBHost string
	DBPort uint
	DBUser string
	DBPassword string
	DBName string
	JWTSecret string
}{
	DBHost: "localhost",
	DBPort: 5432,
	DBUser: "root",
	DBPassword: "secret",
	DBName : "api",
	JWTSecret: "alkdjflaie234lkkajdkslf92efo",
}

