<script setup lang="ts">
import { useHotels, useInfiniteHotels } from "@/services/queries";

import HotelCard from "./HotelCard.vue";
import { useRoute, useRouter } from "vue-router";
import { useSearchStore } from "@/store/searchStore";
import { useFilterStore } from "@/store/filterStore";
import { watch, ref, reactive, onMounted } from "vue";
import HotelCardSkeleton from "./animations/HotelCardSkeleton.vue";
import { Button } from "@/components/ui/button";

const route = useRoute();
const router = useRouter();
const searchStore = useSearchStore();
const filterStore = useFilterStore();

searchStore.loadSearchParams();
filterStore.loadFilterParams();

const queryParams = reactive({
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
  limit: "1",
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
  const query = {
    location: queryParams.location,
    checkIn: queryParams.checkIn,
    checkOut: queryParams.checkOut,
    adults: queryParams.adults,
    children: queryParams.children,
    minPrice: queryParams.minPrice,
    maxPrice: queryParams.maxPrice,
    rating: queryParams.rating,
    amenities: queryParams.amenities,
  };

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

// const { isLoading, isFetching, data, isError, refetch } =
//   useHotels(queryParams);

const handleUpdateFilter = (query: any) => {
  queryParams.location = query.location;
  queryParams.checkIn = query.checkIn;
  queryParams.checkOut = query.checkOut;
  queryParams.adults = query.adults;
  queryParams.children = query.children;
  queryParams.minPrice = query.minPrice;
  queryParams.maxPrice = query.maxPrice;
  queryParams.rating = query.rating;
  queryParams.amenities = query.amenities;
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
</script>
<template>
  <div>
    <div
      v-if="isLoading || isFetching"
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"
    >
      <HotelCardSkeleton v-for="index in 12" :key="index" />
    </div>
    <div v-else-if="isError">Error loading hotels</div>
    <div v-else>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="(page, index) in data?.pages" :key="index">
          <HotelCard
            v-for="item in page.pageData"
            :key="item.id"
            :hotel="item"
          />
        </div>
      </div>
    </div>
    <div class="flex items-center justify-center mt-6">
      <Button
        :v-if="hasNextPage"
        :disabled="isFetching || !hasNextPage"
        @click="nextpage"
        >{{ isFetching ? "Loading" : "Load More Data" }}</Button
      >
    </div>
  </div>
</template>
