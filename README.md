# OSRM Builder

OSRM Builder implements Multi-Level Dijkstra (MLD) pre-processing pipeline to build a OSRM Server using Docker.

## Features

* Fetch OpenStreetMap extracts (OSM Data file) from [Geofabrik](http://download.geofabrik.de/).
* Pre-process the extract with the car profile.
* Upload pre-process files to Google Storage.

# Install using Docker

You will need to install Docker and create a Google Service Account to use a storage in order to upload the OSRM Builder files.

```
$ docker build --no-cache osrm/osrm-server-builder:v1 .
```

## Usage & Example

```
$ 
docker run -v /tmp:/secrets --env GEOFABRIK_URL=http://download.geofabrik.de/europe/monaco-latest.osm.pbf --env BUCKET_NAME=osrm-monaco-server --env GOOGLE_APPLICATION_CREDENTIALS=/secrets/services_account.json osrm/osrm-server-builder:v1

time="2023-01-13T03:41:22Z" level=info msg="Fetching OSM data file: http://download.geofabrik.de/europe/monaco-latest.osm.pbf"

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
 92  477k   92  439k    0     0   253k      0  0:00:01  0:00:01 --:--:--  254k
100  477k  100  477k    0     0   271k      0  0:00:01  0:00:01 --:--:--  271k

time="2023-01-13T03:41:24Z" level=info msg="Start OSRM pre-processing"
time="2023-01-13T03:41:24Z" level=info msg="Extracting data."

[info] Parsed 0 location-dependent features with 0 GeoJSON polygons
[info] Using script /opt/car.lua
[info] Input file: map.osm.pbf
[info] Profile: car.lua
[info] Threads: 5
[info] Parsing in progress..
[info] input file generated by osmium/1.14.0
[info] timestamp: 2023-01-11T21:21:07Z
[info] Using profile api version 4
[info] Found 3 turn restriction tags:
[info]   motorcar
[info]   motor_vehicle
[info]   vehicle
[info] Parse relations ...
[info] Parse ways and nodes ...
[info] Using profile api version 4
[info] Using profile api version 4
[info] Using profile api version 4
[info] Using profile api version 4
[info] Parsing finished after 0.881897 seconds
[info] Raw input contains 28092 nodes, 4538 ways, and 57 relations, 43 restrictions
[info] Sorting used nodes        ... ok, after 0.002391s
[info] Erasing duplicate nodes   ... ok, after 9.9e-05s
[info] Sorting all nodes         ... ok, after 0.001275s
[info] Building node id map      ... ok, after 0.000149s
[info] Confirming/Writing used nodes     ... ok, after 0.002081s
[info] Writing barrier nodes     ... ok, after 0s
[info] Writing traffic light nodes     ... ok, after 0s
[info] Processed 5783 nodes
[info] Sorting edges by start    ... ok, after 0.002708s
[info] Setting start coords      ... ok, after 0.000754s
[info] Sorting edges by target   ... ok, after 0.002328s
[info] Computing edge weights    ... ok, after 0.009259s
[info] Sorting edges by renumbered start ... ok, after 0.004116s
[info] Writing used edges       ... ok, after 0.000873s -- Processed 6035 edges
[info] Writing way meta-data     ... ok, after 0.000257s -- Metadata contains << 1010 entries.
[info] Sorting used ways         ... ok, after 0.000802s
[info] Collecting start/end information on 0 maneuver overrides...ok, after 2.8e-05s
[info] Collecting start/end information on 0 maneuver overrides...ok, after 8e-06s
[info] Collecting start/end information on 43 restrictions...ok, after 0.000327s
[info] Collecting start/end information on 43 restrictions...ok, after 0.000545s
[info] writing street name index ... ok, after 0.001496s
[info] extraction finished after 0.926101s
[info] Generating edge-expanded graph representation
[info] .
 10% 
.
 20% 
.
 30% 
.
 40% 
.
 50% 
.
 60% 
.
 70% 
.
 80% 
.
 90% 
.
 100%
[info] Node compression ratio: 0.172748
[info] Edge compression ratio: 0.203861
[info]  graph compression removed 768 annotations of 1010 in 0.002438 seconds
[info] Find segregated edges in node-based graph ...
[info] ok, after 0.035637s
[info] Segregated edges count = 25
[info] Writing nodes for nodes-based and edges-based graphs ...
[info] Geometry successfully removed:
  compressed edges: 2500
  compressed geometries: 12046
  longest chain length: 48
  cmpr ratio: 0.207538
  avg chain length: 4.8184
[info] Generating edge expanded nodes ... 
[info] .
 10% 
.
 20% 
.
 30% 
.
 40% 
.
 50% 
.
 60% 
.
 70% 
.
 80% 
.
 90% 
.
 100%
[info] Expanding via-way turn restrictions ... 
[info] 
[info] Generated 1742 nodes (0 of which are duplicates)  and 6009 segments in edge-expanded graph
[info] Generating edge-expanded edges 
[info] .
 10% 
.
 20% 
.
 30% 
.
 40% 
.
 50% 
.
 60% 
.
 70% 
.
 80% 
.
 90% 
.
 100% 

[info] Sorting and writing 0 maneuver overrides...
[info] done.
[info] Renumbering turns
[info] Writing 0 conditional turn penalties...
[info] Generated 6009 edge based node segments
[info] Node-based graph contains 1742 edges
[info] Edge-expanded graph ...
[info]   contains 2536 edges
[info] Timing statistics for edge-expanded graph:
[info] Renumbering edges: 0.000469s
[info] Generating nodes: 0.002725s
[info] Generating edges: 0.13732s
[info] Generating guidance turns 
[info] .
 10% 
.
 20% 
.
 30% 
.
 40% 
.
 50% 
.
 60% 
.
 70% 
.
 80% 
.
 90% 
.
 100% 

[info] done.
[info] Created 18 entry classes and 504 Bearing Classes
[info] Handled: 11 of 50 lanes: 22 %.
[info] Assigned 3246 turn instruction types:
[info]   new name: 184 (5.67%)
[info]   continue: 261 (8.04%)
[info]   turn: 928 (28.59%)
[info]   fork: 64 (1.97%)
[info]   end of road: 298 (9.18%)
[info]   enter roundabout: 54 (1.66%)
[info]   enter and exit roundabout: 6 (0.18%)
[info]   enter rotary: 4 (0.12%)
[info]   enter and exit rotary: 4 (0.12%)
[info]   enter roundabout turn: 24 (0.74%)
[info]   (noturn): 493 (15.19%)
[info]   (suppressed): 735 (22.64%)
[info]   roundabout: 6 (0.18%)
[info]   exit roundabout: 67 (2.06%)
[info]   rotary: 3 (0.09%)
[info]   exit rotary: 9 (0.28%)
[info]   exit roundabout turn: 22 (0.68%)
[info]   (stay on roundabout): 84 (2.59%)
[info] Assigned 3246 turn instruction modifiers:
[info]   uturn: 221 (6.81%)
[info]   sharp right: 66 (2.03%)
[info]   right: 659 (20.30%)
[info]   slight right: 231 (7.12%)
[info]   straight: 1195 (36.81%)
[info]   slight left: 246 (7.58%)
[info]   left: 535 (16.48%)
[info]   sharp left: 93 (2.87%)
[info] Guidance turn annotations took 0.186372s
[info] Writing Intersection Classification Data
[info] ok, after 0.002246s
[info] Writing Turns and Lane Data...
[info] ok, after 0.004278s
[info] Saving edge-based node weights to file.
[info] Done writing. (0.000859)
[info] Computing strictly connected components ...
[info] Found 108 SCC (1 large, 107 small)
[info] SCC run took: 0.000902667s
[info] Building r-tree ...
[info] Constructing r-tree of 6009 segments build on-top of 5783 coordinates
[info] finished r-tree construction in 0.013665 seconds
[info] Writing edge-based-graph edges       ... 
[info] ok, after 0.001615s
[info] Processed 2536 edges
[info] Expansion: 12067 nodes/sec and 3635 edges/sec
[info] To prepare the data for routing, run: ./osrm-contract "/tmp/osrm-data/map.osrm"
[info] RAM: peak bytes used: 92774400

time="2023-01-13T03:41:25Z" level=info msg="Partitioning graph."

[info] Computing recursive bisection
[info] Loaded compressed node based graph: 2450 edges, 5783 nodes
[info]  running partition: 128 1.2 0.25 10 1000 # max_cell_size balance boundary cuts small_component_size
[info] Found 4800 SCC (0 large, 4800 small)
[info] SCC run took: 0.00137021s
[info] Full bisection done in 4e-06s
[info] Loaded node based graph to edge based graph mapping
[info] Loaded edge based graph for mapping partition ids: 5046 edges, 1742 nodes
[info] Fixed 0 unconnected nodes
[info] Edge-based-graph annotation:
[info]   level 1 #cells 1 bit size 1
[info]   level 2 #cells 1 bit size 1
[info]   level 3 #cells 1 bit size 1
[info]   level 4 #cells 1 bit size 1
[info] Renumbered data in 0.014563 seconds
[info] MultiLevelPartition constructed in 0.001692 seconds
[info] CellStorage constructed in 0.000652 seconds
[info] MLD data writing took 0.018952 seconds
[info] Cells statistics per level
[info] Level 1 #cells 1 #boundary nodes 0, sources: avg. 0, destinations: avg. 0, entries: 0 (0 bytes)
[info] Level 2 #cells 1 #boundary nodes 0, sources: avg. 0, destinations: avg. 0, entries: 0 (0 bytes)
[info] Level 3 #cells 1 #boundary nodes 0, sources: avg. 0, destinations: avg. 0, entries: 0 (0 bytes)
[info] Level 4 #cells 1 #boundary nodes 0, sources: avg. 0, destinations: avg. 0, entries: 0 (0 bytes)
[info] Bisection took 0.069613 seconds.
[info] RAM: peak bytes used: 30416896

time="2023-01-13T03:41:25Z" level=info msg="Customizing cells."

[info] Loaded edge based graph: 5046 edges, 1742 nodes
[info] Loading partition data took 0.034759 seconds
[info] Cells customization took 0.002164 seconds
[info] Cells statistics per level
[info] Level 1 #cells 1 #boundary nodes 0, sources: avg. 0, destinations: avg. 0, entries: 0 (0 bytes)
[info] Level 2 #cells 1 #boundary nodes 0, sources: avg. 0, destinations: avg. 0, entries: 0 (0 bytes)
[info] Level 3 #cells 1 #boundary nodes 0, sources: avg. 0, destinations: avg. 0, entries: 0 (0 bytes)
[info] Level 4 #cells 1 #boundary nodes 0, sources: avg. 0, destinations: avg. 0, entries: 0 (0 bytes)
[info] Unreachable nodes statistics per level
[info] Unreachable nodes statistics per level
[info] Unreachable nodes statistics per level
[info] Unreachable nodes statistics per level
[info] MLD customization writing took 0.00415 seconds
[info] Graph writing took 0.003042 seconds
[info] RAM: peak bytes used: 22904832
time="2023-01-13T03:41:25Z" level=info msg="Uploadling files to bucket osrm-monaco-server-v1673581285"
time="2023-01-13T03:41:25Z" level=info msg="get file names from root dir builder: /tmp/osrm-data"
time="2023-01-13T03:41:25Z" level=info msg="creating bucket: osrm-monaco-server-v1673581285"
time="2020-09-12T17:58:32Z" level=info msg="OSRM version osrm-monaco-server-v1673581285 created successfully"
```

Use the new OSRM builder version on osrm-server-monaco-car Github Project in order to do a rolling update in kubernetes 

In Helm Chart values replace the old version to osrm-monaco-server-v1673581285

```
osrm:
  builder:
    version: osrm-monaco-server-v1673581285

```
