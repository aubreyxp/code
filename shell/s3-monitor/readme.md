## 功能 
* 监控海康s3上传、下载、list是否正常 

## 步骤 
* 需要安装aws cli工具, 参考 https://docs.aws.amazon.com/zh_cn/cli/latest/userguide/install-linux.html
* 配置公私钥, 参考 https://docs.aws.amazon.com/zh_cn/cli/latest/userguide/cli-chap-configure.html#cli-quick-configuration
* 修改aubreycron里面脚本的实际放置路径 
* 修改s3-monitor-check.sh和s3-monitor-upload.sh里面endpoint和bucket变量的值
* sh 里面设置aws的路径 
* 查看如果没有crontab才能执行crontab aubreycron, 否则-e 
