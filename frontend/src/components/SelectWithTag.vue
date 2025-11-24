<template>
  <div class="select-row">
    <select class="input-class image-select" v-model="selectModel.image">
      <option disabled class="option-placeholder" value="">{{ props.label }}</option>
      <option v-for="item in images" :key="item.image_id" :value="item.image_id">
        {{ item.image }}
      </option>
    </select>
    <select class="input-class tag-select" v-model="selectModel.tag">
      <option disabled class="option-placeholder" value="">Tag</option>
      <option v-for="tag in tags" :key="tag.tag_id" :value="tag.TagName">
        {{ tag.TagName }}
      </option>
    </select>
  </div>
</template>
<script setup lang="ts">
import { ImageOption } from "../types/ImageOption";
import { TagOption } from "../types/TagOption";
import { computed } from "vue";
import { ImageWithTag } from "../types/ImageWithTag";

const props = defineProps<{
  images: ImageOption[],
  tags: TagOption[],
  label: string,
  modelValue: ImageWithTag
}>();

const emit = defineEmits(["update:modelValue"]);

const selectModel = computed({
  get: () => props.modelValue,
  set: (val: ImageWithTag) => emit("update:modelValue", val),
});

</script>

<style lang="css">
.option-placeholder {
  display: none;
}

.select-row {
  display: flex;
  gap: 10px;
  margin-top: 10px;
}

.image-select {
  flex: 7;
  min-width: 260px;
}

.tag-select {
  flex: 3;
  max-width: 147px;
}

.input-class {
  background-color: black;
  color: white;
  border-radius: 4px;
  margin-top: 10px;
  min-height:  30px;
}
</style>