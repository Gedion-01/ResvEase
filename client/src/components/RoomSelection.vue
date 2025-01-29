<script setup lang="ts">
import { ref, defineProps, reactive, watch } from "vue";
import { useFilterStore } from "@/store/filterStore";
import { useRoomBookingStore } from "@/store/bookingStore";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";
import {
  Bed,
  Users,
  Wifi,
  Coffee,
  Utensils,
  ParkingCircle,
  Dog,
  Martini,
  Thermometer,
  Baby,
  Plane,
  WashingMachine,
  Anchor,
  Music,
  PocketKnife,
  Spade,
  Dumbbell,
  Umbrella,
  Bath,
  Mountain,
  Waves,
  Building,
  ShowerHead,
} from "lucide-vue-next";

import { useRooms } from "@/services/queries";
import { useRoute, useRouter } from "vue-router";
import type { Hotel, Room } from "@/types/hotel";

const props = defineProps<{
  hotel: Hotel;
}>();

const filterStore = useFilterStore();
const bookingStore = useRoomBookingStore();

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

const { data, isLoading, isFetching, refetch } = useRooms(
  props.hotel.id,
  queryParams
);

const selectedRoom = ref<string>("");
const currentImageIndex = ref<{ [key: string]: number }>({});

const animationDirection = ref<"slide-left" | "slide-right">("slide-left");

const handleImageNavigation = (roomId: string, direction: "prev" | "next") => {
  animationDirection.value =
    direction === "next" ? "slide-left" : "slide-right";
  const currentIndex = currentImageIndex.value[roomId] || 0;
  const room = data.value?.data.find((room) => room.id === roomId);
  if (!room) return;

  const newIndex =
    (currentIndex + (direction === "next" ? 1 : -1) + room.images.length) %
    room.images.length;

  currentImageIndex.value[roomId] = newIndex;
};

const getAmenityIcon = (amenity: string) => {
  switch (amenity) {
    case "Free Wi-Fi":
      return Wifi;
    case "Swimming Pool":
      return PocketKnife; // Using PocketKnife as a substitute for Pool
    case "Spa":
      return Spade; // Using Spade as a substitute for Spa
    case "Gym":
      return Dumbbell;
    case "Free Parking":
      return ParkingCircle;
    case "Breakfast Included":
      return Coffee;
    case "Pet Friendly":
      return Dog;
    case "Rooftop Bar":
      return Martini;
    case "Outdoor swimming pool":
      return PocketKnife; // Using PocketKnife as a substitute for Pool
    case "Sauna":
      return Thermometer;
    case "24/7 Room Service":
      return Utensils;
    case "Kids' Play Area":
      return Baby;
    case "Airport Shuttle":
      return Plane;
    case "Laundry Service":
      return WashingMachine;
    case "Beach Access":
      return Umbrella;
    case "Water Sports":
      return Anchor;
    case "Live Entertainment":
      return Music;
    case "Air conditioning":
      return Thermometer;
    case "Private bathroom":
      return Bath;
    case "Mountain view":
      return Mountain;
    case "Ocean view":
      return Waves;
    case "City view":
      return Building;
    case "Shower":
      return ShowerHead;
    default:
      return null;
  }
};

const initialFilters = reactive({
  priceRange: {
    minPrice: queryParams.minPrice,
    maxPrice: queryParams.maxPrice,
  },
  filters: [
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
  ],
});

const filters = ref([...initialFilters.filters]);

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

  if (activeFilters.value.includes("price-range")) {
    // Use the initial filter price range
    queryParams.minPrice = initialFilters.priceRange.minPrice.toString();
    queryParams.maxPrice = initialFilters.priceRange.maxPrice.toString();
    query.minPrice = initialFilters.priceRange.minPrice.toString();
    query.maxPrice = initialFilters.priceRange.maxPrice.toString();
  } else {
    // Reset to the global store's price range
    filterStore.rest();
    const [storeMin, storeMax] = filterStore.priceRange;
    queryParams.minPrice = storeMin.toString();
    queryParams.maxPrice = storeMax.toString();
    query.minPrice = storeMin.toString();
    query.maxPrice = storeMax.toString();
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

const reserveRoom = (room: Room) => {
  bookingStore.setRoomBookingDetails(room);
  bookingStore.setHotelBookingDetails(props.hotel);
  router.push("/booking");
};
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
      <ScrollArea class="w-full whitespace-nowrap rounded-md">
        <div class="flex w-max space-x-2 py-4">
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
    <div v-if="isFetching">Loading...</div>
    <div
      v-else-if="!data?.data || data.data.length === 0"
      class="text-center text-gray-500"
    >
      No rooms available.
    </div>
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
            <CardTitle class="whitespace-wrap">{{ room.name }}</CardTitle>
            <p class="text-sm text-muted-foreground mt-1">
              {{ room.description }}
            </p>
          </div>
          <Badge
            :variant="room.availableCount < 3 ? 'destructive' : 'secondary'"
            class="whitespace-nowrap"
          >
            {{ room.availableCount }} rooms left
          </Badge>
        </div>
      </CardHeader>
      <CardContent>
        <div class="grid md:grid-cols-2 gap-6">
          <div class="relative">
            <Dialog>
              <Carousel class="relative cursor-pointer">
                <CarouselContent>
                  <CarouselItem
                    class="relative aspect-video"
                    v-for="(image, index) in room.images"
                    :key="index"
                  >
                    <DialogTrigger asChild>
                      <img
                        :src="image"
                        class="object-cover w-full h-full rounded-lg"
                        :alt="`${room.name} view ${index + 1}`"
                      />
                    </DialogTrigger>
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
              <DialogContent class="max-w-4xl">
                <div class="grid gap-4">
                  <div class="relative w-full h-64 md:h-auto">
                    <img
                      :src="room.images[currentImageIndex[room.id] || 0]"
                      :alt="`${room.name} view ${
                        currentImageIndex[room.id] + 1 || 1
                      }`"
                      class="object-contain w-full h-full rounded-lg"
                    />
                  </div>
                  <div class="grid grid-cols-3 gap-2">
                    <div
                      v-for="(roomImg, index) in room.images"
                      :key="roomImg"
                      :class="[
                        'relative cursor-pointer',
                        index === (currentImageIndex[room.id] || 0)
                          ? 'ring-2 ring-primary rounded-lg'
                          : '',
                      ]"
                      @click="currentImageIndex[room.id] = index"
                    >
                      <img
                        :src="roomImg"
                        :alt="`${room.name} view ${index + 1}`"
                        class="object-cover w-full h-full rounded-lg"
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
                  :is="getAmenityIcon(amenity)"
                  class="h-4 w-4 text-blue-500"
                />
                {{ amenity }}
              </Badge>
            </div>
            <div class="flex justify-between items-center mt-4">
              <div>
                <span class="text-2xl font-bold">{{ room.price }}</span>
                <span class="text-muted-foreground">/night</span>
              </div>
              <Button @click="reserveRoom(room)" variant="default">
                Reserve</Button
              >
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  </div>
</template>
