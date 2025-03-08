import * as net from "net";

interface RPCRequest {
  id: number;
  method: string;
  params: any[];
}

interface RPCResponse {
  id: number;
  result?: any;
  error?: string;
}

function rpcCall(method: string, params: any[]): Promise<any> {
  return new Promise((resolve, reject) => {
    const client = new net.Socket();
    client.connect(3000, "localhost", () => {
      console.log(`Connected to RPC server, calling ${method}...`);
      const request: RPCRequest = {
        id: 1, //TODO: must change
        method,
        params,
      };
      client.write(JSON.stringify(request));
    });

    client.on("data", (data) => {
      const response: RPCResponse = JSON.parse(data.toString().trim());
      if (response.error) {
        reject(new Error(response.error));
      } else {
        resolve(response.result);
      }
      client.destroy();
    });

    client.on("error", (err) => reject(err));
    client.on("close", () => console.log("Connection closed."));
  });
}

(async () => {
  try {
    const multiplyResult = await rpcCall("Calculator.Multiply", [
      { A: 6, B: 7 },
    ]);
    console.log("Multiply Result:", multiplyResult);

    const divideResult = await rpcCall("Calculator.Divide", [{ A: 10, B: 2 }]);
    console.log("Divide Result:", divideResult);

    const HelloResult = await rpcCall("Greeter.SayHello", [{ H: "Shy" }]);
    console.log("Hello Result:", HelloResult);

    const ByeResult = await rpcCall("Greeter.SayBye", [{ B: "Gun" }]);
    console.log("Bye Result:", ByeResult);
  } catch (error) {
    console.error("RPC Error:", error);
  }
})();
