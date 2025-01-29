<script setup lang="ts">
import BookingFlow from "@/components/BookingFlow.vue";
import HotelDetails from "@/components/HotelDetails.vue";
import RoomSelection from "@/components/RoomSelection.vue";
import SearchSection from "@/components/SearchSection.vue";
import { useHotel } from "@/services/queries";
import { onMounted, onUnmounted, ref } from "vue";
import { useRoute } from "vue-router";

// Define the type for the route parameters
interface RouteParams {
  id: string; // Example: if the route is '/items/:id'
}

const route = useRoute();

// Access the parameter and cast it to the type
const id = route.params.id as RouteParams["id"];

const searchSection = ref<HTMLElement | null>(null);
const isSticky = ref(false);

const handleScroll = () => {
  if (!searchSection.value) return;

  const rect = searchSection.value.getBoundingClientRect();
  isSticky.value = rect.top <= 0;
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});

const { data, isLoading, error } = useHotel(id);
if (error) {
  console.error("Error fetching hotel data:", error);
}
</script>

<template>
  <div class="">
    <div
      ref="searchSection"
      :class="[
        'sticky top-0 z-20 pb-4 transition-all ease-in-out duration-200 ',
        isSticky ? '' : 'container',
      ]"
    >
      <SearchSection :is-sticky="isSticky" />
    </div>
    <div
      class="container mx-auto px-4 py-8 grid grid-cols-1 md:grid-cols-3 gap-8"
    >
      <div class="md:col-span-2" v-if="data">
        <HotelDetails :hotel="data" />
        <div class="mt-8"><RoomSelection :hotel="data" /></div>
      </div>
      <div class="md:sticky md:top-24 h-fit"></div>
    </div>
  </div>
</template>
