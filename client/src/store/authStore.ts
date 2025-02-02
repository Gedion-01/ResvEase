import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { useRouter } from "vue-router";
import apiClient from "@/services/api";

export const useAuthStore = defineStore("auth", () => {
  const token = ref(localStorage.getItem("authToken") || null);
  const user = ref(JSON.parse(localStorage.getItem("user") || "{}"));
  const redirectUrl = ref<string | null>(null);
  const router = useRouter();

  const isAuthenticated = computed(() => !!token.value);

  const setRedirectUrl = (url: string) => {
    redirectUrl.value = url;
  };

  const login = async (credential: { email: string; password: string }) => {
    try {
      const response = await apiClient.post("/auth", credential);
      token.value = response.data.token;
      user.value = response.data.user;

      localStorage.setItem("authToken", response.data.token);
      localStorage.setItem("user", JSON.stringify(response.data.user));

      router.push(redirectUrl.value || "/");
      redirectUrl.value = null;
    } catch (error) {
      console.error("An error occurred while logging in:", error);
      throw error;
    }
  };

  const logout = () => {
    token.value = null;
    user.value = {};
    localStorage.removeItem("authToken");
    localStorage.removeItem("user");
    router.push("/auth?tab=login");
  };

  const isTokenExpired = () => {
    if (!token.value) return true;

    try {
      const payload = JSON.parse(atob(token.value.split(".")[1])); // Decode JWT payload
      const expiry = payload.exp; // Expiration time in seconds
      return Date.now() >= expiry * 1000; // Compare with current time
    } catch (error) {
      return true;
    }
  };

  return {
    token,
    user,
    isAuthenticated,
    setRedirectUrl,
    login,
    logout,
    isTokenExpired,
  };
});
