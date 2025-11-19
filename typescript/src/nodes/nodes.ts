import { Effect } from "effect";
import { ProxmoxClient } from "../client.js";
import { NodesResponse } from "./node-schema.js"

declare module "../client.js" {
	interface ProxmoxClient {
		nodes(): Effect.Effect<NodesResponse, Error, never>;
	}
}

ProxmoxClient.prototype.nodes = function () {
	return Effect.tryPromise({
		try: async () => {
			const response = await this.axiosInstance.get("/nodes");
			return response.data;
		},
		catch: (error) => {
			return new Error(`[${this.uuid}] Nodes request failed: ${error}`);
		},
	}).pipe(
		Effect.retry({ times: 3 }),
		Effect.timeout(10000),

		Effect.tapError((error) => Effect.sync(() => console.error(error.message))),
	);
};
