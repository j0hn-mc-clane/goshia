# Goshia - Minimal CD

## Principle

The idea is that you can deploy this application on the servers that require automatic deployment.
When running your devops pipeline you can simple execute a curl command on the /deploy endpoint which will clone or pull your defined repository and run a specific shell script. In my case, this was sufficient for what we needed, a simple trigger. Currently, I use it for a simple docker compose up, but you are free to use this however you want.

Example of a run.sh file (you can specify whatever you want):

```sh
#!/bin/sh

case $1 in
  staging)
    docker compose -f docker-compose.yml -f docker-compose.staging.yml -f docker-compose.workers.yml up -d
    ;;
  dev)
    docker compose -f docker-compose.yml -f docker-compose.dev.yml -f docker-compose.workers.yml up -d
    ;;
    ;;
  local)
    docker compose -f docker-compose.yml -f docker-compose.dev.yml -f docker-compose.workers.yml up -d
    ;;
  *)
    echo "did not find environment match for $1"
    ;;
esac

```

To configure this app, simply copy the configuration.yml to configuration.local.yml and substitute the placeholders.

The application is served on port 8080 with endpoint GET /deploy.
