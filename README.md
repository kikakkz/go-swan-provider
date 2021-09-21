# Swan Provider Tool Guide

## Features:

This miner tool listens to the tasks that come from Swan platform. It provides the following functions:

* Download tasks automatically using Aria2 for downloading service.
* Import deals once download tasks completed.
* Synchronize deal status with Swan platform so that client will know the status changes in realtime.

## Prerequisite
- Lotus-miner
- Aria2

## Config
** rename config.toml.example to config.toml **
```shell
port="8888"
dev = true

[aria2]
aria2_download_dir = ""
aria2_host = "127.0.0.1"
aria2_port = 6800
aria2_secret = "my_aria2_secret"

[main]
api_url = "https://api.filswan.com"
miner_fid = "f0xxxx"
expected_sealing_time = 1920    # 1920 epoch or 16 hours
import_interval = 600           # 600 seconds or 10 minutes
scan_interval = 600             # 600 seconds or 10 minutes
api_key = ""
access_token = ""
api_heartbeat_interval = 600    # 600 seconds or 10 minutes

[bid]
bid_mode = 1   #0:manual, 1:auto
start_epoch = 276570  #current block height + 2880(1 day)
price = "0.01"
verified_price = "0.05"
min_piece_size = "8 Gib"
max_piece_size = "32 Gib"
auto_bid_task_per_day= 20
```

## Installation

Install miner tool and aria2
```shell
sudo apt install aria2
```

## How to use

### Step 1. Run Aria2 as System Service

#### Step 1.1 Set up Aria2:

```shell
sudo mkdir /etc/aria2
# Change user authority to current user
sudo chown $USER:$USER /etc/aria2/
# Create a session file
touch /etc/aria2/aria2.session
# Checkout the source and install 
git clone https://github.com/filswan/swan-provider

cd swan-provider

# Copy config file and service file
cp config/aria2.conf /etc/aria2/
sudo cp aria2c.service /etc/systemd/system/
# Modify the aria2c.service file in /etc/systemd/system/

# Set to start Aria2 automatically
sudo systemctl enable aria2c.service
# Start Aria2
sudo systemctl start aria2c.service
```
If modify user is nessecary while the service has been started, service should be reloaded before start.
```shell
sudo systemctl daemon-reload
sudo systemctl start aria2c.service
```


#### Step 1.2 Test Aria2 service from log (Optional)

Check if Aria2 service is successfullly started

```shell
journalctl -u aria2c.service -f
```
The output will be like:

```shell
Jul 30 03:00:00 systemd[1]: Started Aria2c download manager.
Jul 30 03:00:00 aria2c[2433312]: 07/30 03:00:00 [NOTICE] IPv4 RPC: listening on TCP port 6800
```

The Aira2 service will listen on certain port if installed and started correctly.

### Step 2. Compile Provider
```shell
1. for linux
GOOS=linux GOARCH=amd64 go build -v ./

2. for mac
env GOOS=darwin GOARCH=amd64 go build -v ./
```

### Step 3. Start Swan Provider
```shell
swan-provider
```

The deal status will be synchronized on the filwan.com, both client and miner will know the status changes in realtime.
