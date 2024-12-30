import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import SearchVue from "@/views/SearchView.vue";
import HotelView from "@/views/HotelView.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "Home",
      component: HomeView,
    },
    {
      path: "/search",
      name: "Search",
      component: SearchVue,
    },
    {
      path: "/hotel/:id",
      name: "Hotel",
      component: HotelView,
    },
  ],
});

export default router;
