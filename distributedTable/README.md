# Distributed Table

This is an implementation for a distributed table, storing Potatos (see ../statefulTable for more info). Server nodes gossip about new writes, updating their own prospective states. Locks are placed on state updates, and communicated during read/write requests as to prevent conflicts.  
