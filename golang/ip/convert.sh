cat a.log|awk -F':' '{print $2}' > a.log.1
cat a.log.1|awk -F'"' '{print $2}' > ip.txt
rm a.log.1

