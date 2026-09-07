import { library, config } from '@fortawesome/fontawesome-svg-core'
import '@fortawesome/fontawesome-svg-core/styles.css'
import {
  faArrowRight,
  faBars,
  faBell,
  faBuilding,
  faBox,
  faCalendarDays,
  faCheck,
  faChevronDown,
  faCircleQuestion,
  faClock,
  faCreditCard,
  faEnvelope,
  faGraduationCap,
  faHouse,
  faLocationDot,
  faLock,
  faPhone,
  faRoute,
  faShieldHalved,
  faStar,
  faTruck,
  faUserGroup,
  faVanShuttle,
  faWarehouse,
  faZap,
} from '@fortawesome/free-solid-svg-icons'
import { faApple, faGooglePlay } from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

config.autoAddCss = false

library.add(
  faArrowRight,
  faBars,
  faBell,
  faBuilding,
  faBox,
  faCalendarDays,
  faCheck,
  faChevronDown,
  faCircleQuestion,
  faClock,
  faCreditCard,
  faEnvelope,
  faGraduationCap,
  faHouse,
  faLocationDot,
  faLock,
  faPhone,
  faRoute,
  faShieldHalved,
  faStar,
  faTruck,
  faUserGroup,
  faVanShuttle,
  faWarehouse,
  faZap,
  faApple,
  faGooglePlay,
)

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('FontAwesomeIcon', FontAwesomeIcon)
})
