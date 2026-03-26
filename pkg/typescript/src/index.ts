import * as fs from 'fs';
import * as path from 'path';
import { ClusterModule } from './modules/cluster';
import { LxcModule } from './modules/lxc';
import { NodesModule } from './modules/nodes';
import { TasksModule } from './modules/tasks';
import { VersionModule } from './modules/version';
import { loadWasm } from './wasm/loader';

export * from './types/cluster';
export * from './types/lxc';
export * from './types/nodes';
export * from './types/tasks';
export * from './types/version';

export interface ProxmoxSDKOptions {
    insecure?: boolean;
}

export class ProxmoxSDK {
    public readonly cluster = new ClusterModule();
    public readonly version = new VersionModule();
    public readonly nodes = new NodesModule();
    public readonly tasks = new TasksModule();
    public readonly lxc = new LxcModule();

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
