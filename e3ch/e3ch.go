package e3ch

import (
	"crypto/tls"

	"github.com/HughNian/nmid-discovery/conf"
	client "github.com/soyking/e3ch"
	"go.etcd.io/etcd/client/pkg/v3/transport"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewE3chClient(config *conf.Config) (*client.EtcdHRCHYClient, error) {
	var tlsConfig *tls.Config
	var err error
	if config.CertFile != "" && config.KeyFile != "" && config.CAFile != "" {
		tlsInfo := transport.TLSInfo{
			CertFile:           config.CertFile,
			KeyFile:            config.KeyFile,
			TrustedCAFile:      config.CAFile,
			InsecureSkipVerify: config.SkipVerifyTLS,
		}
		tlsConfig, err = tlsInfo.ClientConfig()
		if err != nil {
			return nil, err
		}
	}

	clt, err := clientv3.New(clientv3.Config{
		Endpoints:   config.EtcdEndPoints,
		Username:    config.EtcdUsername,
		Password:    config.EtcdPassword,
		TLS:         tlsConfig,
		DialTimeout: config.EtcdDialTimeout,
	})
	if err != nil {
		return nil, err
	}

	client, err := client.New(clt, config.EtcdRootKey, config.DirValue)
	if err != nil {
		return nil, err
	}
	return client, client.FormatRootKey()
}

func CloneE3chClient(username, password string, client *client.EtcdHRCHYClient) (*client.EtcdHRCHYClient, error) {
	clt, err := clientv3.New(clientv3.Config{
		Endpoints: client.EtcdClient().Endpoints(),
		Username:  username,
		Password:  password,
	})
	if err != nil {
		return nil, err
	}
	return client.Clone(clt), nil
}
