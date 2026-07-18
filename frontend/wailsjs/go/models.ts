export namespace models {
	
	export class Host {
	    IP: string;
	    MAC: string;
	    Hostname: string;
	    Vendor: string;
	    Online: boolean;
	    RTT: number;
	
	    static createFrom(source: any = {}) {
	        return new Host(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.IP = source["IP"];
	        this.MAC = source["MAC"];
	        this.Hostname = source["Hostname"];
	        this.Vendor = source["Vendor"];
	        this.Online = source["Online"];
	        this.RTT = source["RTT"];
	    }
	}
	export class Link {
	    from: string;
	    to: string;
	    type: string;
	    latency: number;
	
	    static createFrom(source: any = {}) {
	        return new Link(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = source["from"];
	        this.to = source["to"];
	        this.type = source["type"];
	        this.latency = source["latency"];
	    }
	}
	export class Network {
	    cidr: string;
	    interface: string;
	    gateway: string;
	    hosts: Host[];
	
	    static createFrom(source: any = {}) {
	        return new Network(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cidr = source["cidr"];
	        this.interface = source["interface"];
	        this.gateway = source["gateway"];
	        this.hosts = this.convertValues(source["hosts"], Host);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Node {
	    id: string;
	    label: string;
	    type: string;
	    ip: string;
	    mac: string;
	    vendor: string;
	    online: boolean;
	    rtt: number;
	
	    static createFrom(source: any = {}) {
	        return new Node(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	        this.type = source["type"];
	        this.ip = source["ip"];
	        this.mac = source["mac"];
	        this.vendor = source["vendor"];
	        this.online = source["online"];
	        this.rtt = source["rtt"];
	    }
	}
	export class Topology {
	    nodes: Node[];
	    links: Link[];
	    networks: Network[];
	
	    static createFrom(source: any = {}) {
	        return new Topology(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nodes = this.convertValues(source["nodes"], Node);
	        this.links = this.convertValues(source["links"], Link);
	        this.networks = this.convertValues(source["networks"], Network);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

