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
    adults: 1,
    children: 0,
  }),
  actions: {
    loadBookingData() {
      const storedBooking = localStorage.getItem("Booking");
      if (storedBooking) {
        const parsedBooking = JSON.parse(storedBooking);
        if (parsedBooking.room) {
          this.room = parsedBooking.room;
        }
        if (parsedBooking.hotel) {
          this.hotel = parsedBooking.hotel;
        }
        if (typeof parsedBooking.adults === "number") {
          this.adults = parsedBooking.adults;
        }
        if (typeof parsedBooking.children === "number") {
          this.children = parsedBooking.children;
        }
      }
    },
    saveToLocalStorage() {
      const bookingData = {
        room: this.room,
        hotel: this.hotel,
        adults: this.adults,
        children: this.children,
      };
      localStorage.setItem("Booking", JSON.stringify(bookingData));
    },
    setRoomBookingDetails(room: Room) {
      this.room = room;
      this.saveToLocalStorage();
    },
    setHotelBookingDetails(hotel: Hotel) {
      this.hotel = hotel;
      this.saveToLocalStorage();
    },
    setGuests(adults: number, children: number) {
      this.adults = adults;
      this.children = children;
      this.saveToLocalStorage();
    },
  },
});
