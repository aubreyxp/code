#ifndef THREAD_POOL_H
#define THREAD_POOL_H

#include <vector>
#include <functional>
#include <iostream>
#include <sstream>
#include <fstream>
#include <queue>
#include <atomic>
#include <future>
#include <stdexcept>
#include "map_reduce.h" 
#include "utils.h" 

#define  THREADPOOL_MAX_NUM 16
//#define  THREADPOOL_AUTO_GROW

class ThreadPool
{
    vector<thread> _pool;
    queue<Task*> _tasks;
    mutex _lock;
    condition_variable _task_cv;
    atomic<bool> _run{ true };
    atomic<int>  _idlThrNum{ 0 };

public:
    inline ThreadPool(unsigned short size = 5) { addThread(size); }
    inline ~ThreadPool()
    {
        _run=false;
        _task_cv.notify_all(); 
        for (thread& thread : _pool) {
            if(thread.joinable())
                thread.join(); 
        }
    }

public:

    int idlCount() { return _idlThrNum; }
    int thrCount() { return _pool.size(); }
private:
    //添加指定数量的线程
    void addThread(unsigned short size)
    {
        for (; _pool.size() < THREADPOOL_MAX_NUM && size > 0; --size)
        {   
            _pool.emplace_back( [this]{ 
                while (_run)
                {
                    Task* task; 
                    {
                        unique_lock<mutex> lock{ _lock };
                        _task_cv.wait(lock, [this]{
                                return !_run || !_tasks.empty();
                        }); 
                        if (!_run && _tasks.empty())
                            return;
                        task = _tasks.front(); 
                        _tasks.pop();
                    }
                    _idlThrNum--;
		    if (task->phase == MapPhase) {
			   //std::cout << "enter map" << std::endl;
			    std::ifstream t(task->mapFile);
			    std::stringstream buffer;
			    buffer << t.rdbuf();
			    std::string contents(buffer.str());

			    ofstream outfiles[task->nReduce];
			    for (int i = 0; i < task->nReduce; i++) {
				string rpath = reduceName(task->dataDir, task->jobName, task->taskNumber, i);
				//std::cout << rpath << std::endl;
				outfiles[i].open(rpath);
			    }

			    //std::cout << "enter map phase 1" << std::endl;

			    //outfile.open("test")
			    //outfile << "test" << std::endl;
			    std::vector<KeyValue> results = task->mapF(task->mapFile, contents);
			    //std::cout << "enter map phase 2" << std::endl;
			    std::hash<std::string> str_hash;
			    for (int i = 0; i < results.size(); i++) {
				    // if type is int, may be negative number, take care of it
				    uint32_t rand = str_hash(results[i].key);
				    int chose = rand % (task->nReduce);
				    outfiles[chose] << toJson(results[i]) << std::endl;
			    }
			    
			    for (int i = 0; i < task->nReduce; i++) {
				outfiles[i].close();
			    }
			    task->done = true;
			   //std::cout << "over map" << std::endl;
		    } else {
			   std::cout << "enter reduce" << std::endl;
			    // read from every map  file written with json formate
			    ifstream infiles[task->nMap];
			    std::map<string, vector<string> > contentsMp;
			    for (int i = 0; i < task->nMap; i++) { // get all map with the reduce task.taskNumber 
				string rpath = reduceName(task->dataDir, task->jobName, i, task->taskNumber);
				//std::cout << "rpath:" << rpath << std::endl;
				infiles[i].open(rpath);
				string data;
				KeyValue kv;
				while (std::getline(infiles[i], data)) {
					//std::cout << "data:" << data << std::endl;
					kv = fromJson(data);
					auto iter = contentsMp.find(kv.key);
					if(iter == contentsMp.end()) {
						std::vector<string> tmp;
						tmp.push_back(kv.value);
						contentsMp[kv.key] = tmp;
					} else {
						iter->second.push_back(kv.value);
					}
				}
			    }

			    string mn = mergeName(task->dataDir, task->jobName, task->taskNumber);
			    ofstream resultFile;
			    resultFile.open(mn);
			    for (auto it = contentsMp.begin(); it != contentsMp.end(); it++) {
				    string result = task->reduceF(it->first, it->second);
				    resultFile << result << std::endl;
			    }
			    resultFile.close();
			    task->done = true;
		    }
                    _idlThrNum++;
                }
            });
            _idlThrNum++;
        }
    }

public:
    std::vector<string> submit(std::string jobName, std::string dataDir, MapF mapF, ReduceF reduceF, std::vector<string> mapFiles, int nReduce);
};

#endif
