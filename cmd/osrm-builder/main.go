package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mig-elgt/osrm-builder/geofabrik"
	"github.com/mig-elgt/osrm-builder/gstorage"
	"github.com/mig-elgt/osrm-builder/osrm"
	"github.com/sirupsen/logrus"
)

var (
	bucketName   = os.Getenv("BUCKET_NAME")
	geofabrikURL = os.Getenv("GEOFABRIK_URL")
)

const osrmBuildDataPath = "/tmp/osrm-data"

func main() {
	// Fetch OSM data from geofabrik
	g := geofabrik.New()
	logrus.Infof("Fetching OSM data file: %v", geofabrikURL)
	if err := g.Fetch(geofabrikURL, osrmBuildDataPath+"/map.osm.pbf"); err != nil {
		logrus.Fatalf("failed to get osm map file: %v", err)
	}

	// OSRM pre-processing
	osm := osrm.New()
	logrus.Info("Start OSRM pre-processing")
	logrus.Info("Extracting data.")
	if err := osm.Extract(osrmBuildDataPath + "/map.osm.pbf"); err != nil {
		logrus.Fatalf("failed extract: %v", err)
	}
	logrus.Info("Partitioning graph.")
	if err := osm.Partition(osrmBuildDataPath + "/map.osm.pbf"); err != nil {
		logrus.Fatalf("failed partition: %v", err)
	}
	logrus.Info("Customizing cells.")
	if err := osm.Customize(osrmBuildDataPath + "/map.osm.pbf"); err != nil {
		logrus.Fatalf("failed customize: %v", err)
	}

	// Upload OSRM files
	store := gstorage.New()
	bucket := fmt.Sprintf("%v-v%v", bucketName, time.Now().UTC().Unix())
	logrus.Infof("Uploadling files to bucket %v", bucket)
	if err := store.Upload(bucket, osrmBuildDataPath); err != nil {
		logrus.Fatalf("failed upload: %v", err)
	}
	logrus.Infof("OSRM version %v created successfully", bucket)
}
