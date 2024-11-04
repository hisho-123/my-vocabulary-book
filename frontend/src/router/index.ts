import { createRouter, createWebHistory } from "vue-router";
import Test from "../pages/Test.vue";
import Testest from "../pages/Testest.vue";

const routes = [
  {
    path: "/test",
    name: "Test",
    component: Test,
  },
  {
    path: "/testest",
    name: "Testest",
    component: Testest,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
