package config

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)
	}
	return &config{
		db: &db{
			host: envMap["DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("load db port failed: %v", err)
				}
				return p
			}(),
			protocol: envMap["DB_PROTOCOL"],
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			database: envMap["DB_DATABASES"],
			sslMode:  envMap["DB_SSL_MODE"],
			maxConnections: func() int {
				maxConn, err := strconv.Atoi(envMap["DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("load Max Connection number failed: %v", err)
				}
				return maxConn
			}(),
		},
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("load app port failed: %v", err)
				}
				return p
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("load read timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_WRITE_TIMEOUT"])
				if err != nil {
					log.Fatalf("load write timeout s failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				bodyLimit, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("load app body limit failed: %v", err)
				}
				return bodyLimit
			}(),
			fileLimit: func() int {
				fileLimit, err := strconv.Atoi(envMap["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("load app file limit failed: %v", err)
				}
				return fileLimit
			}(),
			gcpbucket: envMap["APP_GCP_BUCKET"],
		},
		jwt: &jwt{
			adminKey:  envMap["JWT_ADMIN_KEY"],
			secretKey: envMap["JWT_SECRET_KEY"],
			apiKey:    envMap["JWT_API_KEY"],
			accessExpiresAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("load access expires at failed: %v", err)
				}
				return t
			}(),
			refreshExpiresAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("load refresh expires at failed: %v", err)
				}
				return t
			}(),
		},
	}
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	JWT() IJWTConfig
}

type config struct {
	db  *db
	app *app
	jwt *jwt
}

type IDbConfig interface {
	Url() string
	MaxOpenConns() int
}

func (c *config) Db() IDbConfig {
	return c.db
}

func (d *db) Url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.host,
		d.port,
		d.username,
		d.password,
		d.database,
		d.sslMode,
	)
}

func (d *db) MaxOpenConns() int {
	return d.maxConnections
}

type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

func (c *config) App() IAppConfig {
	return c.app
}

type IAppConfig interface {
	Url() string
	AppName() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GCPBucket() string
	Host() string
	Port() int
}

func (a *app) Url() string                 { return fmt.Sprintf("%s:%d", a.host, a.port) } // host:port
func (a *app) AppName() string             { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) FileLimit() int              { return a.fileLimit }
func (a *app) GCPBucket() string           { return a.gcpbucket }
func (a *app) Host() string                { return a.host }
func (a *app) Port() int                   { return a.port }

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
	gcpbucket    string
}

func (c *config) JWT() IJWTConfig {
	return c.jwt
}

type IJWTConfig interface {
	SecretKey() []byte
	AdminKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}

type jwt struct {
	secretKey        string
	apiKey           string
	adminKey         string
	accessExpiresAt  int //sec
	refreshExpiresAt int //sec
}

func (j *jwt) SecretKey() []byte          { return []byte(j.secretKey) }
func (j *jwt) AdminKey() []byte           { return []byte(j.adminKey) }
func (j *jwt) ApiKey() []byte             { return []byte(j.apiKey) }
func (j *jwt) AccessExpiresAt() int       { return j.accessExpiresAt }
func (j *jwt) RefreshExpiresAt() int      { return j.refreshExpiresAt }
func (j *jwt) SetJwtAccessExpires(t int)  { j.accessExpiresAt = t }
func (j *jwt) SetJwtRefreshExpires(t int) { j.refreshExpiresAt = t }
