## step
* First, base on map-reduce, make a framework to do Concurrency tasks
* Second, make a map funtion to deal with the split url's files, counter ervery url in the file. 
* Third, transfer all key-value results of map fuction to reduce files, use hash(key)%nReduce to local files. 
* Every ruduce task gets all reduce files in all map task with special reduce number. Every url is received by one reduce task, then counter the number of the url. Then output all urls with number in local reduce file.
* In the second round, map fuction sorts number of url desc and output top10 to the fix reduce file.
* At last, reduce fuction sorts all the output in step 5 , print out the top10 urls.

## usage 
* make clean  // clean the temp object 
* make cleand // clean temp test data
