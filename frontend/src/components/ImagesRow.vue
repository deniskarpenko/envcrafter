<template>
  <div class="image-row">
    <div class="image-row-header">
      <h4 @click="imageCollapseHandler"> {{ title }} {{ arrow }} </h4>
      <img src="/images/icons/gear.svg" @click="settingClickHandler(type)">
    </div>
    <div class="image-list" :class="{ hidden: !isOpen }">
      <button
          v-for="image in images"
          :class="{ selected: selectedImage?.image === image.image }"
          @click="onImageClick(image)"
          >
        <img :src="getSrcForImage(image)" :alt="image.image" />
      </button>
    </div>
    <div class="tag-list" :class="{ hidden: !isOpen }">
      <div v-for="tagRow in chunkedTags" class="tag-row">
        <span
            v-for="tag in tagRow"
            class="tag"
            :class="{selected : selectedTag?.TagName === tag.TagName }"
            @click="onTagClick(tag)"
        >{{ tag.TagName }}</span>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {ImageOption} from "../types/ImageOption";
import {TagOption} from "../types/TagOption";
import { ref, computed } from "vue";
import {ImageWithTag} from "../types/ImageWithTag";
import {ImageTypes} from "../types/ImageTypes";

interface Props {
  images: ImageOption[];
  tags: TagOption[];
  title: string;
  modelValue?: ImageWithTag
  type: ImageTypes
}
const props = defineProps<Props>();

const emit = defineEmits<{
  'update:modelValue': [value: ImageWithTag]
  'click:settings': [type: ImageTypes]
}>();

const isOpen = ref(true);

const arrow = computed(() => {
  return isOpen.value ? " ▼" : " ▲";
});

const selectedImage = ref<ImageOption | null>(null);

const selectedTag = ref<TagOption | null>(null);

const imageCollapseHandler = () => {
  isOpen.value = !isOpen.value;
};

const getSrcForImage = (image: ImageOption): string => {
  return  `/images/tech/${image.docker_image}.svg`;
};

const updateModel = () => {
  const newValue: ImageWithTag = {
    image: selectedImage.value,
    tag: selectedTag.value
  };
  emit('update:modelValue', newValue);
};

const onImageClick = (image: ImageOption): void => {
  if (!isOpen.value) {
    return;
  }

  selectedImage.value = image;
  updateModel();
}

const onTagClick = (tag: TagOption): void => {
  if (!isOpen.value) {
    return;
  }

  selectedTag.value = tag;

 updateModel();
}

const chunkedTags = computed(() => {
  const chunks = [];

  for (let i = 0; i < props.tags.length; i +=5 ) {
    chunks.push(props.tags.slice(i, i + 5));
  }

  return chunks;
});

const settingClickHandler = (type: ImageTypes): void => {
  emit('click:settings', type);
}

</script>
<style>
.image-row > h4 {
  background-color: #656262;
  margin-bottom: 0px;
  text-align: left;
  width: 80px;
  padding-left: 10px;
  padding-right: 5px;
  border-radius: 2px;
  cursor: pointer;
}
.image-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  padding: 10px;
  background-color: #656262;
}

.image-list > button {
  cursor: pointer;
  border: 2px solid transparent;
  border-radius: 4px;
  padding: 5px;
}

.image-list > button:hover {
  transform: scale(1.05);
  border-color: #ccc;
}

.image-list > button.selected {
  transform: scale(1.1);
  border-color: #007bff;
  background-color: #a0c5ef;
  box-shadow: 0 0 10px rgba(0, 123, 255, 0.3);
}

.image-list > button img {
  display: block;
  max-width: 100%;
  height: auto;
}

.hidden {
  visibility: hidden;
}

.tag-list {
  margin-top: 10px;
}

.tag-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  gap: 10px;
}

.tag {
  flex: 1;
  text-align: center;
  padding: 5px 10px;
  background-color: #706363;
  border-radius: 4px;
  border: 1px solid #ddd;
  font-size: 14px;
  cursor: pointer;
}
.tag:hover {
  transform: scale(1.05);
  border-color: #ccc;
}

.tag.selected {
  transform: scale(1.1);
  border-color: #007bff;
  background-color: #a0c5ef;
  box-shadow: 0 0 10px rgba(0, 123, 255, 0.3);
}
.image-row-header {
  display: flex;
  cursor: pointer;
}

.image-row-header > img {
  max-width: 20px;
}
</style>