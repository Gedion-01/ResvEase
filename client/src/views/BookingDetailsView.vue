<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import {
  Calendar,
  MapPin,
  Users,
  Bed,
  Coffee,
  Star,
  Clock,
} from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";
import { Separator } from "@/components/ui/separator";
import { Badge } from "@/components/ui/badge";
import { useBooking } from "@/services/queries";

const route = useRoute();
const router = useRouter();

const id = route.params.id as string;

const { data } = useBooking(id);

type Status = "pending" | "confirmed" | "cancelled";

const getStatusVariant = (status: Status) => {
  switch (status) {
    case "pending":
      return "default";
    case "confirmed":
      return "outline";
    case "cancelled":
      return "destructive";
  }
};

const getStatusText = (status: Status) => {
  switch (status) {
    case "pending":
      return "Upcoming";
    case "confirmed":
      return "Completed";
    case "cancelled":
      return "Cancelled";
  }
};

const goToBookings = () => {
  router.push("/bookings");
};

const goToHome = () => {
  router.push("/");
};
function formatDate(dateString: string) {
  return new Date(dateString).toISOString().split("T")[0];
}
</script>

<template>
  <div v-if="data" class="container mx-auto px-4 py-8">
    <Card class="max-w-3xl mx-auto">
      <CardHeader>
        <CardTitle class="text-2xl">Booking Details</CardTitle>
        <CardDescription>Booking ID: {{ data.id }}</CardDescription>
      </CardHeader>
      <CardContent class="space-y-4">
        <div class="relative">
          <Carousel class="relative cursor-pointer">
            <CarouselContent>
              <CarouselItem
                class="relative aspect-video"
                v-for="(image, index) in data.room.images"
                :key="index"
              >
                <DialogTrigger asChild>
                  <img
                    :src="image"
                    class="object-cover w-full h-full rounded-t-lg"
                    :alt="`${data.room.name} view ${index + 1}`"
                  />
                </DialogTrigger>
              </CarouselItem>
            </CarouselContent>
            <CarouselPrevious
              variant="secondary"
              size="icon"
              class="absolute left-2 top-1/2 -translate-y-1/2 h-8 w-8 rounded-lg"
              @click.stop
            />
            <CarouselNext
              variant="secondary"
              size="icon"
              class="absolute right-2 top-1/2 -translate-y-1/2 h-8 w-8 rounded-lg"
              @click.stop
            />
          </Carousel>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div class="space-y-2">
            <div class="flex items-center">
              <MapPin class="mr-2 h-4 w-4" />
              <span class="font-semibold">Hotel:</span>
            </div>
            <p>{{ data.hotel.name }}</p>
          </div>
          <div class="space-y-2">
            <div class="flex items-center">
              <Bed class="mr-2 h-4 w-4" />
              <span class="font-semibold">Room Type:</span>
            </div>
            <p>{{ data.room.name }}</p>
          </div>
          <div class="space-y-2">
            <div class="flex items-center">
              <Calendar class="mr-2 h-4 w-4" />
              <span class="font-semibold">Check-in:</span>
            </div>
            <p>{{ formatDate(data.fromDate) }}</p>
          </div>
          <div class="space-y-2">
            <div class="flex items-center">
              <Calendar class="mr-2 h-4 w-4" />
              <span class="font-semibold">Check-out:</span>
            </div>
            <p>{{ formatDate(data.tillDate) }}</p>
          </div>
          <div class="space-y-2">
            <div class="flex items-center">
              <Users class="mr-2 h-4 w-4" />
              <span class="font-semibold">Guests:</span>
            </div>
            <p>{{ data.numPersons }}</p>
          </div>
          <div class="space-y-2">
            <div class="flex items-center">
              <Star class="mr-2 h-4 w-4 text-yellow-400" />
              <span class="font-semibold">Hotel Rating:</span>
            </div>
            <p>{{ data.hotel.rating.toFixed(1) }} / 5</p>
          </div>
          <div class="space-y-2">
            <div class="flex items-center">
              <Clock class="mr-2 h-4 w-4" />
              <span class="font-semibold">Booked At:</span>
            </div>
            <p>{{ new Date(data.createdAt).toLocaleString() }}</p>
          </div>
        </div>
        <Separator />
        <div class="space-y-2">
          <div class="flex items-center">
            <Coffee class="mr-2 h-4 w-4" />
            <span class="font-semibold">Amenities:</span>
          </div>
          <ul class="list-disc list-inside">
            <li v-for="(amenity, index) in data.room.amenities" :key="index">
              {{ amenity }}
            </li>
          </ul>
        </div>
        <Separator />
        <div class="flex justify-between items-center text-lg font-semibold">
          <span>Total Price:</span>
          <span>${{ data.totalPrice / 100 }}</span>
        </div>
        <div class="flex justify-between items-center">
          <span class="font-semibold">Status:</span>
          <Badge :variant="getStatusVariant(data.status)">
            {{ getStatusText(data.status) }}
          </Badge>
        </div>
      </CardContent>
      <CardFooter class="flex justify-center space-x-4">
        <Button @click="goToBookings">Back to Bookings</Button>
        <Button variant="outline" @click="goToHome">Home</Button>
      </CardFooter>
    </Card>
  </div>
</template>
