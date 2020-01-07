from threading import Thread

i = 0

def increasing():
  global i 
  for(j=0; j < 1000000; j++)
      i++
   
   
 def decreasing():
  global i
  for(j=0; j < 1000000; j++)
      i--
      
      
 def main():

    thread_1 = Thread(target = increasing(), args = (),)
    thread_2 = Thread(target = decreasing(), args = (),)
    
    thread1.start()
    thread2.start()
    
    thread1.join()
    thread2.join()
    
    print(i)


main()
