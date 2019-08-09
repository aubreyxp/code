#!/bin/bash

# cront根据具体任务调整执行频率
# 搜寻最近一定时间段内error日志, 出错就报警(发短信和邮件), 每次crontable时间范围不一样, 所以报警不会重复 

cd `dirname $0`

# 确定查找日志文件范围
grep_op="$1" #哪一个功能模块 down_check|list_check|upload
time_interval="$2" # 时间类型: H|M

if [ "$grep_op" != "down_check" ] &&  [ "$grep_op" != "list_check" ] && [ "$grep_op" != "upload" ]; then
 echo "usage: log-monitor.sh down_check|list_check|upload H|M"
 exit
fi

if [ "$time_interval" != "H" ] &&  [ "$time_interval" != "M" ]; then
 echo "usage: log-monitor.sh down_check|list_check|upload H|M"
 exit
fi

if [[ $grep_op == "down_check" ]];then
  log_file_prefix="./log/down_check/s3_check_" 
elif [[ $grep_op == "list_check" ]];then
  log_file_prefix="./log/list_check/s3_check_"
elif [[ $grep_op == "upload" ]];then
  log_file_prefix="./log/upload/s3_upload_"
fi
suffix=`date  +%Y_%m`
search_file="$log_file_prefix$suffix.log"

#s3_check_2019_08.log
#s3_check_2019_08.log
if [ $time_interval = 'M' ]; then
  time_grep=$(date "+%F %H:%M"  -d "1 minute ago") # 每分钟运行 但是检查上一分钟的日志 
elif [ $time_interval = 'H' ];then
  time_grep=$(date "+%F %H"  ) # crontable最好设置在整点的5-10分钟后运行
else
  exit
fi
#echo $time_grep

# 发短信
#url="http://172.18.181.129:8008/zabbix_ops/monitor/send_alart?content=jifefe&group=385&type=13"
grep -s "$time_grep" "$search_file" |grep -s error|grep -s $grep_op | while read line; do
    echo "$line"
    cont="err_haikang_s3_$grep_op"
    url="http://172.18.181.129:8008/zabbix_ops/monitor/send_alart?content=$cont&group=385&type=13"
    curl  "$url"
    python /data/niki/send_email/aubrey.py "$line" 
done

# 发邮件

