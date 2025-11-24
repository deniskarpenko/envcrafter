<template>
  <div class="wrapper">
    <Tabs :tabs="tabsConfig">
      <template #images>
        <div  v-for="config in imageRowConfig">
          <images-row
              :key="config.title"
              :title="config.title"
              :images="config.images"
              :tags="config.tags"
              :type="config.type"
              @update:model-value="handleImageRowUpdate(config, $event)"
              @click:settings="showSettings"
          ></images-row>
        </div>
        <settings-dialog
            title="Settings"
            :model-value="isShowSettingsDialog"
            :container-config="selectedContainer"
            @close="handleClose"
        ></settings-dialog>
      </template>
    </Tabs>
    <div class="buttons-containers" @click="handleBuild">
      <button>Build</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import {reactive, ref, onMounted, nextTick, watch} from "vue";
import {GetAllImages} from "../../wailsjs/go/main/App";
import {ContainerConfig, Project} from "../types/Application";
import {ImageOption} from "../types/ImageOption";
import Tabs from "../components/Tabs.vue";
import ImagesRow from "../components/ImagesRow.vue";
import {useImageManagers} from "../composables/ImageManager";
import SettingsDialog from "../components/dialog/SettingsDialog.vue";
import {ImageWithTag} from "../types/ImageWithTag";
import {ImageTypes} from "../types/ImageTypes";

const emit = defineEmits<{
  (event: 'build', app: Project): void;
}>();


const tabsConfig = [
  { id: 'images', title: 'Images', slot: 'images' },
  { id: 'result', title: 'Result', slot: 'result' },
];

const appModel = reactive<Project>({
  backend: null,
  sql: null,
  nosql: null,
  web: null,
});

const imagesOptions = ref<ImageOption[]>([]);
const isShowSettingsDialog = ref(false);
const selectedRow = ref<ImageTypes | null>(null);
const selectedContainer = ref<ContainerConfig | null>(null);


const { configs: imageRowConfig } = useImageManagers(imagesOptions);

const handleImageRowUpdate = (config: any, value: ImageWithTag) => {
  if (config.updateHandler) {
    config.updateHandler(value);
  }

  if (!config.type || !(config.type in appModel)) {
    return;
  }

  const propertyName = config.type as keyof Project;

  if (appModel[propertyName] === null) {
    appModel[propertyName] = {
      image: value,
      config: {
        ports: [],
        volumes: [],
        envFiles: [],
        envs: []
      }
    };
  } else {
    appModel[propertyName]!.image = value;
  }
};

const showSettings = async (type: ImageTypes) => {
  await nextTick();
  selectedRow.value = type;
  isShowSettingsDialog.value = true;

  const propertyName = selectedRow.value as keyof Project;

  if (appModel[propertyName] === null || appModel[propertyName] === undefined) {
    selectedContainer.value = null;
    return;
  }

  selectedContainer.value = appModel[propertyName]!.config;
};

const handleClose = async (config: ContainerConfig) => {
  await nextTick();

  if (selectedRow.value === null || !(selectedRow.value in appModel)) {
    isShowSettingsDialog.value = false;
    return;
  }

  const propertyName = selectedRow.value as keyof Project;

  if (appModel[propertyName] === null) {
    appModel[propertyName] = {
      image: null,
      config: config
    };
  } else {
    appModel[propertyName]!.config = config;
  }

  isShowSettingsDialog.value = false;

  selectedContainer.value = config;
};

const handleBuild = () => {
  emit("build", appModel);
}

onMounted(async () => {
  imagesOptions.value = await GetAllImages();
});
</script>

<style lang="css" scoped>
.wrapper {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 20px;
  gap: 20px;
}

.buttons-containers {
  display: flex;
  justify-content: center;
}
</style>