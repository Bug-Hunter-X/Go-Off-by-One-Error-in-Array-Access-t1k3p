# Go Off-by-One Error Example

This repository demonstrates a common off-by-one error in Go when accessing array elements.  The code attempts to access an array element beyond its bounds, leading to a runtime panic. The solution shows how to correct the loop to avoid this error.

## Bug

The `bug.go` file contains the buggy code. The for loop iterates one element beyond the array's valid index range.

## Solution

The `bugSolution.go` file provides the corrected code. The for loop is modified to iterate within the valid index range of the array, preventing the out-of-bounds access.