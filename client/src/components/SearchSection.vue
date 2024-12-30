<script setup lang="ts">
import type { DateRange } from "radix-vue";
import { RangeCalendar } from "@/components/ui/range-calendar";
import { getLocalTimeZone } from "@internationalized/date";
import { type Ref, ref, computed, reactive } from "vue";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { List, Map } from "lucide-vue-next";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Search } from "lucide-vue-next";
import { format } from "date-fns";

defineProps({
  isSticky: {
    default: false,
    type: Boolean,
    Required: true,
  },
});

const location = ref("");
const guests = reactive({
  adults: 2,
  children: 0,
});

const viewMode = ref<"list" | "map">("list");

const timeZone = getLocalTimeZone();
const start = undefined;
const end = undefined;

const value = ref({
  start,
  end,
}) as Ref<DateRange, DateRange> | undefined;

const formattedStart = computed(() => {
  if (value?.value.start) {
    return format(value?.value.start.toDate(timeZone), "iii, MMM dd");
  }
  return "";
});
const formattedEnd = computed(() => {
  if (value?.value.end) {
    return format(value?.value.end.toDate(timeZone), "iii, MMM dd");
  }
  return "";
});

const calculateNights = computed(() => {
  if (value?.value.start && value.value.end) {
    const startDate = value.value.start.toDate(timeZone);
    const endDate = value.value.end.toDate(timeZone);
    return Math.ceil(
      (endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60 * 24)
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
  guests.children = Math.max(1, guests.children - 1);
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
        <Popover>
          <PopoverTrigger asChild>
            <Button
              variant="outline"
              class="flex items-start justify-between w-full h-full border rounded-md p-1 bg-white cursor-pointer hover:border-gray-300 transition-all"
            >
              <!-- Check-in Section -->
              <div
                class="flex flex-col justify-between px-3 py-1 relative group text-left h-full"
              >
                <div class="text-xs text-gray-500 mb-0.5">Check-in</div>
                <div
                  :class="[
                    'font-medium',
                    value?.end ? 'text-gray-900' : 'text-gray-400',
                  ]"
                >
                  {{ value?.start ? formattedStart : "Select date" }}
                </div>
              </div>

              <!-- Nights Counter -->
              <div
                class="flex items-center justify-center h-full text-sm text-gray-500"
              >
                {{ calculateNights }} nights
              </div>

              <!-- Check-out Section -->
              <div
                class="flex flex-col justify-between px-3 py-1 relative group text-left h-full"
              >
                <div class="text-xs text-gray-500 mb-0.5">Check-out</div>
                <div
                  :class="[
                    'font-medium',
                    value?.end ? 'text-gray-900' : 'text-gray-400',
                  ]"
                >
                  {{ value?.end ? formattedEnd : "Select date" }}
                </div>
              </div>
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0" align="center">
            <div class="flex">
              <RangeCalendar v-model="value" class="rounded-md border" />
            </div>
          </PopoverContent>
        </Popover>
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
                    value?.end ? 'text-gray-900' : 'text-gray-400',
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
        <Button class="w-full h-full bg-primary hover:bg-primary/90">
          <Search class="" />
          Search
        </Button>
      </div>
    </div>
    <div class="mt-4 flex items-center justify-between container">
      <div class="flex gap-2 overflow-x-auto pb-2">
        <Badge
          v-for="filter in quickFilters"
          :key="filter"
          variant="outline"
          class="cursor-pointer hover:bg-primary hover:text-primary-foreground"
        >
          {{ filter }}
        </Badge>
      </div>
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
