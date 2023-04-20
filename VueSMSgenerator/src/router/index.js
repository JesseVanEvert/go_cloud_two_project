import { createRouter, createWebHistory } from 'vue-router'

import Home from '../components/Home.vue';

import GetAllClasses from '../components/GetAllClasses.vue';

import GetAllLecturers from '../components/GetAllLecturers.vue';

import GetStudent from '../components/GetStudentById.vue';

import FrontPage from '../components/FrontPage.vue';
import sendMessage from '../components/SendMessage.vue';



const routes = [
  { path: '/', component: Home },
  { path: '/frontpage', component: FrontPage },
  
  { path: '/getallclasses', component: GetAllClasses },
  
  { path: '/getalllecturers', component: GetAllLecturers },
  
  { path: '/getstudent', component: GetStudent }, 
  
  { path: '/sendMessage', component: sendMessage },
];

const router = createRouter({
  history: createWebHistory('/'), // Replace with your actual URL
  routes: routes
})
export default router