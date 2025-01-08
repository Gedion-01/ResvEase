<script setup lang="ts">
import { MapPin, Star, Bed } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { RouterLink } from "vue-router";
import { Badge } from "@/components/ui/badge";
import type { Hotel } from "@/types/hotel";
import { useSearchStore } from "@/store/searchStore";
import { useFilterStore } from "@/store/filterStore";

defineProps<{
  hotel: Hotel;
}>();

const searchStore = useSearchStore();
const filterStore = useFilterStore();

searchStore.loadSearchParams();
filterStore.loadFilterParams();

const queryParams = {
  location: searchStore.location,
  checkIn: searchStore.checkIn,
  checkOut: searchStore.checkOut,
  adults: searchStore.adults.toString(),
  children: searchStore.children.toString(),
  minPrice: filterStore.priceRange[0].toString(),
  maxPrice: filterStore.priceRange[1].toString(),
};

const queryString = new URLSearchParams(queryParams).toString();
console.log(queryString);

const testArray = [1, 2, 3, 4, 5];
</script>
<template>
  <div
    class="bg-card text-card-foreground rounded-lg overflow-hidden shadow-lg transition-all duration-300 hover:shadow-xl hover:scale-105"
  >
    <div class="relative h-48">
      <img
        :src="hotel.images[0]"
        :alt="hotel.name"
        width="400"
        height="200"
        class="w-full h-48 object-cover"
      />
      <div
        class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"
      ></div>
      <div class="absolute bottom-0 left-0 right-0 p-4">
        <h3 class="font-bold text-lg text-white mb-1 line-clamp-1">
          {{ hotel.name }}
        </h3>
        <div class="flex items-center text-white">
          <MapPin class="h-4 w-4 mr-1" />
          <span class="text-sm line-clamp-1">{{ hotel.location }}</span>
        </div>
      </div>
    </div>
    <div class="p-4">
      <div class="flex items-center justify-between mb-2">
        <div class="flex items-center">
          <Star
            v-for="(val, index) in testArray"
            :key="index"
            :class="[
              'h-5 w-5',
              index < Math.floor(hotel.rating)
                ? 'text-yellow-400 fill-current'
                : 'text-gray-300',
            ]"
          />
          <span class="ml-2 text-sm text-muted-foreground">{{
            hotel.rating.toFixed(1)
          }}</span>
        </div>
        <span class="text-lg font-semibold text-primary"
          >From ${{ Math.min(...hotel.prices) }}</span
        >
      </div>
      <!--
      <div class="flex items-center gap-2 mb-4">
        <Badge variant="secondary" class="flex items-center">
          <Bed class="h-3 w-3 mr-1" />
          {{ hotel.bedrooms }}
          {{ hotel.bedrooms === 1 ? "Bedroom" : "Bedrooms" }}
        </Badge>
        <Badge variant="outline" class="flex items-center">
          {{ hotel.bedTypes.join(", ") }}
        </Badge>
      </div>
      -->
      <RouterLink :to="`/hotel/${hotel.id}/rooms?${queryString}`">
        <Button class="w-full">View Details</Button>
      </RouterLink>
    </div>
  </div>
</template>
