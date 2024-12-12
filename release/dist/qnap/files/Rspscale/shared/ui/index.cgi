#!/bin/sh
CONF=/etc/config/qpkg.conf
QPKG_NAME="Rspscale"
QPKG_ROOT=$(/sbin/getcfg ${QPKG_NAME} Install_Path -f ${CONF} -d"")
exec "${QPKG_ROOT}/rspscale" --socket=/tmp/ropsoft7/rspscaled.sock web --cgi --prefix="/cgi-bin/qpkg/Rspscale/index.cgi/"
