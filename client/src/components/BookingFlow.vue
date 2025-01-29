<script setup lang="ts">
import { type Ref, ref, reactive, computed, watch } from "vue";
import { useRouter } from "vue-router";
import { format, addDays } from "date-fns";
import { useRoomBookingStore } from "@/store/bookingStore";
import type { DateRange } from "radix-vue";
import { RangeCalendar } from "@/components/ui/range-calendar";
import { getLocalTimeZone, parseDate } from "@internationalized/date";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";

import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";

import { Checkbox } from "@/components/ui/checkbox";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";

import {
  CreditCard,
  Building,
  Coffee,
  Car,
  Wifi,
  Calendar,
  PocketKnife,
  Spade,
  Utensils,
  Dumbbell,
  Umbrella,
  ParkingCircle,
  Dog,
  Martini,
  Thermometer,
  Baby,
  Plane,
  WashingMachine,
  Anchor,
  Music,
  MapPin,
  Star,
  Minus,
  Bath,
  Mountain,
  Waves,
  ShowerHead,
} from "lucide-vue-next";

interface Room {
  id: string;
  name: string;
  price: number;
  image: string;
  bedType: string;
}

interface BookingFlowProps {
  room: Room;
  checkIn: Date;
  checkOut: Date;
}

import { CheckCircle2 } from "lucide-vue-next";
import { useSearchStore } from "@/store/searchStore";

interface BookingFlowProps {
  room: Room;
  checkIn: Date;
  checkOut: Date;
}

const props = defineProps<BookingFlowProps>();

const router = useRouter();
const step = ref(1);
const loading = ref(false);

const bookingStore = useRoomBookingStore();
const searchStore = useSearchStore();

const timeZone = getLocalTimeZone();
const start = searchStore.checkIn
  ? parseDate(searchStore.checkIn)
  : parseDate(new Date().toISOString().split("T")[0]);
const end = searchStore.checkOut
  ? parseDate(searchStore.checkOut)
  : parseDate(addDays(new Date(), 1).toISOString().split("T")[0]);

const value = ref({
  start,
  end,
}) as Ref<DateRange>;

const bookingData = reactive({
  firstName: "",
  lastName: "",
  email: "",
  phone: "",
  specialRequests: "",
  addons: [] as string[],
  paymentMethod: "credit-card",
  cardNumber: "",
  expiryDate: "",
  cvv: "",
  saveCard: false,
});

const numberOfNights = computed(() =>
  Math.ceil(
    ((value.value.end ? value.value.end.toDate(timeZone).getTime() : 0) -
      (value.value.start ? value.value.start.toDate(timeZone).getTime() : 0)) /
      (1000 * 60 * 60 * 24)
  )
);
const basePrice = computed(
  () => bookingStore.room.price * numberOfNights.value
);
const taxRate = 0.12;
const taxes = computed(() => basePrice.value * taxRate);
const addonsPrice = computed(() => bookingData.addons.length * 25);
const totalPrice = computed(
  () => basePrice.value + taxes.value + addonsPrice.value
);

const addons = [
  {
    id: "breakfast",
    title: "Breakfast Included",
    description: "Start your day with our delicious breakfast buffet",
    price: 25,
    icon: Coffee,
  },
  {
    id: "parking",
    title: "Parking Space",
    description: "Secure parking space for your vehicle",
    price: 25,
    icon: Car,
  },
  {
    id: "wifi",
    title: "Premium Wi-Fi",
    description: "High-speed internet access",
    price: 25,
    icon: Wifi,
  },
];

const handleAddonToggle = (addonId: string) => {
  const index = bookingData.addons.indexOf(addonId);
  if (index > -1) {
    bookingData.addons.splice(index, 1);
  } else {
    bookingData.addons.push(addonId);
  }
};

const handleInputChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  (bookingData[target.name as keyof typeof bookingData] as string) =
    target.value;
};

const validateStep = () => {
  switch (step.value) {
    case 1:
      return !!(
        bookingData.firstName &&
        bookingData.lastName &&
        bookingData.email &&
        bookingData.phone
      );
    case 2:
      return true; // Add-ons are optional
    case 3:
      return !!(
        bookingData.cardNumber &&
        bookingData.expiryDate &&
        bookingData.cvv
      );
    default:
      return false;
  }
};

const handleSubmit = async () => {
  loading.value = true;
  try {
    // Simulate API call
    await new Promise((resolve) => setTimeout(resolve, 2000));
    // Show success and redirect
    router.push("/booking-confirmation");
  } catch (error) {
    console.error("Booking failed:", error);
  } finally {
    loading.value = false;
  }
};

const nextStep = () => {
  if (step.value === 3) {
    handleSubmit();
  } else {
    step.value += 1;
  }
};

const formatDate = (date: Date) => format(date, "MMM dd, yyyy");

const findAddon = (id: string) => addons.find((addon) => addon.id === id);

const formattedStart = computed(() => {
  if (value?.value.start) {
    return format(value?.value.start.toDate(timeZone), "iii, MMM dd");
  }
  return "";
});
const formattedEnd = computed(() => {
  if (value?.value.end) {
    return format(value?.value.end.toDate(timeZone), "iii, MMM dd");
  }
  return "";
});

watch(
  value,
  (newValue) => {
    searchStore.checkIn = newValue.start
      ? format(newValue.start.toDate(timeZone), "yyyy-MM-dd")
      : "";
    searchStore.checkOut = newValue.end
      ? format(newValue.end.toDate(timeZone), "yyyy-MM-dd")
      : "";
    searchStore.setSearchParams({
      location: searchStore.location,
      checkIn: searchStore.checkIn,
      checkOut: searchStore.checkOut,
      adults: searchStore.adults,
      children: searchStore.children,
    });
  },
  { deep: true }
);

const getAmenityIcon = (amenity: string) => {
  switch (amenity) {
    case "Free Wi-Fi":
      return Wifi;
    case "Swimming Pool":
      return PocketKnife; // Using PocketKnife as a substitute for Pool
    case "Spa":
      return Spade; // Using Spade as a substitute for Spa
    case "Gym":
      return Dumbbell;
    case "Free Parking":
      return ParkingCircle;
    case "Breakfast Included":
      return Coffee;
    case "Pet Friendly":
      return Dog;
    case "Rooftop Bar":
      return Martini;
    case "Outdoor swimming pool":
      return PocketKnife; // Using PocketKnife as a substitute for Pool
    case "Sauna":
      return Thermometer;
    case "24/7 Room Service":
      return Utensils;
    case "Kids' Play Area":
      return Baby;
    case "Airport Shuttle":
      return Plane;
    case "Laundry Service":
      return WashingMachine;
    case "Beach Access":
      return Umbrella;
    case "Water Sports":
      return Anchor;
    case "Live Entertainment":
      return Music;
    case "Air conditioning":
      return Thermometer;
    case "Private bathroom":
      return Bath;
    case "Mountain view":
      return Mountain;
    case "Ocean view":
      return Waves;
    case "City view":
      return Building;
    case "Shower":
      return ShowerHead;
    default:
      return null;
  }
};
</script>

<template>
  <div class="max-w-3xl mx-auto py-8">
    <div
      class="mb-8 px-4 py-6 bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg shadow-md"
    >
      <div class="flex justify-between relative">
        <div
          v-for="(label, index) in ['Guest Info', 'Extras', 'Payment']"
          :key="label"
          class="flex flex-col justify-between relative z-10"
        >
          <div
            :class="[
              'w-10 h-10 sm:w-12 sm:h-12 rounded-full flex items-center justify-center mb-2 sm:mb-3 transition-all duration-300 ease-in-out transform',
              step > index + 1
                ? 'bg-green-500 text-white scale-105 sm:scale-110'
                : step === index + 1
                ? 'bg-primary text-primary-foreground'
                : 'bg-white text-gray-400 border-2 border-gray-200',
            ]"
          >
            <component
              :is="step > index + 1 ? CheckCircle2 : 'span'"
              :class="[
                step > index + 1
                  ? 'w-5 h-5 sm:w-6 sm:h-6'
                  : 'text-base sm:text-lg font-semibold',
                step === index + 1 ? 'animate-pulse' : '',
              ]"
            >
              {{ step > index + 1 ? "" : index + 1 }}
            </component>
          </div>
          <span
            :class="[
              'text-xs sm:text-sm font-medium transition-all duration-300 ease-in-out text-center',
              step >= index + 1 ? 'text-primary' : 'text-gray-500',
            ]"
          >
            {{ label }}
          </span>
        </div>

        <!-- Progress Line -->
        <div
          class="absolute top-5 sm:top-6 left-0 w-full h-0.5 bg-gray-200 rounded"
        >
          <div
            class="h-full bg-primary rounded transition-all duration-500 ease-in-out"
            :style="{ width: `${((step - 1) / 2) * 100}%` }"
          ></div>
        </div>
      </div>
    </div>

    <Card class="mb-6 overflow-hidden">
      <CardHeader
        class="bg-gradient-to-r from-blue-500 to-purple-600 text-white p-0"
      >
        <div class="relative h-48 w-full">
          <img
            :src="bookingStore.hotel.images[0]"
            class="absolute inset-0 w-full h-full object-cover"
            :alt="bookingStore.hotel.name"
          />
          <div class="absolute inset-0 bg-black bg-opacity-50"></div>
          <div class="absolute bottom-0 left-0 p-6 space-y-2">
            <CardTitle class="text-3xl font-bold text-white">{{
              bookingStore.hotel.name
            }}</CardTitle>
            <div class="flex items-center text-white/90">
              <MapPin class="h-5 w-5 mr-2" />
              <span>{{ bookingStore.hotel.location }}</span>
            </div>
          </div>
        </div>
      </CardHeader>
      <CardContent class="p-6">
        <div class="grid gap-6 md:grid-cols-2">
          <!-- Hotel Information -->
          <div class="space-y-4">
            <div class="flex items-center">
              <div class="flex">
                <Star
                  v-for="(val, index) in [1, 2, 3, 4, 5]"
                  :key="index"
                  :class="[
                    'h-5 w-5',
                    index < Math.floor(bookingStore.hotel.rating)
                      ? 'text-yellow-400 fill-current'
                      : 'text-gray-300',
                  ]"
                />
              </div>
              <span class="ml-2 text-sm font-medium text-gray-600">
                {{ bookingStore.hotel.rating.toFixed(1) }}
              </span>
            </div>
            <h4 class="text-xl font-semibold text-gray-800">
              {{ bookingStore.room.name }}
            </h4>
            <div class="grid grid-cols-2 gap-3">
              <div
                v-for="amenity in bookingStore.room.amenities"
                :key="amenity"
                class="flex items-center bg-gray-100 rounded-full px-3 py-1"
              >
                <component
                  :is="getAmenityIcon(amenity)"
                  class="h-4 w-4 text-blue-500"
                />
                <span class="ml-2 text-sm text-gray-700">{{ amenity }}</span>
              </div>
            </div>
          </div>

          <!-- Price Information -->
          <div class="space-y-4 flex flex-col justify-between">
            <div class="md:text-right">
              <div class="text-3xl font-bold text-gray-800">
                ${{ bookingStore.room.price }}
              </div>
              <div class="text-sm text-gray-600">per night</div>
            </div>
            <div class="flex md:justify-end items-end">
              <Popover>
                <PopoverTrigger class="" asChild>
                  <Button
                    variant="outline"
                    class="justify-start text-left font-normal"
                  >
                    <Calendar class="w-5 h-5 mr-2 text-gray-500" />
                    <div class="flex items-center">
                      <div
                        :class="[
                          'font-medium',
                          value?.start ? 'text-gray-900' : 'text-gray-400',
                        ]"
                      >
                        {{ value?.start ? formattedStart : "Check-in" }}
                      </div>
                      <Minus class="mx-2 h-4 w-4 text-gray-400" />
                      <div
                        :class="[
                          'font-medium',
                          value?.end ? 'text-gray-900' : 'text-gray-400',
                        ]"
                      >
                        {{ value?.end ? formattedEnd : "Check-out" }}
                      </div>
                    </div>
                  </Button>
                </PopoverTrigger>
                <PopoverContent class="w-auto p-0" align="center">
                  <RangeCalendar v-model="value" class="rounded-md border" />
                </PopoverContent>
              </Popover>
            </div>
            <div class="text-sm text-gray-600 md:text-right">
              {{ numberOfNights }} night{{ numberOfNights > 1 ? "s" : "" }}
            </div>
          </div>
        </div>

        <Separator class="my-6" />

        <!-- Selected Extras -->
        <div v-if="bookingData.addons.length > 0" class="space-y-2 mb-6">
          <h4 class="font-semibold text-gray-800">Selected Extras</h4>
          <div
            v-for="addonId in bookingData.addons"
            :key="addonId"
            class="flex justify-between text-sm"
          >
            <span class="text-gray-600">{{ findAddon(addonId)?.title }}</span>
            <span class="font-medium">${{ findAddon(addonId)?.price }}</span>
          </div>
        </div>

        <!-- Price Breakdown -->
        <div class="space-y-2">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">Subtotal</span>
            <span class="font-medium">${{ basePrice + addonsPrice }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600"
              >Taxes & Fees ({{ (taxRate * 100).toFixed(0) }}%)</span
            >
            <span class="font-medium">${{ taxes.toFixed(2) }}</span>
          </div>
          <Separator class="my-2" />
          <div class="flex justify-between text-lg font-bold text-gray-800">
            <span>Total</span>
            <span>${{ totalPrice.toFixed(2) }}</span>
          </div>
        </div>
      </CardContent>
    </Card>
    <!-- Step Content -->
    <Card>
      <CardHeader>
        <CardTitle>
          {{ step === 1 && "Guest Information" }}
          {{ step === 2 && "Add Extras to Your Stay" }}
          {{ step === 3 && "Payment Details" }}
        </CardTitle>
        <CardDescription>
          {{ step === 1 && "Please provide your contact information" }}
          {{ step === 2 && "Enhance your stay with these additional services" }}
          {{
            step === 3 &&
            "Enter your payment information to complete the booking"
          }}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <transition name="fade" mode="out-in">
          <div :key="step" class="motion-content">
            <div v-if="step === 1" class="grid gap-4">
              <div class="grid grid-cols-2 gap-4">
                <div class="space-y-2">
                  <Label for="firstName">First Name</Label>
                  <Input id="firstName" v-model="bookingData.firstName" />
                </div>
                <div class="space-y-2">
                  <Label for="lastName">Last Name</Label>
                  <Input id="lastName" v-model="bookingData.lastName" />
                </div>
              </div>
              <div class="space-y-2">
                <Label for="email">Email</Label>
                <Input id="email" type="email" v-model="bookingData.email" />
              </div>
              <div class="space-y-2">
                <Label for="phone">Phone Number</Label>
                <Input id="phone" type="tel" v-model="bookingData.phone" />
              </div>
              <div class="space-y-2">
                <Label for="specialRequests">Special Requests (Optional)</Label>
                <Input
                  id="specialRequests"
                  v-model="bookingData.specialRequests"
                />
              </div>
            </div>

            <div v-if="step === 2" class="grid gap-4">
              <div
                v-for="addon in addons"
                :key="addon.id"
                class="flex items-start space-x-4 p-4 border rounded-lg cursor-pointer hover:bg-accent"
                @click="handleAddonToggle(addon.id)"
              >
                <div class="">
                  <div
                    :class="[
                      'rounded-full p-2',
                      bookingData.addons.includes(addon.id)
                        ? 'bg-primary text-primary-foreground'
                        : 'bg-muted',
                    ]"
                  >
                    <addon.icon class="w-4 h-4" />
                  </div>
                </div>
                <div class="flex-1">
                  <div class="flex justify-between">
                    <h4 class="font-semibold">{{ addon.title }}</h4>
                    <span>${{ addon.price }}</span>
                  </div>
                  <p class="text-sm text-muted-foreground">
                    {{ addon.description }}
                  </p>
                </div>
                <Checkbox
                  :checked="bookingData.addons.includes(addon.id)"
                  @change="handleAddonToggle(addon.id)"
                />
              </div>
            </div>

            <div v-if="step === 3" class="space-y-6">
              <RadioGroup v-model="bookingData.paymentMethod">
                <div class="flex items-center space-x-2">
                  <RadioGroupItem value="credit-card" id="credit-card" />
                  <Label for="credit-card" class="flex items-center">
                    <CreditCard class="w-4 h-4 mr-2" />
                    Credit Card
                  </Label>
                </div>
                <div class="flex items-center space-x-2">
                  <RadioGroupItem value="debit-card" id="debit-card" />
                  <Label for="debit-card" class="flex items-center">
                    <Building class="w-4 h-4 mr-2" />
                    Debit Card
                  </Label>
                </div>
              </RadioGroup>

              <div class="space-y-4">
                <div class="space-y-2">
                  <Label for="cardNumber">Card Number</Label>
                  <Input
                    id="cardNumber"
                    v-model="bookingData.cardNumber"
                    placeholder="0000 0000 0000 0000"
                  />
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <div class="space-y-2">
                    <Label for="expiryDate">Expiry Date</Label>
                    <Input
                      id="expiryDate"
                      v-model="bookingData.expiryDate"
                      placeholder="MM/YY"
                    />
                  </div>
                  <div class="space-y-2">
                    <Label for="cvv">CVV</Label>
                    <Input
                      id="cvv"
                      type="password"
                      v-model="bookingData.cvv"
                      maxLength="4"
                    />
                  </div>
                </div>
                <div class="flex items-center space-x-2">
                  <Checkbox id="saveCard" v-model="bookingData.saveCard" />
                  <Label for="saveCard">Save card for future bookings</Label>
                </div>
              </div>
            </div>
          </div>
        </transition>
      </CardContent>
      <CardFooter class="flex justify-between">
        <Button
          variant="outline"
          @click="step > 1 && (step -= 1)"
          :disabled="step === 1"
          >Back</Button
        >
        <Button @click="nextStep" :disabled="!validateStep() || loading">
          {{
            loading
              ? "Processing..."
              : step === 3
              ? `Pay $${totalPrice.toFixed(2)}`
              : "Continue"
          }}
        </Button>
      </CardFooter>
    </Card>
  </div>
</template>
