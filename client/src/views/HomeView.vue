<script setup lang="ts">
import { useRouter } from "vue-router";
import HeroSection from "@/components/HeroSection.vue";
import HotelList from "@/components/HotelList.vue";
import FeaturedDestinations from "@/components/FeaturedDestinations.vue";
import Footer from "@/components/Footer.vue";
import { Button } from "@/components/ui/button";

const router = useRouter();

const navigateToSearch = () => {
  router.push("/search");
};
</script>

<template>
  <div class="min-h-screen flex flex-col bg-slate-100">
    <HeroSection />

    <main class="flex-grow">
      <section class="py-12">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 class="text-3xl font-bold text-center mb-8">Featured Hotels</h2>
          <HotelList :limit="4" />
          <div class="text-center mt-8">
            <Button @click="navigateToSearch" size="lg"
              >View More Hotels</Button
            >
          </div>
        </div>
      </section>

      <FeaturedDestinations />

      <SpecialOffers />
    </main>

    <Footer />
  </div>
</template>
