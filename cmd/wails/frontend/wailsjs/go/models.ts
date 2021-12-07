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

export class CpeResult {


    static createFrom(source: any = {}) {
        return new CpeResult(source);
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