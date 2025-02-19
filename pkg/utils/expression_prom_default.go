package utils

import (
	"fmt"
)

// todo: later we change these templates to configurable like prometheus-adapter
const (
	// WorkloadCpuUsageExprTemplate is used to query workload cpu usage by promql,  param is namespace,workload-name,duration str
	WorkloadCpuUsageExprTemplate = `sum(irate(container_cpu_usage_seconds_total{namespace="%s",pod=~"^%s-%s"}[%s]))`
	// WorkloadMemUsageExprTemplate is used to query workload mem usage by promql, param is namespace, workload-name
	WorkloadMemUsageExprTemplate = `sum(container_memory_working_set_bytes{namespace="%s",pod=~"^%s-%s"})`

	// following is node exporter metric for node cpu/memory usage
	// NodeCpuUsageExprTemplate is used to query node cpu usage by promql,  param is node name which prometheus scrape, duration str
	NodeCpuUsageExprTemplate = `sum(count(node_cpu_seconds_total{mode="idle",instance=~"(%s)(:\\d+)?"}) by (mode, cpu)) - sum(irate(node_cpu_seconds_total{mode="idle",instance=~"(%s)(:\\d+)?"}[%s]))`
	// NodeMemUsageExprTemplate is used to query node cpu memory by promql,  param is node name, node name which prometheus scrape
	NodeMemUsageExprTemplate = `sum(node_memory_MemTotal_bytes{instance=~"(%s)(:\\d+)?"} - node_memory_MemAvailable_bytes{instance=~"(%s)(:\\d+)?"})`

	// PodCpuUsageExprTemplate is used to query pod cpu usage by promql,  param is namespace,pod, duration str
	PodCpuUsageExprTemplate = `sum(irate(container_cpu_usage_seconds_total{container!="POD",namespace="%s",pod="%s"}[%s]))`
	// PodMemUsageExprTemplate is used to query pod cpu usage by promql,  param is namespace,pod
	PodMemUsageExprTemplate = `sum(container_memory_working_set_bytes{container!="POD",namespace="%s",pod="%s"})`

	// ContainerCpuUsageExprTemplate is used to query container cpu usage by promql,  param is namespace,pod,container duration str
	ContainerCpuUsageExprTemplate = `irate(container_cpu_usage_seconds_total{container!="POD",namespace="%s",pod=~"^%s-%s",container="%s"}[%s])`
	// ContainerMemUsageExprTemplate is used to query container cpu usage by promql,  param is namespace,pod,container
	ContainerMemUsageExprTemplate = `container_memory_working_set_bytes{container!="POD",namespace="%s",pod=~"^%s-%s",container="%s"}`

	CustumerExprTemplate = `sum(%s{%s})`
)

const (
	RegMatchesPodName = `[a-z0-9]+-[a-z0-9]{5}$`
)

func GetCustumerExpression(metricName string, labels string) string {
	return fmt.Sprintf(CustumerExprTemplate, metricName, labels)
}

func GetWorkloadCpuUsageExpression(namespace string, name string) string {
	return fmt.Sprintf(WorkloadCpuUsageExprTemplate, namespace, name, RegMatchesPodName, "3m")
}

func GetWorkloadMemUsageExpression(namespace string, name string) string {
	return fmt.Sprintf(WorkloadMemUsageExprTemplate, namespace, name, RegMatchesPodName)
}

func GetContainerCpuUsageExpression(namespace string, workloadName string, containerName string) string {
	return fmt.Sprintf(ContainerCpuUsageExprTemplate, namespace, workloadName, RegMatchesPodName, containerName, "3m")
}

func GetContainerMemUsageExpression(namespace string, workloadName string, containerName string) string {
	return fmt.Sprintf(ContainerMemUsageExprTemplate, namespace, workloadName, RegMatchesPodName, containerName)
}

func GetPodCpuUsageExpression(namespace string, name string) string {
	return fmt.Sprintf(PodCpuUsageExprTemplate, namespace, name, "3m")
}

func GetPodMemUsageExpression(namespace string, name string) string {
	return fmt.Sprintf(PodMemUsageExprTemplate, namespace, name)
}

func GetNodeCpuUsageExpression(nodeName string) string {
	return fmt.Sprintf(NodeCpuUsageExprTemplate, nodeName, nodeName, "3m")
}

func GetNodeMemUsageExpression(nodeName string) string {
	return fmt.Sprintf(NodeMemUsageExprTemplate, nodeName, nodeName)
}
