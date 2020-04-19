package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jamiealquiza/envy"
	client "github.com/philhug/go-automower/pkg/automower"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var MOWER_MODES = []string{"AUTO", "HOME"}
var MOWER_STATE = []string{"OK_CHARGING", "OK_CUTTING", "OK_LEAVING", "OK_SEARCHING", "PARKED_PARKED_SELECTED", "PARKED_TIMER", "OK_CUTTING_NOT_AUTO"}

func main() {
	var username = flag.String("username", "", "Username for AMC")
	var password = flag.String("password", "", "Password for AMC")
	envy.Parse("AUTOMOWER") // Expose environment variables.
	flag.Parse()
	if *username == "" || *password == "" {
		flag.Usage()
		os.Exit(1)
	}

	c, err := client.NewClientWithUserAndPassword(*username, *password)
	if err != nil {
		log.Panic(err)
	}

	_, err = c.Mowers()
	if err != nil {
		log.Panic(err)
	}

	// Called on each collector.Collect.
	stats := func() ([]client.Mower, error) {
		mowers, err := c.Mowers()
		if err != nil {
			return nil, fmt.Errorf("failed to open get mowers: %v", err)
		}

		return mowers, nil
	}

	// Make Prometheus client aware of our collector.
	co := newCollector(stats)
	prometheus.MustRegister(co)

	// Set up HTTP handler for metrics.
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	// Start listening for HTTP connections.
	const addr = ":8888"
	log.Printf("starting automower exporter on %q", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("cannot start automower exporter: %s", err)
	}

}

var _ prometheus.Collector = &collector{}

// A collector is a prometheus.Collector for Linux CPU stats.
type collector struct {
	// Possible metric descriptions.
	Info               *prometheus.Desc
	BetteryPercentage  *prometheus.Desc
	MowerOperatingMode *prometheus.Desc
	MowerStatus        *prometheus.Desc

	// A parameterized function used to gather metrics.
	stats func() ([]client.Mower, error)
}

// newCollector constructs a collector using a stats function.
func newCollector(stats func() ([]client.Mower, error)) prometheus.Collector {
	return &collector{
		Info: prometheus.NewDesc(
			// Name of the metric.
			"mower_info",
			// The metric's help text.
			"Is Mower alive",
			// The metric's variable label dimensions.
			[]string{"id", "name", "model", "status", "mode"},
			// The metric's constant label dimensions.
			nil,
		),
		BetteryPercentage: prometheus.NewDesc(
			// Name of the metric.
			"mower_battery_percentage",
			// The metric's help text.
			"Percentage of battery",
			// The metric's variable label dimensions.
			[]string{"id", "name", "model"},
			// The metric's constant label dimensions.
			nil,
		),
		MowerOperatingMode: prometheus.NewDesc(
			// Name of the metric.
			"mower_mode",
			// The metric's help text.
			"Current mower status",
			// The metric's variable label dimensions.
			[]string{"id", "name", "model", "mode"},
			// The metric's constant label dimensions.
			nil,
		),
		MowerStatus: prometheus.NewDesc(
			// Name of the metric.
			"mower_status",
			// The metric's help text.
			"Current mower status",
			// The metric's variable label dimensions.
			[]string{"id", "name", "model", "status"},
			// The metric's constant label dimensions.
			nil,
		),

		stats: stats,
	}
}

// Describe implements prometheus.Collector.
func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	// Gather metadata about each metric.
	ds := []*prometheus.Desc{
		c.Info,
		c.BetteryPercentage,
		c.MowerOperatingMode,
	}
	for range MOWER_STATE {
		ds = append(ds, c.MowerStatus)
	}

	for _, d := range ds {
		ch <- d
	}
}

// Collect implements prometheus.Collector.
func (c *collector) Collect(ch chan<- prometheus.Metric) {
	// Take a stats snapshot.  Must be concurrency safe.
	stats, err := c.stats()
	if err != nil {
		// If an error occurs, send an invalid metric to notify
		// Prometheus of the problem.
		ch <- prometheus.NewInvalidMetric(c.Info, err)
		ch <- prometheus.NewInvalidMetric(c.BetteryPercentage, err)
		for range MOWER_MODES {
			ch <- prometheus.NewInvalidMetric(c.MowerOperatingMode, err)
		}
		for range MOWER_STATE {
			ch <- prometheus.NewInvalidMetric(c.MowerStatus, err)
		}
		return
	}

	for _, s := range stats {
		connected := 0
		if s.Status.Connected {
			connected = 1
		}
		ch <- prometheus.MustNewConstMetric(
			c.Info,
			prometheus.GaugeValue,
			float64(connected),
			s.ID,
			s.Name,
			s.Model,
			s.Status.MowerStatus,
			s.Status.OperatingMode,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BetteryPercentage,
			prometheus.GaugeValue,
			float64(s.Status.BatteryPercent),
			s.ID,
			s.Name,
			s.Model,
		)

		for _, mode := range MOWER_MODES {
			var val float64 = 0
			if s.Status.OperatingMode == mode {
				val = 1
			}
			ch <- prometheus.MustNewConstMetric(
				c.MowerOperatingMode,
				prometheus.GaugeValue,
				val,
				s.ID,
				s.Name,
				s.Model,
				mode,
			)
		}

		for _, status := range MOWER_STATE {
			var val float64 = 0
			if s.Status.MowerStatus == status {
				val = 1
			}
			ch <- prometheus.MustNewConstMetric(
				c.MowerStatus,
				prometheus.GaugeValue,
				val,
				s.ID,
				s.Name,
				s.Model,
				status,
			)
		}
	}
}
