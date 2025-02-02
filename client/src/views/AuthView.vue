<script setup lang="ts">
import { computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import SignupForm from "@/components/auth/SignupForm.vue";
import LoginForm from "@/components/auth/LoginForm.vue";

const route = useRoute();
const router = useRouter();

const activeTab = computed({
  get: () => (route.query.tab === "signup" ? "signup" : "login"),
  set: (tab) => router.push({ query: { tab } }),
});
</script>

<template>
  <div
    class="container mx-auto px-4 py-8 min-h-screen flex items-center justify-center"
  >
    <div class="w-full max-w-md">
      <Tabs v-model:modelValue="activeTab" class="w-full">
        <TabsList class="grid w-full grid-cols-2 mb-8">
          <TabsTrigger value="login">Login</TabsTrigger>
          <TabsTrigger value="signup">Sign Up</TabsTrigger>
        </TabsList>
        <TabsContent value="login">
          <LoginForm />
        </TabsContent>
        <TabsContent value="signup">
          <SignupForm />
        </TabsContent>
      </Tabs>
    </div>
  </div>
</template>
