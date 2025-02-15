<script setup lang="ts">
import { useRouter } from "vue-router";
import { Hotel, MapPin, Calendar, Loader2 } from "lucide-vue-next";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

const router = useRouter();
</script>

<template>
  <div class="min-h-[50vh] flex items-center justify-center p-4">
    <Card class="w-full">
      <CardHeader>
        <CardTitle
          class="text-2xl font-bold text-center flex items-center justify-center"
        >
          <Hotel class="mr-2 h-8 w-8 text-primary" />
          No Hotels Found
        </CardTitle>
      </CardHeader>
      <CardContent class="text-center space-y-4">
        <p class="text-lg text-muted-foreground">
          We couldn't find any hotels matching your current filters. Don't
          worry, we've got some suggestions!
        </p>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mt-6">
          <div class="flex flex-col items-center p-4 bg-muted rounded-lg">
            <MapPin class="h-8 w-8 text-primary mb-2" />
            <h3 class="font-semibold">Expand Your Search</h3>
            <p class="text-sm text-muted-foreground">
              Try searching in nearby areas
            </p>
          </div>
          <div class="flex flex-col items-center p-4 bg-muted rounded-lg">
            <Calendar class="h-8 w-8 text-primary mb-2" />
            <h3 class="font-semibold">Adjust Dates</h3>
            <p class="text-sm text-muted-foreground">
              Check different dates for more options
            </p>
          </div>
          <div class="flex flex-col items-center p-4 bg-muted rounded-lg">
            <Loader2 class="h-8 w-8 text-primary mb-2 animate-spin" />
            <h3 class="font-semibold">Reset Filters</h3>
            <p class="text-sm text-muted-foreground">
              Clear current filters and start over
            </p>
          </div>
        </div>
      </CardContent>
    </Card>
  </div>
</template>

<style scoped>
@keyframes float {
  0% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
  100% {
    transform: translateY(0px);
  }
}

.animate-float {
  animation: float 3s ease-in-out infinite;
}
</style>
