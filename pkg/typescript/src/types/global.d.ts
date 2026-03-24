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
        nodes: {
            GetNodes: () => Promise<{
                Data: {
                    Node: string; Status: string; Uptime?: number; SslFingerprint?: string;
                    CPU?: number; Level?: string; MaxCPU?: number; MaxMEM?: number; MEM?: number;
                }[]
            }>;
        };
        tasks: {
            GetTaskStatus: (node: string, upid: string) => Promise<{
                Data: {
                    ID: string; Node: string; PID: number; PStart: number; StartTime: number;
                    Type: string; UPID: string; User: string; Status: string; ExitStatus?: string;
                }
            }>;
            DeleteTask: (node: string, upid: string) => Promise<{ Data: string }>;
        };
    };

    var Go: {
        new (): { importObject: WebAssembly.Imports; run(instance: WebAssembly.Instance): Promise<void>; };
    };
}

export {};