/*
A linked list is a linear data structure that includes a series of connected nodes.
each node stores the data and the address of the next node

The power of a linked list comes from the ability to break the chain
and rejoin it.
E.g. if you wanted to put an element 4 between 1 and 2,
the steps would be:

Create a new struct node and allocate memory to it.
Add its data value as 4
Point its next pointer to the struct node containing 2 as the data value
Change the next pointer of "1" to the node we just created.

Linked List Applications
Dynamic memory allocation
Implemented in stack and queue
In undo functionality of softwares
Hash tables, Graphs
*/
package singly
