cd `dirname $0`
suffix=`date '+%Y_%m'` 
down_prefix="./log/down_check/s3_check_"
last_time=`tail -1 "$down_prefix$suffix.log" |awk '{print $1, $2}'`
current_time=`date '+%Y-%m-%d %H:%M:%S'`
echo $last_time 
echo $current_time
last_sec=`date -d  "$last_time" +%s`
current_sec=`date -d  "$current_time" +%s`
interval=`expr $current_sec - $last_sec`  #计算2个时间的差 
echo $interval
if [ $interval -gt 4000 ];then
  echo "over time"
  # TODO: send mail and message
else
  echo "less time"
  # TODO: send mail and message
fi
#tail -1 upload/s3_upload_2019_08.log |awk '{print $1, $2}'
