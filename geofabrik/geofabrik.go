package geofabrik

import (
	"os"
	"os/exec"

	osrmbuilder "github.com/mig-elgt/osrm-builder"
	"github.com/pkg/errors"
)

type geofabrik struct{}

func New() osrmbuilder.MapFetcher {
	return &geofabrik{}
}

// url = http://download.geofabrik.de/asia/bhutan-latest.osm.pbf
// output = build/map.osm.pbf
// Fetch fetchs a file from geofabrik site that holds OSM data.
// The data includes information which is irrelevant to routing,
// such as positions of public waste baskets.
func (g *geofabrik) Fetch(url string, output string) error {
	cmd := exec.Command("curl", url, "--create-dirs", "-o", output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not fetch file")
	}
	return nil
}
