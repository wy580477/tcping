# tcping

tcping is like [tcping.exe](https://elifulkerson.com/projects/tcping.php), but written with Golang.

## Usage

- The default count of ping is 4.

- If the port is omitted, the default port is 80.

- The default interval of ping is 1s.

- The default timeout of ping is 1s.

### ping tcp

```bash
> tcping google.com 443
Ping tcp://google.com:443 - Connected - time=15.425732ms
Ping tcp://google.com:443 - Connected - time=2.628025ms
Ping tcp://google.com:443 - Connected - time=2.400356ms
Ping tcp://google.com:443 - Connected - time=1.967587ms

Ping statistics tcp://google.com:443
	4 probes sent.
	4 successful, 0 failed.
Approximate trip times:
	Minimum = 1.967587ms, Maximum = 15.425732ms, Average = 5.605425ms
```

### ping http

```bash
tcping http://microsoft.com
Ping http://microsoft.com:80(20.112.52.29) connected - time=474.944472ms dns=19.575171ms status=301
Ping http://microsoft.com:80(20.103.85.33) connected - time=554.906472ms dns=9.883287ms status=301
Ping http://microsoft.com:80(20.81.111.85) connected - time=548.337763ms dns=10.241747ms status=301
Ping http://microsoft.com:80(20.53.203.50) connected - time=419.761726ms dns=11.08323ms status=301

Ping statistics http://microsoft.com:80
        4 probes sent.
        4 successful, 0 failed.
Approximate trip times:
        Minimum = 419.761726ms, Maximum = 554.906472ms, Average = 499.487608ms
```

### ping https

```bash
tcping https://microsoft.com
Ping https://microsoft.com:443(20.84.181.62) connected - time=794.45726ms dns=8.955196ms status=301
Ping https://microsoft.com:443(20.81.111.85) connected - time=795.754879ms dns=8.816674ms status=301
Ping https://microsoft.com:443(20.81.111.85) connected - time=812.379029ms dns=10.356469ms status=301
Ping https://microsoft.com:443(20.84.181.62) connected - time=710.339939ms dns=9.642587ms status=301

Ping statistics https://microsoft.com:443
        4 probes sent.
        4 successful, 0 failed.
Approximate trip times:
        Minimum = 710.339939ms, Maximum = 812.379029ms, Average = 778.232776ms
```
