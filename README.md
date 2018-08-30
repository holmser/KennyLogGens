# KennyLogGens

![Kenny Loggins](images/kennyloggins_small.jpg)

KennyLogGens is a utility for generating log files for testing log
shipping systems written in Go  and using Cobra.  Welcome to the Danger Zone.

```
Usage:
  klog [flags]
  klog [command]

Available Commands:
  cook        Cook CPU
  gen         Generate logs to syslog
  help        Help about any command

Flags:
  -h, --help          help for klog
  -s, --seconds int   number of seconds between log entries (default 1)

Use "klog [command] --help" for more information about a command.

```


```
10:16:41.087 logGen ▶ INFO 	Revvin' up your engine
10:16:42.090 logGen ▶ INFO 	Listen to her howlin' roar
10:16:43.090 logGen ▶ INFO 	Metal under tension
10:16:44.091 logGen ▶ INFO 	Beggin' you to touch and go
10:16:45.096 logGen ▶ ERROR 	Highway to the danger zone
10:16:46.100 logGen ▶ ERROR 	Ride into the danger zone
```
