#ifndef __H_UTILS__
#define __H_UTILS__

#include <iostream>
#include <string>
using namespace std;

string reduceName(string dataDir, string jobName , int mapTask, int reduceTask); 
string mergeName(string dataDir, string jobName , int reduceTask);  

template <class Container>
void split(const std::string& str, Container& cont, char delim = '\n')
{
    std::stringstream ss(str);
    std::string token;
    while (std::getline(ss, token, delim)) {
        cont.push_back(token);
    }
}


#endif

