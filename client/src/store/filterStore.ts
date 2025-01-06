import { GalleryThumbnailsIcon } from "lucide-vue-next";
import { defineStore } from "pinia";

export const useFilterStore = defineStore("filter", {
  state: () => ({
    priceRange: [0, 1000],
    starRating: 0,
    selectedAmenities: [] as string[],
  }),
  actions: {
    setPriceRange(range: [number, number]) {
      this.priceRange = range;
    },
    setStarRating(rating: number) {
      this.starRating = rating;
    },
    toggleAmenity(amenity: string) {
      if (this.selectedAmenities.includes(amenity)) {
        this.selectedAmenities = this.selectedAmenities.filter(
          (a) => a !== amenity
        );
      } else {
        this.selectedAmenities.push(amenity);
      }
    },
    loadFilterParams() {
      const params = localStorage.getItem("filterParams");
      if (params) {
        const parsedParams = JSON.parse(params);
        this.priceRange = parsedParams.priceRange;
        this.starRating = parsedParams.starRating;
        this.selectedAmenities = parsedParams.selectedAmenities;
      }
    },
    saveFilterParams() {
      const params = {
        priceRange: this.priceRange,
        starRating: this.starRating,
        selectedAmenities: this.selectedAmenities,
      };
      localStorage.setItem("filterParams", JSON.stringify(params));
    },
  },
});
