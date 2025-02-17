<script setup lang="ts">
import { useInfiniteHotels } from "@/services/queries";

import HotelCard from "./HotelCard.vue";
import { useRoute, useRouter } from "vue-router";
import { useSearchStore } from "@/store/searchStore";
import { useFilterStore } from "@/store/filterStore";
import { watch, reactive, computed } from "vue";
import HotelCardSkeleton from "./animations/HotelCardSkeleton.vue";
import { Button } from "@/components/ui/button";
import NoHotelsFound from "./404/NoHotelsFound.vue";
import ErrorDisplay from "./Error/ErrorDisplay.vue";

interface QueryParams {
  location: string;
  checkIn: string;
  checkOut: string;
  adults: string;
  children: string;
  minPrice: string;
  maxPrice: string;
  rating: string;
  amenities?: string;
  page: string;
  limit: string;
}

const props = defineProps<{
  limit?: string;
}>();

const route = useRoute();
const router = useRouter();
const searchStore = useSearchStore();
const filterStore = useFilterStore();

searchStore.loadSearchParams();
filterStore.loadFilterParams();

const queryParams = reactive<QueryParams>({
  location: searchStore.location,
  checkIn: searchStore.checkIn,
  checkOut: searchStore.checkOut,
  adults: searchStore.adults.toString(),
  children: searchStore.children.toString(),
  minPrice: filterStore.priceRange[0].toString(),
  maxPrice: filterStore.priceRange[1].toString(),
  rating: filterStore.starRating.toString(),
  amenities: filterStore.selectedAmenities.join(","),
  page: "1",
  limit: props.limit || "6",
});

// if (JSON.stringify(route.query) !== JSON.stringify(queryParams)) {
//   router.replace({
//     path: route.path,
//     query: {
//       ...queryParams,
//     },
//   });
// }

const updateQueryParams = () => {
  const query: Record<string, string> = {
    location: queryParams.location,
    checkIn: queryParams.checkIn,
    checkOut: queryParams.checkOut,
    adults: queryParams.adults,
    children: queryParams.children,
    minPrice: queryParams.minPrice,
    maxPrice: queryParams.maxPrice,
    rating: queryParams.rating,
  };

  if (
    queryParams.amenities &&
    queryParams.amenities !== "undefined" &&
    queryParams.amenities !== ""
  ) {
    query.amenities = queryParams.amenities;
  }

  if (JSON.stringify(route.query) !== JSON.stringify(query)) {
    router.replace({
      path: route.path,
      query,
    });
  }
};

updateQueryParams();

const {
  isLoading,
  isFetching,
  data,
  fetchNextPage,
  hasNextPage,
  isError,
  refetch,
} = useInfiniteHotels(queryParams);

const handleUpdateFilter = (query: any) => {
  queryParams.location = query.location;
  queryParams.checkIn = query.checkIn;
  queryParams.checkOut = query.checkOut;
  queryParams.adults = query.adults;
  queryParams.children = query.children;
  queryParams.minPrice = query.minPrice;
  queryParams.maxPrice = query.maxPrice;
  queryParams.rating = query.rating;
  queryParams.amenities = query.amenities || "";
  refetch();
};

watch(
  () => route.query,
  (newQuery) => {
    handleUpdateFilter(newQuery);
  },
  { deep: true }
);

const nextpage = () => {
  fetchNextPage();
};

const hotels = computed(() => {
  if (!data.value) return [];
  if (props.limit) {
    return data.value.pages.flatMap((page) => page.pageData).slice(0, 4);
  }
  return data.value.pages.flatMap((page) => page.pageData);
});
</script>
<template>
  <div>
    <div
      v-if="isLoading || isFetching"
      :class="[
        'grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6',
        limit ? 'md:grid-cols-4 lg:grid-cols-4' : '',
      ]"
    >
      <HotelCardSkeleton v-if="limit" v-for="index in 4" :key="index + limit" />
      <HotelCardSkeleton v-else v-for="index in 12" :key="index" />
    </div>
    <div v-else-if="isError">
      <ErrorDisplay error-type="generic" />
    </div>
    <div v-else-if="hotels.length === 0">
      <NoHotelsFound />
    </div>
    <div v-else-if="route.name == 'SearchResults'">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <HotelCard v-for="item in hotels" :key="item?.id" :hotel="item!" />
      </div>
    </div>

    <div v-else>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <HotelCard v-for="item in hotels" :key="item?.id" :hotel="item!" />
      </div>
    </div>
    <div
      v-if="route.name == 'SearchResults'"
      class="flex items-center justify-center mt-6"
    >
      <Button
        :v-if="hasNextPage"
        :disabled="isFetching || !hasNextPage"
        @click="nextpage"
        >{{ isFetching ? "Loading" : "View more" }}</Button
      >
    </div>
  </div>
</template>
