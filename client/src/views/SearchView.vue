<script setup lang="ts">
import { onMounted, onUnmounted, ref } from "vue";
import SearchSection from "@/components/SearchSection.vue";
import SearchFilters from "@/components/SearchFilters.vue";
import HotelList from "@/components/HotelList.vue";

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

const viewMode = ref<"list" | "map">("list");
</script>

<template>
  <section class="min-h-screen bg-slate-100 flex flex-col">
    <div :class="['py-1 flex-grow']">
      <div
        ref="searchSection"
        :class="[
          'sticky top-0 z-20 pb-4 transition-all ease-in-out duration-200 ',
          isSticky ? '' : 'container',
        ]"
      >
        <SearchSection :is-sticky="isSticky" />
      </div>
      <div class="flex flex-col md:flex-row gap-8 mt-4 container">
        <aside
          class="w-full md:w-1/4 md:sticky md:top-[232px] self-start group"
        >
          <div class="space-y-6 group-sticky:bg-red-500">
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
