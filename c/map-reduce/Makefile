CC = gcc
CXX = g++
AS = yasm
CTAGS = ctags
INCLUDES=-Iinclude
LIBS=

CFLAGS = -g -Wall -fPIC -c -std=c++11  $(INCLUDES)
OFLAGS =  -std=c++11 -g -Wall  

SRCDIR = .
SRC_JSON_DIR = src/lib_json
SRCS = $(wildcard $(SRCDIR)/*.cc $(SRCDIR)/*.cpp $(SRC_JSON_DIR)/*.cc $(SRC_JSON_DIR)/*.cpp)
OBJS = $(addsuffix .o, $(basename $(SRCS)))
HEADS = $(addsuffix .h, $(basename $(SRCS)))

APP = mr
all : $(APP)

show :
	@echo $(SRCS)

$(APP) : $(OBJS)
	$(CXX) $(OFLAGS) -o $@ $^ $(LIBS)

%.o:%.cc
	$(CXX) $(CFLAGS) -o $@ $<

%.o:%.c
	$(CC) $(CFLAGS) -o $@ $<

%.o:%.cpp
	$(CXX) $(CFLAGS) -o $@ $<

clean :
	rm -Rf *.o $(APP) $(SRCDIR)/*.o $(SRC_JSON_DIR)/*.o 

cleand :
	rm  ./data/mrtmp/*


