<script setup lang="ts">
import { useAuthStore } from "./store/authStore";
import { computed, onMounted } from "vue";
import { RouterView, useRoute } from "vue-router";
import Navbar from "@/components/Navbar.vue";
import Toaster from "@/components/ui/toast/Toaster.vue";

const authStore = useAuthStore();

onMounted(() => {
  authStore.initializeTokenCheck();
});
const route = useRoute();
const hideNavbar = computed(() => route.meta.hideNavbar === true);
</script>

<template>
  <Navbar v-if="!hideNavbar" />
  <RouterView />
  <Toaster />
</template>
