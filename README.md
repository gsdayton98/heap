# heap --  immutable go heap

Provides operations to re-arrange a slice into a heap, pop an entry from the
heap, and push an entry onto the heap.  Differs from the normal go heap package
in that the underlying container is not changed.  When a push or pop operation
changes the size of the underlying container, it returns a new container.
