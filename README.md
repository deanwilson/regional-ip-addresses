# Regional IP Addresses

A cli tool for fetching and displaying IP addresses by country

## Introduction

I wrote this tool to provide source data for blocking IP addresses by
geographical regions. As an example I don't need `ssh` on my VPS to be
reachable from outside of the UK.

This command fetches the source data from [RIPE](https://www.ripe.net/)
so please don't run it too often.

### ipsets

Due to the number of networks assigned to some countries this command
has the ability to output the networks in an
[ipset](https://www.linuxjournal.com/content/advanced-firewall-configurations-ipset)
compatible format. `ipset ` is a much faster and more efficient way for iptables
tables to match against a large number of addresses.

    go run main.go -ipset -ipset-name test01 -ipset-header | head
    
    ipset create test01 hash:net
    ipset -A test01 2.24.0.0/13
    ipset -A test01 2.56.188.0/22
    ipset -A test01 2.56.196.0/22
    ipset -A test01 2.57.16.0/22
    ipset -A test01 2.57.20.0/22
    # ... snip ...
    
 This `ipset` can then be used in iptables, for example to drop
 traffic from any networks contained in the set. 

    iptables -I INPUT -m set --match-set test01 src -j DROP
    
## Examples

Fetch and display the Great Britain IPv4 Address ranges

    go run main.go -source web

Run against the sample data file from this repository

    go run main.go -source file -data sample-data/example.json

### Author

 * [Dean Wilson](https://www.unixdaemon.net)
