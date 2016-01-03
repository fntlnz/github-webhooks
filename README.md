# GitHub Webhooks [![Build Status](https://travis-ci.org/fntlnz/github-webhooks.svg?branch=master)](https://travis-ci.org/fntlnz/github-webhooks)

Execute shell commands on the server when a GitHub web hook event is fired on your repo.

## Usage

### Run the server

The `server` command starts an HTTP webserver that can take requests from the GitHub WebHooks service.
This command uses the default configuration path `/etc/github-webhooks.json` but you can provide a custom one with the `-c, --configuration` parameter.

```
github-webhooks server
```

```
github-webhooks server -c /my/custom/configuration.json
```

### Configuration file example

GitHub WebHooks is configured trough a json file with the following structure:

Note that **host** and **port** are optional (defaults are 0.0.0.0 and 3091) and that if you want to **target hooks
to a specific branch** you have to add another node to the configuration specifying the
branch name after the repository name, for example if you want to accept hooks
from `fntlnz/dockerfiles` master branch you have to use `fntlnz/dockerfiles/master` as repo name.

As you have probably noted each repository have a node called events,
here you have to specify the event to listen (for example `push`) and on each event a list of commands to execute.

The **path** node allows you to overwrite the `$PATH` environment variable,
if not set default `$PATH` is used.

```json
{
    "host": "0.0.0.0",
    "port": "3091",
    "path": "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin",
    "repositories": {
        "fntlnz/dockerfiles": {
            "events": {
                "ping": ["touch ping-on-any-branch.txt"]
            }
        },
        "fntlnz/dockerfiles/master": {
            "events": {
                "push": ["touch push-on-master-branch.txt"]
            }
        }
    }
}
```

## Security

It's recommended to configure your firewall to accept only requests coming
from trusted servers. In this case trusted servers are the GitHub ones.

To know which are ip addresses of GitHub hooks servers take a look at [https://api.github.com/meta](https://api.github.com/meta)

### iptables
Here's the iptables configuration to create a specific chain to allow GitHub Webhooks servers to access the service.

```bash
# Create a chain named GitHubWebHooks
iptables -N GitHubWebHooks
 
# Enable the 3091 port on the chain

iptables -I INPUT -s 0/0 -p tcp --dport 3091 -j GitHubWebHooks

# Enable trusted ips on GitHubWebHooks  chain 

iptables -I GitHubWebHooks -s 192.30.252.0/22 -j ACCEPT
 
# Drop the public from GitHubWebHooks chain
iptables -A GitHubWebHooks -s 0/0 -j DROP

# Save
service iptables save

# Restart
service iptables restart

```
