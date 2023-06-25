package rewriter

import (
	"errors"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/dallasmarlow/kubeconfig_tunnel_rewriter/kubeconfig"
)

func findClusterEntry(config kubeconfig.KubeConfig, clusterName string) (int, error) {
	for idx, cluster := range config.Clusters {
		if cluster.Name == clusterName {
			return idx, nil
		}
	}

	return 0, errors.New("specified cluster name not found in config: " + clusterName)
}

func parseServerHost(server string) (string, error) {
	serverURL, err := url.Parse(server)
	if err != nil {
		return "", err
	}

	if strings.Contains(serverURL.Host, ":") {
		host, _, err := net.SplitHostPort(serverURL.Host)
		if err != nil {
			return "", err
		}

		return host, nil
	}

	return serverURL.Host, nil
}

func Process(
	config kubeconfig.KubeConfig,
	clusterName, tunnelAddr string,
	tunnelPort int) (kubeconfig.KubeConfig, error) {
	idx, err := findClusterEntry(config, clusterName)
	if err != nil {
		return kubeconfig.KubeConfig{}, err
	}

	cluster := config.Clusters[idx]
	if cluster.Cluster.Server == "" {
		return kubeconfig.KubeConfig{}, errors.New("specified cluster entry does not contain server URL")
	}

	serverHost, err := parseServerHost(cluster.Cluster.Server)
	if err != nil {
		return kubeconfig.KubeConfig{}, err
	}

	cluster.Cluster.Server = "https://" + tunnelAddr + ":" + strconv.Itoa(tunnelPort)
	cluster.Cluster.TlsServerName = serverHost
	config.Clusters[idx] = cluster

	return config, nil
}
