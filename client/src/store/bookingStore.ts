import type { Hotel, Room } from "@/types/hotel";
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
    hotel: {} as Hotel,
  }),
  actions: {
    setRoomBookingDetails(room: Room) {
      this.room = room;
    },
    setHotelBookingDetails(hotel: Hotel) {
      this.hotel = hotel;
    },
  },
});
