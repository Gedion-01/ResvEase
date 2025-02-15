<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { CheckCircle, Calendar, Clipboard, Home } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { useToast } from "@/components/ui/toast/use-toast";

const route = useRoute();
const router = useRouter();
const { toast } = useToast();

const sessionId = ref("");

const truncatedSessionId = computed(() => {
  if (sessionId.value.length > 20) {
    return sessionId.value.slice(0, 25) + "..." + sessionId.value.slice(-10);
  }
  return sessionId.value;
});

onMounted(() => {
  sessionId.value = (route.query.session_id as string) || "";
  if (sessionId.value) {
    triggerConfetti();
  }
});

const triggerConfetti = () => {
  import("canvas-confetti").then((confetti) => {
    confetti.default({
      particleCount: 100,
      spread: 70,
      origin: { y: 0.6 },
    });
  });
};

const copySessionId = () => {
  navigator.clipboard.writeText(sessionId.value);
  toast({
    title: "Copied!",
    description: "Session ID copied to clipboard",
  });
};

const viewBookingDetails = () => {
  router.push("/bookings");
};

const goToHomepage = () => {
  router.push("/");
};
</script>

<template>
  <div
    class="min-h-screen bg-gradient-to-br from-green-50 to-teal-100 flex items-center justify-center p-4"
  >
    <Card class="w-full max-w-lg">
      <CardHeader class="space-y-1">
        <CardTitle
          class="text-2xl font-bold text-center flex items-center justify-center text-green-600"
        >
          <CheckCircle class="mr-2 h-8 w-8" />
          Payment Successful!
        </CardTitle>
      </CardHeader>
      <CardContent class="space-y-4">
        <div class="p-6 bg-green-100 rounded-lg border border-green-200">
          <p class="text-center text-gray-700">
            Great news! Your payment has been processed successfully. Your
            booking is now confirmed.
          </p>
        </div>
        <div v-if="sessionId" class="space-y-2">
          <div
            class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
          >
            <span class="text-sm font-medium text-gray-600">Session ID:</span>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger
                  class="text-sm text-gray-800 truncate max-w-[150px]"
                >
                  {{ truncatedSessionId }}
                </TooltipTrigger>
              </Tooltip>
            </TooltipProvider>
            <Button variant="ghost" size="icon" @click="copySessionId">
              <Clipboard class="h-4 w-4" />
            </Button>
          </div>
        </div>
      </CardContent>
      <CardFooter class="flex flex-col space-y-2">
        <Button
          @click="viewBookingDetails"
          class="w-full bg-teal-500 hover:bg-teal-600"
        >
          <Calendar class="mr-2 h-4 w-4" /> View Booking Details
        </Button>
        <Button @click="goToHomepage" variant="outline" class="w-full">
          <Home class="mr-2 h-4 w-4" /> Return to Homepage
        </Button>
      </CardFooter>
    </Card>
  </div>
</template>
