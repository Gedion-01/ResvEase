import { defineStore } from "pinia";

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
        this.location = parsedParams.location;
        this.checkIn = parsedParams.checkIn;
        this.checkOut = parsedParams.checkOut;
        this.adults = parsedParams.adults;
        this.children = parsedParams.children;
      }
    },
  },
});
