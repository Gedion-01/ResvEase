<script setup lang="ts">
import { computed } from "vue";
import { useRouter } from "vue-router";
import { AlertCircle, RefreshCcw, Home, LifeBuoy } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

const props = defineProps<{
  errorType: "network" | "server" | "notFound" | "generic";
  errorMessage?: string;
}>();

const router = useRouter();

const errorDetails = computed(() => {
  switch (props.errorType) {
    case "network":
      return {
        title: "Network Error",
        description:
          "Oops! It seems like there's a problem with your internet connection.",
        icon: RefreshCcw,
        action: { label: "Try Again", handler: () => window.location.reload() },
      };
    case "server":
      return {
        title: "Server Error",
        description:
          "We're experiencing some technical difficulties. Our team has been notified.",
        icon: AlertCircle,
        action: { label: "Go to Homepage", handler: () => router.push("/") },
      };
    case "notFound":
      return {
        title: "Page Not Found",
        description:
          "The page you're looking for doesn't exist or has been moved.",
        icon: Home,
        action: { label: "Go to Homepage", handler: () => router.push("/") },
      };
    default:
      return {
        title: "Unexpected Error",
        description:
          "Something went wrong. Please try again or contact support.",
        icon: AlertCircle,
        action: { label: "Go to Homepage", handler: () => router.push("/") },
      };
  }
});

const contactSupport = () => {
  // Implement logic to contact support or show contact information
  router.push("/contact");
};
</script>

<template>
  <div class="min-h-[50vh] flex items-center justify-center p-4">
    <Card class="w-full max-w-2xl">
      <CardHeader>
        <CardTitle
          class="text-2xl font-bold text-center flex items-center justify-center text-destructive"
        >
          <component :is="errorDetails.icon" class="mr-2 h-8 w-8" />
          {{ errorDetails.title }}
        </CardTitle>
      </CardHeader>
      <CardContent class="text-center space-y-4">
        <Alert variant="destructive">
          <AlertCircle class="h-4 w-4" />
          <AlertTitle>Error</AlertTitle>
          <AlertDescription>
            {{ props.errorMessage || errorDetails.description }}
          </AlertDescription>
        </Alert>
      </CardContent>
      <CardFooter
        class="flex flex-col sm:flex-row justify-center space-y-2 sm:space-y-0 sm:space-x-2"
      >
        <Button @click="errorDetails.action.handler" class="w-full sm:w-auto">
          {{ errorDetails.action.label }}
        </Button>
        <Button
          @click="contactSupport"
          variant="outline"
          class="w-full sm:w-auto"
        >
          Contact Support
        </Button>
      </CardFooter>
    </Card>
  </div>
</template>

<style scoped>
@keyframes pulse {
  0%,
  100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>
