#!/bin/sh
CONF=/etc/config/qpkg.conf
QPKG_NAME="Rspscale"
QPKG_ROOT=`/sbin/getcfg ${QPKG_NAME} Install_Path -f ${CONF}`
QPKG_PORT=`/sbin/getcfg ${QPKG_NAME} Service_Port -f ${CONF}`
export QNAP_QPKG=${QPKG_NAME}
set -e

case "$1" in
  start)
    ENABLED=$(/sbin/getcfg ${QPKG_NAME} Enable -u -d FALSE -f ${CONF})
    if [ "${ENABLED}" != "TRUE" ]; then
        echo "${QPKG_NAME} is disabled."
        exit 1
    fi
    mkdir -p /home/httpd/cgi-bin/qpkg
    ln -sf ${QPKG_ROOT}/ui /home/httpd/cgi-bin/qpkg/${QPKG_NAME}
    mkdir -p -m 0755 /tmp/rspscale
    if [ -e /tmp/ropsoft7/rspscaled.pid ]; then
        PID=$(cat /tmp/ropsoft7/rspscaled.pid)
        if [ -d /proc/${PID}/ ]; then
          echo "${QPKG_NAME} is already running."
          exit 0
        fi
    fi
    ${QPKG_ROOT}/rspscaled --port ${QPKG_PORT} --statedir=${QPKG_ROOT}/state --socket=/tmp/ropsoft7/rspscaled.sock 2> /dev/null &
    echo $! > /tmp/ropsoft7/rspscaled.pid
    ;;

  stop)
    if [ -e /tmp/ropsoft7/rspscaled.pid ]; then
      PID=$(cat /tmp/ropsoft7/rspscaled.pid)
      kill -9 ${PID} || true
      rm -f /tmp/ropsoft7/rspscaled.pid
    fi
    ;;

  restart)
    $0 stop
    $0 start
    ;;
  remove)
    ;;

  *)
    echo "Usage: $0 {start|stop|restart|remove}"
    exit 1
esac

exit 0
