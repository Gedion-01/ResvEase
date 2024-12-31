<script setup lang="ts">
import { ref, defineProps } from "vue";
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

interface RoomSelectionProps {
  hotelId: string;
  onRoomSelect: (roomId: string) => void;
}

defineProps<{
  hotelId: string;
  onRoomSelect: (roomId: string) => void;
}>();

const selectedRoom = ref("");
const currentImageIndex = ref<{ [key: string]: number }>({});

// This would typically come from an API based on the hotelId
const roomTypes: RoomType[] = [
  {
    id: "1",
    name: "Deluxe Room",
    description: "Spacious room with city view",
    price: 200,
    capacity: 2,
    amenities: ["Free Wi-Fi", "Breakfast Included", "Room Service"],
    available: 5,
    images: ["/room1-1.jpg", "/room1-2.jpg", "/room1-3.jpg"].map(
      (img) => `/placeholder.svg?text=${img}`
    ),
  },
  {
    id: "2",
    name: "Executive Suite",
    description: "Luxury suite with separate living area",
    price: 350,
    capacity: 3,
    amenities: ["Free Wi-Fi", "Breakfast Included", "Room Service", "Mini Bar"],
    available: 3,
    images: ["/room2-1.jpg", "/room2-2.jpg", "/room2-3.jpg"].map(
      (img) => `/placeholder.svg?text=${img}`
    ),
  },
  {
    id: "3",
    name: "Family Room",
    description: "Perfect for families with children",
    price: 400,
    capacity: 4,
    amenities: [
      "Free Wi-Fi",
      "Breakfast Included",
      "Room Service",
      "Kids Area",
    ],
    available: 2,
    images: ["/room3-1.jpg", "/room3-2.jpg", "/room3-3.jpg"].map(
      (img) => `/placeholder.svg?text=${img}`
    ),
  },
];

const handleImageNavigation = (roomId: string, direction: "prev" | "next") => {
  const room = roomTypes.find((r) => r.id === roomId);
  if (!room) return;

  if (!room || !room.images || room.images.length === 0) return;

  const currentIndex = currentImageIndex.value[roomId] || 0;
  const newIndex =
    (currentIndex + (direction === "next" ? 1 : -1) + room.images.length) %
    room.images.length;

  currentImageIndex.value = {
    ...currentImageIndex.value,
    [roomId]: newIndex,
  };
};

type Amenity = "Free Wi-Fi" | "Breakfast Included" | "Room Service" | string;

interface Room {
  amenities: Amenity[];
}

const room: Room = {
  amenities: [
    "Free Wi-Fi",
    "Breakfast Included",
    "Room Service",
    "Other Amenity",
  ],
};

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
      v-for="room in roomTypes"
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
          <Badge :variant="room.available < 3 ? 'destructive' : 'secondary'">
            {{ room.available }} rooms left
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
                      currentImageIndex[room.id] || 1
                    }`"
                    fill
                    class="object-cover rounded-lg"
                  />
                  <div
                    class="absolute inset-0 flex items-center justify-between p-2"
                  >
                    <Button
                      variant="secondary"
                      size="icon"
                      class="opacity-80 hover:opacity-100"
                      @click="
                        (e) => {
                          e.stopPropagation();
                          handleImageNavigation(room.id, 'prev');
                        }
                      "
                    >
                      <ChevronLeft class="h-4 w-4" />
                    </Button>
                    <Button
                      variant="secondary"
                      size="icon"
                      class="opacity-80 hover:opacity-100"
                      @click="
                        (e) => {
                          e.stopPropagation();
                          handleImageNavigation(room.id, 'next');
                        }
                      "
                    >
                      <ChevronRight class="h-4 w-4" />
                    </Button>
                  </div>
                </div>
              </DialogTrigger>
              <DialogContent class="max-w-4xl">
                <div class="grid gap-4">
                  <div class="relative aspect-video">
                    <Image
                      :src="room.images[currentImageIndex[room.id] || 0]"
                      :alt="`${room.name} view ${
                        currentImageIndex[room.id] || 1
                      }`"
                      fill
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
                      @click="
                        () =>
                          (currentImageIndex = {
                            ...currentImageIndex,
                            [room.id]: index,
                          })
                      "
                    >
                      <Image
                        :src="roomImg"
                        :alt="`${room.name} view ${index + 1}`"
                        fill
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
                    onRoomSelect(room.id);
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
