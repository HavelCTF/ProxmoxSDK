import { ClusterModule } from './modules/cluster';
import { VersionModule } from './modules/version';
import { NodesModule } from './modules/nodes';
import { loadWasm } from './wasm/loader';

export * from './types/cluster';
export * from './types/version';
export * from './types/nodes';

export interface ProxmoxSDKOptions {
    insecure?: boolean;
}

export class ProxmoxSDK {
    public readonly cluster = new ClusterModule();
    public readonly version = new VersionModule();
    public readonly nodes = new NodesModule();

    static async create(
        host: string,
        token: string,
        uuid: string,
        wasmSource: string | Buffer | Uint8Array,
        options?: ProxmoxSDKOptions,
    ): Promise<ProxmoxSDK> {
        if (options?.insecure && typeof process !== 'undefined') {
            process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';
        }
        await loadWasm(wasmSource);
        globalThis.ProxmoxWASM.initClient(host, token, uuid);
        return new ProxmoxSDK();
    }
}
