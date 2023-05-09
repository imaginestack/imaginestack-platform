/*
Copyright 2019 The ImagineKube Authors.

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

package constants

const (
	APIVersion = "v1alpha1"

	KubeSystemNamespace           = "kube-system"
	OpenPitrixNamespace           = "openpitrix-system"
	KubesphereDevOpsNamespace     = "imaginekube-devops-system"
	IstioNamespace                = "istio-system"
	ImagineKubeMonitoringNamespace = "imaginekube-monitoring-system"
	ImagineKubeLoggingNamespace    = "imaginekube-logging-system"
	ImagineKubeNamespace           = "imaginekube-system"
	ImagineKubeControlNamespace    = "imaginekube-controls-system"
	PorterNamespace               = "porter-system"
	IngressControllerNamespace    = ImagineKubeControlNamespace
	AdminUserName                 = "admin"
	IngressControllerPrefix       = "imaginekube-router-"
	ImagineKubeConfigName          = "imaginekube-config"
	ImagineKubeConfigMapDataKey    = "imaginekube.yaml"

	ClusterNameLabelKey               = "imaginekube.com/cluster"
	NameLabelKey                      = "imaginekube.com/name"
	WorkspaceLabelKey                 = "imaginekube.com/workspace"
	NamespaceLabelKey                 = "imaginekube.com/namespace"
	DisplayNameAnnotationKey          = "imaginekube.com/alias-name"
	ChartRepoIdLabelKey               = "application.imaginekube.com/repo-id"
	ChartApplicationIdLabelKey        = "application.imaginekube.com/app-id"
	ChartApplicationVersionIdLabelKey = "application.imaginekube.com/app-version-id"
	CategoryIdLabelKey                = "application.imaginekube.com/app-category-id"
	DanglingAppCleanupKey             = "application.imaginekube.com/app-cleanup"
	CreatorAnnotationKey              = "imaginekube.com/creator"
	UsernameLabelKey                  = "imaginekube.com/username"
	DevOpsProjectLabelKey             = "imaginekube.com/devopsproject"
	KubefedManagedLabel               = "kubefed.io/managed"

	UserNameHeader = "X-Token-Username"

	AuthenticationTag = "Authentication"
	UserTag           = "User"
	GroupTag          = "Group"

	WorkspaceMemberTag     = "Workspace Member"
	DevOpsProjectMemberTag = "DevOps Project Member"
	NamespaceMemberTag     = "Namespace Member"
	ClusterMemberTag       = "Cluster Member"

	GlobalRoleTag        = "Global Role"
	ClusterRoleTag       = "Cluster Role"
	WorkspaceRoleTag     = "Workspace Role"
	DevOpsProjectRoleTag = "DevOps Project Role"
	NamespaceRoleTag     = "Namespace Role"

	OpenpitrixTag            = "OpenPitrix Resources"
	OpenpitrixAppInstanceTag = "App Instance"
	OpenpitrixAppTemplateTag = "App Template"
	OpenpitrixCategoryTag    = "Category"
	OpenpitrixAttachmentTag  = "Attachment"
	OpenpitrixRepositoryTag  = "Repository"
	OpenpitrixManagementTag  = "App Management"
	// HelmRepoMinSyncPeriod min sync period in seconds
	HelmRepoMinSyncPeriod = 180

	CleanupDanglingAppOngoing = "ongoing"
	CleanupDanglingAppDone    = "done"

	DevOpsCredentialTag  = "DevOps Credential"
	DevOpsPipelineTag    = "DevOps Pipeline"
	DevOpsWebhookTag     = "DevOps Webhook"
	DevOpsJenkinsfileTag = "DevOps Jenkinsfile"
	DevOpsScmTag         = "DevOps Scm"
	DevOpsJenkinsTag     = "Jenkins"

	ToolboxTag      = "Toolbox"
	RegistryTag     = "Docker Registry"
	GitTag          = "Git"
	TerminalTag     = "Terminal"
	MultiClusterTag = "Multi-cluster"

	WorkspaceTag     = "Workspace"
	NamespaceTag     = "Namespace"
	DevOpsProjectTag = "DevOps Project"
	UserResourceTag  = "User's Resources"

	NamespaceResourcesTag = "Namespace Resources"
	ClusterResourcesTag   = "Cluster Resources"
	ComponentStatusTag    = "Component Status"

	GatewayTag = "Gateway"

	NetworkTopologyTag = "Network Topology"

	ImagineKubeMetricsTag = "ImagineKube Metrics"
	ClusterMetricsTag    = "Cluster Metrics"
	NodeMetricsTag       = "Node Metrics"
	NamespaceMetricsTag  = "Namespace Metrics"
	PodMetricsTag        = "Pod Metrics"
	PVCMetricsTag        = "PVC Metrics"
	IngressMetricsTag    = "Ingress Metrics"
	ContainerMetricsTag  = "Container Metrics"
	WorkloadMetricsTag   = "Workload Metrics"
	WorkspaceMetricsTag  = "Workspace Metrics"
	ComponentMetricsTag  = "Component Metrics"
	CustomMetricsTag     = "Custom Metrics"

	LogQueryTag      = "Log Query"
	EventsQueryTag   = "Events Query"
	AuditingQueryTag = "Auditing Query"

	ClusterMetersTag   = "Cluster Meters"
	NodeMetersTag      = "Node Meters"
	WorkspaceMetersTag = "Workspace Meters"
	NamespaceMetersTag = "Namespace Meters"
	WorkloadMetersTag  = "Workload Meters"
	PodMetersTag       = "Pod Meters"
	ServiceMetricsTag  = "ServiceName Meters"

	ApplicationReleaseName = "meta.helm.sh/release-name"
	ApplicationReleaseNS   = "meta.helm.sh/release-namespace"

	ApplicationName    = "app.kubernetes.io/name"
	ApplicationVersion = "app.kubernetes.io/version"
	AlertingTag        = "Alerting"

	NotificationTag             = "Notification"
	NotificationSecretNamespace = "imaginekube-monitoring-federated"
	NotificationManagedLabel    = "notification.imaginekube.com/managed"

	DashboardTag = "Dashboard"
)

var (
	SystemNamespaces = []string{ImagineKubeNamespace, ImagineKubeLoggingNamespace, ImagineKubeMonitoringNamespace, OpenPitrixNamespace, KubeSystemNamespace, IstioNamespace, KubesphereDevOpsNamespace, PorterNamespace}
)
