/* Do not change, this code is generated from Golang structs */

export {};

export class User {


    static createFrom(source: any = {}) {
        return new User(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}

export class Backup {


    static createFrom(source: any = {}) {
        return new Backup(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class Asset {


    static createFrom(source: any = {}) {
        return new Asset(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}


export class QueryOptions {


    static createFrom(source: any = {}) {
        return new QueryOptions(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}


export class Manufacturer {


    static createFrom(source: any = {}) {
        return new Manufacturer(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}





export class Report {


    static createFrom(source: any = {}) {
        return new Report(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}

export class Cpe {
    deprecated: boolean;
    cpe23uri: string;
    lastModifiedDate: string;
    vulnerabilities: string[];

    static createFrom(source: any = {}) {
        return new Cpe(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.deprecated = source["deprecated"];
        this.cpe23uri = source["cpe23uri"];
        this.lastModifiedDate = source["lastModifiedDate"];
        this.vulnerabilities = source["vulnerabilities"];
    }
}
export class cveReferenceData {
    url: string;
    name: string;
    refsource: string;
    tags: string[];

    static createFrom(source: any = {}) {
        return new cveReferenceData(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.url = source["url"];
        this.name = source["name"];
        this.refsource = source["refsource"];
        this.tags = source["tags"];
    }
}
export class cveReference {
    reference_data: cveReferenceData[];

    static createFrom(source: any = {}) {
        return new cveReference(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.reference_data = this.convertValues(source["reference_data"], cveReferenceData);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
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
export class cveDescriptionData {
    lang: string;
    value: string;

    static createFrom(source: any = {}) {
        return new cveDescriptionData(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.lang = source["lang"];
        this.value = source["value"];
    }
}
export class cveDescription {
    description_data: cveDescriptionData[];

    static createFrom(source: any = {}) {
        return new cveDescription(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.description_data = this.convertValues(source["description_data"], cveDescriptionData);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
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
export class cveMeta {
    ID: string;
    ASSIGNER: string;

    static createFrom(source: any = {}) {
        return new cveMeta(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.ID = source["ID"];
        this.ASSIGNER = source["ASSIGNER"];
    }
}
export class Cve {
    CVE_data_meta: cveMeta;
    description: cveDescription;
    references: cveReference;

    static createFrom(source: any = {}) {
        return new Cve(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.CVE_data_meta = this.convertValues(source["CVE_data_meta"], cveMeta);
        this.description = this.convertValues(source["description"], cveDescription);
        this.references = this.convertValues(source["references"], cveReference);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
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