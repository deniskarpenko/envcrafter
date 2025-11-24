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

