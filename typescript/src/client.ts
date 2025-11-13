import type { AxiosInstance } from "axios";
import axios from "axios";
import * as https from "https";

export class ProxmoxClient {
	protected axiosInstance: AxiosInstance;
	protected uuid: string;

	constructor(baseUrl: string, apiToken: string, uuid: string) {
		// Remove trailing slash from baseUrl
		baseUrl = baseUrl.replace(/\/$/, "");

		this.axiosInstance = axios.create({
			baseURL: `${baseUrl}/api2/json`,
			headers: {
				Authorization: apiToken,
			},
			httpsAgent: new https.Agent({ rejectUnauthorized: false }),
		});
		this.uuid = uuid;
	}
}
