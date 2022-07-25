package hard

// Question: flip one cell from '0' to '1' and make a large island

// Solution:
// 1. Go through all grids and count the size of each island
// 2. the counter should be mapped into a unique key for all connected cells
// 3. Go through all grids again to find cell connecting two or more islands
// 4. And then find the maximum size of connection island

// ps: Use a grid to mark the same island with a incremental index
