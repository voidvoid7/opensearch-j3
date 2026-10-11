#!/bin/bash
sudo sysctl -w vm.max_map_count=262144     # Linux hosts; OpenSearch won't start without it
docker compose up -d --build
curl 'localhost:9200/_cat/indices/app-logs-*?v'
curl 'localhost:9200/app-logs-*/_search?q=level:ERROR&size=2&pretty'