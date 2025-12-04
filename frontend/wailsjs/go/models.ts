export namespace main {
	
	export class ContainerConfig {
	    ports: string[];
	    volumes: string[];
	    envFiles: number[][];
	    envs: string[];
	
	    static createFrom(source: any = {}) {
	        return new ContainerConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ports = source["ports"];
	        this.volumes = source["volumes"];
	        this.envFiles = source["envFiles"];
	        this.envs = source["envs"];
	    }
	}
	export class ImageWithTag {
	    name: string;
	    tag: string;
	
	    static createFrom(source: any = {}) {
	        return new ImageWithTag(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.tag = source["tag"];
	    }
	}
	export class ImageWithTagConfig {
	    image?: ImageWithTag;
	    config?: ContainerConfig;
	
	    static createFrom(source: any = {}) {
	        return new ImageWithTagConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.image = this.convertValues(source["image"], ImageWithTag);
	        this.config = this.convertValues(source["config"], ContainerConfig);
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
	export class ServiceConfig {
	    backend?: ImageWithTagConfig;
	    sql?: ImageWithTagConfig;
	    nosql?: ImageWithTagConfig;
	    web?: ImageWithTagConfig;
	
	    static createFrom(source: any = {}) {
	        return new ServiceConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.backend = this.convertValues(source["backend"], ImageWithTagConfig);
	        this.sql = this.convertValues(source["sql"], ImageWithTagConfig);
	        this.nosql = this.convertValues(source["nosql"], ImageWithTagConfig);
	        this.web = this.convertValues(source["web"], ImageWithTagConfig);
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
	export class ProjectConfig {
	    services: ServiceConfig[];
	
	    static createFrom(source: any = {}) {
	        return new ProjectConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.services = this.convertValues(source["services"], ServiceConfig);
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

export namespace models {
	
	export class Image {
	    image_id: string;
	    image: string;
	    docker_image: string;
	    image_type: string;
	    is_dockerfile: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Image(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.image_id = source["image_id"];
	        this.image = source["image"];
	        this.docker_image = source["docker_image"];
	        this.image_type = source["image_type"];
	        this.is_dockerfile = source["is_dockerfile"];
	    }
	}
	export class Tag {
	    TagName: string;
	    tag_id: number;
	    image_id: number;
	
	    static createFrom(source: any = {}) {
	        return new Tag(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.TagName = source["TagName"];
	        this.tag_id = source["tag_id"];
	        this.image_id = source["image_id"];
	    }
	}

}

