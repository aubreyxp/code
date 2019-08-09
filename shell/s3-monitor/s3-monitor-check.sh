#!/bin/bash

# 前提: 每个小时执行一次

PATH=/usr/bin:$PATH

# 可使用相对路径
cd `dirname $0`

op=$1
if [[ $op == "down_check" ]];then
  download_dir="./log/down_check/down_check"
  record_file="./log/down_check/down_check.txt"
  log_file_prefix="./log/down_check/s3_check_" 
  aws_op="cp"
elif [[ $op == "list_check" ]];then
  download_dir="./log/list_check/list_check"
  record_file="./log/list_check/list_check.txt"
  log_file_prefix="./log/list_check/s3_check_" 
  aws_op="ls"
else
  echo "usage: sh s3-monitor-check.sh down_check|list_check"
  exit
fi

# 存放下载文件的地方
if [ ! -d "$download_dir" ]; then
  mkdir $download_dir
fi

# 记录上一次上传文件的时间 
if [ ! -f "$record_file" ]; then 
 touch "$record_file" 
fi 

# 记录日志文件 
if [ ! -d "./log/down_check" ]; then 
 mkdir "./log/down_check" 
fi 

if [ ! -d "./log/list_check" ]; then 
 mkdir "./log/list_check" 
fi 

function Log() {
  now=`date "+%F %H:%M:%S"`
  func=$1
  level=$2
  msg=$3
  suffix=`date  +%Y_%m`
  echo "$now $func-$level $msg" >> $log_file_prefix$suffix.log
}
#example: Log test debug "hello"

upload_file_tmp=./file-temp/tmp.dat # 上传的文件模版
function upload_for_check() {
  time_str=$1
  aws s3 cp $upload_file_tmp  s3://$bucket/s3_monitor/$op/$time_str/ --endpoint-url $endpoint
  if [ $? -ne 0 ]; then
    echo "upload to s3_monitor/$op/$time_str/tmp.dat failed"
    Log $op error "upload to s3_monitor/$op/$time_str/tmp.dat failed"
  else
    echo "upload  s3_monitor/$op/$time_str/tmp.dat succed"
    Log $op info "upload  s3_monitor/$op/$time_str/tmp.dat succed"
  fi
  echo $time_str > $record_file 
}

# 设置bucket相关变量 
bucket='store-gw'
endpoint="http://s3-cn-sh2.ufileos.com"
now_time_str=`date +%Y%m%d%H` 
last_time_str=`date -d "1 hour ago" +%Y%m%d%H`

last_mark=`cat $record_file` # 上一次什么时候上传 
# 如果是第一执行的时候处理
if [ "$last_mark" != "$last_time_str" ];then 
   # 第一次启动, 上传前一个小时文件, 方便下载, 后面程序行为可保持一致
   echo "first $op, need upload last file for first $op or something wrong."
   Log $op info "first $op, need upload last file for first $op or something wrong."
   upload_for_check $last_time_str
fi

# 执行命令 
if [ $aws_op = 'cp' ]; then
  cmd="aws s3 $aws_op  s3://$bucket/s3_monitor/$op/$last_time_str/tmp.dat ./$download_dir/tmp_$last_time_str.dat --endpoint-url $endpoint "
elif [ $aws_op = 'ls' ];then
  cmd="aws s3 $aws_op  s3://$bucket/s3_monitor/$op/$last_time_str/tmp.dat  --endpoint-url $endpoint "
fi
ret=`eval $cmd 2>&1`

if [ $? -ne 0 ]; then
  echo "check $aws_op  s3_monitor/$op/$last_time_str/tmp failed"
  Log $op error "check cmd:\"$cmd\"  err_msg:\"$ret\""
else
  echo "check $aws_op  s3_monitor/down/$last_time_str/tmp succed"
  Log $op info "check $aws_op  cmd:\"$cmd\" succed"
fi


# 上传供一个小时候后检查
upload_for_check $now_time_str

