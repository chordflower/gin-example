package app

import (
	"time"

	"emperror.dev/errors"
	"github.com/chordflower/gin_example/internal/utils"
)

// Configuration describes the server configuration
type Configuration struct {
	// The address to be used by this server
	Address string `json:"address" yaml:"address" toml:"address"`
	// Configuration related to databases
	Databases *Databases `json:"databases" yaml:"databases" toml:"databases"`
	// Some server parameters
	Parameters *Parameters `json:"parameters" yaml:"parameters" toml:"parameters"`
	// The port used by this server
	Port int64 `json:"port" yaml:"port" toml:"port"`
	// Parameters related to TLS configuration
	Security *Security `json:"security" yaml:"security" toml:"security"`
	// The type of address of this server
	Type Type `json:"type" yaml:"type" toml:"type"`
}

// NewWithDefaults creates a configuration with all defaults of the configuration, overwriting any existing value
func NewWithDefaults() *Configuration {
	return &Configuration{
		Address:    "127.0.0.1",
		Type:       Socket,
		Port:       20038,
		Databases:  newDatabaseWithDefaults(),
		Parameters: newParametersWithDefaults(),
		Security:   newSecurityWithDefaults(),
	}
}

// Validate is a method to validate if the configuration is correct or not
func (c *Configuration) Validate() error {
	v := utils.NewValidator()
	v.IsNotEmpty(string(c.Type), "The server type must not be empty")
	v.IsPort(c.Port, "The server port must be a valid port")
	return errors.WrapIf(errors.Combine(v.AllValid(), c.Security.Validate(), c.Databases.Validate(), c.Parameters.Validate()), "Validation error in configuration root")
}

// Databases contains the configuration related to databases
type Databases struct {
	// Configuration related to the relational database
	Relational *Relational `json:"relational" yaml:"relational" toml:"relational"`
	// Configuration related to the redis database
	Session *Session `json:"session" yaml:"session" toml:"session"`
}

func newDatabaseWithDefaults() *Databases {
	return &Databases{
		Relational: newRelationalWithDefaults(),
		Session:    newSessionWithDefaults(),
	}
}

// Validate is a method to validate if the configuration is correct or not
func (c *Databases) Validate() error {
	return errors.WrapIf(errors.Combine(c.Relational.Validate(), c.Session.Validate()), "Validation error in the database key")
}

// Relational contains the configuration related to the relational database
type Relational struct {
	// Location of the PEM certificate to use
	Certificate string `json:"certificate" yaml:"certificate" toml:"certificate"`
	// The database to use
	Database string `json:"database" yaml:"database" toml:"database"`
	// The host used by postgresql
	Host string `json:"host" yaml:"host" toml:"host"`
	// Location of the PEM key to use
	Key string `json:"key" yaml:"key" toml:"key"`
	// The username password to use
	Password string `json:"password" yaml:"password" toml:"password"`
	// The port used by the postgresql server
	Port uint64 `json:"port" yaml:"port" toml:"port"`
	// Whether or not to enable TLS
	TLS bool `json:"tls" yaml:"tls" toml:"tls"`
	// The type of host of the postgresql server
	Type Type `json:"type" yaml:"type" toml:"type"`
	// The username to use
	Username string `json:"username" yaml:"username" toml:"username"`
}

func newRelationalWithDefaults() *Relational {
	return &Relational{
		Certificate: "./postgresql-certificate.pem",
		Host:        "127.0.0.1",
		Database:    "example",
		Key:         "./postgresql-key.pem",
		Password:    "example123",
		Port:        5432,
		TLS:         false,
		Type:        Socket,
		Username:    "example",
	}
}

// Validate is a method to validate if the configuration is correct or not
func (c *Relational) Validate() error {
	v := utils.NewValidator()
	v.IsNotEmpty(c.Database, "The postgresql database must not be empty")
	v.IsPort(c.Port, "The postgresql port must be a valid port")
	v.Check(c.TLS && len(c.Certificate) != 0, "When TLS is active the certificate must not be blank")
	v.Check(c.TLS && len(c.Key) != 0, "When TLS is active the key must not be blank")
	v.IsNotEmpty(string(c.Type), "The postgresql type must not be empty")
	return errors.WrapIf(nil, "Validation error in the relational key")
}

// Session contains the configuration related to the redis database
type Session struct {
	// Location of the PEM certificate to use
	Certificate string `json:"certificate" yaml:"certificate" toml:"certificate"`
	// The database to use
	Database uint64 `json:"database" yaml:"database" toml:"database"`
	// The host used by redis
	Host string `json:"host" yaml:"host" toml:"host"`
	// Location of the PEM key to use
	Key string `json:"key" yaml:"key" toml:"key"`
	// The username password to use
	Password string `json:"password" yaml:"password" toml:"password"`
	// The port used by the redis server
	Port uint64 `json:"port" yaml:"port" toml:"port"`
	// Whether or not to enable TLS
	TLS bool `json:"tls" yaml:"tls" toml:"tls"`
	// The type of host of the redis server
	Type Type `json:"type" yaml:"type" toml:"type"`
	// The username to use
	Username string `json:"username" yaml:"username" toml:"username"`
}

func newSessionWithDefaults() *Session {
	return &Session{
		Certificate: "./redis-certificate.pem",
		Database:    12,
		Host:        "127.0.0.1",
		Key:         "./redis-key.pem",
		Password:    "example123",
		Port:        6789,
		TLS:         false,
		Type:        Socket,
		Username:    "example",
	}
}

// Validate is a method to validate if the configuration is correct or not
func (c *Session) Validate() error {
	v := utils.NewValidator()
	v.IsBetweenNumbers(c.Database, 0, 15, "The redis database must be between 0 and 15")
	v.IsPort(c.Port, "The redis port must be a valid port")
	v.Check(c.TLS && len(c.Certificate) != 0, "When TLS is active the certificate must not be blank")
	v.Check(c.TLS && len(c.Key) != 0, "When TLS is active the key must not be blank")
	v.IsNotEmpty(string(c.Type), "The redis type must not be empty")
	return errors.WrapIf(nil, "Validation error in the session key")
}

// Parameters contains some server parameters
type Parameters struct {
	// The idle timeout to be used by the server, in golang Duration string format
	IdleTimeout string `json:"idle_timeout" yaml:"idle_timeout" toml:"idle_timeout"`
	// The maximum size of the headers in bytes
	MaxHeaderBytes uint64 `json:"max_header_bytes" yaml:"max_header_bytes" toml:"max_header_bytes"`
	// The read header timeout to be used by the server (head only), in golang Duration string format
	ReadHeaderTimeout string `json:"read_header_timeout" yaml:"read_header_timeout" toml:"read_header_timeout"`
	// The read timeout to be used by the server (headers+body), in golang Duration string format
	ReadTimeout string `json:"read_timeout" yaml:"read_timeout" toml:"read_timeout"`
	// Where to place the uploaded files
	UploadFolder string `json:"upload_folder" yaml:"upload_folder" toml:"upload_folder"`
	// The write timeout to be used by the server, in golang Duration string format
	WriteTimeout string `json:"write_timeout" yaml:"write_timeout" toml:"write_timeout"`
}

func newParametersWithDefaults() *Parameters {
	return &Parameters{
		IdleTimeout:       "2s",
		MaxHeaderBytes:    2097152,
		ReadHeaderTimeout: "10s",
		ReadTimeout:       "100s",
		UploadFolder:      "./upload",
		WriteTimeout:      "2m",
	}
}

// Validate is a method to validate if the configuration is correct or not
func (c *Parameters) Validate() error {
	v := utils.NewValidator()
	v.IsDuration(c.IdleTimeout, "The idle timeout must be a valid golang duration")
	v.IsPositive(c.MaxHeaderBytes, "The maximum header bytes must be a positive number")
	v.IsDuration(c.ReadHeaderTimeout, "The read header timeout must be a valid golang duration")
	v.IsDuration(c.ReadTimeout, "The read timeout must be a valid golang duration")
	v.IsNotEmpty(c.UploadFolder, "The upload folder must not be empty")
	v.IsDuration(c.WriteTimeout, "The write timeout must be a valid golang duration")
	return errors.WrapIf(nil, "Validation error in the parameters key")
}

// IdleTimeoutDuration is a method that returns the idle timeout as a duration
func (c *Parameters) IdleTimeoutDuration() (ret time.Duration, err error) {
	ret, err = time.ParseDuration(c.IdleTimeout)
	return
}

// ReadHeaderTimeoutDuration is a method that returns the read header timeout as a duration
func (c *Parameters) ReadHeaderTimeoutDuration() (ret time.Duration, err error) {
	ret, err = time.ParseDuration(c.ReadHeaderTimeout)
	return
}

// ReadTimeoutDuration is a method that returns the read timeout as a duration
func (c *Parameters) ReadTimeoutDuration() (ret time.Duration, err error) {
	ret, err = time.ParseDuration(c.ReadTimeout)
	return
}

// WriteTimeoutDuration is a method that returns the write timeout as a duration
func (c *Parameters) WriteTimeoutDuration() (ret time.Duration, err error) {
	ret, err = time.ParseDuration(c.WriteTimeout)
	return
}

// Security contains the parameters related to TLS configuration
type Security struct {
	// Location of the PEM certificate to use
	Certificate string `json:"certificate" yaml:"certificate" toml:"certificate"`
	// Location of the PEM key to use
	Key string `json:"key" yaml:"key" toml:"key"`
	// Whether or not to enable TLS
	TLS bool `json:"tls" yaml:"tls" toml:"tls"`
}

func newSecurityWithDefaults() *Security {
	return &Security{
		TLS:         false,
		Certificate: "./certificate.pem",
		Key:         "./key.pem",
	}
}

// Validate is a method to validate if the configuration is correct or not
func (c *Security) Validate() error {
	v := utils.NewValidator()
	v.Check(c.TLS && len(c.Certificate) != 0, "When TLS is active the certificate must not be blank")
	v.Check(c.TLS && len(c.Key) != 0, "When TLS is active the key must not be blank")
	return errors.WrapIf(v.AllValid(), "Validation error in the security key")
}

// Type is the address type
type Type string

const (
	// Socket describes that the address is a valid TCP address
	Socket Type = "socket"
	// Unix describes that the address is the path to a local unix socket
	Unix Type = "unix"
)
