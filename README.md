#github webhooks

Allows you to execute commands on the server as a result of a github webhook request.

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
