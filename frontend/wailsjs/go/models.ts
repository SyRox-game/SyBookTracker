export namespace main {
	
	export class Tracker {
	    id: number;
	    date: string;
	    status: string;
	    wordstatus: number;
	    writtenwords: number;
	
	    static createFrom(source: any = {}) {
	        return new Tracker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.date = source["date"];
	        this.status = source["status"];
	        this.wordstatus = source["wordstatus"];
	        this.writtenwords = source["writtenwords"];
	    }
	}

}

