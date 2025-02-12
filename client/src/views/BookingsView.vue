<script setup lang="ts">
import { useAuthStore } from "@/store/authStore";
import { useRouter } from "vue-router";
import {
  Calendar,
  CreditCard,
  ChevronRight,
  MapPin,
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
import { Badge } from "@/components/ui/badge";
import { useUserBookings } from "@/services/queries";

type Status = "pending" | "confirmed" | "cancelled";

const authStore = useAuthStore();

const userID = authStore.user?.id;
const { data } = useUserBookings(userID);

const router = useRouter();

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

const viewDetails = (id: string) => {
  router.push(`/booking-details/${id}`);
};

function formatDate(dateString: string) {
  return new Date(dateString).toISOString().split("T")[0];
}
</script>

<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-6">Your Bookings</h1>
    <div class="grid gap-6">
      <Card v-for="booking in data?.data" :key="booking.id" class="w-full">
        <CardHeader>
          <CardTitle class="text-xl">{{ booking.hotel.name }}</CardTitle>
          <CardDescription>{{ booking.room.name }}</CardDescription>
        </CardHeader>
        <CardContent>
          <div class="flex flex-col md:flex-row gap-4">
            <div class="relative w-full md:w-1/3 aspect-video">
              <img
                :src="booking.room.images[0]"
                :alt="booking.room.name"
                layout="fill"
                object-fit="cover"
                class="rounded-lg w-full h-full"
              />
            </div>
            <div class="flex-1">
              <div class="flex justify-between items-center mb-4">
                <div class="flex items-center">
                  <Calendar class="mr-2 h-5 w-5" />
                  <span>
                    {{ formatDate(booking.fromDate) }} -
                    {{ formatDate(booking.fromDate) }}
                  </span>
                </div>
                <Badge :variant="getStatusVariant(booking.status)">
                  {{ getStatusText(booking.status) }}
                </Badge>
              </div>
              <div class="flex items-center mb-2">
                <MapPin class="mr-2 h-5 w-5" />
                <span>{{ booking.hotel.location }}</span>
              </div>
              <div class="flex items-center mb-2">
                <Star class="mr-2 h-5 w-5 text-yellow-400" />
                <span>{{ booking.hotel.rating.toFixed(1) }}</span>
              </div>
              <div class="flex items-center mb-2">
                <Clock class="mr-2 h-5 w-5" />
                <span
                  >Booked on:
                  {{ new Date(booking.updatedAt).toLocaleString() }}</span
                >
              </div>

              <div class="flex justify-between items-center">
                <div class="flex items-center">
                  <CreditCard class="mr-2 h-5 w-5" />
                  <span>Total: ${{ booking.totalPrice / 100 }}</span>
                </div>
              </div>
            </div>
          </div>
        </CardContent>
        <CardFooter class="flex justify-end">
          <Button variant="ghost" @click="viewDetails(booking.id)">
            View Details <ChevronRight class="ml-2 h-4 w-4" />
          </Button>
        </CardFooter>
      </Card>
    </div>
  </div>
</template>
