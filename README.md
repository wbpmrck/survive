##what is this

this is a part-time project,it can be a turn-base or real-time SLG game server.

####TODO

* sandbox

    > build basic sandbox machanism
    
    > use javascript binding
  
    
* rpc

    > define interfaces of io message actions(uncouple the "client" and link library)
    
    > use link as default rpc support
    
    > define router:
    
    >> router take message input from rpcClient,and decide where the message to dispatch:to a local messageQueue,or a remote
     endpoint
     
     >> router may will have to encapsulate the AgentMessage into a package,which has route information
     (rpc feature should consider in client side(maybe in a driver),server should not depend on a clients response(in
     a time slice,do this will block other options,waiting is expensive)
     
* demo

    > use console and channel to simulate client,"connect" to server, and send/recv messages
    
* data persistence

      
    > SQLite support
    
    > Redis support
    
    > MongoDB support 