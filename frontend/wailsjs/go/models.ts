export namespace models {
	
	export class Request {
	    id: string;
	    name: string;
	    method: string;
	    url: string;
	    headers: Record<string, string>;
	    body: string;
	    type: string;
	    query_vars?: Record<string, any>;
	
	    static createFrom(source: any = {}) {
	        return new Request(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.method = source["method"];
	        this.url = source["url"];
	        this.headers = source["headers"];
	        this.body = source["body"];
	        this.type = source["type"];
	        this.query_vars = source["query_vars"];
	    }
	}
	export class Response {
	    status_code: number;
	    status: string;
	    headers: Record<string, string>;
	    body: string;
	    time: number;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status_code = source["status_code"];
	        this.status = source["status"];
	        this.headers = source["headers"];
	        this.body = source["body"];
	        this.time = source["time"];
	        this.error = source["error"];
	    }
	}

}

