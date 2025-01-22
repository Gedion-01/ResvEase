<script setup lang="ts">
import BookingFlow from "@/components/BookingFlow.vue";
</script>
<template>
  <BookingFlow
    :room="{
      id: '1',
      name: 'Standard Room',
      price: 100,
      image: 'https://via.placeholder.com/300',
      bedType: 'Double',
    }"
    :checkIn="new Date()"
    :checkOut="new Date(new Date().setDate(new Date().getDate() + 1))"
  />
</template>
