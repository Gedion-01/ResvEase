<script setup lang="ts">
import { useAuthStore } from "./store/authStore";
import { onMounted } from "vue";
import { RouterView } from "vue-router";
import Navbar from "@/components/Navbar.vue";
import TestVIew from "./views/TestVIew.vue";

const authStore = useAuthStore();

onMounted(() => {
  authStore.initializeTokenCheck();
});
</script>

<template>
  <Navbar />
  <RouterView />
  <!-- <TestVIew /> -->
</template>
