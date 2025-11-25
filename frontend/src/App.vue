<template>
  <Application @build="handleBuild"></Application>
</template>
<script lang="ts" setup>
import Application from "./layers/Application.vue";
import {Project} from "./types/Application";
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

