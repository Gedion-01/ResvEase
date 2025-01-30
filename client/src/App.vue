<script setup lang="ts">
import { RouterView } from "vue-router";
import Navbar from "@/components/Navbar.vue";
import TestVIew from "./views/TestVIew.vue";
</script>

<template>
  <Navbar />
  <RouterView />
  <!-- <TestVIew /> -->
</template>
