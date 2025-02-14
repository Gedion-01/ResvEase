<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import SignupForm from "@/components/auth/SignupForm.vue";
import LoginForm from "@/components/auth/LoginForm.vue";

const route = useRoute();
const router = useRouter();

const activeTab = computed({
  get: () => (route.query.tab === "signup" ? "signup" : "login"),
  set: (tab) => router.push({ query: { tab } }),
});

const featuredHotels = [
  {
    name: "Luxury Resort & Spa",
    image:
      "https://images.unsplash.com/photo-1590523277543-a94d2e4eb00b?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    location: "Maldives",
  },
  {
    name: "Urban Oasis Hotel",
    image:
      "https://cf.bstatic.com/xdata/images/hotel/max1024x768/532330328.jpg?k=c4a818e4799cb35e6c08e2c164ac6e74259a1de03ce5f35bff138b822219933d&o=&hp=1",
    location: "New York",
  },
  {
    name: "Mountain Lodge Retreat",
    image:
      "https://media.cntraveler.com/photos/5bb26d6d625211259a97383c/16:9/w_3599,h_2024,c_limit/Manor-Vail-Lodge__2018_Manor-Vail-Lodge-Winter-Exterior.jpg",
    location: "Swiss Alps",
  },
  {
    name: "Beachfront Paradise",
    image:
      "https://cf.bstatic.com/xdata/images/hotel/max1024x768/547748864.jpg?k=bebf03dc9ba5e5327d76078e90c45ffa9049ec5b16851d0a03b350dc830d147b&o=&hp=1",
    location: "Bali",
  },
];

const currentHotelIndex = ref(0);

const changeHotel = () => {
  currentHotelIndex.value =
    (currentHotelIndex.value + 1) % featuredHotels.length;
};

setInterval(changeHotel, 5000);

const goToHome = () => {
  router.push("/");
};
</script>

<template>
  <div class="min-h-screen flex bg-gray-100 w-full">
    <div class="hidden lg:flex lg:w-1/2 relative overflow-hidden">
      <div
        v-for="(hotel, index) in featuredHotels"
        :key="index"
        :class="{
          'absolute inset-0 bg-cover bg-center transition-opacity duration-1000': true,
          'opacity-100': index === currentHotelIndex,
          'opacity-0': index !== currentHotelIndex,
        }"
        :style="{ backgroundImage: `url(${hotel.image})` }"
      >
        <div class="absolute inset-0 bg-black bg-opacity-40"></div>
        <div class="absolute bottom-0 left-0 right-0 p-8 text-white">
          <h3 class="text-2xl font-bold mb-2">{{ hotel.name }}</h3>
          <p class="text-lg">{{ hotel.location }}</p>
        </div>
      </div>
    </div>

    <div class="w-full lg:w-1/2 flex items-center justify-center p-8">
      <div class="w-full max-w-md">
        <div class="text-center mb-8">
          <!-- <img src="/logo.png" alt="HotelHub Logo" class="mx-auto h-16 mb-4" /> -->
          <h2 class="text-3xl font-bold text-gray-800">
            Welcome to BookMyStay
          </h2>
          <p class="text-gray-600">
            Your gateway to extraordinary stays worldwide
          </p>
        </div>

        <Tabs v-model:modelValue="activeTab" class="w-full">
          <TabsList class="grid w-full grid-cols-2 mb-8">
            <TabsTrigger value="login" class="flex items-center justify-center">
              <i class="fas fa-sign-in-alt mr-2"></i> Login
            </TabsTrigger>
            <TabsTrigger
              value="signup"
              class="flex items-center justify-center"
            >
              <i class="fas fa-user-plus mr-2"></i> Sign Up
            </TabsTrigger>
          </TabsList>
          <TabsContent value="login">
            <LoginForm />
          </TabsContent>
          <TabsContent value="signup">
            <SignupForm />
          </TabsContent>
        </Tabs>

        <div class="mt-8 text-center">
          <p class="text-sm text-gray-600">
            By continuing, you agree to BookMyStay's Terms of Service and
            Privacy Policy.
          </p>
          <button
            @click="goToHome"
            class="text-blue-600 hover:text-blue-800 transition duration-300 flex items-center justify-center mx-auto mt-8"
          >
            <i class="fas fa-home mr-2"></i> Back to Home
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css");
</style>
