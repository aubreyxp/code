
cat split.log | awk  'function foo(name) {print name}{foo($0)}'
#cat split.log | awk  'echo $0'
