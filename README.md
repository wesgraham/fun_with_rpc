# Fun With JSON-RPC
**A repository full of personal projects using JSON-RPC's over HTTP.**

The first project, `incriment points`, allows a client to submit an RPC to a server to incriment and decriment the price of an item.

The second project `stateful table`, allows a client to submit CRUD requests on a table full of Potatoes maintained by a server. The server gracefully shuts down when it experiences a sigint or sigterm, saving state into a file called "file.tmp". This file is loaded into state before our server is rebooted.  
