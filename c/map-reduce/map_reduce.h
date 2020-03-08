#ifndef __H_MAP_REDUCE_H__
#define __H_MAP_REDUCE_H__

#include <string>
#include <vector>
#include <atomic>
#include "json/json.h"
using namespace std;

typedef struct {
	string key;
	string value;
} KeyValue;



typedef string (*ReduceF)(string key, vector<string> values);
typedef vector<KeyValue> (*MapF)(string  filename,  string contents);


typedef string TaskPhase;
const TaskPhase MapPhase = "map";
const TaskPhase ReducePhase = "reduce";

typedef  struct {
   string phase;
   string dataDir;
   string jobName;
   string mapFile;
   int taskNumber;
   int  nReduce;
   int  nMap;
   MapF mapF;
   ReduceF  reduceF;
   bool done;
} Task;

std::string toJson(KeyValue& kv); 
KeyValue fromJson(string& str); 


#endif
