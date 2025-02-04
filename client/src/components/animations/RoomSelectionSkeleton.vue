<script setup lang="ts">
import { Skeleton } from "@/components/ui/skeleton";
import { Card, CardHeader, CardContent, CardTitle } from "@/components/ui/card";
</script>

<template>
  <Card class="cursor-pointer transition-colors">
    <CardHeader class="p-0">
      <div class="relative">
        <Skeleton class="w-full aspect-video rounded-t-lg rounded-b-none" />
      </div>
    </CardHeader>
    <CardContent class="p-4">
      <CardTitle class="mb-2">
        <Skeleton class="h-6 w-3/4" />
      </CardTitle>
      <Skeleton class="h-4 w-full mb-2" />
      <Skeleton class="h-4 w-5/6 mb-4" />
      <div class="flex flex-wrap justify-between gap-3 mb-4">
        <Skeleton class="h-5 w-24" />
        <Skeleton class="h-5 w-24" />
      </div>
      <div class="flex flex-wrap gap-2 mb-4">
        <Skeleton class="h-6 w-20" />
        <Skeleton class="h-6 w-24" />
        <Skeleton class="h-6 w-16" />
      </div>
      <div class="flex justify-between items-center">
        <Skeleton class="h-8 w-24" />
        <Skeleton class="h-10 w-24" />
      </div>
    </CardContent>
  </Card>
</template>
