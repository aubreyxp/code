# 思路: 主要是利用最后一条日志时间和当前时间做对比, 如果差距较大报警, 如果正常在每天固定时刻报告正常

cd `dirname $0`

# 确定查找日志文件范围
grep_op="$1" #哪一个功能模块 down_check|list_check|upload

if [ "$grep_op" != "down_check" ] &&  [ "$grep_op" != "list_check" ] && [ "$grep_op" != "upload" ]; then
 echo "usage: log-monitor.sh down_check|list_check|upload H|M"
 exit
fi


if [[ $grep_op == "down_check" ]];then
  log_file_prefix="./log/down_check/s3_check_" 
  max_interval=7200
elif [[ $grep_op == "list_check" ]];then
  log_file_prefix="./log/list_check/s3_check_"
  max_interval=7200
elif [[ $grep_op == "upload" ]];then
  log_file_prefix="./log/upload/s3_upload_"
  max_interval=120
fi
suffix=`date  +%Y_%m`
search_file="$log_file_prefix$suffix.log"

# 检查服务是否存活
last_time=`tail -1 "$search_file" |awk '{print $1, $2}'`
current_time=`date '+%Y-%m-%d %H:%M:%S'`
last_sec=`date -d  "$last_time" +%s`
current_sec=`date -d  "$current_time" +%s`
interval=`expr $current_sec - $last_sec`  #计算2个时间的差 
if [ $interval -gt $max_interval ];then
  over_msg="service haikang_wave_test $grep_op seems not right: [$search_file seems not right, no append content in last $max_interval seconds]"
  python /data/niki/send_email/hk_wave_test.py "hk_s3_wave_test_${grep_op}_fail" "$over_msg"
  cont="hk_s3_wave_test_${grep_op}_fail"
  url="http://172.18.181.129:8008/zabbix_ops/monitor/send_alart?content=$cont&group=385&type=13"
  curl  "$url"
else
  current_hour=`date '+%H:%M'`
  if [ "$current_hour" = "11:30" ]; then
	  suc_msg="service haikang_wave_test about $grep_op is normal!"
	  python /data/niki/send_email/hk_wave_test.py "hk_s3_wave_test_${grep_op}_normal" "$suc_msg"
      cont="hk_s3_wave_test_${grep_op}_normal"
      url="http://172.18.181.129:8008/zabbix_ops/monitor/send_alart?content=$cont&group=385&type=13"
      curl  "$url"
  fi 
fi
