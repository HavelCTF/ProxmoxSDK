import { loadWasm } from './wasm/loader';
import { ClusterModule } from './modules/cluster';

export * from './types/cluster';

export interface ProxmoxSDKOptions {
    insecure?: boolean;
}

export class ProxmoxSDK {
    public readonly cluster = new ClusterModule();

    static async create(host: string, token: string, uuid: string, wasmSource: string | Buffer | Uint8Array, options?: ProxmoxSDKOptions): Promise<ProxmoxSDK> {
        if (options?.insecure && typeof process !== 'undefined') {
            process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';
        }
        await loadWasm(wasmSource);
        globalThis.ProxmoxWASM.initClient(host, token, uuid);
        return new ProxmoxSDK();
    }
}