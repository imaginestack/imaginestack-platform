/*
Copyright 2023 The ImagineKube Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v2"

	networkv1alpha1 "imaginekube.com/api/network/v1alpha1"

	"imaginekube.com/imaginekube/pkg/apiserver/authentication"
	"imaginekube.com/imaginekube/pkg/apiserver/authentication/oauth"
	"imaginekube.com/imaginekube/pkg/apiserver/authorization"
	"imaginekube.com/imaginekube/pkg/models/terminal"
	"imaginekube.com/imaginekube/pkg/simple/client/alerting"
	"imaginekube.com/imaginekube/pkg/simple/client/auditing"
	"imaginekube.com/imaginekube/pkg/simple/client/cache"
	"imaginekube.com/imaginekube/pkg/simple/client/devops/jenkins"
	"imaginekube.com/imaginekube/pkg/simple/client/edgeruntime"
	"imaginekube.com/imaginekube/pkg/simple/client/events"
	"imaginekube.com/imaginekube/pkg/simple/client/gateway"
	"imaginekube.com/imaginekube/pkg/simple/client/gpu"
	"imaginekube.com/imaginekube/pkg/simple/client/k8s"
	"imaginekube.com/imaginekube/pkg/simple/client/kubeedge"
	"imaginekube.com/imaginekube/pkg/simple/client/ldap"
	"imaginekube.com/imaginekube/pkg/simple/client/logging"
	"imaginekube.com/imaginekube/pkg/simple/client/metering"
	"imaginekube.com/imaginekube/pkg/simple/client/monitoring/prometheus"
	"imaginekube.com/imaginekube/pkg/simple/client/multicluster"
	"imaginekube.com/imaginekube/pkg/simple/client/network"
	"imaginekube.com/imaginekube/pkg/simple/client/notification"
	"imaginekube.com/imaginekube/pkg/simple/client/openpitrix"
	"imaginekube.com/imaginekube/pkg/simple/client/s3"
	"imaginekube.com/imaginekube/pkg/simple/client/servicemesh"
	"imaginekube.com/imaginekube/pkg/simple/client/sonarqube"
)

func newTestConfig() (*Config, error) {
	var conf = &Config{
		DevopsOptions: &jenkins.Options{
			Host:           "http://ks-devops.imaginekube-devops-system.svc",
			Username:       "jenkins",
			Password:       "imaginekube",
			MaxConnections: 10,
		},
		SonarQubeOptions: &sonarqube.Options{
			Host:  "http://sonarqube.imaginekube-devops-system.svc",
			Token: "ABCDEFG",
		},
		KubernetesOptions: &k8s.KubernetesOptions{
			KubeConfig: "/Users/zry/.kube/config",
			Master:     "https://127.0.0.1:6443",
			QPS:        1e6,
			Burst:      1e6,
		},
		ServiceMeshOptions: &servicemesh.Options{
			IstioPilotHost:            "http://istio-pilot.istio-system.svc:9090",
			JaegerQueryHost:           "http://jaeger-query.istio-system.svc:80",
			ServicemeshPrometheusHost: "http://prometheus-k8s.imaginekube-monitoring-system.svc",
		},
		LdapOptions: &ldap.Options{
			Host:            "http://openldap.imaginekube-system.svc",
			ManagerDN:       "cn=admin,dc=example,dc=org",
			ManagerPassword: "P@88w0rd",
			UserSearchBase:  "ou=Users,dc=example,dc=org",
			GroupSearchBase: "ou=Groups,dc=example,dc=org",
			InitialCap:      10,
			MaxCap:          100,
			PoolName:        "ldap",
		},
		CacheOptions: &cache.Options{
			Type:    "redis",
			Options: map[string]interface{}{},
		},
		S3Options: &s3.Options{
			Endpoint:        "http://minio.openpitrix-system.svc",
			Region:          "us-east-1",
			DisableSSL:      false,
			ForcePathStyle:  false,
			AccessKeyID:     "ABCDEFGHIJKLMN",
			SecretAccessKey: "OPQRSTUVWXYZ",
			SessionToken:    "abcdefghijklmn",
			Bucket:          "ssss",
		},
		OpenPitrixOptions: &openpitrix.Options{
			S3Options: &s3.Options{
				Endpoint:        "http://minio.openpitrix-system.svc",
				Region:          "",
				DisableSSL:      false,
				ForcePathStyle:  false,
				AccessKeyID:     "ABCDEFGHIJKLMN",
				SecretAccessKey: "OPQRSTUVWXYZ",
				SessionToken:    "abcdefghijklmn",
				Bucket:          "app",
			},
			ReleaseControllerOptions: &openpitrix.ReleaseControllerOptions{
				MaxConcurrent: 10,
				WaitTime:      30 * time.Second,
			},
		},
		NetworkOptions: &network.Options{
			EnableNetworkPolicy: true,
			NSNPOptions: network.NSNPOptions{
				AllowedIngressNamespaces: []string{},
			},
			WeaveScopeHost: "weave-scope-app.weave",
			IPPoolType:     networkv1alpha1.IPPoolTypeNone,
		},
		MonitoringOptions: &prometheus.Options{
			Endpoint: "http://prometheus.imaginekube-monitoring-system.svc",
		},
		LoggingOptions: &logging.Options{
			Host:        "http://elasticsearch-logging.imaginekube-logging-system.svc:9200",
			IndexPrefix: "elk",
			Version:     "6",
		},
		AlertingOptions: &alerting.Options{
			Endpoint: "http://alerting-client-server.imaginekube-alerting-system.svc:9200/api",

			PrometheusEndpoint:       "http://prometheus-operated.imaginekube-monitoring-system.svc",
			ThanosRulerEndpoint:      "http://thanos-ruler-operated.imaginekube-monitoring-system.svc",
			ThanosRuleResourceLabels: "thanosruler=thanos-ruler,role=thanos-alerting-rules",
		},
		NotificationOptions: &notification.Options{
			Endpoint: "http://notification.imaginekube-alerting-system.svc:9200",
		},
		AuthorizationOptions: authorization.NewOptions(),
		AuthenticationOptions: &authentication.Options{
			AuthenticateRateLimiterMaxTries: 5,
			AuthenticateRateLimiterDuration: 30 * time.Minute,
			JwtSecret:                       "xxxxxx",
			LoginHistoryMaximumEntries:      100,
			MultipleLogin:                   false,
			OAuthOptions: &oauth.Options{
				Issuer:            oauth.DefaultIssuer,
				IdentityProviders: []oauth.IdentityProviderOptions{},
				Clients: []oauth.Client{{
					Name:                         "imaginekube-console-client",
					Secret:                       "xxxxxx-xxxxxx-xxxxxx",
					RespondWithChallenges:        true,
					RedirectURIs:                 []string{"http://ks-console.imaginekube-system.svc/oauth/token/implicit"},
					GrantMethod:                  oauth.GrantHandlerAuto,
					AccessTokenInactivityTimeout: nil,
				}},
				AccessTokenMaxAge:            time.Hour * 24,
				AccessTokenInactivityTimeout: 0,
			},
		},
		MultiClusterOptions: multicluster.NewOptions(),
		EventsOptions: &events.Options{
			Host:        "http://elasticsearch-logging-data.imaginekube-logging-system.svc:9200",
			IndexPrefix: "ks-logstash-events",
			Version:     "6",
		},
		AuditingOptions: &auditing.Options{
			Host:        "http://elasticsearch-logging-data.imaginekube-logging-system.svc:9200",
			IndexPrefix: "ks-logstash-auditing",
			Version:     "6",
		},
		KubeEdgeOptions: &kubeedge.Options{
			Endpoint: "http://edge-watcher.kubeedge.svc/api/",
		},
		EdgeRuntimeOptions: &edgeruntime.Options{
			Endpoint: "http://edgeservice.kubeedge.svc/api/",
		},
		MeteringOptions: &metering.Options{
			RetentionDay: "7d",
		},
		GatewayOptions: &gateway.Options{
			WatchesPath: "/etc/imaginekube/watches.yaml",
			Namespace:   "imaginekube-controls-system",
		},
		GPUOptions: &gpu.Options{
			Kinds: []gpu.GPUKind{},
		},
		TerminalOptions: &terminal.Options{
			Image:   "alpine:3.15",
			Timeout: 600,
		},
	}
	return conf, nil
}

func saveTestConfig(t *testing.T, conf *Config) {
	content, err := yaml.Marshal(conf)
	if err != nil {
		t.Fatalf("error marshal config. %v", err)
	}
	err = os.WriteFile(fmt.Sprintf("%s.yaml", defaultConfigurationName), content, 0640)
	if err != nil {
		t.Fatalf("error write configuration file, %v", err)
	}
}

func cleanTestConfig(t *testing.T) {
	file := fmt.Sprintf("%s.yaml", defaultConfigurationName)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Log("file not exists, skipping")
		return
	}

	err := os.Remove(file)
	if err != nil {
		t.Fatalf("remove %s file failed", file)
	}

}

func TestGet(t *testing.T) {
	conf, err := newTestConfig()
	if err != nil {
		t.Fatal(err)
	}
	saveTestConfig(t, conf)
	defer cleanTestConfig(t)

	conf2, err := TryLoadFromDisk()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(conf, conf2); diff != "" {
		t.Fatal(diff)
	}
}

func TestStripEmptyOptions(t *testing.T) {
	var config Config

	config.CacheOptions = &cache.Options{Type: ""}
	config.DevopsOptions = &jenkins.Options{Host: ""}
	config.MonitoringOptions = &prometheus.Options{Endpoint: ""}
	config.SonarQubeOptions = &sonarqube.Options{Host: ""}
	config.LdapOptions = &ldap.Options{Host: ""}
	config.NetworkOptions = &network.Options{
		EnableNetworkPolicy: false,
		WeaveScopeHost:      "",
		IPPoolType:          networkv1alpha1.IPPoolTypeNone,
	}
	config.ServiceMeshOptions = &servicemesh.Options{
		IstioPilotHost:            "",
		ServicemeshPrometheusHost: "",
		JaegerQueryHost:           "",
	}
	config.S3Options = &s3.Options{
		Endpoint: "",
	}
	config.AlertingOptions = &alerting.Options{
		Endpoint:            "",
		PrometheusEndpoint:  "",
		ThanosRulerEndpoint: "",
	}
	config.LoggingOptions = &logging.Options{Host: ""}
	config.NotificationOptions = &notification.Options{Endpoint: ""}
	config.MultiClusterOptions = &multicluster.Options{Enable: false}
	config.EventsOptions = &events.Options{Host: ""}
	config.AuditingOptions = &auditing.Options{Host: ""}
	config.KubeEdgeOptions = &kubeedge.Options{Endpoint: ""}
	config.EdgeRuntimeOptions = &edgeruntime.Options{Endpoint: ""}

	config.stripEmptyOptions()

	if config.CacheOptions != nil ||
		config.DevopsOptions != nil ||
		config.MonitoringOptions != nil ||
		config.SonarQubeOptions != nil ||
		config.LdapOptions != nil ||
		config.NetworkOptions != nil ||
		config.ServiceMeshOptions != nil ||
		config.S3Options != nil ||
		config.AlertingOptions != nil ||
		config.LoggingOptions != nil ||
		config.NotificationOptions != nil ||
		config.MultiClusterOptions != nil ||
		config.EventsOptions != nil ||
		config.AuditingOptions != nil ||
		config.KubeEdgeOptions != nil ||
		config.EdgeRuntimeOptions != nil {
		t.Fatal("config stripEmptyOptions failed")
	}
}
