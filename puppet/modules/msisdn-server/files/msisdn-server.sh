#!/bin/sh
### BEGIN INIT INFO
# Provides:          msisdn-server.sh
# Required-Start:    $all
# Required-Stop:     $all
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Start APP
# Description:       This script starts APP (binary)
### END INIT INFO

APP=/compiled/msisdn-server
PID=/var/run/msisdn-server.pid

[ -x "$APP" ] || exit 0

case "$1" in

  start)
    start-stop-daemon --start --quiet --oknodo --background --make-pidfile \
                      --pidfile $PID --chuid vagrant --exec $APP
    ;;

  stop)
   start-stop-daemon --stop --quiet --oknodo --retry TERM/5/KILL/5 \
                     --pidfile $PID --exec $APP
    ;;

esac

exit 0
