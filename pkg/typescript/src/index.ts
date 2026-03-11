import { loadWasm } from './wasm/loader';
import { ClusterModule } from './modules/cluster';

export * from './types/cluster';

export class ProxmoxSDK {
    public readonly cluster = new ClusterModule();

    static async create(host: string, token: string, uuid: string, wasmSource: string | Buffer | Uint8Array): Promise<ProxmoxSDK> {
        await loadWasm(wasmSource);
        globalThis.ProxmoxWASM.initClient(host, token, uuid);
        return new ProxmoxSDK();
    }
}