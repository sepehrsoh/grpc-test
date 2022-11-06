import * as grpcWeb from 'grpc-web';

import * as messages_hello_pb from '../messages/hello_pb';


export class HelloClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  sayHello(
    request: messages_hello_pb.HelloRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: messages_hello_pb.HelloReply) => void
  ): grpcWeb.ClientReadableStream<messages_hello_pb.HelloReply>;

}

export class HelloPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  sayHello(
    request: messages_hello_pb.HelloRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<messages_hello_pb.HelloReply>;

}

