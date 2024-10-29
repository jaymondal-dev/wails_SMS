export namespace sms {
	
	export class ClassInfo {
	    id: number;
	    name: string;
	    year: number;
	    TableName: string;
	    // Go type: time
	    createdAt: any;
	
	    static createFrom(source: any = {}) {
	        return new ClassInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.year = source["year"];
	        this.TableName = source["TableName"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
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
	export class Student12 {
	    id: number;
	    name: string;
	    roll: number;
	    studentId: string;
	    motherName: string;
	    fatherName: string;
	    guardianName: string;
	    // Go type: time
	    createdAt: any;
	    term1atcm: number;
	    term1atco: number;
	    term1aips: number;
	    term1aimpc: number;
	    term2atcm: number;
	    term2atco: number;
	    term2aips: number;
	    term2aimpc: number;
	    term3atcm: number;
	    term3atco: number;
	    term3aips: number;
	    term3aimpc: number;
	
	    static createFrom(source: any = {}) {
	        return new Student12(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.roll = source["roll"];
	        this.studentId = source["studentId"];
	        this.motherName = source["motherName"];
	        this.fatherName = source["fatherName"];
	        this.guardianName = source["guardianName"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.term1atcm = source["term1atcm"];
	        this.term1atco = source["term1atco"];
	        this.term1aips = source["term1aips"];
	        this.term1aimpc = source["term1aimpc"];
	        this.term2atcm = source["term2atcm"];
	        this.term2atco = source["term2atco"];
	        this.term2aips = source["term2aips"];
	        this.term2aimpc = source["term2aimpc"];
	        this.term3atcm = source["term3atcm"];
	        this.term3atco = source["term3atco"];
	        this.term3aips = source["term3aips"];
	        this.term3aimpc = source["term3aimpc"];
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
	export class Student345 {
	    id: number;
	    name: string;
	    roll: number;
	    studentId: string;
	    motherName: string;
	    fatherName: string;
	    guardianName: string;
	    // Go type: time
	    createdAt: any;
	    term1fl: number;
	    term1sl: number;
	    term1m: number;
	    term1e: number;
	    term1hpeth: number;
	    term1hpepr: number;
	    term1aweth: number;
	    term1awepr: number;
	    term2fl: number;
	    term2sl: number;
	    term2m: number;
	    term2e: number;
	    term2hpeth: number;
	    term2hpepr: number;
	    term2aweth: number;
	    term2awepr: number;
	    term3fl: number;
	    term3sl: number;
	    term3m: number;
	    term3e: number;
	    term3hpeth: number;
	    term3hpepr: number;
	    term3aweth: number;
	    term3awepr: number;
	
	    static createFrom(source: any = {}) {
	        return new Student345(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.roll = source["roll"];
	        this.studentId = source["studentId"];
	        this.motherName = source["motherName"];
	        this.fatherName = source["fatherName"];
	        this.guardianName = source["guardianName"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.term1fl = source["term1fl"];
	        this.term1sl = source["term1sl"];
	        this.term1m = source["term1m"];
	        this.term1e = source["term1e"];
	        this.term1hpeth = source["term1hpeth"];
	        this.term1hpepr = source["term1hpepr"];
	        this.term1aweth = source["term1aweth"];
	        this.term1awepr = source["term1awepr"];
	        this.term2fl = source["term2fl"];
	        this.term2sl = source["term2sl"];
	        this.term2m = source["term2m"];
	        this.term2e = source["term2e"];
	        this.term2hpeth = source["term2hpeth"];
	        this.term2hpepr = source["term2hpepr"];
	        this.term2aweth = source["term2aweth"];
	        this.term2awepr = source["term2awepr"];
	        this.term3fl = source["term3fl"];
	        this.term3sl = source["term3sl"];
	        this.term3m = source["term3m"];
	        this.term3e = source["term3e"];
	        this.term3hpeth = source["term3hpeth"];
	        this.term3hpepr = source["term3hpepr"];
	        this.term3aweth = source["term3aweth"];
	        this.term3awepr = source["term3awepr"];
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

