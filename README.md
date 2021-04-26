# GoMemoize
Very basic demonstration of concurrency in Go by calculating the fibonacci sequence using memoization, worker threads, and shared memory. Memoization is the process of remembering
already calculated values into a datastructure (in my example, a map) and checking to see if a certain value has been calculated already, if it has not, add it to the datastructure.
If so, use the already calculated value in the next calculation. This removes any redundant calculations that are encountered when computing smaller subproblems. 
