blitzfer
========

Go based filesystem level analytics tool. 

----

- Architecture Overview

Pretty simple really. When all is said and done, this will stand up, ElasticSearch and Kibana containers. A blitzfer container will populate elasticsearch with filesystem metadata. Kibana will visualize said data, once the data exists in elasticsearch. The data can then be reviewed and leveraged for many purposes, such as migration efforts, directory hash balancing, etc.

----

- TODO
-- Containerize App
-- Add YAML config support to the App.
-- Daemonize / scheduled fs metadata aggregation.
-- DockerCompose everything for simple deployment.
