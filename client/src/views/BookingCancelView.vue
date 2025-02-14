<script setup lang="ts">
import { XCircle, ArrowLeft, CreditCard, Calendar } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { useRouter } from "vue-router";

const router = useRouter();

const goBack = () => {
  router.back();
};

const retryPayment = () => {
  console.log("Retrying payment...");
};

const modifyBooking = () => {
  console.log("Modifying booking...");
};
</script>

<template>
  <div
    class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4"
  >
    <Card class="w-full max-w-md">
      <CardHeader class="space-y-1">
        <CardTitle
          class="text-2xl font-bold text-center flex items-center justify-center"
        >
          <XCircle class="mr-2 h-6 w-6 text-yellow-500" />
          Payment Cancelled
        </CardTitle>
      </CardHeader>
      <CardContent class="space-y-4">
        <div class="p-6 bg-yellow-50 rounded-lg border border-yellow-200">
          <p class="text-center text-gray-700">
            You've cancelled your payment process. No worries - your booking is
            still available, and no charges have been made to your account.
          </p>
        </div>
        <div class="space-y-2">
          <h3 class="font-semibold text-gray-700">
            What would you like to do?
          </h3>
          <ul class="list-disc list-inside text-sm text-gray-600 space-y-1">
            <li>Return to your booking and try the payment again</li>
            <li>Modify your booking details if needed</li>
            <li>Take some time to think and come back later</li>
          </ul>
        </div>
      </CardContent>
      <CardFooter class="flex flex-col space-y-2">
        <Button
          @click="retryPayment"
          class="w-full bg-green-500 hover:bg-green-600"
        >
          <CreditCard class="mr-2 h-4 w-4" /> Retry Payment
        </Button>
        <Button
          @click="modifyBooking"
          variant="outline"
          class="w-full border-indigo-200 text-indigo-600 hover:bg-indigo-50"
        >
          <Calendar class="mr-2 h-4 w-4" /> Modify Booking
        </Button>
        <Button
          @click="goBack"
          variant="ghost"
          class="w-full text-gray-600 hover:text-gray-800"
        >
          <ArrowLeft class="mr-2 h-4 w-4" /> Go Back
        </Button>
      </CardFooter>
    </Card>
  </div>
</template>
