This repository contains a Go program that demonstrates memory usage when working with maps in Go. Specifically, it compares the memory consumption of maps using fixed-size arrays versus maps using pointers to arrays. The code evaluates how memory usage changes after allocating a map, adding elements, deleting elements, and running garbage collection (GC).

Overview
The program explores two scenarios:

Maps with Fixed-Size Arrays: map[int][128]byte
Maps with Pointers to Arrays: map[int]*[128]byte
The goal is to observe and compare the memory consumption in both cases. The program includes a simple benchmark to measure memory usage at different stages of map manipulation.

Code Explanation
The program performs the following steps:

Map Allocation: Creates an empty map.
Adding Elements: Inserts 10 million elements into the map.
Deleting Elements: Removes 5 million elements from the map.
Running GC: Manually triggers garbage collection to clean up unused memory.
Memory Measurement: Prints the memory usage at each step.
