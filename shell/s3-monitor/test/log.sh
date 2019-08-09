log_file_prefix=./test
function Log() {
  now=`date "+%F %H:%M:%S"`
  func=$1
  level=$2
  msg=$3
  suffix=`date  +%Y_%m`
  echo "$now $func-$level $msg" >> $log_file_prefix$suffix.log
}
log test debug "hello"
