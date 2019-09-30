# ADDNS - Ark Dynamic DNS

A simple application to update the IP address of an A record of a domain you own in Linode. The application is written in [Go](https://golang.org/)

## Set up the Config File
Copy and rename the example config file 'addns.json.example' to 'addns.json'. You can place this in the same folder as 
the binary, or in your home folder. The application will look in the same folder first, then in your home folder.

``` 
cp addns.json.example addns.json
```

Make sure you stick to the types defined in the config file. The access token is a string. The ID's are integers so don't wrap them in quotes.

## Personal Access Token

To update the DNS of one of your records you need to generate a Personal Access Token from [Linode's Cloud Manager](https://cloud.linode.com/profile/tokens)

Place the Token in the config file, and find the DomainId and RecordId from your cloud manager.

## Compiling

Move in to the addns folder and run the following to build the binary. 

``` 
go build
```

## Running

Just run the binary. On success we simply print out the results that are returned from Linode. 

```
matt@silverark:~$ ./addns

Query Result: {"id": 13860830, "ttl_sec": 300, "priority": 0, "port": 0, "type": "A", "protocol": null, "weight": 0, "name": "ddns", "tag":
null, "service": null, "target": "89.242.168.63"}
 
```

If there are errors, they will be printed out too.

``` 
matt@silverark6:~$ ./addns
Query Result: {"errors": [{"reason": "Invalid OAuth Token"}]}
```
