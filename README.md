# check-ftp

[![Build Status](https://cloud.drone.io/api/badges/ch1aki/check-ftp/status.svg)](https://cloud.drone.io/ch1aki/check-ftp)

FTP connection check plugin for mackerel.io agent.

## Synopsis
```
check-ftp -H ftp.example.com -P 21 -w 3 -c 5 -t 10
```

## Setting for mackerel-agent

If there are no problems in the execution result, add a setting in mackerel-agent.conf .

```
[plugin.checks.check-ftp-sample]
command = ["check-ftp", "-H", "ftp.example.com", "-P", "21", "-w", "3", "-c", "5", "-t", "10"]
```

## Usage
### Options

```
  -H, --host=                 Hostname (default: localhost)
  -P, --port=                 Port (default: 21)
  -u, --user=                 FTP username (default: anonymous)
  -p, --password=             FTP password (default: anonymous)
  -w, --warning=              Warning threshold (sec)
  -c, --critical=             Critical threshold (sec)
  -t, --timeout=              Timeout (sec) (default: 10)
  -s, --ftps                  Use FTPS
  -i, --implicit-mode         Connects directly using TLS
      --no-check-certificate  Do not check certificate
```