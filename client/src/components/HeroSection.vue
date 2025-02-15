<!-- <script setup lang="ts"></script>
<template>
  <div className="relative w-full h-[60vh] mb-8">
    <img
      src="https://images.unsplash.com/photo-1582719508461-905c673771fd?q=80&w=1925&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
      alt="Luxury hotel room"
      layout="fill"
      objectFit="cover"
      priority
    />
    <div
      className="absolute inset-0 bg-black bg-opacity-50 flex flex-col justify-center items-center text-white"
    >
      <h1 className="text-4xl md:text-6xl font-bold mb-4 text-center">
        Find Your Perfect Stay
      </h1>
      <p className="text-xl md:text-2xl mb-8 text-center">
        Discover luxury accommodations worldwide
      </p>
    </div>
  </div>
</template> -->
<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";

import SearchSection from "@/components/SearchSection.vue";

const parallaxOffset = ref(0);

const handleScroll = () => {
  parallaxOffset.value = window.pageYOffset * 0.5;
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
</script>

<template>
  <div class="relative w-full h-[80vh] overflow-hidden">
    <div
      class="absolute inset-0 bg-cover bg-center"
      :style="{
        backgroundImage: `url('https://images.unsplash.com/photo-1582719508461-905c673771fd?q=80&w=1925&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D')`,
        transform: `translateY(${parallaxOffset}px)`,
      }"
      aria-hidden="true"
    ></div>
    <div
      class="absolute inset-0 bg-gradient-to-r from-black to-transparent opacity-70"
    ></div>
    <div
      class="absolute inset-0 flex flex-col justify-center items-center text-white p-4"
    >
      <h1
        class="text-4xl md:text-6xl font-bold mb-4 text-center animate-fade-in-up"
      >
        <span
          class="bg-clip-text text-transparent bg-gradient-to-r from-white to-blue-300"
        >
          Find Your Perfect Stay
        </span>
      </h1>
      <p
        class="text-xl md:text-2xl mb-8 text-center max-w-2xl animate-fade-in-up animation-delay-200"
      >
        Discover luxury accommodations worldwide for unforgettable experiences
      </p>
      <div
        class="bg-white/10 backdrop-blur-md rounded-lg p-4 w-full max-w-6xl animate-fade-in-up animation-delay-400"
      >
        <SearchSection />
      </div>
    </div>
  </div>
</template>

<style scoped>
.animate-fade-in-up {
  animation: fadeInUp 0.6s ease-out forwards;
  opacity: 0;
}

.animation-delay-200 {
  animation-delay: 0.2s;
}

.animation-delay-400 {
  animation-delay: 0.4s;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
