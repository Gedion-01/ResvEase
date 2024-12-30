<script setup lang="ts">
import { ref } from "vue";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import SearchSection from "@/components/SearchSection.vue";
import { List, Map } from "lucide-vue-next";
import SearchFilters from "@/components/SearchFilters.vue";
import HotelList from "@/components/HotelList.vue";

const quickFilters = [
  "Free cancellation",
  "Breakfast included",
  "Pool",
  "Beach access",
  "Pet friendly",
  "5-star hotels",
];

const viewMode = ref<"list" | "map">("list");
</script>

<template>
  <section class="min-h-screen bg-background flex flex-col">
    <div class="container py-8 flex-grow">
      <div class="sticky top-14 z-20 bg-background pb-4">
        <SearchSection />
        <div class="mt-4 flex items-center justify-between">
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
      <div class="flex flex-col md:flex-row gap-8 mt-4">
        <aside class="w-full md:w-1/4 md:sticky md:top-[260px] self-start">
          <div class="space-y-6">
            <SearchFilters />
          </div>
        </aside>
        <main class="w-full md:w-3/4">
          <HotelList v-if="viewMode === 'list'" />
          <div
            v-else-if="viewMode === 'map'"
            class="h-[calc(100vh-12rem)] rounded-lg overflow-hidden"
          >
            <iframe
              src="https://www.google.com/maps/embed"
              width="100%"
              height="100%"
              style="border: 0"
              allowfullscreen
              loading="lazy"
              referrerpolicy="no-referrer-when-downgrade"
            />
          </div>
        </main>
      </div>
    </div>
  </section>
</template>
