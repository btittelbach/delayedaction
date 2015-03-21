#!/bin/zsh --extendedglob
# Example Script that forces redetection of USB Devices on log message
# (note: does _not_ usb-reset devices, you need a PPPS per-port-power-switching hub for that and toggle power/pm_qos_no_power_off at least)
local USBPROD="1234"
local USBVEND="5678"
local LOGFILE=/var/log/kern.log
local LOGREGEX="Some Kernel Problem"

if [[ $1 == "doit" ]]; then
	DEVICES=( /sys/bus/usb/devices/*/idProduct(e:'[[ "$(<$REPLY)" == "$USBPROD" && "$(<${REPLY/%idProduct/idVendor})" == "$USBVEND" ]]'::A:s/idProduct/authorized/) )
	for fa ($DEVICES) echo 0 >| "$fa"
	sleep 2
	for fa ($DEVICES) echo 1 >| "$fa"
else
	tail -f $LOGFILE | egrep --line-buffered "$LOGREGEX" | /usr/local/bin/goruncmdwhenstdinputstops -t 5 ${0:a} doit
fi
exit 0
