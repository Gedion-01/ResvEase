<script setup lang="ts">
import HotelCard from "./HotelCard.vue";
// This would typically come from an API or database
const mockHotels = [
  {
    id: "1",
    name: "Luxury Resort & Spa",
    image: "/hotel1.jpg",
    rating: 4.5,
    prices: [200, 250, 300],
    location: "Maldives",
    bedTypes: ["King", "Queen"],
    bedrooms: 2,
  },
  {
    id: "2",
    name: "City Center Hotel",
    image: "/hotel2.jpg",
    rating: 4.2,
    prices: [150, 180, 220],
    location: "New York",
    bedTypes: ["Queen", "Twin"],
    bedrooms: 1,
  },
  {
    id: "3",
    name: "Mountain Retreat",
    image: "/hotel3.jpg",
    rating: 4.8,
    prices: [280, 350, 400],
    location: "Swiss Alps",
    bedTypes: ["King"],
    bedrooms: 3,
  },
  {
    id: "4",
    name: "Beachfront Paradise",
    image: "/hotel4.jpg",
    rating: 4.6,
    prices: [220, 270, 320],
    location: "Bali",
    bedTypes: ["King", "Queen"],
    bedrooms: 2,
  },
  {
    id: "5",
    name: "Historic Downtown Inn",
    image: "/hotel5.jpg",
    rating: 4.3,
    prices: [180, 220, 260],
    location: "Prague",
    bedTypes: ["Queen", "Twin"],
    bedrooms: 1,
  },
  {
    id: "6",
    name: "Ski Lodge & Resort",
    image: "/hotel6.jpg",
    rating: 4.7,
    prices: [300, 350, 400],
    location: "Aspen",
    bedTypes: ["King", "Queen", "Twin"],
    bedrooms: 4,
  },
];
</script>
<template>
  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
    <HotelCard v-for="hotel in mockHotels" :key="hotel.id" :hotel="hotel" />
  </div>
</template>
