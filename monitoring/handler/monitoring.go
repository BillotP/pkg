package handler

import (
	"fmt"

	"github.com/upfluence/goutils/Godeps/_workspace/src/github.com/upfluence/base/monitoring"
	"github.com/upfluence/goutils/monitoring/metric"
)

type MonitoringHandler struct {
	prefix  string
	metrics map[monitoring.MetricID]metric.Metric
}

func NewMonitoringHandler(
	prefix string,
	metrics map[monitoring.MetricID]metric.Metric,
) *MonitoringHandler {
	return &MonitoringHandler{prefix, metrics}
}

// Can't understand why thrift inteface declares a []string instead of
// []monitoring.MetricsID
func (m *MonitoringHandler) Collect(metrics []string) (
	monitoring.Metrics,
	error,
) {
	promises := map[monitoring.MetricID]<-chan float64{}
	results := monitoring.Metrics{}

	for _, id := range metrics {
		metricId := monitoring.MetricID(id) // <- what the F*** thrift
		if metric, ok := m.metrics[metricId]; ok {
			promises[metricId] = metric.Collect()
		} else {
			return nil, &monitoring.UnknownMetric{metricId}
		}
	}

	for id, promise := range promises {
		results[monitoring.MetricID(fmt.Sprintf("%s.%s", m.prefix, id))] = <-promise
	}

	return results, nil
}