<script setup lang="ts"></script>
<template>
  <div className="relative w-full h-[60vh] mb-8">
    <image
      src="/hero-image.jpg"
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
</template>
