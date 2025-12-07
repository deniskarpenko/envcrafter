import { main } from "../../wailsjs/go/models";
import ProjectToGo = main.Project;
import ServiceToGo = main.Service;
import ImageWithTagToGo = main.ImageWithTag;
import ContainerConfigToGo = main.ContainerConfig;
import { ContainerConfig, Project, Service } from "../types/Application";
import { ImageWithTag } from "../types/ImageWithTag";

export const convertFromVueToGoProject = (project: Project): ProjectToGo => {
    const projectToGo = new ProjectToGo();

    if (project.backend) {
        projectToGo.backend = convertFromVueToGoService(project.backend);
    }

    return projectToGo;
}

export const convertFromVueToGoService = (service: Service): ServiceToGo => {
    return new ServiceToGo({
        image: service.image ? convertFromVueToGoImageWithTag(service.image): null,
        config: convertFromVueToGoContainerConfig(service.config)
    });;
}

export const convertFromVueToGoImageWithTag = (image: ImageWithTag): ImageWithTagToGo => {
    return new ImageWithTagToGo({
        name: image.image?.image,
        tag: image.tag?.TagName
    });;
}

export const convertFromVueToGoContainerConfig = (config: ContainerConfig): ContainerConfigToGo => {
    return new ContainerConfigToGo({
        ports: config.ports,
        volumes: config.volumes,
        envs: config.envs,
        envFiles: config.envFiles
    });
}