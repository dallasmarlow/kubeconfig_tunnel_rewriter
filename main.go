package main

import (
	"errors"
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/util/homedir"

	"github.com/dallasmarlow/kubeconfig_tunnel_rewriter/kubeconfig"
	"github.com/dallasmarlow/kubeconfig_tunnel_rewriter/rewriter"
)

var (
	ClusterName, KubeConfigPath string

	clusterNameFlag = flag.String("cluster-name", "", "cluster name to rewrite configuration for")
	kubeconfigFlag  = flag.String("kubeconfig", "", "kubeconfig file path")
	tunnelAddrFlag  = flag.String("tunnel-addr", "127.0.0.1", "tunnel IP address or hostname")
	tunnelPortFlag  = flag.Int("tunnel-port", 8443, "tunnel TCP port number")
)

func main() {
	flag.Parse()

	if *clusterNameFlag == "" {
		log.Fatalln(errors.New("flag `-cluster-name` must not be empty"))
	} else {
		ClusterName = *clusterNameFlag
	}

	if *kubeconfigFlag == "" {
		if home := homedir.HomeDir(); home != "" {
			KubeConfigPath = filepath.Join(home, ".kube", "config")
		} else {
			log.Fatalln(errors.New("unable to detect home directory for default kubeconfig path"))
		}
	} else {
		KubeConfigPath = *kubeconfigFlag
	}

	config, err := kubeconfig.Load(KubeConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	processedConfig, err := rewriter.Process(
		config,
		ClusterName,
		*tunnelAddrFlag,
		*tunnelPortFlag,
	)
	if err != nil {
		log.Fatalln(err)
	}

	if err := kubeconfig.Write(processedConfig, KubeConfigPath); err != nil {
		log.Fatalln(err)
	}
}
