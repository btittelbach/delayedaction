#!/bin/zsh --extendedglob
# (c) Bernhard Tittelbach, xro@realraum.at, 2015
#
# this script allows you to react to a bunch of temporally related events and trigger only one event in response
#
#

zmodload zsh/system || exit 3
local LOCKNAME
local SHOWHELP

DELAY=("t" "10")
zparseopts -D -E "n:=LOCKNAME" "t:=DELAY" "h=SHOWHELP"
if [[ -n "$SHOWHELP" ]]; then
	print "Everytime $0 is called, a timer is reset (or started if this is the first call)"
	print "If the timer expires, <cmd> is called and $0 exits"
	print "This allows you to react to a bunch of temporally related events but trigger only one event in response"
	print ""
	print "$0 [-h] [-t n] [-n <name>] <cmd>"
	print "\t -h \t show this help and exit"
	print "\t -t <x>\t timer timeout in seconds"
	print "\t -n <name> \t use name for lock instead of autogenerating it from cmd"
	exit 0
fi

local CMDMD5=${(w)$(echo -n "$*" | md5sum)[1]}
local LOCKFILE=/tmp/delayedaction.${LOCKNAME[2]:-$CMDMD5}
touch $LOCKFILE
zsystem flock -t 1 $LOCKFILE || exit 0

while inotifywait -r -t "${DELAY[2]}" -q -q "$LOCKFILE"; do
  sleep 1
done

$*