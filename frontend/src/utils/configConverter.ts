import {Project, ImageWithTagConfig, ContainerConfig} from "../types/Application";
import {main} from "../../wailsjs/go/models";
import ProjectConfig = main.ProjectConfig;
import GoImageWithTagConfig = main.ImageWithTagConfig;
import GoImageWithTag= main.ImageWithTag;
import GoContainerConfig = main.ContainerConfig;
import {ImageWithTag} from "../types/ImageWithTag";

export const toProjectConfig = (config: Project): ProjectConfig => {
    const projectConfig = new ProjectConfig();

    if (config.backend) {
        projectConfig.backend = toGoImageWithTagConfig(config.backend);
    }

    if (config.sql) {
        projectConfig.sql = toGoImageWithTagConfig(config.sql);
    }

    if (config.nosql) {
        projectConfig.nosql = toGoImageWithTagConfig(config.nosql);
    }

    if (config.web) {
        projectConfig.web = toGoImageWithTagConfig(config.web);
    }

    return projectConfig;
}

export const toGoImageWithTagConfig = (config: ImageWithTagConfig): GoImageWithTagConfig => {
    const imageWithTagConfig = new GoImageWithTagConfig();

    if (config.image) {
        imageWithTagConfig.image = toGoImageWithTag(config.image);
    }


    imageWithTagConfig.config = toGoContainerConfig(config.config);

    return imageWithTagConfig;
}

export const toGoImageWithTag = (imageWithTag: ImageWithTag): GoImageWithTag => {
    return new GoImageWithTag({
        name: imageWithTag.image?.image,
        tag: imageWithTag.tag?.TagName,
    });
}

export const toGoContainerConfig = (config: ContainerConfig): GoContainerConfig => {
    return new GoContainerConfig({
        ports: config.ports,
        volumes: config.volumes,
        envFiles: config.envFiles,
        envs: config.envs,
    });
}


