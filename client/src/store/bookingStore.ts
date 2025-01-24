import type { Room } from "@/types/hotel";
import { defineStore } from "pinia";

export const useRoomBookingStore = defineStore("roomBooking", {
  state: () => ({
    room: {
      id: "",
      name: "",
      description: "",
      price: 0,
      capacity: 0,
      bedType: "",
      bedRooms: 0,
      availableCount: 0,
      amenities: [],
      images: [],
    } as Room,
    hotel: {
      name: "",
      rating: 0,
    },
  }),
  actions: {
    setRoomBookingDetails(room: Room) {
      this.room = room;
    },
    setHotelBookingDetails(name: string, rating: number) {
      this.hotel = {
        name,
        rating,
      };
    },
  },
});
