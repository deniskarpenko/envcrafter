<template>
  <div class="tabs">
    <div class="tab-headers">
      <div
          v-for="(tab, index) in tabs"
          :key="tab.id"
          :class="{ active: activeTab === index }"
          @click="setActiveTab(index)"
      >
        {{ tab.title }}
      </div>
    </div>
    <div class="tab-content">
      <slot :name="tabs[activeTab].slot" />
    </div>
  </div>
</template>
<script setup lang="ts">
import {Tab} from "../types/Tabs/Tab";
import {computed, onUnmounted, ref, watch} from "vue";

interface Props {
  tabs: Tab[],
  defaultTab?: number
}

interface Emits {
  tabChanged: [index: number, tab: Tab]
}

const props = withDefaults(defineProps<Props>(), {
  defaultTab: 0
});

const emit = defineEmits<Emits>();

const activeTab = ref(props.defaultTab);

const setActiveTab = (index: number) => {
  if (props.tabs[index]?.disabled) return
  activeTab.value = index
  emit('tabChanged', index, props.tabs[index])
}

const currentTab = computed(() => props.tabs[activeTab.value])

watch(() => props.defaultTab, (newDefault) => {
  activeTab.value = newDefault;
}, { immediate: true });

onUnmounted(() => activeTab.value = props.defaultTab);
</script>
<style lang="css">
.tabs {
  background-color: transparent;
}
.tab-headers {
  display: flex;
  gap: 2px;
}

.active {
  box-shadow: 0 0 10px rgba(0, 123, 255, 0.3);
  background-color: #a0c5ef;
}

.tab-headers > div {
  border: 2px solid;
  padding: 10px;
  cursor: pointer;
}
.tab-content {
  background-color: transparent;
}
</style>