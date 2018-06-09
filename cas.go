package service

import (
	"time"

	"errors"

	"github.com/go-yaml/yaml"
	"github.com/gocql/gocql"
)

// Cassandra wrapper
type CassandraCfg struct {
	Cluster  []string `yaml:"cluster"`
	Keyspace string   `yaml:"keyspace"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	NumConns int      `yaml:"numConns"`
}

// cas defines a Cassandra session and
// it's configuration.
type cas struct {
	cfg     CassandraCfg
	Session *gocql.Session
}

// CasRows used for simple result sets not bound to a
// specific type
type CasRows []map[string]interface{}

func (c *cas) Query(query string) (CasRows, error) {
	// call external libs for business logic here
	q := c.Session.Query(query)
	err := q.Exec()
	if err != nil {
		return nil, err
	}

	itr := q.Iter()
	defer itr.Close()

	ret := make(CasRows, 0)

	for {
		// New map each iteration
		row := make(map[string]interface{})
		if !itr.MapScan(row) {
			break
		}
		ret = append(ret, row)
	}

	if err := itr.Close(); err != nil {
		return nil, err
	}

	return ret, nil
}

// CassandraFromCfg takes a configuration map
func CassandraFromCfg(cfg Cfg) (*cas, error) {
	cas := &cas{}
	ccfg := CassandraCfg{}

	if cfg["cassandra"] == nil {
		return cas, errors.New("no cassandra configuration found")
	}

	// get yaml
	d, err := yaml.Marshal(cfg["cassandra"])
	if err != nil {
		return cas, err
	}

	// use the yaml to hydrate the configuration
	err = yaml.Unmarshal(d, &ccfg)
	if err != nil {
		return cas, err
	}

	return Cassandra(ccfg)
}

// Cassandra produces a cassandra object with session
func Cassandra(cfg CassandraCfg) (*cas, error) {

	cluster := gocql.NewCluster(cfg.Cluster...)
	cluster.DisableInitialHostLookup = true
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.Compressor = &gocql.SnappyCompressor{}
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{NumRetries: 3}
	cluster.Consistency = gocql.LocalQuorum
	cluster.Timeout = 10 * time.Second

	if cfg.Username != "" && cfg.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: cfg.Username,
			Password: cfg.Password,
		}
	}

	cluster.Keyspace = cfg.Keyspace
	cluster.NumConns = cfg.NumConns

	session, err := cluster.CreateSession()

	return &cas{cfg: cfg, Session: session}, err
}
