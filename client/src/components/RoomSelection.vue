<script setup lang="ts">
import { ref, defineProps, reactive } from "vue";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import {
  Bed,
  Users,
  Wifi,
  Coffee,
  Utensils,
  ChevronLeft,
  ChevronRight,
} from "lucide-vue-next";
import { useRooms } from "@/services/queries";
import { useRoute } from "vue-router";

interface RoomType {
  id: string;
  name: string;
  description: string;
  price: number;
  capacity: number;
  amenities: string[];
  available: number;
  images: string[];
}

const props = defineProps<{
  hotelId: string;
  onRoomSelect: (roomId: string) => void;
}>();

const route = useRoute();
const queryParams = {
  location: route.query.location as string,
  checkIn: route.query.checkIn as string,
  checkOut: route.query.checkOut as string,
  adults: route.query.adults as string,
  children: route.query.children as string,
  roomCapacity: (
    parseInt(route.query.adults as string, 10) +
    parseInt(route.query.children as string, 10)
  ).toString(),
};

const { data, isLoading } = useRooms(props.hotelId, queryParams);

const selectedRoom = ref<string>("");
const currentImageIndex = ref<{ [key: string]: number }>({});

const handleImageNavigation = (roomId: string, direction: "prev" | "next") => {
  const currentIndex = currentImageIndex.value[roomId] || 0;
  const room = data.value?.data.find((room) => room.id === roomId);
  if (!room) return;

  const newIndex =
    (currentIndex + (direction === "next" ? 1 : -1) + room.images.length) %
    room.images.length;

  currentImageIndex.value[roomId] = newIndex;
};

type Amenity = "Free Wi-Fi" | "Breakfast Included" | "Room Service" | string;

interface Room {
  amenities: Amenity[];
}

const getAmenityIcon = (amenity: Amenity) => {
  switch (amenity) {
    case "Free Wi-Fi":
      return Wifi;
    case "Breakfast Included":
      return Coffee;
    case "Room Service":
      return Utensils;
    default:
      return null;
  }
};
</script>

<template>
  <div class="space-y-6">
    <h2 class="text-2xl font-bold">Select Your Room</h2>
    <Card
      v-for="room in data?.data"
      :key="room.id"
      :class="[
        'cursor-pointer transition-colors',
        selectedRoom === room.id ? 'border-primary' : '',
      ]"
    >
      <CardHeader>
        <div class="flex justify-between items-start">
          <div>
            <CardTitle>{{ room.name }}</CardTitle>
            <p class="text-sm text-muted-foreground mt-1">
              {{ room.description }}
            </p>
          </div>
          <Badge
            :variant="room.availableCount < 3 ? 'destructive' : 'secondary'"
          >
            {{ room.availableCount }} rooms left
          </Badge>
        </div>
      </CardHeader>
      <CardContent>
        <div class="grid md:grid-cols-2 gap-6">
          <div class="relative aspect-video">
            <Dialog>
              <DialogTrigger asChild>
                <div class="relative w-full h-full cursor-pointer">
                  <img
                    :src="room.images[currentImageIndex[room.id] || 0]"
                    :alt="`${room.name} view ${
                      currentImageIndex[room.id] + 1 || 1
                    }`"
                    class="object-cover rounded-lg"
                  />
                  <div
                    class="absolute inset-0 flex items-center justify-between p-2"
                  >
                    <Button
                      variant="secondary"
                      size="icon"
                      class="opacity-80 hover:opacity-100"
                      @click.stop="handleImageNavigation(room.id, 'prev')"
                    >
                      <ChevronLeft class="h-4 w-4" />
                    </Button>
                    <Button
                      variant="secondary"
                      size="icon"
                      class="opacity-80 hover:opacity-100"
                      @click.stop="handleImageNavigation(room.id, 'next')"
                    >
                      <ChevronRight class="h-4 w-4" />
                    </Button>
                  </div>
                </div>
              </DialogTrigger>
              <DialogContent class="max-w-4xl">
                <div class="grid gap-4">
                  <div class="relative aspect-video">
                    <img
                      :src="room.images[currentImageIndex[room.id] || 0]"
                      :alt="`${room.name} view ${
                        currentImageIndex[room.id] + 1 || 1
                      }`"
                      class="object-contain"
                    />
                  </div>
                  <div class="grid grid-cols-3 gap-2">
                    <div
                      v-for="(roomImg, index) in room.images"
                      :key="roomImg"
                      :class="[
                        'relative aspect-video cursor-pointer',
                        index === (currentImageIndex[room.id] || 0)
                          ? 'ring-2 ring-primary'
                          : '',
                      ]"
                      @click="currentImageIndex[room.id] = index"
                    >
                      <img
                        :src="roomImg"
                        :alt="`${room.name} view ${index + 1}`"
                        class="object-cover rounded"
                      />
                    </div>
                  </div>
                </div>
              </DialogContent>
            </Dialog>
          </div>
          <div class="space-y-4">
            <div class="flex items-center gap-2">
              <Bed class="h-5 w-5" />
              <span>Sleeps {{ room.capacity }}</span>
              <Users class="h-5 w-5 ml-4" />
              <span>Max {{ room.capacity }} guests</span>
            </div>
            <div class="flex flex-wrap gap-2">
              <Badge
                v-for="amenity in room.amenities"
                :key="amenity"
                variant="outline"
                class="flex items-center gap-1"
              >
                <component
                  v-if="getAmenityIcon(amenity)"
                  :is="getAmenityIcon(amenity)"
                  class="h-4 w-4"
                />
                {{ amenity }}
              </Badge>
            </div>
            <div class="flex justify-between items-center mt-4">
              <div>
                <span class="text-2xl font-bold">{{ room.price }}</span>
                <span class="text-muted-foreground">/night</span>
              </div>
              <Button
                @click="
                  () => {
                    selectedRoom = room.id;
                    props.onRoomSelect(room.id);
                  }
                "
                :variant="selectedRoom === room.id ? 'default' : 'outline'"
              >
                Select Room
              </Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  </div>
</template>
