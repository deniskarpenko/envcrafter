<template>
  <Application @build="handleBuild"></Application>
</template>
<script lang="ts" setup>
import { Build } from "../wailsjs/go/main/App";
import Application from "./layers/Application.vue";
import {Project} from "./types/Application";
const handleBuild = (project : Project): void => {
  console.log(project);

  Build(project);
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

