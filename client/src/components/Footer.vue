<script setup lang="ts">
import { Facebook, Twitter, Instagram, Youtube } from "lucide-vue-next";
const currentYear = new Date().getFullYear();
</script>

<template>
  <footer class="bg-gray-900 text-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
        <div>
          <h3 class="text-lg font-semibold mb-4">About Us</h3>
          <ul class="space-y-2">
            <li><a href="#" class="hover:text-gray-300">Our Story</a></li>
            <li><a href="#" class="hover:text-gray-300">Team</a></li>
            <li><a href="#" class="hover:text-gray-300">Careers</a></li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">Support</h3>
          <ul class="space-y-2">
            <li><a href="#" class="hover:text-gray-300">FAQs</a></li>
            <li><a href="#" class="hover:text-gray-300">Contact Us</a></li>
            <li><a href="#" class="hover:text-gray-300">Privacy Policy</a></li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">Services</h3>
          <ul class="space-y-2">
            <li><a href="#" class="hover:text-gray-300">Hotel Booking</a></li>
            <li><a href="#" class="hover:text-gray-300">Flight Booking</a></li>
            <li><a href="#" class="hover:text-gray-300">Car Rental</a></li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">Follow Us</h3>
          <div class="flex space-x-4">
            <a href="#" class="hover:text-gray-300"><Facebook /></a>
            <a href="#" class="hover:text-gray-300"><Twitter /></a>
            <a href="#" class="hover:text-gray-300"><Instagram /></a>
            <a href="#" class="hover:text-gray-300"><Youtube /></a>
          </div>
        </div>
      </div>
      <div class="mt-8 pt-8 border-t border-gray-700 text-center">
        <p>&copy; {{ currentYear }} BookMyStay. All rights reserved.</p>
      </div>
    </div>
  </footer>
</template>
