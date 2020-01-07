### C code:
- Uses the module pthread where the pthread class is definied to run concurrent operations that increase and decrease i


### Python code:
- Uses the Thread module to bypass GIL such that the function increasing and decreasing run concurrently

### Go code:
- Can be solved in different ways, but I've used functions that take in two different channels and executing two goroutines such that both functions run at the same time till the threshold of the two functions are satisfied
