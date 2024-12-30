<script setup lang="ts">
import HotelDetails from "@/components/HotelDetails.vue";

// This would typically come from an API or database
const mockHotel = {
  id: "1",
  name: "Luxury Resort & Spa",
  images: ["/hotel1.jpg", "/hotel2.jpg", "/hotel3.jpg"],
  rating: 4.5,
  price: 250,
  location: "Maldives",
  description:
    "Experience luxury and relaxation in our beachfront resort. Enjoy stunning ocean views, world-class amenities, and impeccable service.",
  amenities: [
    "Free Wi-Fi",
    "Swimming Pool",
    "Spa",
    "Restaurant",
    "Fitness Center",
    "Beach Access",
  ],
};
</script>

<template>
  <div class="container mx-auto px-4 py-8">
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      <div class="md:col-span-2">
        <HotelDetails :hotel="mockHotel" />
        <!-- <div class="mt-8">
          <RoomSelection hotelId={id} onRoomSelect={(roomId) => {
          console.log('Selected room:', roomId) // In a real app, this would
          update the booking form with the selected room }} />
        </div> -->
      </div>
      <div class="md:sticky md:top-20 h-fit">
        <BookingForm hotel="{mockHotel}" />
      </div>
    </div>
  </div>
</template>
