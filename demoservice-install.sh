#!/bin/bash
srv=demo
if [ "$2" != "" ]; then
  srv=$2
fi
case "$1" in
  -i)
    if [ ! -d /home/demo ];then
      useradd demo
      mkdir -p /home/demo
      chown -R demo:demo /home/demo
    fi
    working=/home/demo/demoservice
    working=${working//\//\\\/}
    if [ ! -f /etc/systemd/system/demoservice.service ];then
      sed "s/WORKING/$working/g" demoservice.service > /etc/systemd/system/demoservice.service
    fi
    if [ ! -f /home/demo/conf/demoservice.properties ];then
      mkdir -p /home/demo/conf/
      cp -f conf/demoservice.properties /home/demo/conf/
    fi
    rm -rf /home/demo/demoservice
    mkdir -p /home/demo/demoservice
    cp -rf * /home/demo/demoservice/
    chown -R demo:demo /home/demo
    systemctl enable demoservice.service
    systemctl start demoservice.service
    ;;
  *)
    echo "Usage: ./demoservice-install.sh -i"
    ;;
esac
