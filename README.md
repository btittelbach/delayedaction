# delayedaction
Delayed Action is a collection of tools that let you react to a bunch of temporally related events with just one action

## Problem

You have a bunch of events. E.g. read/writes on a file or directory or several specific lines in a logfile.
In reaction to these, you want to run a command but just once, not for every read/write/logline.

## My Solution

Your event triggers a timer. On each event the timer is reset. If events finally stop coming, the timer expires and your action runs.

## Considerations

Sometimes it might be useful, to run the command right away and then inhibit further actions as events keep coming in. I've never needed this so far. That is not what this script does right now. Pulls welcome

# Tools:

## delayaction.sh

React to multiple command call events with one command call

Every time ```./delayaction.sh``` is called, a timer is reset (or started if this is the first call)
If the timer expires, ```<cmd>``` is called and ```./delayaction.sh``` exits
This allows you to react to a bunch of temporally related events but trigger only one event in response

### Usage
```
./delayaction.sh [-h] [-t n] [-n <name>] <cmd>
         -h      show this help and exit
         -t <x>  timer timeout in seconds
         -n <name>       use name for lock instead of autogenerating it from cmd
```


## goruncmdwhenstdinputstops

React to stdinput events within a given duration with one command call

Will listen to stdin. Each time input arrives it resets (or starts if this is the first char on stdin) a timer.
If no more input arrives for the given duration, the given ```<cmd>``` is executed.

### Install

```
go build -o goruncmdwhenstdinputstops
mv goruncmdwhenstdinputstops /usr/local/bin
```

### Usage

```
Usage:
./goruncmdwhenstdinputstops [-d <duration>] <cmd>
Will execute cmd after <duration> has elapsed without any new input on stdin

Options:
  -delay=10s: after which events have stopped coming, execute cmd. e.g.: 2m
```


## wait_until_no_inotify_modifications.sh

Blocks until no more changes have occurred in a given directory for a given duration.

Include it into your own scripts at your convenience

### Usage

```
./wait_until_no_inotify_modifications.sh <duration> <directory>
```
