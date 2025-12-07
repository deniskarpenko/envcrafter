import { main } from "../../wailsjs/go/models";
import ProjectToGo = main.Project;
import ServiceToGo = main.Service
import { Project, Service } from "../types/Application";

export const convertFromVueToGoProject = (project: Project): ProjectToGo => {
    const projectToGo = new ProjectToGo();

    if (project.backend) {
        projectToGo.backend = convertFromVueToGoService(project.backend);
    }

    return projectToGo;
}

export const convertFromVueToGoService = (service: Service): ServiceToGo => {
    const serviceToGo = new ServiceToGo({
        
    });



    return serviceToGo;
}