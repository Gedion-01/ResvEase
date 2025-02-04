<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { ref, computed, reactive, watch } from "vue";
import { useSearchStore } from "@/store/searchStore";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { List, Map } from "lucide-vue-next";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Search } from "lucide-vue-next";
import { addDays, format, isBefore, startOfDay } from "date-fns";
import { getLocalTimeZone } from "@internationalized/date";
import { useQueryClient } from "@tanstack/vue-query";
import { useFilterStore } from "@/store/filterStore";
import { ElDatePicker } from "element-plus";

const searchStore = useSearchStore();
searchStore.loadSearchParams();

const filterStore = useFilterStore();
filterStore.loadFilterParams();

const props = defineProps({
  isSticky: {
    default: false,
    type: Boolean,
    Required: true,
  },
});

const router = useRouter();
const route = useRoute();
const queryClient = useQueryClient();

const location = ref(searchStore.location);
const guests = reactive({
  adults: searchStore.adults,
  children: searchStore.children,
});

const viewMode = ref<"list" | "map">("list");

console.log("searchStore", searchStore.checkIn, searchStore.checkOut);

const timeZone = getLocalTimeZone();

const value = ref<[Date, Date]>([
  new Date(searchStore.checkIn),
  new Date(searchStore.checkOut),
]);
const shortcuts = [
  {
    text: "Next week",
    value: () => {
      const start = new Date();
      const end = new Date();
      end.setTime(end.getTime() + 7 * 24 * 3600 * 1000);
      return [start, end];
    },
  },
  {
    text: "Next month",
    value: () => {
      const start = new Date();
      const end = new Date();
      end.setTime(end.getTime() + 30 * 24 * 3600 * 1000);
      return [start, end];
    },
  },
  {
    text: "Next 3 months",
    value: () => {
      const start = new Date();
      const end = new Date();
      end.setTime(end.getTime() + 90 * 24 * 3600 * 1000);
      return [start, end];
    },
  },
];
const handleDateChange = (val: [Date, Date]) => {
  value.value = val;
};
console.log("value", value.value);
const formattedStart = computed(() => {
  if (value.value && value.value[0]) {
    return format(value.value[0], "iii, MMM dd");
  }
  return "";
});

const formattedEnd = computed(() => {
  if (value.value && value.value[1]) {
    return format(value.value[1], "iii, MMM dd");
  }
  return "";
});

const calculateNights = computed(() => {
  if (value.value && value.value[0] && value.value[1]) {
    const startDate = value.value[0];
    const endDate = value.value[1];
    return Math.ceil(
      (endDate.valueOf() - startDate.valueOf()) / (1000 * 60 * 60 * 24)
    );
  }
  return 0;
});

const decrementAdults = () => {
  guests.adults = Math.max(1, guests.adults - 1);
};
const incrementAdults = () => {
  guests.adults++;
};

const decrementChildren = () => {
  guests.children = Math.max(0, guests.children - 1);
};
const incrementChildren = () => {
  guests.children++;
};

const quickFilters = [
  "Free cancellation",
  "Breakfast included",
  "Pool",
  "Beach access",
  "Pet friendly",
  "5-star hotels",
];

const handleSearch = async () => {
  const query = {
    location: location.value,
    checkIn:
      value?.value && value.value[0]
        ? format(value.value[0], "yyyy-MM-dd")
        : format(new Date(), "yyyy-MM-dd"),
    checkOut:
      value?.value && value.value[1]
        ? format(value.value[1], "yyyy-MM-dd")
        : format(addDays(new Date(), 1), "yyyy-MM-dd"),
    adults: guests.adults.toString(),
    children: guests.children.toString(),
    roomCapacity: (guests.adults + guests.children).toString(),
    minPrice: filterStore.priceRange[0]?.toString() || "0",
    maxPrice: filterStore.priceRange[1]?.toString() || "1000",
    rating: filterStore.starRating?.toString() || "0",
    amenities: filterStore.selectedAmenities.join(",") || "",
  };
  searchStore.setSearchParams({
    location: location.value,
    checkIn: query.checkIn,
    checkOut: query.checkOut,
    adults: guests.adults,
    children: guests.children,
  });

  if (route.name === "Hotel") {
    router.replace({ query });

    const id = route.params.id as string;
  } else {
    router.replace({ name: "SearchResults", query });
  }
};

watch(location, (newLocation) => {
  searchStore.location = newLocation;
  searchStore.setSearchParams({
    location: newLocation,
    checkIn: searchStore.checkIn,
    checkOut: searchStore.checkOut,
    adults: searchStore.adults,
    children: searchStore.children,
  });
});

watch(
  value,
  (newValue) => {
    searchStore.checkIn =
      newValue && newValue[0] ? format(newValue[0], "yyyy-MM-dd") : "";
    searchStore.checkOut =
      newValue && newValue[1] ? format(newValue[1], "yyyy-MM-dd") : "";
    searchStore.setSearchParams({
      location: searchStore.location,
      checkIn: searchStore.checkIn,
      checkOut: searchStore.checkOut,
      adults: searchStore.adults,
      children: searchStore.children,
    });
  },
  { deep: true }
);

watch(
  guests,
  (newGuests) => {
    searchStore.adults = newGuests.adults;
    searchStore.children = newGuests.children;
    searchStore.setSearchParams({
      location: searchStore.location,
      checkIn: searchStore.checkIn,
      checkOut: searchStore.checkOut,
      adults: searchStore.adults,
      children: searchStore.children,
    });
  },
  { deep: true }
);

const isDateUnavailable = (date: Date) => {
  return isBefore(date, startOfDay(new Date()));
};

const size = ref<"default" | "large" | "small">("default");
</script>

<template>
  <div
    :class="[
      'w-full mx-auto bg-background rounded-lg shadow-lg p-4',
      'bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60',
    ]"
  >
    <div class="grid grid-cols-1 md:grid-cols-11 gap-4 container">
      <div class="md:col-span-4 relative">
        <div className="w-full relative">
          <div
            className="absolute left-4 top-2 font-medium text-xs text-muted-foreground pointer-events-none"
          >
            Destination
          </div>
          <Input
            v-model="location"
            type="text"
            placeholder="City, airport, region, landmark or property name"
            className="w-full px-4 pt-7 pb-1 h-auto border rounded-md font-medium placeholder:text-sm"
          />
        </div>
        <button
          @click="location = ''"
          class="absolute right-2 top-1/2 -translate-y-1/2"
        >
          <X class="h-4 w-4 text-gray-500" />
        </button>
      </div>

      <div class="md:col-span-3">
        <div class="flex h-full">
          <el-date-picker
            v-model="value"
            type="daterange"
            range-separator="To"
            start-placeholder="Start date"
            end-placeholder="End date"
            :disabled-date="isDateUnavailable"
            class="rounded-md"
            @change="handleDateChange"
          />
        </div>
      </div>

      <div class="md:col-span-2">
        <Popover>
          <PopoverTrigger asChild>
            <Button
              variant="outline"
              class="flex items-start justify-between w-full h-full border rounded-md p-1 bg-white cursor-pointer hover:border-gray-300 transition-all"
            >
              <div
                class="flex flex-col justify-between px-3 py-1 relative group text-left h-full"
              >
                <div class="text-xs text-gray-500 mb-0.5">Guests</div>
                <div
                  :class="[
                    'font-medium',
                    value?.[1] ? 'text-gray-900' : 'text-gray-400',
                  ]"
                >
                  {{ guests.adults + guests.children }} Guest{{
                    guests.adults + guests.children !== 1 ? "s" : ""
                  }}
                </div>
              </div>
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-80">
            <div class="space-y-4">
              <div class="flex justify-between items-center">
                <div>
                  <p class="font-medium">Adults</p>
                  <p class="text-sm text-muted-foreground">Ages 18 or above</p>
                </div>
                <div class="flex items-center space-x-2">
                  <Button
                    variant="outline"
                    size="icon"
                    @click="decrementAdults"
                  >
                    -
                  </Button>
                  <span class="w-8 text-center">{{ guests.adults }}</span>
                  <Button
                    variant="outline"
                    size="icon"
                    @click="incrementAdults"
                  >
                    +
                  </Button>
                </div>
              </div>
              <div class="flex justify-between items-center">
                <div>
                  <p class="font-medium">Children</p>
                  <p class="text-sm text-muted-foreground">Ages 0-17</p>
                </div>
                <div class="flex items-center space-x-2">
                  <Button
                    variant="outline"
                    size="icon"
                    @click="decrementChildren"
                  >
                    -
                  </Button>
                  <span class="w-8 text-center">{{ guests.children }}</span>
                  <Button
                    variant="outline"
                    size="icon"
                    @click="incrementChildren"
                  >
                    +
                  </Button>
                </div>
              </div>
            </div>
          </PopoverContent>
        </Popover>
      </div>

      <div class="md:col-span-2">
        <Button
          class="w-full h-full bg-primary hover:bg-primary/90"
          @click="handleSearch"
        >
          <Search class="" />
          {{ route.name === "Hotel" ? "Update" : "Search" }}
        </Button>
      </div>
    </div>
    <div
      v-if="route.fullPath == '/' || route.fullPath.includes('/search')"
      class="mt-4 flex items-center justify-between container"
    >
      <div class="flex items-center gap-2">
        <Button
          :variant="viewMode === 'list' ? 'default' : 'outline'"
          size="sm"
          @click="() => (viewMode = 'list')"
        >
          <List class="h-4 w-4 mr-1" />
          List
        </Button>
        <Button
          :variant="viewMode === 'map' ? 'default' : 'outline'"
          size="sm"
          @click="() => (viewMode = 'map')"
        >
          <Map class="h-4 w-4 mr-1" />
          Map
        </Button>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
