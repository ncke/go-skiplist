# go-skiplist
A skip list data structure in go.
## What is a skip list?
A skip list is a probabilistic data structure based on linked lists. It provides O(log n) search complexity and O(log n) insertion complexity, which is to say that it is a fast map. Skip lists do not require re-balancing after insertion so they remain available for concurrent access. Similar data structures, such as the binary tree, must be locked while re-balancing occurs.
## See also.
[1] Open Data Structures (n.d.) 'Skiplists'. Available at https://opendatastructures.org/newhtml/ods/latex/skiplists.html (Accessed 18 August 2022).

[2] Pugh, W. (1990) 'Skip Lists: a probabilistic alternative to balanced trees', Communications of the ACM, 33(6), pp. 668-676. Available at https://dl.acm.org/doi/pdf/10.1145/78973.78977 (Accessed 18 August 2022).