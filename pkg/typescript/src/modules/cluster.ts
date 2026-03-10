import https from 'https';
import { ClusterNextIdResponse } from '../types/cluster';

// In Node we can disable TLS certificate validation globally using env var
if (typeof process !== 'undefined' && process.env) {
    process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';
}

// create an agent for Node just in case fetch honors it
const nodeAgent = ((): https.Agent | undefined => {
    if (typeof process !== 'undefined' && process.release && process.release.name === 'node') {
        return new https.Agent({ rejectUnauthorized: false });
    }
    return undefined;
})();

export class ClusterModule {
    private host: string;
    private token: string;

    constructor(host: string, token: string) {
        this.host = host.replace(/\/+$/g, ''); // trim trailing slash
        this.token = token;
    }

    public async GetNextId(): Promise<ClusterNextIdResponse> {
        const url = `${this.host}/api2/json/cluster/nextid`;
        // build init object separately to avoid TS complaining about `agent`
        const init: any = {
            headers: {
                Authorization: this.token,
                Accept: "application/json",
            },
        };
        if (nodeAgent) {
            init.agent = nodeAgent;
        }
        const res = await fetch(url, init);

        if (!res.ok) {
            const body = await res.text();
            throw new Error(`HTTP ${res.status}: ${body}`);
        }

        const body = await res.json();
        // Proxmox returns { data: "<vmid>" }
        const vmid = body?.data;
        if (typeof vmid !== 'string') {
            throw new Error("unexpected response: " + JSON.stringify(body));
        }
        const result: ClusterNextIdResponse = { id: vmid };
        return result;
    }
}