# Cloudflare-Helper
This tool helps you to update DNS record on cloudflare  
This is useful if you want host a server from home but your public ip address changes from time to time (aka. Dynamic IP address)

## How to
1. First you need a API Key from cloudflare known as a Bearer Token, you can learn how to acquire one [here](https://support.cloudflare.com/hc/en-us/articles/200167836-Managing-API-Tokens-and-Keys)
2. Download zip according to your platform
3. Modify config.yml from example folder
4.  ```bash
    cloudflare-helper -config /path/to/config.yml
    ```
- Use docker
```bash
docker run \
    -v /path/to/config.yml:/etc/cfh/config.yml \
    tdl3/cloudflare-helper
```
### Note
- You can find zone id on you cloudflare overview page
- Since v0.0.2 you can omit `dns_id` field in config file
