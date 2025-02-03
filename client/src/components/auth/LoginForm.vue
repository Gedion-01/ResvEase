<script setup lang="ts">
import { useAuthStore } from "@/store/authStore";
import { ref } from "vue";
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
import { Alert, AlertTitle, AlertDescription } from "@/components/ui/alert";

import type { AxiosError } from "axios";
import { LoaderCircle, TriangleAlert } from "lucide-vue-next";

const authStore = useAuthStore();
const isLoading = ref(false);
const errorMessage = ref("");

const loginSchema = toTypedSchema(
  z.object({
    email: z.string().email({ message: "Invalid email address" }),
    password: z
      .string()
      .min(8, { message: "Password must be at least 8 characters" }),
  })
);

const { handleSubmit, ...loginForm } = useForm({
  validationSchema: loginSchema,
  initialValues: {
    email: "",
    password: "",
  },
});

const onLoginSubmit = handleSubmit(async (values) => {
  console.log("Login data:", values);
  isLoading.value = true;
  errorMessage.value = "";
  const timeout = setTimeout(() => {
    isLoading.value = false;
    console.error("Login request timed out");
  }, 20000);
  try {
    await authStore.login(values);
    clearTimeout(timeout);
  } catch (error: unknown) {
    const axiosError = error as AxiosError;
    if (axiosError.response && axiosError.response.status === 401) {
      errorMessage.value = "Invalid email or password";
    } else {
      errorMessage.value = "An error occurred while logging in";
    }
  } finally {
    isLoading.value = false;
    clearTimeout(timeout);
  }
});
</script>

<template>
  <form @submit.prevent="onLoginSubmit" class="space-y-6">
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
            placeholder="Enter your password"
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
      <span v-else>Login</span>
    </Button>
  </form>
</template>
