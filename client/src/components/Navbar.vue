<script setup lang="ts">
import { RouterLink } from "vue-router";
import { LogOut, Moon, Sun, User } from "lucide-vue-next";
import { useAuthStore } from "@/store/authStore";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";

const authStore = useAuthStore();

const handleLogout = () => {
  authStore.logout();
};
</script>
<template>
  <header
    class="top-0 z-50 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 mx-auto"
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
            <Popover v-if="authStore.isAuthenticated">
              <PopoverTrigger asChild>
                <Button
                  variant="outline"
                  class="ml-4 flex items-center justify-center rounded-full h-9 w-9 p-0"
                >
                  <User class="h-5 w-5" />
                </Button>
              </PopoverTrigger>
              <PopoverContent class="w-48">
                <div class="flex flex-col space-y-2">
                  <RouterLink
                    to="/profile"
                    class="flex items-center space-x-2 p-2 hover:bg-gray-100 rounded-md"
                  >
                    <User class="h-4 w-4" />
                    <span>Profile</span>
                  </RouterLink>
                  <button
                    @click="handleLogout"
                    class="flex items-center space-x-2 p-2 hover:bg-gray-100 rounded-md"
                  >
                    <LogOut class="h-4 w-4" />
                    <span>Logout</span>
                  </button>
                </div>
              </PopoverContent>
            </Popover>
          </nav>
        </div>
      </div>
    </div>
  </header>
</template>
