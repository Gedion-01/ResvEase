import "./assets/main.css";
import router from "@/router";
import { VueQueryPlugin } from "@tanstack/vue-query";

import { createApp } from "vue";
import { createPinia } from "pinia";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import App from "./App.vue";

const app = createApp(App);
const pinia = createPinia();

app.use(pinia);
app.use(router);
app.use(VueQueryPlugin);

app.use(ElementPlus);

app.mount("#app");
