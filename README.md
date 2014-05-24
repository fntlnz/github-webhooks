# github webhooks

Allows you to execute commands on the server as a result of a github webhook request.

## Usage

```
github-webhooks -configuration=github-webhooks.json
```

### Configuration file example

Note that **port** is optional (default is 3091) and that if you want to **target hooks
to a specific branch** you have to add another node to the configuration specifying the
branch name after the repository name, for example if you want to accept hooks
from `majinbuu/statik` master branch you have to use `majinbuu/statik/master` as repo name.

As you have probably noted each repository have a node called events,
here you have to specify the event to listen (for example `push`) and on each event a list of commands to execute.

```json
{
    "port": "3091",
    "repositories": {
        "majinbuu/statik": {
            "events": {
                "ping": ["touch ping-on-any-branch.txt"]
            }
        },
        "majinbuu/statik/master": {
            "events": {
                "push": ["touch push-on-master-branch.txt"]
            }
        }
    }
}
```


## Installation

### From source

```
go get github.com/majinbuu/github-webhooks
```

### Pre-built binaries

Them are coming soon!

## Security

It's recommended to configure your firewall to accept only requests coming
from trusted servers. In this case trusted servers are the GitHub ones.

To know which are ip addresses of GitHub hooks servers take a look at [https://api.github.com/meta](https://api.github.com/meta)

### iptables

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
