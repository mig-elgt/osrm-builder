package osrmbuilder

// MapFetcher defines and interface to fetch the OSM map data.
type MapFetcher interface {
	// url = http://download.geofabrik.de/asia/bhutan-latest.osm.pbf
	// output = build/map.osm.pbf

	// Fetch fetchs a file from geofabrik site that holds OSM data.
	// The data includes information which is irrelevant to routing,
	// such as positions of public waste baskets.
	Fetch(url, output string) error
}

// Builder describes an interface to hold a set of methods
// abouth the OSRM server pre-processing.
type Builder interface {
	// Extract extracts a graph out of the OpenStreetMap base map.
	Extract(osmPath string) error

	// Partition partitions a map graph recursively into cells.
	Partition(osrmPath string) error

	// Customize customizes the cells by calculating routing weights for all cells.
	Customize(osrmPath string) error
}

// Uploader describes an interface to upload files as result
// of OSRM builder.
type Uploader interface {
	Upload(bucket, root string) error
}
