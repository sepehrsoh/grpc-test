# grpc-test

grpc-test is a simple gRPC server and client implementation that utilizes a proto file to generate code for Go, JavaScript, and TypeScript and etc. The project also includes a REST gateway, which wraps around the gRPC server, enabling RESTful requests. Additionally, it uses Buf to manage protocol buffers.

## Usage

Follow the steps below to build and run the gRPC server and client:

### Prerequisites

Before you begin, ensure you have the following installed on your system:

1. Go (https://golang.org/)
2. Buf (https://docs.buf.build/installation)

### Building and Running the Server

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/grpc-test.git
   ```

2. Use the provided Makefile to build the server and client:

   ```bash
   make all
   ```

3. Run the gRPC server:

   ```bash
   ./grpc-server
   ```

   The server will be up and running, listening on port 8000 for gRPC requests and port 9000 for RESTful requests.

### RESTful Requests

The REST gateway wraps around the gRPC server, allowing you to make RESTful requests instead of gRPC calls directly. You can use tools like `curl` or Postman to interact with the server.

For example, to make a POST request to the server using `curl`:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"Name": "Hello gRPC!"}' http://localhost:9000/hi
```

### Generating Code

The project uses protocol buffers, and Buf is used to manage the proto files and code generation. Make sure to use Buf to lint and generate the proto files as needed.

## Contributing

Feel free to contribute to this project by submitting pull requests or reporting issues. Your contributions are greatly appreciated!

## License

This project is licensed under the [MIT License](LICENSE).

Happy gRPC testing!
