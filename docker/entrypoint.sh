#!/bin/sh
set -xe
# copy default configure
if [ ! -f /srv/conf/demoservice.properties ];then
    mkdir -p /srv/conf
    cp -f /srv/demoservice/conf/demoservice.properties /srv/conf/demoservice.properties
fi

cd /srv/demoservice/
/srv/demoservice/demo /srv/conf/demoservice.properties
echo "service is done"