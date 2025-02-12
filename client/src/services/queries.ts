import { useQuery } from "@tanstack/vue-query";
import {
  getBooking,
  getHotel,
  getHotels,
  getRooms,
  getUserBookings,
} from "./api";

export function useHotel(hotelId: string) {
  return useQuery({
    queryKey: ["hotel", hotelId],
    queryFn: () => getHotel(hotelId),
  });
}

export function useHotels(queryParams: Record<string, any>) {
  return useQuery({
    queryKey: ["hotels"],
    queryFn: () => getHotels(queryParams),
  });
}

export function useRooms(hotelId: string, queryParams: Record<string, any>) {
  return useQuery({
    queryKey: ["rooms", hotelId],
    queryFn: () => getRooms(hotelId, queryParams),
  });
}

export function useUserBookings(userID: string) {
  return useQuery({
    queryKey: ["bookings"],
    queryFn: () => getUserBookings(userID),
  });
}

export function useBooking(bookingId: string) {
  return useQuery({
    queryKey: ["booking", bookingId],
    queryFn: () => getBooking(bookingId),
  });
}
