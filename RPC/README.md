### Links

- [RPC Go client and server](https://medium.com/@shivambhadani_/introduction-to-rpc-in-go-building-rpc-client-and-server-with-golang-5794675e9a12)

- [RPC vs REST](https://stackoverflow.com/questions/75439599/rpc-vs-restful-endpoints)

- [RPC vs REST (Better version)](https://aws.amazon.com/compare/the-difference-between-rpc-and-rest/)
  - In Remote Procedure Call (RPC), the client uses HTTP **POST** to call a specific function by name. Client-side developers must know the function name and parameters in advance for RPC to work.
  - In REST, clients and servers use HTTP verbs like **GET**, **POST**, **PATCH**, **PUT**, **DELETE**, and **OPTIONS** to perform operations. Developers only need to know the server resource URLs and don't have to be concerned with individual function names.
  - Modern implementations of RPC, such as gRPC, are now more popular. For some use cases, gRPC performs better than RPC and REST. It allows **streaming** client-server communications rather than the **request-and-response** data exchange pattern.
  - REST can pass any data format and multiple formats, like JSON and XML, within the same API. However, with RPC APIs, the data format is selected by the server and fixed during implementation. You can have specific JSON RPC or XML RPC implementations, and the client has no flexibility. _(REST cannot use protobuf)_
  - REST systems must always be stateless, but RPC systems can be stateful or stateless, depending on design.

- [About RPC](https://medium.com/@lelianto.eko/understanding-rpc-and-how-it-works-a-case-study-of-grpc-framework-6e7235cd3d5f)
  - The packaged data is sent to the server over a network protocol, such as HTTP or TCP. _(Both REST (is) and RPC (can be) are some framework over HTTP)_
