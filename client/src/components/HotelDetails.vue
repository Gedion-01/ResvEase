<script setup lang="ts">
import { ref } from "vue";
import { Button } from "@/components/ui/button";
import {
  Carousel,
  type CarouselApi,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";

import {
  Wifi,
  PocketKnife,
  Spade,
  Utensils,
  Dumbbell,
  Umbrella,
  MapPin,
  Star,
  ParkingCircle,
  Coffee,
  Dog,
  Martini,
  Thermometer,
  Baby,
  Plane,
  WashingMachine,
  Anchor,
  Music,
} from "lucide-vue-next";
import type { FunctionalComponent } from "vue";
import type { LucideProps } from "lucide-vue-next";

defineProps<{ hotel: HotelProps }>();

const selectedIndex = ref(0);

const emblaMainApi = ref<CarouselApi>();

function onSelect() {
  if (!emblaMainApi.value) return;
  selectedIndex.value = emblaMainApi.value.selectedScrollSnap();
}

function onThumbClick(index: number) {
  if (!emblaMainApi.value) return;
  emblaMainApi.value.scrollTo(index);
  selectedIndex.value = index;
}

function handleCarouselInit(val: CarouselApi) {
  emblaMainApi.value = val;
  onSelect();
  emblaMainApi.value?.on("select", onSelect);
  emblaMainApi.value?.on("reInit", onSelect);
}

type AmenityIconType = FunctionalComponent<LucideProps>;

const amenityIcons: { [key: string]: AmenityIconType } = {
  "Free Wi-Fi": Wifi,
  "Swimming Pool": PocketKnife,
  Spa: Spade,
  Restaurant: Utensils,
  "Fitness Center": Dumbbell,
  "Beach Access": Umbrella,
  "Free Parking": ParkingCircle,
  "Breakfast Included": Coffee,
  "Pet Friendly": Dog,
  "Rooftop Bar": Martini,
  "Outdoor swimming pool": PocketKnife,
  Sauna: Thermometer,
  "24/7 Room Service": Utensils,
  "Kids' Play Area": Baby,
  "Airport Shuttle": Plane,
  "Laundry Service": WashingMachine,
  "Water Sports": Anchor,
  "Live Entertainment": Music,
};

const getAmenityIcon = (amenity: string): AmenityIconType => {
  return amenityIcons[amenity] || MapPin;
};

interface HotelProps {
  id: string;
  name: string;
  images: string[];
  rating: number;
  location: string;
  description: string;
  amenities: string[];
}

const testArray = [1, 2, 3, 4, 5];
</script>
<template>
  <div>
    <h1 class="text-3xl font-bold mb-4">{{ hotel.name }}</h1>
    <div class="flex items-center mb-4">
      <MapPin class="h-5 w-5 mr-2" />
      <span>{{ hotel.location }}</span>
    </div>
    <Carousel
      class="relative w-full max-w-3xl mb-4"
      @init-api="handleCarouselInit"
    >
      <CarouselContent>
        <CarouselItem v-for="(img, index) in hotel.images" :key="img">
          <img
            :src="img"
            class="rounded-lg h-[400px] w-[800px] object-cover"
            :alt="hotel.name"
          />
        </CarouselItem>
      </CarouselContent>
      <CarouselPrevious
        variant="secondary"
        size="icon"
        class="absolute left-2 top-1/2 -translate-y-1/2 h-8 w-8 rounded-lg"
        @click.stop
      />
      <CarouselNext
        variant="secondary"
        size="icon"
        class="absolute right-2 top-1/2 -translate-y-1/2 h-8 w-8 rounded-lg"
        @click.stop
      />
    </Carousel>

    <div class="flex justify-center mt-4">
      <Button
        v-for="(img, index) in hotel.images"
        :key="img"
        :variant="selectedIndex === index ? 'default' : 'secondary'"
        size="sm"
        class="mx-1"
        @click="onThumbClick(index)"
      >
        {{ index + 1 }}
      </Button>
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
