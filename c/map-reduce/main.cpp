#include "stdio.h"
#include "map_reduce.h"
#include "threadpool.h"

vector<KeyValue> CountMp(string filename, string contents){
	std::cout << "enter CountMp" << std::endl;
	//std::cout << "size" << contents.size() <<  std::endl;

	std::vector<string> lines;
	split(contents, lines);


	std::map<string, int> mp;
	auto it = lines.begin();
	for (;it != lines.end(); it++) {
		if (*it == "") {
			continue;
		}
		//std::cout << *it << std::endl;
		
		auto mpIt = mp.find(*it);
		if (mpIt == mp.end()) {
			mp[*it] = 1;
		} else {
			mp[*it]++;
		}
	}
	
	std::vector<KeyValue> results;
	auto mpIt = mp.begin();
	for (; mpIt != mp.end(); mpIt++) {
		KeyValue kv;
		kv.key = mpIt->first;
		kv.value = std::to_string(mpIt->second);
		results.push_back(kv);
		//std::cout << kv.key <<  " " << kv.value << std::endl;
	}
	
	std::cout << "leave CountMp" << std::endl;

	//return std::move(results);
	return (results);
}

string CountReduce (string key, vector<string> values) {
	auto it = values.begin();
	int total = 0;
	for (; it != values.end();it++) {
		if (*it == "") {
			continue;
		}
		int cnt = std::stoi(*it);
		total += cnt;
	}

	std::stringstream ss;
	ss << key << " " << total;
	return ss.str();
}

bool Cmpare(const KeyValue &a, const KeyValue &b)
{
	     int aV = std::atoi(a.value.c_str());
	     int bV = std::atoi(b.value.c_str());
	     return aV > bV;
}

vector<KeyValue> Top10Mp(string filename, string contents){
	std::cout << "enter Top10Mp" << std::endl;

	std::vector<string> lines;
	split(contents, lines);

	std::vector<KeyValue> list;
	auto it = lines.begin();
	for (;it != lines.end(); it++) {
		if (*it == "") {
			continue;
		}
		//std::cout << *it << std::endl;
		std::vector<string> strs;
		split(*it, strs, ' ');
		if(strs.size() != 2) {
			continue;
		}

		KeyValue kv;
		kv.key = strs[0];
		kv.value = strs[1];
		list.push_back(kv);
	}

	std::sort(list.begin(), list.end(), Cmpare); 

	std::vector<KeyValue> results;
	int count  = 0;
	auto itVec = list.begin();
	for (; itVec != list.end(); itVec++) {
		std::stringstream ss;
		ss << (*itVec).key << " " << (*itVec).value;

		KeyValue kv;
		kv.key = "";
		kv.value = ss.str();
		results.push_back(kv);

		count++;
		if (count == 10) {
			break;
		}
	}
	
	return results;
}

string Top10Reduce (string key, vector<string> values) {
	//std::sort(list.begin(), list.end(), Cmpare); 
	std::vector<KeyValue> list;
	auto it = values.begin();
	for (;it != values.end(); it++) {
		if (*it == "") {
			continue;
		}
		std::vector<string> strs;
		split(*it, strs, ' ');
		if(strs.size() != 2) {
			continue;
		}

		KeyValue kv;
		kv.key = strs[0];
		kv.value = strs[1];
		list.push_back(kv);
	}

	std::sort(list.begin(), list.end(), Cmpare); 

	int count  = 0;
	auto itVec = list.begin();
	std::stringstream ss;
	for (; itVec != list.end(); itVec++) {
		ss << (*itVec).key << " " << (*itVec).value << std::endl;
		count++;
		if (count == 10) {
			break;
		}
	}

	return ss.str();
}


int main() {
	ThreadPool tp;

	std::vector<string> mapFiles;
	mapFiles.push_back("./data/inputMapFile0");
	mapFiles.push_back("./data/inputMapFile1");
	mapFiles.push_back("./data/inputMapFile2");
	mapFiles.push_back("./data/inputMapFile3");
	mapFiles.push_back("./data/inputMapFile4");
	mapFiles = tp.submit("test-mr-roud-1", "./data", CountMp, CountReduce, mapFiles, 4);
	vector<string> ret = tp.submit("test-mr-roud-2", "./data", Top10Mp, Top10Reduce, mapFiles, 1);
	std::cout << "result file " << ret[0] << std::endl;
}
