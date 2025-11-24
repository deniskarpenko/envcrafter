import {ImageManager, type ImageRowConfig} from "./ImageManager";
import {ImageTypes} from "../../types/ImageTypes";
import {ImageOption} from "../../types/ImageOption";
import {computed, type Ref} from "vue";

export function useImageManagers(allImages: Ref<ImageOption[]>) {
    const managers = {
        backend: new ImageManager(ImageTypes.Backend, "Backend", allImages),
        sql: new ImageManager(ImageTypes.SQL, "SQL", allImages),
        nosql: new ImageManager(ImageTypes.NOSQL, "NOSQL", allImages),
        web: new ImageManager(ImageTypes.WEB, "Web", allImages),
    };

    const configs = computed((): ImageRowConfig[] =>
        Object.values(managers).map(manager => manager.getConfig())
    );

    return {managers, configs};
}