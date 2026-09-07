import { library, config } from '@fortawesome/fontawesome-svg-core'
import '@fortawesome/fontawesome-svg-core/styles.css'
import {
  faArrowLeft,
  faBolt,
  faCalendarDays,
  faClipboard,
  faEnvelope,
  faEye,
  faGaugeHigh,
  faGlobe,
  faPen,
  faPlus,
  faRightFromBracket,
} from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

config.autoAddCss = false
library.add(
  faArrowLeft,
  faBolt,
  faCalendarDays,
  faClipboard,
  faEnvelope,
  faEye,
  faGaugeHigh,
  faGlobe,
  faPen,
  faPlus,
  faRightFromBracket,
)

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('FontAwesomeIcon', FontAwesomeIcon)
})
