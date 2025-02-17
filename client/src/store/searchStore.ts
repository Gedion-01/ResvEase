import { defineStore } from "pinia";
import { addDays, format, isBefore, parseISO, startOfDay } from "date-fns";

export const useSearchStore = defineStore("search", {
  state: () => ({
    location: "",
    checkIn: "",
    checkOut: "",
    adults: 2,
    children: 0,
  }),
  actions: {
    setSearchParams(params: {
      location: string;
      checkIn: string;
      checkOut: string;
      adults: number;
      children: number;
    }) {
      this.location = params.location;
      this.checkIn = params.checkIn;
      this.checkOut = params.checkOut;
      this.adults = params.adults;
      this.children = params.children;
      localStorage.setItem("searchParams", JSON.stringify(params));
    },
    loadSearchParams() {
      const params = localStorage.getItem("searchParams");
      if (params) {
        const parsedParams = JSON.parse(params);
        const storedCheckIn = parsedParams.checkIn
          ? parseISO(parsedParams.checkIn)
          : null;
        const storedCheckOut = parsedParams.checkOut
          ? parseISO(parsedParams.checkOut)
          : null;
        const today = startOfDay(new Date());

        if (storedCheckIn && isBefore(storedCheckIn, today)) {
          this.checkIn = format(today, "yyyy-MM-dd");
          this.checkOut = format(addDays(today, 1), "yyyy-MM-dd");
        } else {
          this.checkIn = parsedParams.checkIn;
          this.checkOut = parsedParams.checkOut;
        }

        this.location = parsedParams.location;
        this.adults = parsedParams.adults;
        this.children = parsedParams.children;
      }
    },
  },
});
