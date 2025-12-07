import {ImageWithTag} from "./ImageWithTag";


export interface ContainerConfig {
    ports: string[];
    volumes: string[];
    envFiles: File[];
    envs: string[];

}

export interface Service {
    image: ImageWithTag | null;
    config: ContainerConfig;
}

export interface Project {
    backend: Service | null;
    sql: Service | null;
    nosql: Service | null;
    web: Service | null;
}

export interface Application {
    projects: Project[];
}
