import {ImageOption} from "./ImageOption";
import {TagOption} from "./TagOption";

export interface ImageWithTag {
    image: ImageOption | null;
    tag: TagOption | null;
}