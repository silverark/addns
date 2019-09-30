# ADDNS - Ark Dynamic DNS

A simple application to update the IP address of an A record of a domain you own in Linode. the application is written in [Go](https://golang.org/)

## Set up the Config File
Copy and rename the example config file 'addns.json.example' to 'addns.json'. You can place this in the same folder as 
the binary, or in your home folder. The application will look in the same folder first, then in your home folder.

``` 
cp addns.json.example addns.json
```

## Personal Access Token

To update the DNS of one of your records you need to generate a Personal Access Token from [Linode's Cloud Manager](https://cloud.linode.com/profile/tokens)

Place the Token in the config file, and find the DomainId and RecordId from your cloud manager.


