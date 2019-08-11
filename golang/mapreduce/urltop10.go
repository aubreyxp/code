package main

import (
	"fmt"
	"strconv"
	"strings"
)

// URLTop10 .
func URLTop10(nWorkers int) RoundsArgs {
	// YOUR CODE HERE :)
	// And don't forget to document your idea.
	var args RoundsArgs
	// round 1: do url count
	args = append(args, RoundArgs{
		MapFunc:    URLCountMap,
		ReduceFunc: URLCountReduce,
		NReduce:    nWorkers,
	})
	// round 2: sort and get the 10 most frequent URLs
	args = append(args, RoundArgs{
		MapFunc:    URLTop10Map,
		ReduceFunc: ExampleURLTop10Reduce,
		NReduce:    1,
	})
	return args
}

func URLCountMap(filename string, contents string) []KeyValue {
	lines := strings.Split(string(contents), "\n")
	mp := map[string]int{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		mp[l]++
	}

	kvs := make([]KeyValue, 0, len(mp))
	kv := KeyValue{}
	for k, v := range mp {
		kv.Key = k
		kv.Value = strconv.Itoa(v)
		kvs = append(kvs, kv)
	}
	return kvs
}

func URLCountReduce(key string, values []string) string {
	sum := 0
	for _, value := range values {
		count, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		sum += count
	}
	return fmt.Sprintf("%s %d\n", key, sum)
}

func URLTop10Map(filename string, contents string) []KeyValue {
	kvs := []KeyValue{}
	lines := strings.Split(string(contents), "\n")
	cnts := make(map[string]int, len(lines))
	for _, line := range lines {
		//fmt.Println("see line:", line)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		tmp := strings.Split(line, " ")
		//fmt.Println("see tmp:", tmp)
		//continue
		//fmt.Println(tmp)
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		cnts[tmp[0]] = n
	}

	fmt.Printf("top 10 will ....\n")
	us, cs := TopN(cnts, 10)
	kv := KeyValue{}
	for i := range us {
		kv.Key = ""
		kv.Value = fmt.Sprintf("%s %d", us[i], cs[i])
		kvs = append(kvs, kv)
	}
	return kvs
}

/*
func URLTop10Reduce(key string, values []string) string {
}
*/
