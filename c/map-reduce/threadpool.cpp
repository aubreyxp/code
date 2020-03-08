#include "threadpool.h"

std::vector<string> ThreadPool::submit(std::string jobName, std::string dataDir, MapF mapF, ReduceF reduceF, std::vector<string> mapFiles, int nReduce){
	auto size = mapFiles.size();
	std::vector<string> rets;
	if (size == 0) {
		return rets; 
	}
	
	std::cout << "map size:" << size << std::endl;

	std::vector<Task*> mTaskList;
	for (int i = 0; i < size; i++) {
		auto t = new(Task);
		t->dataDir = dataDir;
		t->jobName = jobName;
		t->mapFile = mapFiles[i];
		t->phase = MapPhase;
		t->taskNumber = i;
		t->nReduce = nReduce;
		t->nMap = size;
		t->mapF = mapF;
		t->done = false;
		mTaskList.push_back(t);

		{
			std::lock_guard<std::mutex> lock{ _lock };
			_tasks.push(t);
		}
		_task_cv.notify_one();  

	}

         // wait for all map tasks to complete
	for (;;) {
		auto it = mTaskList.begin();
		bool over = true;
		for (; it != mTaskList.end(); it++) {
			if ((*it)->done == false) {
				over = false;
				break;
			}
		}

		if (!over) { 
			std::this_thread::sleep_for(std::chrono::seconds(1));
		} else {
			break;
		}
	}
	
	// free memory	
	auto it = mTaskList.begin();
	for (; it != mTaskList.end(); it++) {
		delete *it;
	}

	// create reduce task
	std::vector<Task*> rTaskList;
	for (int i = 0; i < nReduce; i++) {
		auto t = new(Task);
		t->dataDir = dataDir;
		t->jobName = jobName;
		t->phase = ReducePhase;
		t->taskNumber = i;
		t->nMap = size;
		t->reduceF = reduceF;
		t->done = false;
		rTaskList.push_back(t);

		{
			std::lock_guard<std::mutex> lock{ _lock };
			_tasks.push(t);
		}
		_task_cv.notify_one();  
	}

	// wait for all reduce tasks to complete
	for (;;) {
		auto it = rTaskList.begin();
		bool over = true;
		for (; it != rTaskList.end(); it++) {
			if ((*it)->done == false) {
				over = false;
				break;
			}
		}

		if (!over) {
			std::this_thread::sleep_for(std::chrono::milliseconds(1));
		} else {
			break;
		}
	}
	
	// get all reduce path
	for (int i = 0; i < nReduce; i++) {
		string mn = mergeName(rTaskList[i]->dataDir, rTaskList[i]->jobName, rTaskList[i]->taskNumber);
		rets.push_back(mn);
	}
	
	// free memory	
	{
		auto it = rTaskList.begin();
		for (; it != rTaskList.end(); it++) {
			delete *it;
		}
	}

	return rets;
}


