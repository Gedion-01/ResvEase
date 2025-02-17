<script setup lang="ts">
import { useAuthStore } from "@/store/authStore";
import { ref, watch } from "vue";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from "@/components/ui/form";
import confetti from "canvas-confetti";
import { CheckCircle } from "lucide-vue-next";
import { Alert, AlertTitle, AlertDescription } from "@/components/ui/alert";
import {
  Dialog,
  DialogContent,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { LoaderCircle, TriangleAlert } from "lucide-vue-next";
import type { AxiosError } from "axios";
import { useRouter } from "vue-router";

const router = useRouter();

const authStore = useAuthStore();
const isLoading = ref(false);
const errorMessage = ref("");
const showDialog = ref(false);
const showConfetti = ref(false);

const signupSchema = toTypedSchema(
  z
    .object({
      firstName: z
        .string()
        .min(2, { message: "FirstName must be at least 2 characters" }),
      lastName: z
        .string()
        .min(2, { message: "LastName must be at least 2 characters" }),
      email: z.string().email({ message: "Invalid email address" }),
      password: z
        .string()
        .min(8, { message: "Password must be at least 8 characters" }),
      confirmPassword: z.string(),
    })
    .refine((data) => data.password === data.confirmPassword, {
      message: "Passwords don't match",
      path: ["confirmPassword"],
    })
);

const { handleSubmit, ...signupForm } = useForm({
  validationSchema: signupSchema,
  initialValues: {
    firstName: "",
    lastName: "",
    email: "",
    password: "",
    confirmPassword: "",
  },
});

const onSignupSubmit = handleSubmit(async (values) => {
  isLoading.value = true;
  errorMessage.value = "";
  const timeout = setTimeout(() => {
    isLoading.value = false;
    console.error("Signup request timed out");
  }, 20000);
  try {
    await authStore.signup(values);
    showDialog.value = true;
  } catch (error: unknown) {
    console.error("Signup error:", error);
    const axiosError = error as AxiosError;
    if (axiosError.response && axiosError.response.status === 409) {
      errorMessage.value = "Email already in use";
    } else {
      errorMessage.value = "An error occurred while signing up";
    }
  } finally {
    isLoading.value = false;
    clearTimeout(timeout);
  }
});

const handleDialogClose = () => {
  showDialog.value = false;
  router.push("/auth?tab=login");
};

watch(showDialog, (newValue) => {
  if (newValue) {
    showConfetti.value = true;
    setTimeout(() => {
      showConfetti.value = false;
    }, 3000);

    confetti({
      particleCount: 100,
      spread: 70,
      origin: { y: 0.6 },
    });
  }
});
</script>

<template>
  <form @submit.prevent="onSignupSubmit" class="space-y-6">
    <FormField v-slot="{ componentField }" name="firstName">
      <FormItem>
        <FormLabel>FirstName</FormLabel>
        <FormControl>
          <Input placeholder="Enter your first name" v-bind="componentField" />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>
    <FormField v-slot="{ componentField }" name="lastName">
      <FormItem>
        <FormLabel>LastName</FormLabel>
        <FormControl>
          <Input placeholder="Enter your last name" v-bind="componentField" />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>
    <FormField v-slot="{ componentField }" name="email">
      <FormItem>
        <FormLabel>Email</FormLabel>
        <FormControl>
          <Input
            type="email"
            placeholder="Enter your email"
            v-bind="componentField"
          />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>
    <FormField v-slot="{ componentField }" name="password">
      <FormItem>
        <FormLabel>Password</FormLabel>
        <FormControl>
          <Input
            type="password"
            placeholder="Create a password"
            v-bind="componentField"
          />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>
    <FormField v-slot="{ componentField }" name="confirmPassword">
      <FormItem>
        <FormLabel>Confirm Password</FormLabel>
        <FormControl>
          <Input
            type="password"
            placeholder="Confirm your password"
            v-bind="componentField"
          />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>
    <div v-if="errorMessage" class="mt-4">
      <Alert variant="destructive">
        <TriangleAlert class="w-4 h-4 mr-2" />
        <AlertTitle>Error</AlertTitle>
        <AlertDescription>{{ errorMessage }}</AlertDescription>
      </Alert>
    </div>
    <Button type="submit" class="w-full" :disabled="isLoading">
      <LoaderCircle class="h-4 w-4 mr-2 animate-spin" v-if="isLoading" />
      <span v-else>Sign Up</span>
    </Button>
  </form>

  <div>
    <Dialog :open="showDialog" @update:open="showDialog = $event">
      <DialogContent class="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle class="text-2xl font-bold text-center"
            >Welcome Aboard! 🎉</DialogTitle
          >
        </DialogHeader>
        <div class="flex flex-col items-center space-y-4 py-6">
          <Transition name="scale" appear>
            <CheckCircle class="w-16 h-16 text-green-500" />
          </Transition>
          <DialogDescription class="text-center text-lg">
            Your account has been created successfully. You're all set to
            explore!
          </DialogDescription>
          <Transition name="fade-slide-up">
            <div
              v-if="showConfetti"
              class="text-xl font-semibold text-purple-600"
            >
              Hooray! 🎊
            </div>
          </Transition>
        </div>
        <div class="flex justify-center">
          <Button
            @click="handleDialogClose"
            class="px-8 py-2 text-lg font-semibold transition-all duration-300 transform hover:scale-105"
          >
            Start Your Journey
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  </div>
</template>
<style scoped>
.open-dialog-btn {
  @apply bg-gradient-to-r from-purple-400 to-pink-500 text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:shadow-lg transition-all duration-300 hover:scale-105;
}

.scale-enter-active,
.scale-leave-active {
  transition: all 0.3s ease;
}

.scale-enter-from,
.scale-leave-to {
  opacity: 0;
  transform: scale(0.9);
}

.fade-slide-up-enter-active,
.fade-slide-up-leave-active {
  transition: all 0.3s ease;
}

.fade-slide-up-enter-from,
.fade-slide-up-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
</style>
