import { useInfiniteQuery, useQuery } from "@tanstack/vue-query";
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

export function useInfiniteHotels(queryParams: Record<string, any>) {
  return useInfiniteQuery({
    queryKey: ["infinite-hotels", queryParams],
    queryFn: ({ pageParam = 1 }) =>
      getHotels({ ...queryParams, page: pageParam }),
    getNextPageParam: (lastPage) => lastPage.cursor,
    initialPageParam: 1,
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

export function useInfiniteRooms(
  hotelId: string,
  queryParams: Record<string, any>
) {
  return useInfiniteQuery({
    queryKey: ["infinite-rooms", queryParams],
    queryFn: ({ pageParam = 1 }) =>
      getRooms(hotelId, { ...queryParams, page: pageParam }),
    getNextPageParam: (lastPage) => lastPage.cursor,
    initialPageParam: 1,
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
