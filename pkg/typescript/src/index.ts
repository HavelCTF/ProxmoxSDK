import * as fs from 'node:fs';
import * as path from 'node:path';
import { ClusterService } from './modules/cluster';
import { getNodes, NodeService } from './modules/nodes';
import { getVersion } from './modules/version';
import { loadWasm } from './wasm/loader';

export * from './types/cluster';
export * from './types/lxc';
export * from './types/nodes';
export * from './types/storage';
export * from './types/tasks';
export * from './types/version';

export interface ProxmoxSDKOptions {
    insecure?: boolean;
}

export class ProxmoxSDK {
    async getVersion() {
        return getVersion();
    }

    async getNodes() {
        return getNodes();
    }

    cluster(): ClusterService {
        return new ClusterService();
    }

    node(nodeName: string): NodeService {
        return new NodeService(nodeName);
    }

    static async create(host: string, token: string, uuid: string, options?: ProxmoxSDKOptions): Promise<ProxmoxSDK> {
        if (options?.insecure && typeof process !== 'undefined') {
            process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';
        }

        const wasmPath = path.join(__dirname, 'main_wasm.wasm');
        const wasmSource = fs.readFileSync(wasmPath);

        await loadWasm(wasmSource);
        globalThis.ProxmoxWASM.initClient(host, token, uuid);

        return new ProxmoxSDK();
    }
}
