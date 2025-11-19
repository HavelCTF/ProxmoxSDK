import { Int, Number, String } from "effect/Schema"

export interface NodesResponse {
	nodes : Node[];
}

interface Node {
	node : String
	status : NodeStatus
	uptime? : Int
	sslfingerprint? : String
	cpu? : Number
	level? : String
	maxcpu? : Int
	maxmem? : Int
	mem? : Int
}

enum NodeStatus {
	Unknown = "unknown",
	Online = "online",
	Offline = "offline",
}
