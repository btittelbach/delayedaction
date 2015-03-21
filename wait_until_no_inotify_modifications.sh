#!/bin/zsh
#(c) Bernhard Tittelbach, xro@realraum.at, 2013

local TIMEOUT=$(($1))
shift
local WPATH="${*}"
while inotifywait -r -t "$TIMEOUT" -q -q -e modify,close,open,attrib,move,create,delete "$WPATH"; do
  sleep 1
done
