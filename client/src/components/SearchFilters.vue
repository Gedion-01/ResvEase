<script setup lang="ts">
import { nextTick, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useFilterStore } from "@/store/filterStore";
import { Slider } from "@/components/ui/slider";
import { Checkbox } from "@/components/ui/checkbox";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Star } from "lucide-vue-next";

const filterStore = useFilterStore();
filterStore.loadFilterParams();

const priceRange = ref(filterStore.priceRange);
const starRating = ref(filterStore.starRating);
const selectedAmenities = ref<string[]>(filterStore.selectedAmenities);

const amenities = [
  "Wi-Fi",
  "Pool",
  "Parking",
  "Breakfast",
  "Gym",
  "Restaurant",
  "Air Conditioning",
  "Pet Friendly",
];

const handleAmenityChange = (amenity: string) => {
  filterStore.toggleAmenity(amenity);
  selectedAmenities.value = filterStore.selectedAmenities;
  filterStore.saveFilterParams();
  executeFilter();
};

const router = useRouter();
const route = useRoute();

const executeFilter = () => {
  const query: Record<string, string> = {
    ...route.query,
    minPrice: priceRange.value[0].toString(),
    maxPrice: priceRange.value[1].toString(),
    rating: starRating.value.toString(),
    // amenities: selectedAmenities.value.join(","),
  };
  if (selectedAmenities.value.length > 0) {
    query.amenities = selectedAmenities.value.join(",");
  } else {
    delete query.amenities;
  }
  // router.replace({ name: "SearchResults", query });
  nextTick(() => {
    router.replace({ name: "SearchResults", query });
  });
};

watch(priceRange, (newRange) => {
  filterStore.setPriceRange(newRange);
  filterStore.saveFilterParams();
  executeFilter();
});

watch(starRating, (newRating) => {
  filterStore.setStarRating(newRating);
  filterStore.saveFilterParams();
  executeFilter();
});
</script>
<template>
  <div class="space-y-6 sticky top-20">
    <div>
      <h2 class="text-lg font-semibold mb-2">Price Range</h2>
      <Slider :min="0" :max="1000" :step="10" v-model="priceRange" />
      <div class="flex justify-between mt-2">
        <span>${{ priceRange[0] }}</span>
        <span>${{ priceRange[1] }}</span>
      </div>
    </div>
    <div>
      <h2 class="text-lg font-semibold mb-2">Star Rating</h2>
      <div class="flex items-center space-x-2">
        <Button
          v-for="rating in 5"
          :key="rating"
          :variant="starRating >= rating ? 'default' : 'outline'"
          size="sm"
          class="p-2"
          @click="starRating = rating"
        >
          <Star
            :class="[
              '`h-4 w-4',
              starRating >= rating ? 'text-yellow-400 fill-current' : '',
            ]"
          />
        </Button>
        <Button
          v-if="starRating > 0"
          variant="ghost"
          size="sm"
          @click="starRating = 0"
          class="text-muted-foreground"
        >
          Clear
        </Button>
      </div>
    </div>
    <div>
      <h2 class="text-lg font-semibold mb-2">Amenities</h2>
      <div class="space-y-2">
        <div
          class="flex items-center"
          v-for="amenity in amenities"
          :key="amenity"
        >
          <Checkbox
            :key="amenity"
            :checked="selectedAmenities.includes(amenity)"
            @update:checked="() => handleAmenityChange(amenity)"
          />
          <Label htmlFor="{amenity}" class="ml-2">{{ amenity }}</Label>
        </div>
      </div>
    </div>
  </div>
</template>
