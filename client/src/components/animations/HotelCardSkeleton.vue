<script setup lang="ts">
import { Skeleton } from "@/components/ui/skeleton";
</script>
<template>
  <div
    class="bg-card text-card-foreground rounded-lg overflow-hidden shadow-lg"
  >
    <div class="relative h-48">
      <Skeleton class="w-full h-48" />
      <div class="absolute bottom-0 left-0 right-0 p-4">
        <Skeleton class="h-6 w-3/4 mb-2" />
        <Skeleton class="h-4 w-1/2" />
      </div>
    </div>
    <div class="p-4">
      <div class="flex items-center justify-between mb-2">
        <Skeleton class="h-5 w-24" />
        <Skeleton class="h-6 w-20" />
      </div>
      <Skeleton class="h-10 w-full" />
    </div>
  </div>
</template>
