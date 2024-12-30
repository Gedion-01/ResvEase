<script setup lang="ts">
import { ref } from "vue";
import { Slider } from "@/components/ui/slider";
import { Checkbox } from "@/components/ui/checkbox";
import { Label } from "@/components/ui/label";

const priceRange = ref([0, 1000]);
const selectedAmenities = ref<string[]>([]);

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
  if (selectedAmenities.value.includes(amenity)) {
    selectedAmenities.value = selectedAmenities.value.filter(
      (a) => a !== amenity
    );
  } else {
    selectedAmenities.value = [...selectedAmenities.value, amenity];
  }
};
</script>
<template>
  <div class="space-y-6 sticky top-20">
    <div>
      <h2 class="text-lg font-semibold mb-2">Price Range</h2>
      <Slider :min="0" :max="1000" :step="10" v-model="priceRange" />
      <div class="flex justify-between mt-2">
        <span>{{ priceRange[0] }}</span>
        <span>{{ priceRange[1] }}</span>
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
