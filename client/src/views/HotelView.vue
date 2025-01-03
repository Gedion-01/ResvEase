<script setup lang="ts">
import HotelDetails from "@/components/HotelDetails.vue";
import RoomSelection from "@/components/RoomSelection.vue";
import { useHotel, useRooms } from "@/services/queries";
import { useRoute } from "vue-router";

// Define the type for the route parameters
interface RouteParams {
  id: string; // Example: if the route is '/items/:id'
}

const route = useRoute();

// Access the parameter and cast it to the type
const id = route.params.id as RouteParams["id"];
console.log(id);

// This would typically come from an API or database
// const mockHotel = {
//   id: "1",
//   name: "Luxury Resort & Spa",
//   images: ["/hotel1.jpg", "/hotel2.jpg", "/hotel3.jpg"],
//   rating: 4.5,
//   price: 250,
//   location: "Maldives",
//   description:
//     "Experience luxury and relaxation in our beachfront resort. Enjoy stunning ocean views, world-class amenities, and impeccable service.",
//   amenities: [
//     "Free Wi-Fi",
//     "Swimming Pool",
//     "Spa",
//     "Restaurant",
//     "Fitness Center",
//     "Beach Access",
//   ],
// };

const mockHotel = {
  id: "1",
  name: "Luxury Resort & Spa",
  images: [
    "https://plus.unsplash.com/premium_photo-1661929519129-7a76946c1d38?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://images.unsplash.com/photo-1561501900-3701fa6a0864?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://images.unsplash.com/photo-1583037189850-1921ae7c6c22?q=80&w=1975&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  ],
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

// const { data, isLoading } = useRooms(id);
const { data, isLoading } = useHotel(id);
</script>

<template>
  <!-- {{ data?.data }} -->
  <div class="container mx-auto px-4 py-8">
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      <div class="md:col-span-2" v-if="data">
        <HotelDetails :hotel="data" />
        <div class="mt-8">
          <RoomSelection
            :hotelId="id"
            :onRoomSelect="
              (roomId) => {
                console.log('Selected room:', roomId);
                // In a real app, this would update the booking form with the selected room
              }
            "
          />
        </div>
      </div>
      <div class="md:sticky md:top-20 h-fit">
        <BookingForm hotel="{mockHotel}" />
      </div>
    </div>
  </div>
</template>
