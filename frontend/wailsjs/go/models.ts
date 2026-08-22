export namespace models {
	
	export class DiscoverySource {
	    Type: string;
	    Value: string;
	
	    static createFrom(source: any = {}) {
	        return new DiscoverySource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Type = source["Type"];
	        this.Value = source["Value"];
	    }
	}
	export class HTTPInfo {
	    port: number;
	    scheme: string;
	    server: string;
	    title: string;
	    statusCode: number;
	    contentType: string;
	    scripts: string[];
	    keywords: string[];
	    fingerprint: string[];
	
	    static createFrom(source: any = {}) {
	        return new HTTPInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.port = source["port"];
	        this.scheme = source["scheme"];
	        this.server = source["server"];
	        this.title = source["title"];
	        this.statusCode = source["statusCode"];
	        this.contentType = source["contentType"];
	        this.scripts = source["scripts"];
	        this.keywords = source["keywords"];
	        this.fingerprint = source["fingerprint"];
	    }
	}
	export class SNMPInfo {
	    Version: string;
	    Community: string;
	    SysDescr: string;
	    SysName: string;
	    SysLocation: string;
	
	    static createFrom(source: any = {}) {
	        return new SNMPInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Community = source["Community"];
	        this.SysDescr = source["SysDescr"];
	        this.SysName = source["SysName"];
	        this.SysLocation = source["SysLocation"];
	    }
	}
	export class UDPService {
	    ip: string;
	    port: number;
	    protocol: string;
	    service: string;
	    info: string;
	
	    static createFrom(source: any = {}) {
	        return new UDPService(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ip = source["ip"];
	        this.port = source["port"];
	        this.protocol = source["protocol"];
	        this.service = source["service"];
	        this.info = source["info"];
	    }
	}
	export class MDNSService {
	    name: string;
	    service: string;
	    host: string;
	    ip: string;
	    port: number;
	    txt: string[];
	
	    static createFrom(source: any = {}) {
	        return new MDNSService(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.service = source["service"];
	        this.host = source["host"];
	        this.ip = source["ip"];
	        this.port = source["port"];
	        this.txt = source["txt"];
	    }
	}
	export class Port {
	    number: number;
	    protocol: string;
	    service: string;
	    open: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Port(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.number = source["number"];
	        this.protocol = source["protocol"];
	        this.service = source["service"];
	        this.open = source["open"];
	    }
	}
	export class Host {
	    IP: string;
	    MAC: string;
	    Hostname: string;
	    Vendor: string;
	    Ports: Port[];
	    HTTP: HTTPInfo[];
	    MDNS: MDNSService[];
	    UDPServices: UDPService[];
	    SNMP: SNMPInfo[];
	    Type: string;
	    Confidence: number;
	    Sources: DiscoverySource[];
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
	        this.Ports = this.convertValues(source["Ports"], Port);
	        this.HTTP = this.convertValues(source["HTTP"], HTTPInfo);
	        this.MDNS = this.convertValues(source["MDNS"], MDNSService);
	        this.UDPServices = this.convertValues(source["UDPServices"], UDPService);
	        this.SNMP = this.convertValues(source["SNMP"], SNMPInfo);
	        this.Type = source["Type"];
	        this.Confidence = source["Confidence"];
	        this.Sources = this.convertValues(source["Sources"], DiscoverySource);
	        this.Online = source["Online"];
	        this.RTT = source["RTT"];
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
	export class Link {
	    from: string;
	    to: string;
	    type: string;
	    latency: number;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new Link(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = source["from"];
	        this.to = source["to"];
	        this.type = source["type"];
	        this.latency = source["latency"];
	        this.status = source["status"];
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
	    hostname: string;
	    vendor: string;
	    sources: DiscoverySource[];
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
	        this.hostname = source["hostname"];
	        this.vendor = source["vendor"];
	        this.sources = this.convertValues(source["sources"], DiscoverySource);
	        this.online = source["online"];
	        this.rtt = source["rtt"];
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
	export class ScanResult {
	    topology: Topology;
	    duration: number;
	
	    static createFrom(source: any = {}) {
	        return new ScanResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topology = this.convertValues(source["topology"], Topology);
	        this.duration = source["duration"];
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

