<script setup lang="ts">
import { ref } from "vue";
import { Button } from "@/components/ui/button";
import {
  Wifi,
  PocketKnife,
  Spade,
  Utensils,
  Dumbbell,
  Umbrella,
  MapPin,
  Star,
} from "lucide-vue-next";
import type { FunctionalComponent } from "vue";
import type { LucideProps } from "lucide-vue-next";

defineProps<{ hotel: HotelProps }>();

type AmenityIconType = FunctionalComponent<LucideProps>;

const amenityIcons: { [key: string]: AmenityIconType } = {
  "Free Wi-Fi": Wifi,
  "Swimming Pool": PocketKnife, // Using PocketKnife as a substitute for Pool
  Spa: Spade, // Using Spade as a substitute for Spa
  Restaurant: Utensils,
  "Fitness Center": Dumbbell,
  "Beach Access": Umbrella,
};

const getAmenityIcon = (amenity: string): AmenityIconType => {
  return amenityIcons[amenity] || MapPin; // Default to MapPin if no icon is found
};

interface HotelProps {
  id: string;
  name: string;
  images: string[];
  rating: number;
  price: number;
  location: string;
  description: string;
  amenities: string[];
}

const currentImage = ref(0);
const testArray = [1, 2, 3, 4, 5];
</script>
<template>
  <div>
    <h1 class="text-3xl font-bold mb-4">{{ hotel.name }}</h1>
    <div class="flex items-center mb-4">
      <MapPin class="h-5 w-5 mr-2" />
      <span>{{ hotel.location }}</span>
    </div>
    <div class="mb-6 relative">
      <img
        :src="hotel.images[currentImage]"
        :alt="hotel.name"
        width="800"
        height="400"
        class="rounded-lg"
      />
      <div class="absolute bottom-4 left-4 right-4 flex justify-center">
        <Button
          v-for="(img, index) in hotel.images"
          :key="img"
          :variant="currentImage === index ? 'default' : 'secondary'"
          size="sm"
          class="mx-1"
          @click="() => (currentImage = index)"
        >
          {{ index + 1 }}
        </Button>
      </div>
    </div>
    <div class="flex items-center mb-4">
      <Star
        v-for="(val, index) in testArray"
        :class="[
          'h-5 w-5',
          index < Math.floor(hotel.rating)
            ? 'text-yellow-400 fill-current'
            : 'text-gray-300',
        ]"
      />
      <span class="ml-2 text-sm text-gray-600">{{
        hotel.rating.toFixed(1)
      }}</span>
    </div>
    <p class="mb-6">{{ hotel.description }}</p>
    <h2 class="text-xl font-bold mb-4">Amenities</h2>
    <div class="grid grid-cols-2 gap-4 mb-6">
      <div
        v-for="amenity in hotel.amenities"
        :key="amenity"
        class="flex items-center"
      >
        <component :is="getAmenityIcon(amenity)" class="h-5 w-5" />
        <span class="ml-2">{{ amenity }}</span>
      </div>
    </div>
  </div>
</template>
