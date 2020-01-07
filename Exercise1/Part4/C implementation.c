#include <pthread.h>
#include <stdio.h>


int i = 0;

//increase
void* thread_1(){ 
   for (int j = 0; j < 1000000; j++){
       i++;
   }
   return NULL;
}

//decrease
void* thread_2(){ 
   for (int j = 0; j < 1000000; j++){
      i--;
   }
   return NULL;
}




int main(){
  
  pthread_t thread_1;
  phtread_t thread_2;

  pthread_create(&thread_1, NULL, thread_1(), NULL);
  pthread_create(&thread_2, NULL, thread_2(), NULL);
  
  pthread_join(thread_1, NULL);
  pthread_join(thread_2, NULL);
  
  printf("%d\n", i);
  
  return 0;
}
