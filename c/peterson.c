#include<stdio.h>
#include <unistd.h>
#include <stdbool.h>
#include <pthread.h>

bool flag[2];
int turn;
volatile int count;

void enter(int process_id) {
    flag[process_id] = true;
    int other = 1 - process_id;

    turn = other;

    while (flag[other] == true && turn == other)
    {
        usleep(1);// busy wait
    }
}

void leave(int process_id) {
    flag[process_id] = false;
}

void* parallel(void *arg){
        int id = *((int*)arg);
        for (int i = 0; i < 100000; i++) {
                enter(id);
                count++;
                leave(id);
        }
        return 0;
}

int main(void)
{
    pthread_t ntid_0;
    int process_id_0 = 0;
    pthread_create(&ntid_0,NULL,parallel,&process_id_0);

    pthread_t ntid_1;
    int process_id_1 = 1;
    pthread_create(&ntid_1,NULL,parallel,&process_id_1);

    pthread_join(ntid_0,NULL);
    pthread_join(ntid_1,NULL);

    printf("count count:%d\n", count); // something wrong, not always right

    return 0;
 }
