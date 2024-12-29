<script setup lang="ts">
import type { DateRange } from 'radix-vue'
import { RangeCalendar } from '@/components/ui/range-calendar'
import { getLocalTimeZone, today } from '@internationalized/date'
import { type Ref, ref, computed, reactive } from "vue";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { format } from "date-fns";

const location = ref("");
const guests = reactive({
  adults: 2,
  children: 0,
}) 

const timeZone = getLocalTimeZone()
const start = undefined
const end = undefined

const value = ref({
  start,
  end,
}) as Ref<DateRange, DateRange> | undefined

const formattedStart = computed(() => {
  if (value?.value.start) {
    return format(value?.value.start.toDate(timeZone), "iii, MMM dd")
  }
  return '' 
})
const formattedEnd = computed(() => {
  if (value?.value.end) {
    return format(value?.value.end.toDate(timeZone), "iii, MMM dd")
  }
  return '' 
})

const calculateNights = computed(() => {
  if (value?.value.start && value.value.end) {
    const startDate = value.value.start.toDate(timeZone)
    const endDate = value.value.end.toDate(timeZone)
    return Math.ceil((endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60 * 24))
  }
  return 0 
})

const decrementAdults = () => {
  guests.adults = Math.max(1, guests.adults - 1)
}
const incrementAdults = () => {
  guests.adults++;
}

const decrementChildren = () => {
  guests.children = Math.max(1, guests.children - 1)
}
const incrementChildren = () => {
  guests.children++;
}
</script>

<template>
    <div class="w-full max-w-6xl mx-auto bg-background rounded-lg shadow-lg p-4">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-4">
        <div class="md:col-span-4 relative">
          <Input
            type="text"
            placeholder="City, airport, region, landmark or property name"
            v-model="location"
            class="w-full pr-8"
          />
            <button
              @click="location = ''"
              class="absolute right-2 top-1/2 -translate-y-1/2"
            >
              <X class="h-4 w-4 text-gray-500" />
            </button>
        </div>
        
        <div class="md:col-span-5">
          <Popover>
            <PopoverTrigger asChild>
               <div class="flex items-center justify-between border rounded-md p-1 bg-white cursor-pointer hover:border-gray-300 transition-all">
          <!-- Check-in Section -->
          <div class="flex-1 px-3 py-1 relative group">
            <div class="text-xs text-gray-500 mb-0.5">Check-in</div>
            <div :class="['font-medium',
 value?.end ? 'text-gray-900' : 'text-gray-400'
            ]">
              {{ value?.start ? formattedStart : 'Select date' }}
            </div>
            <div 
              v-if="value?.start" 
              class="absolute bottom-0 left-3 right-3 h-0.5 bg-blue-500"
            />
          </div>

          <!-- Nights Counter -->
          <div class="px-3 text-sm text-gray-500 self-end pb-1">
            {{ calculateNights }} nights
          </div>

          <!-- Check-out Section -->
          <div class="flex-1 px-3 py-1 relative group">
            <div class="text-xs text-gray-500 mb-0.5">Check-out</div>
            <div :class="[
              'font-medium',
              value?.end ? 'text-gray-900' : 'text-gray-400'
            ]">
              {{ value?.end ? formattedEnd : 'Select date' }}
            </div>
            <div 
              v-if="value?.end" 
              class="absolute bottom-0 left-3 right-3 h-0.5 bg-blue-500"
            />
          </div>
        </div>
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
              <Button variant="outline" class="w-full justify-between">
                <div class="flex items-center">
                  <Users class="mr-2 h-4 w-4" />
                  <span>
                    {{ guests.adults + guests.children}} Guest{{guests.adults + guests.children !== 1 ? 's' : ''}}
                  </span>
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
                    <span class="w-8 text-center">{{guests.adults}}</span>
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
                    <span class="w-8 text-center">{{guests.children}}</span>
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

        <div class="md:col-span-1">
          <Button 
            class="w-full bg-primary hover:bg-primary/90" 
            onClick={handleSearch}
          >
            Search
          </Button>
        </div>
      </div>
    </div>
</template>
