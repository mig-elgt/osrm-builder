package osrm

import (
	"os"
	"os/exec"

	osrmbuilder "github.com/mig-elgt/osrm-builder"
	"github.com/pkg/errors"
)

type osrm struct{}

func New() osrmbuilder.Builder {
	return &osrm{}
}

// Extract extracts a graph out of the OpenStreetMap base map.
func (o *osrm) Extract(osmPath string) error {
	cmd := exec.Command("osrm-extract", "-p", "/opt/car.lua", osmPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not extract osm data")
	}
	return nil
}

// Partition partitions a map graph recursively into cells.
func (o *osrm) Partition(osrmPath string) error {
	cmd := exec.Command("osrm-partition", osrmPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not perform partition map")
	}
	return nil
}

// Customize customizes the cells by calculating routing weights for all cells.
func (o *osrm) Customize(osrmPath string) error {
	cmd := exec.Command("osrm-customize", osrmPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not perform partition map")
	}
	return nil
}
