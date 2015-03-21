# delayedaction
Delayed Action is a collection of tools that let you react to a bunch of temporally related events with just one action

## Problem

You have a bunch of events. E.g. read/writes on a file or directory or several specific lines in a logfile.
In reaction to these, you want to run a command but just once, not for every read/write/logline.

## My Solution

Your event triggers a timer. On each event the timer is reset. If events finally stop coming, the timer expires and your action runs.