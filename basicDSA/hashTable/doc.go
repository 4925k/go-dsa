/*
The Hash table data structure stores elements in key-value pairs where

Key- unique integer that is used for indexing the values
Value - data that are associated with keys.

In a hash table, a new index is processed using the keys. And, the element corresponding to that key is stored in the index. This process is called hashing.
Let k be a key and h(x) be a hash function.
Here, h(k) will give us a new index to store the element linked with k.

When the hash function generates the same index for multiple keys, there will be a conflict (what value to be stored in that index). This is called a hash collision.
We can resolve the hash collision using one of the following techniques.
- Collision resolution by chaining
- Open Addressing: Linear/Quadratic Probing and Double Hashing
*/
package hashtable
