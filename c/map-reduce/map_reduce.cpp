#include "map_reduce.h"
#include <iostream> 

std::string toJson(KeyValue& kv) {
	Json::Value root;
	root["key"] = kv.key;
	root["value"] =  kv.value;

	Json::StreamWriterBuilder builder;
	builder["indentation"] = ""; 
	const std::string json_file = Json::writeString(builder, root);
	return json_file;
}

KeyValue fromJson(string& str) {
	KeyValue kv;
	memset(&kv,0,sizeof(KeyValue));

	Json::CharReaderBuilder builder;
	Json::CharReader * reader = builder.newCharReader();
	Json::Value value;
	string errors;
	bool parsingSuccessful = reader->parse(str.c_str(), str.c_str() + str.size(), &value, &errors);
	delete reader;
	if ( !parsingSuccessful )
	{
		std::cout << str << endl;
		std::cout << errors << endl;
	}

	//std::cout << "enter 1" << std::endl;
	kv.key=value["key"].asString();
	//std::cout << "enter 2" << std::endl;
	kv.value=value["value"].asString();
	//std::cout << "enter 3" << std::endl;

	return  kv;
}


