<script setup lang="ts">
import { RouterLink } from "vue-router";
import { Moon, Sun } from "lucide-vue-next";
import SearchSection from "./SearchSection.vue";
</script>
<template>
  <header
    class="sticky top-0 z-50 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 mx-auto"
  >
    <div class="w-full border-b">
      <div class="flex container h-14 items-center">
        <div class="mr-4 hidden md:flex">
          <RouterLink class="mr-6 flex items-center space-x-2" to="/">
            <span class="hidden font-bold sm:inline-block">BookMyStay</span>
          </RouterLink>
          <nav class="flex items-center space-x-6 text-sm font-medium">
            <RouterLink to="/search">Search</RouterLink>
            <RouterLink to="/deals">Deals</RouterLink>
            <RouterLink to="/account">My Account</RouterLink>
          </nav>
        </div>
        <div
          class="flex flex-1 items-center justify-between space-x-2 md:justify-end"
        >
          <nav class="flex items-center">
            <button
              class="inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-9 py-2 w-9 px-0"
            >
              <Sun
                class="h-[1.5rem] w-[1.3rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
              />
              <Moon
                class="absolute h-[1.5rem] w-[1.3rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
              />
              <span class="sr-only">Toggle theme</span>
            </button>
          </nav>
        </div>
      </div>
    </div>
    <!-- <SearchSection /> -->
  </header>
</template>
