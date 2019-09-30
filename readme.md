# ADDNS - Ark Dynamic DNS

A simple application to update the IP address of an A record of a domain you own in Linode. the application is written in [Go](https://golang.org/)

## Set up the Config File
Copy and rename the example config file 'config.json.example' to 'config.json'.

``` 
cp config.json.example config.json
```

## Personal Access Token

To update the DNS of one of your records you need to generate a Personal Access Token from Linode's Cloud Manager: https://cloud.linode.com/profile/tokens

For convenience you can place the token in a file called '.addns_config.json' in your users directory $HOME/.addns_config.json


