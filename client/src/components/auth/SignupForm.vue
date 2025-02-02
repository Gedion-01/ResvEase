<!-- filepath: /home/gedion/Documents/Projects/Hotel-Reservation-App/client/src/components/SignupForm.vue -->
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
  console.log("Signup data:", values);
  //   isLoading.value = true;
  //   const timeout = setTimeout(() => {
  //     isLoading.value = false;
  //     console.error("Signup request timed out");
  //   }, 20000);
  //   try {
  //     await authStore.signup(values);
  //     clearTimeout(timeout);
  //     isLoading.value = false;
  //   } catch (error) {
  //     clearTimeout(timeout);
  //     isLoading.value = false;
  //     console.error("Signup failed:", error);
  //   }
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
    <Button type="submit" class="w-full" :disabled="isLoading">
      <span v-if="isLoading">Loading...</span>
      <span v-else>Sign Up</span>
    </Button>
  </form>
</template>
