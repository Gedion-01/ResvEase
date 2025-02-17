<script setup lang="ts">
const currentYear = new Date().getFullYear();
</script>

<template>
  <footer class="bg-gray-900 text-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
        <div>
          <h3 class="text-lg font-semibold mb-4">About Us</h3>
          <ul class="space-y-2">
            <li><a href="#" class="hover:text-gray-300">Our Story</a></li>
            <li><a href="#" class="hover:text-gray-300">Team</a></li>
            <li><a href="#" class="hover:text-gray-300">Careers</a></li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">Support</h3>
          <ul class="space-y-2">
            <li><a href="#" class="hover:text-gray-300">FAQs</a></li>
            <li><a href="#" class="hover:text-gray-300">Contact Us</a></li>
            <li><a href="#" class="hover:text-gray-300">Privacy Policy</a></li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">Services</h3>
          <ul class="space-y-2">
            <li><a href="#" class="hover:text-gray-300">Hotel Booking</a></li>
            <li><a href="#" class="hover:text-gray-300">Flight Booking</a></li>
            <li><a href="#" class="hover:text-gray-300">Car Rental</a></li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-semibold mb-4">Follow Us</h3>
          <div class="flex space-x-4">
            <a href="#" class="">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                x="0px"
                y="0px"
                width="32"
                height="32"
                viewBox="0,0,256,256"
              >
                <g
                  class="icon"
                  fill="#fffdfd"
                  fill-rule="nonzero"
                  stroke="none"
                  stroke-width="1"
                  stroke-linecap="butt"
                  stroke-linejoin="miter"
                  stroke-miterlimit="10"
                  stroke-dasharray=""
                  stroke-dashoffset="0"
                  font-family="none"
                  font-weight="none"
                  font-size="none"
                  text-anchor="none"
                  style="mix-blend-mode: normal"
                >
                  <g transform="scale(10.66667,10.66667)">
                    <path
                      d="M2.36719,3l7.0957,10.14063l-6.72266,7.85938h2.64063l5.26367,-6.16992l4.31641,6.16992h6.91016l-7.42187,-10.625l6.29102,-7.375h-2.59961l-4.86914,5.6875l-3.97266,-5.6875zM6.20703,5h2.04883l9.77734,14h-2.03125z"
                    ></path>
                  </g>
                </g>
              </svg>
            </a>
            <a href="#" class="">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                x="0px"
                y="0px"
                width="32"
                height="32"
                viewBox="0,0,256,256"
              >
                <g
                  class="icon"
                  fill="#fffdfd"
                  fill-rule="nonzero"
                  stroke="none"
                  stroke-width="1"
                  stroke-linecap="butt"
                  stroke-linejoin="miter"
                  stroke-miterlimit="10"
                  stroke-dasharray=""
                  stroke-dashoffset="0"
                  font-family="none"
                  font-weight="none"
                  font-size="none"
                  text-anchor="none"
                  style="mix-blend-mode: normal"
                >
                  <g transform="scale(10.66667,10.66667)">
                    <path
                      d="M8,3c-2.8,0 -5,2.2 -5,5v8c0,2.8 2.2,5 5,5h8c2.8,0 5,-2.2 5,-5v-8c0,-2.8 -2.2,-5 -5,-5zM8,5h8c1.7,0 3,1.3 3,3v8c0,1.7 -1.3,3 -3,3h-8c-1.7,0 -3,-1.3 -3,-3v-8c0,-1.7 1.3,-3 3,-3zM17,6c-0.6,0 -1,0.4 -1,1c0,0.6 0.4,1 1,1c0.6,0 1,-0.4 1,-1c0,-0.6 -0.4,-1 -1,-1zM12,7c-2.8,0 -5,2.2 -5,5c0,2.8 2.2,5 5,5c2.8,0 5,-2.2 5,-5c0,-2.8 -2.2,-5 -5,-5zM12,9c1.7,0 3,1.3 3,3c0,1.7 -1.3,3 -3,3c-1.7,0 -3,-1.3 -3,-3c0,-1.7 1.3,-3 3,-3z"
                    ></path>
                  </g>
                </g>
              </svg>
            </a>
            <a href="#" class="">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                x="0px"
                y="0px"
                width="32"
                height="32"
                viewBox="0,0,256,256"
              >
                <g
                  class="icon"
                  fill="#fffdfd"
                  fill-rule="nonzero"
                  stroke="none"
                  stroke-width="1"
                  stroke-linecap="butt"
                  stroke-linejoin="miter"
                  stroke-miterlimit="10"
                  stroke-dasharray=""
                  stroke-dashoffset="0"
                  font-family="none"
                  font-weight="none"
                  font-size="none"
                  text-anchor="none"
                  style="mix-blend-mode: normal"
                >
                  <g transform="scale(10.66667,10.66667)">
                    <path
                      d="M12,2c-5.511,0 -10,4.489 -10,10c0,5.511 4.489,10 10,10c5.511,0 10,-4.489 10,-10c0,-5.511 -4.489,-10 -10,-10zM12,4c4.43012,0 8,3.56988 8,8c0,4.01447 -2.93468,7.31302 -6.78125,7.89844v-5.51367h2.32813l0.36523,-2.36524h-2.69336v-1.29297c0,-0.983 0.32023,-1.85547 1.24023,-1.85547h1.47656v-2.06445c-0.26,-0.035 -0.8087,-0.11133 -1.8457,-0.11133c-2.166,0 -3.43555,1.144 -3.43555,3.75v1.57422h-2.22656v2.36524h2.22656v5.49414c-3.78401,-0.63806 -6.6543,-3.90867 -6.6543,-7.87891c0,-4.43012 3.56988,-8 8,-8z"
                    ></path>
                  </g>
                </g>
              </svg>
            </a>
          </div>
        </div>
      </div>
      <div class="mt-8 pt-8 border-t border-gray-700 text-center">
        <p>&copy; {{ currentYear }} ResvEase. All rights reserved.</p>
      </div>
    </div>
  </footer>
</template>
<style scoped>
a:hover {
  .icon {
    fill: #d1d5db;
    transition-duration: 300ms;
  }
}
</style>
