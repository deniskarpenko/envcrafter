import {ImageWithTag} from "./ImageWithTag";


export interface ContainerConfig {
    ports: string[];
    volumes: string[];
    envFiles: File[];
    envs: string[];

}

export interface ImageWithTagConfig {
    image: ImageWithTag | null;
    config: ContainerConfig;

}

export interface Project {
    backend: ImageWithTagConfig | null;
    sql: ImageWithTagConfig | null;
    nosql: ImageWithTagConfig | null;
    web: ImageWithTagConfig | null;
}

export interface ApplicationConfig {
    projects: Project[];
}
