package tls

import (
	"github.com/konpyutaika/nifikop/api/v1"
	"github.com/konpyutaika/nifikop/pkg/util/clientconfig"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Tls interface {
	clientconfig.Manager
}

type tls struct {
	client     client.Client
	clusterRef v1.ClusterReference
}

func New(client client.Client, clusterRef v1.ClusterReference) Tls {
	return &tls{clusterRef: clusterRef, client: client}
}
