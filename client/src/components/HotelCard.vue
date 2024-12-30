<script setup lang="ts">
import { Star } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { RouterLink } from "vue-router";

interface HotelProps {
  id: string;
  name: string;
  image: string;
  rating: number;
  price: number;
  location: string;
  description: string;
}

defineProps<{
  hotel: HotelProps;
}>();
const testArray = [1, 2, 3, 4, 5];
</script>
<template>
  <div
    class="border rounded-lg overflow-hidden shadow-lg transition-shadow hover:shadow-xl"
  >
    <img
      :src="hotel.image"
      :alt="hotel.name"
      width="{400}"
      height="{200}"
      class="w-full h-48 object-cover"
    />
    <div class="p-4">
      <h3 class="font-bold text-lg mb-2">{{ hotel.name }}</h3>
      <p class="text-gray-600 mb-2">{{ hotel.location }}</p>
      <div class="flex items-center mb-2">
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
      <p class="text-sm mb-4">{{ hotel.description }}</p>
      <div class="flex justify-between items-center">
        <p class="text-lg font-bold">{{ hotel.price }} / night</p>
        <RouterLink :to="'/hotel/${hotel.id}'">
          <Button>View Details</Button>
        </RouterLink>
      </div>
    </div>
  </div>
</template>
