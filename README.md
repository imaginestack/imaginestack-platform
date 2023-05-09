<p align="center">
<a href="https://imaginekube.com/"><img src="docs/images/logo-128.png" alt="banner" width="200px"></a>
</p>

<p align="center">
<b>The container platform tailored for <i>Kubernetes multi-cloud, datacenter, and edge</i> management</b>
</p>

----

## ImagineKube

> English | (README.md)

[ImagineKube](https://imaginekube.com/) is a **distributed operating system for cloud-native application management**, using [Kubernetes](https://kubernetes.io) as its kernel. It provides a plug-and-play architecture, allowing third-party applications to be seamlessly integrated into its ecosystem. ImagineKube is also a multi-tenant container platform with full-stack automated IT operation and streamlined DevOps workflows. It provides developer-friendly wizard web UI, helping enterprises to build out a more robust and feature-rich platform, which includes most common functionalities needed for enterprise Kubernetes strategy, see [Feature List](#features) for details.


Demo environment

 ImagineKube provides you with free, stable, and out-of-the-box managed cluster service. After registration and login, you can easily create a K3s cluster with ImagineKube installed in only 5 seconds and experience feature-rich ImagineKube.


Features
<details>
  <summary><b>🕸 Provisioning Kubernetes Cluster</b></summary>
  Support deploy Kubernetes on any infrastructure, support online and air-gapped installation. <a href="https://imaginekube.com/docs/installing-on-linux/introduction/intro/">Learn more</a>.
  </details>
<details>
  <summary><b>🔗 Kubernetes Multi-cluster Management</b></summary>
  Provide a centralized control plane to manage multiple Kubernetes clusters, and support the ability to propagate an app to multiple K8s clusters across different cloud providers.
  </details>
<details>
  <summary><b>🤖 Kubernetes DevOps</b></summary>
  Provide GitOps-based CD solutions and use Argo CD to provide the underlying support, collecting CD status information in real time. With the mainstream CI engine Jenkins integrated, DevOps has never been easier. <a href="https://imaginekube.com/devops/">Learn more</a>.
  </details>
<details>
  <summary><b>🔎 Cloud Native Observability</b></summary>
  Multi-dimensional monitoring, events and auditing logs are supported; multi-tenant log query and collection, alerting and notification are built-in. <a href="https://imaginekube.com/observability/">Learn more</a>.
  </details>
<details>
  <summary><b>🧩 Service Mesh (Istio-based)</b></summary>
  Provide fine-grained traffic management, observability and tracing for distributed microservice applications, provides visualization for traffic topology. <a href="https://imaginekube.com/service-mesh/">Learn more</a>.
  </details>
<details>
  <summary><b>💻 App Store</b></summary>
  Provide an App Store for Helm-based applications, and offer application lifecycle management on Kubernetes platform. <a href="https://imaginekube.com/docs/pluggable-components/app-store/">Learn more</a>.
  </details>
<details>
  <summary><b>💡 Edge Computing Platform</b></summary>
  ImagineKube integrates <a href="https://kubeedge.io/en/">KubeEdge</a> to enable users to deploy applications on the edge devices and view logs and monitoring metrics of them on the console. <a href="https://imaginekube.com/docs/pluggable-components/kubeedge/">Learn more</a>.
  </details>
<details>
  <summary><b>📊 Metering and Billing</b></summary>
  Track resource consumption at different levels on a unified dashboard, which helps you make better-informed decisions on planning and reduce the cost. <a href="https://imaginekube.com/docs/toolbox/metering-and-billing/view-resource-consumption/">Learn more</a>.
  </details>
<details>
  <summary><b>🗃 Support Multiple Storage and Networking Solutions</b></summary>
  <li>Support GlusterFS, CephRBD, NFS, LocalPV solutions, and provide CSI plugins to consume storage from multiple cloud providers.</li><li>Provide Load Balancer Implementation <a href="https://github.com/imaginekube/openelb">Open