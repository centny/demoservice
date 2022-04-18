#!/bin/bash
##############################
#####Setting Environments#####
echo "Setting Environments"
set -xe
export cpwd=`pwd`
export LD_LIBRARY_PATH=/usr/local/lib:/usr/lib
output=$cpwd/build
#### Package ####
srv_ver=$1
srv_name=demoservice
build=$cpwd/build
output=$cpwd/build/$srv_name-$srv_ver
out_dir=$srv_name-$srv_ver
srv_out=$output/$srv_name
go_path=`go env GOPATH`
go_os=`go env GOOS`
go_arch=`go env GOARCH`

##build normal
cat <<EOF > version.go
package main

const Version = "$srv_ver"
EOF
echo "Build $srv_name normal executor..."
go build -o $srv_out/demo github.com/centny/demoservice
cp -rf conf $srv_out
cp -rf demoservice.service demoservice-install.sh $srv_out
cp -f run_*.sh $output
cp -f docker/Dockerfile docker/entrypoint.sh $output

###
cd $output
out_tar=$srv_name-$go_os-$go_arch-$srv_ver.tar.gz
rm -f $out_tar
tar -czvf $build/$out_tar $srv_name

##normal package
cd $output
out_tar=$srv_name-$go_os-$go_arch-$srv_ver.tar.gz
rm -f $out_tar
tar -czvf $build/$out_tar $srv_name

##docker package
having=`docker image ls -q demoservice:$srv_ver`
if [ "$having" != "" ];then
  docker image rm -f demoservice:$srv_ver
fi
docker build -t demoservice:$srv_ver .
cd $output
docker image save demoservice:$srv_ver -o demoservice-$srv_ver.img
tar -czvf demoservice-docker-$srv_ver.tar.gz demoservice-$srv_ver.img run_*.sh

cd $cpwd

echo "Package $srv_name-$go_os-$go_arch-$srv_ver done..."