import { useQuery } from "@tanstack/vue-query";
import { getHotel, getHotels, getRooms } from "./api";

export function useHotel(hotelId: string) {
  return useQuery({
    queryKey: ["hotel", hotelId],
    queryFn: () => getHotel(hotelId),
  });
}

export function useHotels() {
  return useQuery({
    queryKey: ["hotels"],
    queryFn: getHotels,
  });
}

export function useRooms(hotelId: string, queryParams: Record<string, any>) {
  return useQuery({
    queryKey: ["rooms", hotelId],
    queryFn: () => getRooms(hotelId, queryParams),
  });
}
