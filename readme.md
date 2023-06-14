# Goshia - Minimal CD

## Principle

The idea is that you can deploy this application on the servers that require automatic deployment.
When running your devops pipeline you can simple execute a curl command on the /deploy endpoint which will clone or pull your defined repository and run a specific shell script. In my case, this was sufficient for what we needed, a simple trigger. Currently, I use it for a simple docker compose up, but you are free to use this however you want.

Example of a run.sh file (you can specify whatever you want):

```sh
#!/bin/sh

case $1 in
  staging)
    docker compose up -d
    ;;
  *)
    echo "did not find environment match for $1"
    ;;
esac

```

To configure this app:
1. Clone the repository
2. Copy configuration.template.yml to configuration.yml and fill in your details
3. Run sudo make, this will make sure golang is installed, folders are created and whatnot
4. Ready to go!

Deployment can be done by executing a HTTP GET request on /deploy with queryparameter named branch, specifying which branch to checkout.
