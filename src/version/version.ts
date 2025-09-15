import { Effect } from "effect";
import { ProxmoxClient } from "../client.js";

interface VersionResponse {
	version: string;
	release: string;
	repoid: string;
}

declare module "../client.js" {
	interface ProxmoxClient {
		version(): Effect.Effect<VersionResponse, Error, never>;
	}
}

ProxmoxClient.prototype.version = function () {
	return Effect.tryPromise({
		try: async () => {
			const response = await this.axiosInstance.get("/version");
			return response.data;
		},
		catch: (error) => {
			return new Error(`[${this.uuid}] Version request failed: ${error}`);
		},
	}).pipe(
		Effect.retry({ times: 3 }),
		Effect.timeout(10000),

		Effect.tapError((error) => Effect.sync(() => console.error(error.message))),
	);
};
