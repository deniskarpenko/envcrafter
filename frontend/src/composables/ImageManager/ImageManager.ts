import {ImageTypes} from "../../types/ImageTypes";

import { ref, computed, watch, type ComputedRef,type Ref} from "vue";
import {ImageOption} from "../../types/ImageOption";
import {TagOption} from "../../types/TagOption";
import {ImageWithTag} from "../../types/ImageWithTag";
import {GetTagByImageId} from "../../../wailsjs/go/main/App";

export interface ImageRowConfig {
    title: string;
    type: ImageTypes;
    images: ImageOption[];
    tags: TagOption[];
    model: Ref<ImageWithTag>;
    updateHandler: (value: ImageWithTag) => void;
}

export class ImageManager<T extends ImageTypes> {
    public readonly model: Ref<ImageWithTag>;
    public readonly images: ComputedRef<ImageOption[]>;
    public readonly tags: Ref<TagOption[]>;
    public readonly updateHandler: (value: ImageWithTag) => void;

    constructor(
        public readonly type: T,
        public readonly title: string,
        public readonly allImages: Ref<ImageOption[]>
    ) {
        this.model = ref<ImageWithTag>({ image: null, tag: null });
        this.images = computed(() => this.allImages.value.filter((image) => image.image_type === this.type));
        this.tags = ref<TagOption[]>([]);
        this.updateHandler = (value: ImageWithTag) => {
            this.model.value = value;
        };
        this.setupModelWatcher();
    }

    public getConfig(): ImageRowConfig {
        return {
            title: this.title,
            type: this.type,
            images: this.images.value,
            tags: this.tags.value,
            model: this.model,
            updateHandler: this.updateHandler,
        };
    }

    private setupModelWatcher(): void {
        watch(() => this.model.value?.image, async (newImage) => {
            if (!newImage) {
                this.tags.value = [];
                return;
            }

            try {
                this.tags.value = await GetTagByImageId(parseInt(newImage.image_id));
            } catch (error) {
                console.error(`Error fetching ${this.title.toLowerCase()} tags:`, error);
                this.tags.value = [];
            }
        });
    }

}