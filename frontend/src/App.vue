<template>
  <Application @build="handleBuild"></Application>
</template>
<script lang="ts" setup>
import { Build } from "../wailsjs/go/main/App";
import { main } from "../wailsjs/go/models";
import Application from "./layers/Application.vue";
import {Project} from "./types/Application";
import ProjectForGo = main.Project;
import { convertFromVueToGoProject } from "./utils/converter";
const handleBuild = (project : Project): void => {
  
  const projectToGo = convertFromVueToGoProject(project);

  Build(projectToGo)
}
</script>

