# Hand-In2

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
Packages in our implementation are fictional, sent and represented by our sequence numbers, which the serverside uses to acknowledge that packages are recieved. In terms of data structures, we are using two main functions named server and client, using four channel functions who are responsible for the server-acknowledgement, server-sequence, client acknowledgement and client-sequence
these channels run concurrently, simulation two systems talking to eachother using integers.


b) Does your implementation use threads or processes? Why is it not realistic to use threads?
Our implementation uses threads.
the reason why threads arent realistic, and especially without a middleware, is because packet-loss and delay cannot be simulated correctly. the threads running concurrently can be hardcoded into doing exactly as planned on every run, in contrast to two systems running completely indemendently 
from eachother.


c) How do you handle message re-ordering?
Our implementation does not handle re-ordering since it was made without a middleware, but in the case we had one, we would be able to make it using a reorder request through the middleware, which would be acknowledge in the clientside, and then sent through the middleware to the server, which would then acknowledge recieving it, to then resume and acknowledge the the new incomming packages


d) How do you handle message loss?
same as before, we do not handle it in this implementation, but it could be done using a timeout, where if the message was not recieved within a certain timeframe, the message would be considered "lost" and a request for a re-order would be sent through the middleware to the clientside, such that the package would be resent.


e) Why is the 3-way handshake important?
The 3-way handshake is important because before packages can be sent, we have to make sure that the connection is established corretly. The handshake lets both parties establish their sequence and acknowledgements to eachother, such that they are syncronised, which is used when sending the packages to make sure that they are ordered correctly, and no package is lost or wrong in any way.