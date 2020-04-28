import Vue from 'vue';

import { library } from '@fortawesome/fontawesome-svg-core';
import { faPlayCircle, faCheckCircle, faTimesCircle } from '@fortawesome/free-solid-svg-icons';
import { faGithub, faLinkedin, faChrome } from '@fortawesome/free-brands-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import router from './router';
import App from './App.vue';

Vue.component('fa-icon', FontAwesomeIcon);
library.add(faGithub, faLinkedin, faChrome, faPlayCircle, faCheckCircle, faTimesCircle);


Vue.config.productionTip = false;

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app');
