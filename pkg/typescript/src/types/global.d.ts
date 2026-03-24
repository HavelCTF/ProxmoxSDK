declare global {
    var ProxmoxWASM: {
        initClient: (host: string, token: string, uuid: string) => boolean;
        cluster: {
            GetNextId: () => Promise<{ VMID: string }>;
            GetTasks: () => Promise<{ Data: { UPID: string }[] }>;
        };
        version: {
            GetVersion: () => Promise<{ Data: { Version: string; Release: string; RepoID: string; Console: string; } }>;
        };
    };

    var Go: {
        new (): {
            importObject: WebAssembly.Imports;
            run(instance: WebAssembly.Instance): Promise<void>;
        };
    };
}

export {};