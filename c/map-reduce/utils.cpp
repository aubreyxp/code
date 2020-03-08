#ifndef __H_UTILS__
#define __H_UTILS__

#include <iostream>
#include <string>
#include <sstream>
#include <algorithm>
#include <iterator>
#include "utils.h" 
using namespace std;

string reduceName(string dataDir, string jobName , int mapTask, int reduceTask)  {
	std::stringstream ss;
	ss << dataDir << "/" << "mrtmp/" << jobName << "-" << mapTask << "-" << reduceTask;
	return ss.str();
}

string mergeName(string dataDir, string jobName , int reduceTask)  {
	std::stringstream ss;
	ss << dataDir << "/" << "mrtmp/" << jobName << "-reduce-"  << reduceTask;
	return ss.str();
}

 


#endif

