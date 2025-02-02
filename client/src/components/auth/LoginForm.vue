<!-- filepath: /home/gedion/Documents/Projects/Hotel-Reservation-App/client/src/components/LoginForm.vue -->
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

const authStore = useAuthStore();
const isLoading = ref(false);

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
  const timeout = setTimeout(() => {
    isLoading.value = false;
    console.error("Login request timed out");
  }, 20000);
  try {
    await authStore.login(values);
    clearTimeout(timeout);
    isLoading.value = false;
  } catch (error) {
    clearTimeout(timeout);
    isLoading.value = false;
    console.error("Login failed:", error);
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
    <Button type="submit" class="w-full" :disabled="isLoading">
      <span v-if="isLoading">Loading...</span>
      <span v-else>Login</span>
    </Button>
  </form>
</template>
