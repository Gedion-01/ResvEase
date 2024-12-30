import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeVue.vue";
import SearchVue from "@/views/SearchVue.vue";

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
  ],
});

export default router;
