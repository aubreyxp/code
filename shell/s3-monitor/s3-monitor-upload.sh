#!/bin/bash

cd `dirname $0`
PATH=/usr/bin:$PATH

# 设置变量
bucket='store-gw'
endpoint="http://s3-cn-sh2.ufileos.com"
timestr=`date +%Y%m%d%H%M` 

# 存放日志文件的地方
if [ ! -d "./log/upload" ]; then
  mkdir "./log/upload"
fi

log_file_prefix="./log/upload/s3_upload_"
function Log() {
  now=`date "+%F %H:%M:%S"`
  func=$1
  level=$2
  msg=$3
  suffix=`date  +%Y_%m`
  echo "$now $func-$level $msg" >> $log_file_prefix$suffix.log
}
#example: Log test debug "hello"

upload_file_dir=./file-temp
for i in {0..9}
do
  cmd="aws s3 cp $upload_file_dir/tmp$i.dat  s3://$bucket/s3_monitor/upload/$timestr/ --endpoint-url $endpoint"
  ret=`eval $cmd 2>&1`
  if [ $? -ne 0 ]; then
    echo "upload to s3_monitor/$timestr/tmp$i failed"
    Log upload error "cmd:\"$cmd\" err_msg:\"$ret\""
    exit
  fi
done
echo "upload to s3_monitor/$timestr/tmp[0-9] succeed"
Log uploader info "upload to s3_monitor/$timestr/tmp[0-9] succeed"
