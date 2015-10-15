Webhook Plugin
==============

This is the official webhook plugin. This plugin can be used to call an uri,
possibly a next step in your deployment pipeline.

To enable this plugin, add the following in your `.emerald.yml`

```
*usual emerald config*
...
plugins:
  - plugin_image: emeraldci/docker-image-plugin
    url: https://registry.hub.docker.com/u/emeraldci/api/trigger/<Trigger Token>/
```
