<script setup lang="ts">
import { ref, computed } from "vue";
import { useRouter } from "vue-router";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { CalendarX, Home, AlertTriangle, AlertCircle } from "lucide-vue-next";

const props = defineProps<{
  errorType:
    | "no available rooms for the given date range in this hotel"
    | "room is booked, please choose another room"
    | "failed to create booking"
    | string;
}>();

const isOpen = ref(false);
const router = useRouter();

const openDialog = () => {
  isOpen.value = true;
};

const closeDialog = () => {
  isOpen.value = false;
};

const title = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return "No Available Rooms";
    case "room is booked, please choose another room":
      return "Room Already Booked";
    case "failed to create booking":
      return "Booking Failed";
    default:
      return "An Error Occurred";
  }
});

const description = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return "We couldn't find any available rooms for your selected dates.";
    case "room is booked, please choose another room":
      return "The room you selected has just been booked by another guest.";
    case "failed to create booking":
      return "We encountered an unexpected error while processing your booking.";
    default:
      return "Something went wrong. Please try again later.";
  }
});

const message = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return "Try adjusting your dates or exploring our other amazing properties nearby.";
    case "room is booked, please choose another room":
      return "Don't worry! We have other fantastic rooms available that might suit your needs.";
    case "failed to create booking":
      return "Our team has been notified and is working to resolve the issue. Please try again later.";
    default:
      return "An unexpected error occurred. Please try again later.";
  }
});

const icon = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return CalendarX;
    case "room is booked, please choose another room":
      return Home;
    case "failed to create booking":
      return AlertTriangle;
    default:
      return AlertCircle;
  }
});

const iconClass = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return "text-orange-500";
    case "room is booked, please choose another room":
      return "text-blue-500";
    case "failed to create booking":
      return "text-red-500";
    default:
      return "text-gray-500";
  }
});

const dialogClass = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return "sm:max-w-[425px] border-orange-200";
    case "room is booked, please choose another room":
      return "sm:max-w-[425px] border-blue-200";
    case "failed to create booking":
      return "sm:max-w-[425px] border-red-200";
    default:
      return "sm:max-w-[425px] border-gray-200";
  }
});

const buttonClass = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return "bg-orange-500 hover:bg-orange-600";
    case "room is booked, please choose another room":
      return "bg-blue-500 hover:bg-blue-600";
    case "failed to create booking":
      return "bg-red-500 hover:bg-red-600";
    default:
      return "bg-gray-500 hover:bg-gray-600";
  }
});

const primaryActionText = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return "Modify Search";
    case "room is booked, please choose another room":
      return "View Other Rooms";
    case "failed to create booking":
      return "Try Again";
    default:
      return "OK";
  }
});

const secondaryActionText = computed(() => {
  switch (props.errorType) {
    case "no available rooms for the given date range in this hotel":
      return "Explore Nearby";
    case "room is booked, please choose another room":
      return "Contact Us";
    case "failed to create booking":
      return null;
    default:
      return null;
  }
});

const primaryAction = () => {
  switch (props.errorType) {
    case "room is booked, please choose another room":
      router.go(-1);
      break;
    default:
      closeDialog();
  }

  closeDialog();
};

const secondaryAction = () => {
  closeDialog();
};

defineExpose({ openDialog });
</script>
<template>
  <Dialog :open="isOpen" @update:open="closeDialog">
    <DialogContent :class="dialogClass">
      <DialogHeader>
        <DialogTitle class="text-2xl font-bold text-center">{{
          title
        }}</DialogTitle>
        <DialogDescription class="text-center text-lg mt-2">
          {{ description }}
        </DialogDescription>
      </DialogHeader>
      <div class="flex flex-col items-center space-y-4 py-6">
        <Transition name="bounce" appear>
          <component :is="icon" class="w-16 h-16" :class="iconClass" />
        </Transition>
        <p class="text-center text-gray-600">
          {{ message }}
        </p>
      </div>
      <div class="flex justify-center space-x-2">
        <Button @click="primaryAction" :class="buttonClass">
          {{ primaryActionText }}
        </Button>
        <Button
          v-if="secondaryActionText"
          @click="secondaryAction"
          variant="outline"
        >
          {{ secondaryActionText }}
        </Button>
      </div>
    </DialogContent>
  </Dialog>
</template>

<style scoped>
.bounce-enter-active {
  animation: bounce-in 0.5s;
}
.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.25);
  }
  100% {
    transform: scale(1);
  }
}
</style>
