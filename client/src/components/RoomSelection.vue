<script setup lang="ts">
import { ref, defineProps, reactive, watch } from "vue";
import { useFilterStore } from "@/store/filterStore";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import {
  Bed,
  Users,
  Wifi,
  Coffee,
  Utensils,
  ChevronLeft,
  ChevronRight,
} from "lucide-vue-next";
import { useRooms } from "@/services/queries";
import { useRoute, useRouter } from "vue-router";

const props = defineProps<{
  hotelId: string;
  onRoomSelect: (roomId: string) => void;
}>();

const filterStore = useFilterStore();

const route = useRoute();
const router = useRouter();

const queryParams = reactive({
  location: route.query.location as string,
  checkIn: route.query.checkIn as string,
  checkOut: route.query.checkOut as string,
  adults: route.query.adults as string,
  children: route.query.children as string,
  roomCapacity: (
    parseInt(route.query.adults as string, 10) +
    parseInt(route.query.children as string, 10)
  ).toString(),
  minPrice: route.query.minPrice as string,
  maxPrice: route.query.maxPrice as string,
  roomAmenities: route.query.roomAmenities
    ? (route.query.roomAmenities as string).split(",")
    : [],
});

console.log("Initial queryParams:", queryParams);

const { data, isLoading, isFetching, refetch } = useRooms(
  props.hotelId,
  queryParams
);

const selectedRoom = ref<string>("");
const currentImageIndex = ref<{ [key: string]: number }>({});

const handleImageNavigation = (roomId: string, direction: "prev" | "next") => {
  const currentIndex = currentImageIndex.value[roomId] || 0;
  const room = data.value?.data.find((room) => room.id === roomId);
  if (!room) return;

  const newIndex =
    (currentIndex + (direction === "next" ? 1 : -1) + room.images.length) %
    room.images.length;

  currentImageIndex.value[roomId] = newIndex;
};

type Amenity = "Free Wi-Fi" | "Breakfast Included" | "Room Service" | string;

interface Room {
  amenities: Amenity[];
}

const getAmenityIcon = (amenity: Amenity) => {
  switch (amenity) {
    case "Free Wi-Fi":
      return Wifi;
    case "Breakfast Included":
      return Coffee;
    case "Room Service":
      return Utensils;
    default:
      return null;
  }
};

const filters = ref([
  {
    id: "price-range",
    label: `$${queryParams.minPrice}-$${queryParams.maxPrice}`,
    count: 0,
  },
  { id: "breakfast", label: "Breakfast Included", count: 0 },
  { id: "cancellation", label: "Free Cancellation", count: 0 },
  { id: "double-bed", label: "1 Double Bed", count: 0 },
  { id: "two-beds", label: "2 Beds", count: 0 },
  { id: "prepay", label: "Prepay Online", count: 0 },
  { id: "instant", label: "Instant Confirmation", count: 0 },
]);

const activeFilters = ref<string[]>([]);

if (queryParams.minPrice && queryParams.maxPrice) {
  activeFilters.value.push("price-range");
}
queryParams.roomAmenities.forEach((amenity) => {
  activeFilters.value.push(amenity);
});

const toggleFilter = (filterId: string) => {
  if (activeFilters.value.includes(filterId)) {
    activeFilters.value = activeFilters.value.filter((id) => id !== filterId);
  } else {
    activeFilters.value.push(filterId);
  }
  updateQueryParams();
};
const updateQueryParams = () => {
  const query = { ...route.query };

  // Update queryParams based on activeFilters
  if (activeFilters.value.includes("price-range")) {
    query.minPrice = queryParams.minPrice;
    query.maxPrice = queryParams.maxPrice;
  } else {
    filterStore.rest();
    queryParams.minPrice = filterStore.priceRange[0].toString();
    queryParams.maxPrice = filterStore.priceRange[1].toString();
    query.minPrice = filterStore.priceRange[0].toString();
    query.maxPrice = filterStore.priceRange[1].toString();
  }

  // Update amenities in queryParams
  const amenities = activeFilters.value.filter((filter) =>
    [
      "breakfast",
      "cancellation",
      "double-bed",
      "two-beds",
      "prepay",
      "instant",
    ].includes(filter)
  );
  if (amenities.length > 0) {
    query.roomAmenities = amenities.join(",");
  } else {
    delete query.amenities;
  }

  router.push({ query });
  refetch();
};

const handleUpdateRooms = (query: any) => {
  queryParams.location = query.location;
  queryParams.checkIn = query.checkIn;
  queryParams.checkOut = query.checkOut;
  queryParams.adults = query.adults;
  queryParams.children = query.children;
  (queryParams.roomCapacity = (
    parseInt(route.query.adults as string, 10) +
    parseInt(route.query.children as string, 10)
  ).toString()),
    (queryParams.minPrice = query.minPrice);
  queryParams.maxPrice = query.maxPrice;
  queryParams.roomAmenities = query.roomAmenities
    ? query.roomAmenities.split(",")
    : [];
  refetch();
};

watch(
  () => route.query,
  (newQuery) => {
    handleUpdateRooms(newQuery);
  },
  { deep: true }
);

const calculateFilterCounts = () => {
  const counts = {
    "price-range": 0,
    breakfast: 0,
    cancellation: 0,
    "double-bed": 0,
    "two-beds": 0,
    prepay: 0,
    instant: 0,
  };

  data.value?.data.forEach((room) => {
    if (
      room.price >= parseFloat(queryParams.minPrice) &&
      room.price <= parseFloat(queryParams.maxPrice)
    )
      counts["price-range"]++;
    if (room.amenities.includes("Breakfast Included")) counts["breakfast"]++;
    if (room.amenities.includes("Free Cancellation")) counts["cancellation"]++;
    if (room.bedType === "1 Double Bed") counts["double-bed"]++;
    if (room.bedType === "2 Beds") counts["two-beds"]++;
    // if (room.paymentOption === "Prepay Online") counts["prepay"]++;
    // if (room.confirmationType === "Instant Confirmation") counts["instant"]++;
  });

  filters.value.forEach((filter) => {
    filter.count = counts[filter.id as keyof typeof counts];
  });
};
watch(data, calculateFilterCounts, { immediate: true });
</script>

<template>
  <div class="space-y-6">
    <div class="space-y-4">
      <div class="flex items-center justify-between">
        <h2 class="text-2xl font-bold">Select Your Room</h2>
        <Button
          v-if="activeFilters.length > 0"
          variant="ghost"
          size="sm"
          @click="activeFilters = []"
          class="text-muted-foreground"
        >
          Clear all filters
        </Button>
      </div>
      <ScrollArea class="w-full whitespace-nowrap rounded-md border">
        <div class="flex w-max space-x-2 p-4">
          <Badge
            v-for="filter in filters"
            :key="filter.id"
            :variant="activeFilters.includes(filter.id) ? 'default' : 'outline'"
            class="cursor-pointer"
            @click="toggleFilter(filter.id)"
          >
            {{ filter.label }} ({{ filter.count }})
            <X v-if="activeFilters.includes(filter.id)" class="ml-1 h-3 w-3" />
          </Badge>
        </div>
        <ScrollBar orientation="horizontal" />
      </ScrollArea>
    </div>
    <div v-if="isLoading || isFetching">Loading...</div>
    <Card
      v-else
      v-for="room in data?.data"
      :key="room.id"
      :class="[
        'cursor-pointer transition-colors',
        selectedRoom === room.id ? 'border-primary' : '',
      ]"
    >
      <CardHeader>
        <div class="flex justify-between items-start">
          <div>
            <CardTitle>{{ room.name }}</CardTitle>
            <p class="text-sm text-muted-foreground mt-1">
              {{ room.description }}
            </p>
          </div>
          <Badge
            :variant="room.availableCount < 3 ? 'destructive' : 'secondary'"
          >
            {{ room.availableCount }} rooms left
          </Badge>
        </div>
      </CardHeader>
      <CardContent>
        <div class="grid md:grid-cols-2 gap-6">
          <div class="relative aspect-video">
            <Dialog>
              <DialogTrigger asChild>
                <div class="relative w-full h-full cursor-pointer">
                  <img
                    :src="room.images[currentImageIndex[room.id] || 0]"
                    :alt="`${room.name} view ${
                      currentImageIndex[room.id] + 1 || 1
                    }`"
                    class="object-cover rounded-lg"
                  />
                  <div
                    class="absolute inset-0 flex items-center justify-between p-2"
                  >
                    <Button
                      variant="secondary"
                      size="icon"
                      class="opacity-80 hover:opacity-100"
                      @click.stop="handleImageNavigation(room.id, 'prev')"
                    >
                      <ChevronLeft class="h-4 w-4" />
                    </Button>
                    <Button
                      variant="secondary"
                      size="icon"
                      class="opacity-80 hover:opacity-100"
                      @click.stop="handleImageNavigation(room.id, 'next')"
                    >
                      <ChevronRight class="h-4 w-4" />
                    </Button>
                  </div>
                </div>
              </DialogTrigger>
              <DialogContent class="max-w-4xl">
                <div class="grid gap-4">
                  <div class="relative aspect-video">
                    <img
                      :src="room.images[currentImageIndex[room.id] || 0]"
                      :alt="`${room.name} view ${
                        currentImageIndex[room.id] + 1 || 1
                      }`"
                      class="object-contain"
                    />
                  </div>
                  <div class="grid grid-cols-3 gap-2">
                    <div
                      v-for="(roomImg, index) in room.images"
                      :key="roomImg"
                      :class="[
                        'relative aspect-video cursor-pointer',
                        index === (currentImageIndex[room.id] || 0)
                          ? 'ring-2 ring-primary'
                          : '',
                      ]"
                      @click="currentImageIndex[room.id] = index"
                    >
                      <img
                        :src="roomImg"
                        :alt="`${room.name} view ${index + 1}`"
                        class="object-cover rounded"
                      />
                    </div>
                  </div>
                </div>
              </DialogContent>
            </Dialog>
          </div>
          <div class="space-y-4">
            <div class="flex items-center gap-2">
              <Bed class="h-5 w-5" />
              <span>Sleeps {{ room.capacity }}</span>
              <Users class="h-5 w-5 ml-4" />
              <span>Max {{ room.capacity }} guests</span>
            </div>
            <div class="flex flex-wrap gap-2">
              <Badge
                v-for="amenity in room.amenities"
                :key="amenity"
                variant="outline"
                class="flex items-center gap-1"
              >
                <component
                  v-if="getAmenityIcon(amenity)"
                  :is="getAmenityIcon(amenity)"
                  class="h-4 w-4"
                />
                {{ amenity }}
              </Badge>
            </div>
            <div class="flex justify-between items-center mt-4">
              <div>
                <span class="text-2xl font-bold">{{ room.price }}</span>
                <span class="text-muted-foreground">/night</span>
              </div>
              <Button
                @click="
                  () => {
                    selectedRoom = room.id;
                    props.onRoomSelect(room.id);
                  }
                "
                :variant="selectedRoom === room.id ? 'default' : 'outline'"
              >
                Select Room
              </Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  </div>
</template>
