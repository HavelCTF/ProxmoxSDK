import type {
	Int as EInt,
	Number as ENumber,
	String as EString,
} from "effect/Schema";

export interface NodesResponse {
	nodes: Node[];
}

interface Node {
	node: EString;
	status: NodeStatus;
	uptime?: EInt;
	sslfingerprint?: EString;
	cpu?: ENumber;
	level?: EString;
	maxcpu?: EInt;
	maxmem?: EInt;
	mem?: EInt;
}

enum NodeStatus {
	Unknown = "unknown",
	Online = "online",
	Offline = "offline",
}
