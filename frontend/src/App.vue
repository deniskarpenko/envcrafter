<template>
  <Application @build="handleBuild"></Application>
</template>
<script lang="ts" setup>
import Application from "./layers/Application.vue";
// import {Build} from "../wailsjs/go/main/App";
// import {main} from "../wailsjs/go/models";
// import ApplicationConfig = main.ApplicationConfig;
import {Project} from "./types/Application";
// import ProjectConfig = main.ProjectConfig;
import {toProjectConfig} from "./utils/configConverter";

const handleBuild = (project : Project): void => {
  console.log(project);
 // const app = new ApplicationConfig();
 //
 // const projectConfig = toProjectConfig(project)
 //
 // if (!app.projects) {
 //   app.projects = [];
 // }
 //
 // app.projects.push(projectConfig);
 //
 //  Build(app);
}
</script>

