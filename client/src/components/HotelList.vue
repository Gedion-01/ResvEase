<script setup lang="ts">
import { useHotels } from "@/services/queries";
import HotelCard from "./HotelCard.vue";
import { useRoute, useRouter } from "vue-router";
import { useSearchStore } from "@/store/searchStore";

const route = useRoute();
const router = useRouter();
const searchStore = useSearchStore();

searchStore.loadSearchParams();

const queryParams = {
  location: searchStore.location,
  checkIn: searchStore.checkIn,
  checkOut: searchStore.checkOut,
  adults: searchStore.adults.toString(),
  children: searchStore.children.toString(),
};

if (JSON.stringify(route.query) !== JSON.stringify(queryParams)) {
  router.replace({ query: queryParams });
}

const { isLoading, data, isError } = useHotels();
</script>
<template>
  <div>
    <div v-if="isLoading">Loading...</div>
    <div v-else-if="isError">Error loading hotels</div>
    <div v-else>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <HotelCard v-for="hotel in data?.data" :key="hotel.id" :hotel="hotel" />
      </div>
    </div>
  </div>
</template>
