import { useQuery } from "@tanstack/vue-query";
import { getHotels } from "./api";

export function useHotels() {
  return useQuery({
    queryKey: ["hotels"],
    queryFn: getHotels,
  });
}
