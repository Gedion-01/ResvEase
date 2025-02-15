<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

const destinations = [
  {
    name: "Paris",
    image:
      "https://images.unsplash.com/photo-1502602898657-3e91760cbb34?q=80&w=2073&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    description: "The City of Light",
  },
  {
    name: "Tokyo",
    image:
      "https://plus.unsplash.com/premium_photo-1661902398022-762e88ff3f82?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    description: "Where tradition meets future",
  },
  {
    name: "New York",
    image:
      "https://images.unsplash.com/photo-1500916434205-0c77489c6cf7?q=80&w=1887&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    description: "The Big Apple",
  },
  {
    name: "Bali",
    image:
      "https://images.unsplash.com/photo-1546484475-7f7bd55792da?q=80&w=1887&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    description: "Island of the Gods",
  },
];
</script>

<template>
  <section class="py-12 bg-slate-200">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <h2 class="text-3xl font-bold text-center mb-8">Featured Destinations</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <Card
          v-for="destination in destinations"
          :key="destination.name"
          class="overflow-hidden"
        >
          <img
            :src="destination.image"
            :alt="destination.name"
            class="w-full h-48 object-cover"
          />
          <CardHeader>
            <CardTitle>{{ destination.name }}</CardTitle>
            <CardDescription>{{ destination.description }}</CardDescription>
          </CardHeader>
          <CardFooter>
            <Button variant="outline" class="w-full">Explore</Button>
          </CardFooter>
        </Card>
      </div>
    </div>
  </section>
</template>
