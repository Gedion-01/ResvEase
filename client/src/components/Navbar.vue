<script setup lang="ts">
import { RouterLink, useRoute } from "vue-router";
import { LogOut, Moon, Sun, User } from "lucide-vue-next";
import { useAuthStore } from "@/store/authStore";
import { useSearchStore } from "@/store/searchStore";
import { useFilterStore } from "@/store/filterStore";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { computed, reactive, ref } from "vue";

const isScrolled = ref(false);
const isMobileMenuOpen = ref(false);

const handleScroll = () => {
  isScrolled.value = window.scrollY > 10;
};

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
};

const route = useRoute();

const authStore = useAuthStore();
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
});

const queryString = computed(() => new URLSearchParams(queryParams).toString());

const handleLogout = () => {
  authStore.logout();
};
</script>

<template>
  <header
    :class="[
      'top-0 left-0 right-0 z-50 transition-all duration-300',
      isScrolled ? 'bg-slate-100 shadow-md' : 'bg-slate-100',
      route.name !== 'Home' ? 'border-b border-slate-300' : '',
    ]"
  >
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex h-16 items-center justify-between">
        <div class="flex items-center">
          <RouterLink class="flex items-center space-x-2" to="/">
            <!-- <img
              src="/path-to-your-logo.svg"
            
              class="h-8 w-auto"
            /> -->
            <span
              class="hidden text-xl font-bold text-slate-800 sm:inline-block"
              >ResvEase</span
            >
          </RouterLink>
        </div>

        <nav class="hidden md:flex items-center space-x-6 main-nav">
          <RouterLink
            v-for="(link, index) in ['Search', 'My Bookings', 'Sign in']"
            :key="index"
            :to="
              link === 'Search'
                ? `/search?${queryString}`
                : link === 'My Bookings'
                ? '/bookings'
                : link === 'Sign in'
                ? '/auth'
                : `/${link.toLowerCase()}`
            "
            :class="[
              'text-sm nav-link font-medium transition-colors hover:text-slate-900',
              (authStore.isAuthenticated && link === 'Sign in') ||
              (!authStore.isAuthenticated && link === 'My Bookings')
                ? 'hidden'
                : '',
              route.path.includes(link.toLowerCase())
                ? 'text-slate-900'
                : 'text-slate-600',
            ]"
          >
            {{ link }}
          </RouterLink>
        </nav>

        <div class="flex items-center space-x-4">
          <Button
            variant="ghost"
            size="icon"
            class="relative text-slate-600 hover:text-slate-900"
          >
            <Sun
              class="h-5 w-5 rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
            />
            <Moon
              class="absolute h-5 w-5 rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
            />
            <span class="sr-only">Toggle theme</span>
          </Button>

          <Popover v-if="authStore.isAuthenticated">
            <PopoverTrigger asChild>
              <Button variant="outline" size="icon" class="rounded-full">
                <User class="h-5 w-5" />
              </Button>
            </PopoverTrigger>
            <PopoverContent class="w-56">
              <div class="flex flex-col space-y-1">
                <RouterLink
                  to="/profile"
                  class="flex items-center space-x-2 p-2 hover:bg-slate-100 rounded-md transition-colors"
                >
                  <User class="h-4 w-4" />
                  <span>Profile</span>
                </RouterLink>
                <button
                  @click="handleLogout"
                  class="flex items-center space-x-2 p-2 hover:bg-slate-100 rounded-md transition-colors text-left w-full"
                >
                  <LogOut class="h-4 w-4" />
                  <span>Logout</span>
                </button>
              </div>
            </PopoverContent>
          </Popover>

          <Button
            variant="ghost"
            size="icon"
            class="md:hidden text-slate-600 hover:text-slate-900"
            @click="toggleMobileMenu"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 6h16M4 12h16m-7 6h7"
              />
            </svg>
          </Button>
        </div>
      </div>
    </div>

    <!-- Mobile Menu -->
    <div v-if="isMobileMenuOpen" class="md:hidden bg-slate-50">
      <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
        <RouterLink
          v-for="(link, index) in ['Search', 'Deals', 'My Bookings', 'Sign in']"
          :key="index"
          :to="
            link === 'Search'
              ? `/search?${queryString}`
              : link === 'My Bookings'
              ? '/bookings'
              : link === 'Sign in'
              ? '/auth'
              : `/${link.toLowerCase()}`
          "
          :class="[
            'block px-3 py-2 rounded-md text-base font-medium',
            (authStore.isAuthenticated && link === 'Sign in') ||
            (!authStore.isAuthenticated && link === 'My Bookings')
              ? 'hidden'
              : '',
            route.path.includes(link.toLowerCase())
              ? 'bg-slate-200 text-slate-900'
              : 'text-slate-600 hover:bg-slate-100 hover:text-slate-900',
          ]"
          @click="toggleMobileMenu"
        >
          {{ link }}
        </RouterLink>
      </div>
    </div>
  </header>
</template>

<style scoped>
.router-link-active {
  color: theme("colors.primary.DEFAULT");
}
.nav-link {
  position: relative;
  padding-bottom: 2px;
  transition: color 0.3s ease;
}
.nav-link::after {
  content: "";
  position: absolute;
  bottom: 0px;
  left: 0;
  width: 100%;
  height: 3px;
  background-color: theme("colors.primary.DEFAULT");
  transform: scaleX(0);
  transform-origin: left;
  transition: transform 0.3s ease;
  border-radius: 4px;
}
.nav-link:hover::after,
.nav-link.router-link-active::after {
  transform: scaleX(1);
}
</style>
