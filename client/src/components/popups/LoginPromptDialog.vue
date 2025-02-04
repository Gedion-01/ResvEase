<script setup lang="ts">
import { ref } from "vue";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { LockIcon } from "lucide-vue-next";
import { useAuthStore } from "@/store/authStore";
import { useRouter, useRoute } from "vue-router";

const router = useRouter();
const route = useRoute();

const authStore = useAuthStore();

const isOpen = ref(false);

const openDialog = () => {
  isOpen.value = true;
};

const closeDialog = () => {
  isOpen.value = false;
};

const goToLogin = () => {
  authStore.setRedirectUrl(route.fullPath);
  router.push("/auth?tab=login");
  closeDialog();
};

defineExpose({ openDialog });
</script>

<template>
  <Dialog :open="isOpen" @update:open="closeDialog">
    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle class="text-2xl font-bold text-center"
          >Hold Up! 🛑</DialogTitle
        >
        <DialogDescription class="text-center text-lg mt-2">
          You need to be logged in to reserve a room.
        </DialogDescription>
      </DialogHeader>
      <div class="flex flex-col items-center space-y-4 py-6">
        <Transition name="bounce" appear>
          <LockIcon class="w-16 h-16 text-yellow-500" />
        </Transition>
        <p class="text-center text-gray-600">
          Don't miss out on your perfect stay! Log in to unlock exclusive deals
          and seamless booking.
        </p>
      </div>
      <div class="flex flex-col space-y-2">
        <Button @click="goToLogin" class="w-full text-lg font-semibold">
          Go to Login
        </Button>
        <Button
          @click="closeDialog"
          variant="outline"
          class="w-full text-lg font-semibold"
        >
          Maybe Later
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
