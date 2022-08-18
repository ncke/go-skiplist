# go-skiplist
A skip list data structure in go.
## What are skip lists?
A skip list is a probabilistic data structure based on linked lists. It provides O(log n) search complexity and O(log n) insertion complexity, which is to say that it is a fast map. Skip lists do not require re-balancing after insertion so they remain available for concurrent access. Similar data structures, such as the binary tree, must be locked while re-balancing occurs.
## See also.
[1] Open Data Structures (n.d.) 'Skiplists'. Available at https://opendatastructures.org/newhtml/ods/latex/skiplists.html (Accessed 18 August 2022).
