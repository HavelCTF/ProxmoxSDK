declare global {
    var ProxmoxWASM: {
        initClient: (host: string, token: string, uuid: string) => boolean;
        cluster: {
            GetNextId: () => Promise<unknown>;
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