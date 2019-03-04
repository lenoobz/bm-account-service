// Package conf provides mapping and parsing configuration
package conf

// MongoConfig database config structure
type MongoConfig struct {
	Host       string
	Port       int64
	Username   string
	Password   string
	Dbname     string
	Collection string
}
