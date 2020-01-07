from threading import Thread

i = 0

def increasing():
  global i 
  for j in range(1000000):
      i += 1
   
   
 def decreasing():
  global i
  for j in range(1000000):
      i -= 1
      
      
 def main():

    thread_1 = Thread(target = increasing, args = (),)
    thread_2 = Thread(target = decreasing, args = (),)
    
    thread1.start()
    thread2.start()
    
    thread1.join()
    thread2.join()
    
    print(i)


main()
