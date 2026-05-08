declare global {
    var ProxmoxWASM: {
        initClient: (host: string, token: string, uuid: string) => boolean;
        cluster: {
            GetNextId: () => Promise<{ VMID: string }>;
            GetTasks: () => Promise<{ Data: { UPID: string }[] }>;
        };
        version: {
            GetVersion: () => Promise<{ Data: { Version: string; Release: string; RepoID: string; Console: string } }>;
        };
        nodes: {
            GetNodes: () => Promise<{
                Data: {
                    Node: string;
                    Status: string;
                    Uptime?: number;
                    SslFingerprint?: string;
                    CPU?: number;
                    Level?: string;
                    MaxCPU?: number;
                    MaxMEM?: number;
                    MEM?: number;
                }[];
            }>;
        };
        tasks: {
            GetTasks: (node: string) => Promise<{
                Data: {
                    UPID: string;
                    Node?: string;
                    PID?: number;
                    PStart?: number;
                    StartTime?: number;
                    Type?: string;
                    User?: string;
                    Status?: string;
                }[];
            }>;
            GetTaskStatus: (
                node: string,
                upid: string,
            ) => Promise<{
                Data: {
                    ID: string;
                    Node: string;
                    PID: number;
                    PStart: number;
                    StartTime: number;
                    Type: string;
                    UPID: string;
                    User: string;
                    Status: string;
                    ExitStatus?: string;
                };
            }>;
            DeleteTask: (node: string, upid: string) => Promise<{ Data: string }>;
            Wait: (
                node: string,
                upid: string,
                opts?: { PollIntervalMs?: number; TimeoutMs?: number },
            ) => Promise<{
                Data: {
                    ID: string;
                    Node: string;
                    PID: number;
                    PStart: number;
                    StartTime: number;
                    Type: string;
                    UPID: string;
                    User: string;
                    Status: string;
                    ExitStatus?: string;
                };
            }>;
        };
        lxc: {
            GetLXCs: (node: string) => Promise<{
                LXCs: {
                    Status: string;
                    VMID: number;
                    CPU?: number;
                    CPUS?: number;
                    Disk?: number;
                    DiskRead?: number;
                    DiskWrite?: number;
                    Lock?: string;
                    MaxDisk?: number;
                    MaxMem?: number;
                    MaxSwap?: number;
                    Mem?: number;
                    Name?: string;
                    NetIn?: number;
                    NetOut?: number;
                    PressureCPUSome?: string;
                    PressureIOFull?: string;
                    PressureIOSome?: string;
                    PressureMemoryFull?: string;
                    PressureMemorySome?: string;
                    Tags?: string;
                    Template?: boolean;
                    Uptime?: number;
                }[];
            }>;
            PostLXC: (
                node: string,
                data: { OSTemplate: string; VMID: number; Features?: { Nesting: boolean } },
            ) => Promise<{ Data: string }>;
            StartLXC: (node: string, vmid: number) => Promise<{ Data: string }>;
            StopLXC: (node: string, vmid: number) => Promise<{ Data: string }>;
            DeleteLXC: (node: string, vmid: number) => Promise<{ Data: string }>;
            CloneLXC: (
                node: string,
                vmid: number,
                data: { NewId: number; Target?: string },
            ) => Promise<{ Data: string }>;
            Get: (
                node: string,
                vmid: number,
            ) => Promise<{
                Data: {
                    Name?: string;
                    Status: string;
                    VMID?: number;
                    Uptime?: number;
                    CPUs?: number;
                    CPU?: number;
                    Mem?: number;
                    MaxMem?: number;
                    Disk?: number;
                    MaxDisk?: number;
                    Swap?: number;
                    MaxSwap?: number;
                    NetIn?: number;
                    NetOut?: number;
                    DiskRead?: number;
                    DiskWrite?: number;
                    Lock?: string;
                    Tags?: string;
                    Type?: string;
                };
            }>;
        };
        storage: {
            GetContent: (
                node: string,
                storage: string,
                contentType?: string,
            ) => Promise<{
                Data: {
                    VolID: string;
                    Content: string;
                    Format?: string;
                    Size?: number;
                    Used?: number;
                    CTime?: number;
                    VMID?: number;
                    Notes?: string;
                    Parent?: string;
                }[];
            }>;
            UploadTemplate: (
                node: string,
                storage: string,
                filename: string,
                body: Uint8Array,
            ) => Promise<{ Data: string }>;
        };
    };

    var Go: {
        new (): { importObject: WebAssembly.Imports; run(instance: WebAssembly.Instance): Promise<void> };
    };
}

export {};
