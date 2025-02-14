import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import SearchVue from "@/views/SearchView.vue";
import HotelView from "@/views/HotelView.vue";
import BookView from "@/views/BookView.vue";
import AuthView from "@/views/AuthView.vue";
import BookingsView from "@/views/BookingsView.vue";
import BookingDetailsView from "@/views/BookingDetailsView.vue";
import { useAuthStore } from "@/store/authStore";

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
      name: "SearchResults",
      component: SearchVue,
    },
    {
      path: "/hotel/:id/rooms",
      name: "Hotel",
      component: HotelView,
    },
    {
      path: "/booking",
      name: "Book",
      component: BookView,
      meta: { requiresAuth: true },
    },
    {
      path: "/auth",
      name: "Auth",
      component: AuthView,
      meta: { requiresGuest: true, hideNavbar: true },
    },
    {
      path: "/bookings",
      name: "Bookings",
      component: BookingsView,
      meta: { requiresAuth: true },
    },
    {
      path: "/booking-details/:id",
      name: "BookingDetails",
      component: BookingDetailsView,
      meta: { requiresAuth: true },
    },
  ],
});

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    authStore.setRedirectUrl(to.fullPath);
    next("/auth?tab=login");
  } else if (to.meta.requiresGuest && authStore.isAuthenticated) {
    next("/");
  } else {
    next();
  }
});

export default router;
