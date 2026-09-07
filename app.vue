<script setup lang="ts">
import { Motion } from 'motion-v'
import { adminNavItems, calendarCells, calendarDays } from './data/admin'
import { areas, brand, faqs, navigation, quoteSteps, routePairs, services, trustBadges } from './data/content'

const route = useRoute()
const router = useRouter()
const config = useRuntimeConfig()
const openFaq = ref(0)
const currentHeroSlide = ref(0)
const adminLoading = ref(false)
const loginError = ref('')
const quoteSearch = ref('')
const selectedContent = ref('business-profile')
const adminLogin = reactive({
  username: 'admin',
  password: 'password',
})
const quoteForm = reactive({
  pickupPostcode: '',
  deliveryPostcode: '',
  serviceType: 'House Move',
  preferredDate: '',
})
const showBookingModal = ref(false)
const bookingSaving = ref(false)
const bookingError = ref('')
const bookingForm = reactive({
  quoteRequestId: '',
  customerName: '',
  phoneNumber: '',
  serviceType: 'House Move',
  pickupPostcode: '',
  deliveryPostcode: '',
  bookingDate: '2026-09-07',
  status: 'confirmed',
  notes: '',
})
const heroSlides = [
  { title: 'House removals', image: '/images/hero-removals.png', position: 'object-center' },
  { title: 'Man and van', image: '/images/hero-removals-2.png', position: 'object-center' },
  { title: 'Same-day moves', image: '/images/hero-removals-3.png', position: 'object-center' },
]
const fallbackContentSections: ContentSection[] = [
  {
    id: 'business-profile',
    label: 'Business Profile',
    fields: {
      company: brand.name,
      phone: brand.phone,
      email: brand.email,
      coverage: 'UK-wide removals with same-day availability where possible.',
      promise: 'Clear pricing, careful movers, direct customer support.',
    },
  },
  {
    id: 'quote-workflow',
    label: 'Quote Workflow',
    fields: {
      stages: 'New inquiry -> Qualified -> Quote sent -> Booked -> Dispatch ready',
      responseTarget: 'Respond to new quotes within 60 minutes.',
      requiredDetails: 'Pickup, delivery, service type, preferred date, access notes.',
      followUp: 'Call once, send written confirmation, then schedule the move.',
    },
  },
  {
    id: 'operations',
    label: 'Operations',
    fields: {
      crewPlanning: 'Match crew size and van type to move size, access, and distance.',
      dispatchNotes: 'Capture parking, stairs, lifts, fragile items, and time windows.',
      insurance: 'Confirm insurance and handling notes before booking.',
      reviewRequest: 'Ask every completed customer for a review after the move.',
    },
  },
]
type QuoteRequest = {
  id: string
  fullName: string
  phoneNumber: string
  emailAddress: string
  serviceType: string
  pickupPostcode: string
  deliveryPostcode: string
  preferredDate: string
  flexibleOnDate: boolean
  additionalNotes: string
  status: string
  createdAt: string
}
type Booking = {
  id: string
  quoteRequestId: string
  customerName: string
  phoneNumber: string
  serviceType: string
  pickupPostcode: string
  deliveryPostcode: string
  bookingDate: string
  status: string
  notes: string
  createdAt: string
}
type AdminMessage = {
  id: string
  senderName: string
  senderEmail: string
  senderPhone: string
  subject: string
  body: string
  isRead: boolean
  createdAt: string
}
type ContentSection = {
  id: string
  label: string
  fields: Record<string, string>
}
const quotes = ref<QuoteRequest[]>([])
const bookings = ref<Booking[]>([])
const messages = ref<AdminMessage[]>([])
const contentSections = ref<ContentSection[]>([])
const isAdminRoute = computed(() => route.path.startsWith('/admin'))
const activeAdminView = computed(() => {
  if (route.path.includes('/quotes')) return 'quotes'
  if (route.path.includes('/bookings')) return 'bookings'
  if (route.path.includes('/messages')) return 'messages'
  if (route.path.includes('/content')) return 'content'
  if (route.path.includes('/login')) return 'login'
  return 'dashboard'
})
const apiBase = computed(() => config.public.apiBase as string)
const filteredQuotes = computed(() => {
  const term = quoteSearch.value.trim().toLowerCase()
  if (!term) return quotes.value
  return quotes.value.filter((quote) =>
    [quote.fullName, quote.pickupPostcode, quote.deliveryPostcode, quote.serviceType, quote.status]
      .some((value) => value?.toLowerCase().includes(term)),
  )
})
const unreadMessages = computed(() => messages.value.filter((message) => !message.isRead))
const confirmedBookings = computed(() => bookings.value.filter((booking) => booking.status === 'confirmed'))
const nextBooking = computed(() => bookings.value.find((booking) => booking.bookingDate) || bookings.value[0])
const bookingsByCalendarDay = computed(() => {
  return bookings.value.reduce<Record<string, Booking[]>>((days, booking) => {
    const day = booking.bookingDate?.slice(0, 10) === '2026-09-01'
      ? '1'
      : booking.bookingDate?.startsWith('2026-09-')
        ? String(Number(booking.bookingDate.slice(8, 10)))
        : ''
    if (day) {
      days[day] = [...(days[day] || []), booking]
    }
    return days
  }, {})
})
const todaysJobs = computed(() => bookingsByCalendarDay.value['7'] || [])
const dashboardStats = computed(() => [
  {
    label: 'New Quotes',
    value: quotes.value.filter((quote) => quote.status === 'new').length,
    note: 'Awaiting response',
    icon: 'clipboard',
    tone: 'text-dach-orange',
  },
  {
    label: 'Bookings This Month',
    value: bookings.value.length,
    note: 'Confirmed jobs',
    icon: 'calendar-days',
    tone: 'text-green-600',
  },
  {
    label: 'Unread Messages',
    value: unreadMessages.value.length,
    note: 'Needs attention',
    icon: 'envelope',
    tone: 'text-dach-orange',
  },
  {
    label: 'Next Job',
    value: nextBooking.value?.bookingDate || '-',
    note: nextBooking.value?.customerName || 'No booking scheduled',
    icon: 'truck',
    tone: 'text-dach-muted',
  },
])
const availableContentSections = computed(() => contentSections.value.length ? contentSections.value : fallbackContentSections)
const selectedContentSection = computed(() => availableContentSections.value.find((section) => section.id === selectedContent.value) || availableContentSections.value[0])

let heroTimer: ReturnType<typeof setInterval> | undefined

onMounted(() => {
  heroTimer = setInterval(() => {
    currentHeroSlide.value = (currentHeroSlide.value + 1) % heroSlides.length
  }, 4500)

  if (isAdminRoute.value) {
    loadAdminData()
  }
})

onBeforeUnmount(() => {
  if (heroTimer) clearInterval(heroTimer)
})

watch(isAdminRoute, (value) => {
  if (value) loadAdminData()
})

async function apiGet<T>(path: string) {
  return await $fetch<T>(`${apiBase.value}${path}`)
}

async function loadAdminData() {
  adminLoading.value = true
  try {
    const [quoteData, bookingData, messageData, contentData] = await Promise.all([
      apiGet<QuoteRequest[]>('/api/quotes'),
      apiGet<Booking[]>('/api/bookings'),
      apiGet<AdminMessage[]>('/api/messages'),
      apiGet<ContentSection[]>('/api/content'),
    ])
    quotes.value = quoteData
    bookings.value = bookingData
    messages.value = messageData
    contentSections.value = contentData
    if (!availableContentSections.value.some((section) => section.id === selectedContent.value)) {
      selectedContent.value = availableContentSections.value[0]?.id || 'business-profile'
    }
  } catch (error) {
    loadLocalAdminData()
    if (!contentSections.value.length) {
      contentSections.value = fallbackContentSections
    }
  } finally {
    adminLoading.value = false
  }
}

function loadLocalAdminData() {
  if (!import.meta.client) return
  const savedBookings = window.localStorage.getItem('dach-admin-bookings')
  if (savedBookings) {
    bookings.value = JSON.parse(savedBookings)
  }
}

function saveLocalBookings() {
  if (!import.meta.client) return
  window.localStorage.setItem('dach-admin-bookings', JSON.stringify(bookings.value))
}

function calendarCellBookings(cell: string, index: number) {
  if (index < 1 || index > 30) return []
  return bookingsByCalendarDay.value[cell] || []
}

async function submitQuote() {
  await $fetch(`${apiBase.value}/api/quotes`, {
    method: 'POST',
    body: {
      fullName: 'Website Visitor',
      phoneNumber: '',
      emailAddress: '',
      serviceType: quoteForm.serviceType,
      pickupPostcode: quoteForm.pickupPostcode,
      deliveryPostcode: quoteForm.deliveryPostcode,
      preferredDate: quoteForm.preferredDate,
      flexibleOnDate: true,
      additionalNotes: 'Submitted from landing quote form.',
    },
  })
  quoteForm.pickupPostcode = ''
  quoteForm.deliveryPostcode = ''
  quoteForm.preferredDate = ''
}

function openBookingForm(quote?: QuoteRequest) {
  bookingError.value = ''
  bookingForm.quoteRequestId = quote?.id || ''
  bookingForm.customerName = quote?.fullName || ''
  bookingForm.phoneNumber = quote?.phoneNumber || ''
  bookingForm.serviceType = quote?.serviceType || 'House Move'
  bookingForm.pickupPostcode = quote?.pickupPostcode || ''
  bookingForm.deliveryPostcode = quote?.deliveryPostcode || ''
  bookingForm.bookingDate = quote?.preferredDate?.match(/^\d{4}-\d{2}-\d{2}$/) ? quote.preferredDate : '2026-09-07'
  bookingForm.status = 'confirmed'
  bookingForm.notes = quote?.additionalNotes || ''
  showBookingModal.value = true
}

async function createBooking() {
  bookingSaving.value = true
  bookingError.value = ''
  try {
    const created = await $fetch<Booking>(`${apiBase.value}/api/bookings`, {
      method: 'POST',
      body: bookingForm,
    })
    bookings.value = [...bookings.value, created]
    saveLocalBookings()
    showBookingModal.value = false
  } catch (error) {
    const localBooking: Booking = {
      id: `local_${Date.now()}`,
      quoteRequestId: bookingForm.quoteRequestId,
      customerName: bookingForm.customerName,
      phoneNumber: bookingForm.phoneNumber,
      serviceType: bookingForm.serviceType,
      pickupPostcode: bookingForm.pickupPostcode,
      deliveryPostcode: bookingForm.deliveryPostcode,
      bookingDate: bookingForm.bookingDate,
      status: bookingForm.status,
      notes: bookingForm.notes,
      createdAt: new Date().toISOString(),
    }
    bookings.value = [...bookings.value, localBooking]
    saveLocalBookings()
    showBookingModal.value = false
  } finally {
    bookingSaving.value = false
  }
}

async function loginAdmin() {
  loginError.value = ''
  try {
    await $fetch(`${apiBase.value}/api/auth/login`, {
      method: 'POST',
      body: adminLogin,
    })
    await router.push('/admin/dashboard')
  } catch (error) {
    loginError.value = 'Invalid credentials or backend offline.'
  }
}
</script>

<template>
  <div class="hidden">
    <NuxtPage />
  </div>

  <main v-if="!isAdminRoute" class="min-h-screen bg-white text-dach-black">
    <header class="sticky top-0 z-50 border-b border-dach-line bg-white/95 backdrop-blur">
      <div class="section-wrap flex h-20 items-center justify-between">
        <a href="#" class="flex items-center">
          <img src="/images/logo.jpg" alt="Dach Removals" class="h-12 w-auto object-contain" />
        </a>

        <nav class="hidden items-center gap-8 text-sm font-semibold lg:flex">
          <a v-for="item in navigation" :key="item.href" :href="item.href" class="transition hover:text-dach-orange">
            {{ item.label }}
          </a>
        </nav>

        <div class="flex items-center gap-3">
          <a :href="`tel:${brand.phone}`" class="hidden items-center gap-3 rounded-full bg-dach-cream px-5 py-3 text-sm font-semibold md:flex">
            <FontAwesomeIcon icon="phone" />
            {{ brand.phone }}
          </a>
          <a href="/admin" class="rounded-full bg-dach-orange px-5 py-3 text-sm font-semibold text-white transition hover:bg-dach-black">Log In</a>
        </div>
      </div>
    </header>

    <section class="relative isolate overflow-hidden bg-dach-black text-white">
      <img
        v-for="(slide, index) in heroSlides"
        :key="slide.title"
        :src="slide.image"
        :alt="slide.title"
        class="absolute inset-0 -z-20 h-full w-full object-cover transition duration-1000"
        :class="[slide.position, currentHeroSlide === index ? 'scale-100 opacity-100' : 'scale-105 opacity-0']"
      />
      <div class="absolute inset-0 -z-10 bg-gradient-to-r from-dach-black/88 via-dach-black/58 to-dach-black/18" />
      <div class="absolute inset-x-0 bottom-0 -z-10 h-40 bg-gradient-to-t from-dach-black/60 to-transparent" />

      <Motion
        as="div"
        class="section-wrap relative grid min-h-[700px] items-center gap-12 py-20 lg:grid-cols-[0.92fr_430px]"
        :initial="{ opacity: 0 }"
        :animate="{ opacity: 1 }"
        :transition="{ duration: 0.5, ease: 'easeOut' }"
      >
        <Motion
          as="div"
          class="max-w-2xl"
          :initial="{ opacity: 0, y: 24 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.55, ease: 'easeOut' }"
        >
          <div class="mb-7 flex flex-wrap gap-2">
            <span v-for="badge in trustBadges" :key="badge" class="rounded-full border border-white/15 bg-white/10 px-3 py-2 text-sm font-medium backdrop-blur">{{ badge }}</span>
          </div>
          <h1 class="max-w-3xl text-5xl font-bold leading-[1.02] tracking-tight md:text-7xl">
            UK removals, <span class="text-dach-orange">made simple.</span>
          </h1>
          <p class="mt-5 max-w-xl text-lg leading-8 text-white/82">Fast quotes. Careful movers. Clear pricing.</p>
          <div class="mt-8 flex flex-wrap gap-3">
            <a href="#quote" class="rounded-full bg-dach-orange px-5 py-4 font-semibold text-white transition hover:bg-white hover:text-dach-black">
              Get a Quote <FontAwesomeIcon icon="arrow-right" class="ml-2" />
            </a>
            <a :href="`tel:${brand.phone}`" class="rounded-full border border-white/25 px-5 py-4 font-semibold text-white transition hover:bg-white hover:text-dach-black">
              <FontAwesomeIcon icon="phone" class="mr-2" />{{ brand.phone }}
            </a>
          </div>
          <div class="mt-10 flex items-center gap-3">
            <button
              v-for="(slide, index) in heroSlides"
              :key="`dot-${slide.title}`"
              class="h-1.5 w-10 transition"
              :class="currentHeroSlide === index ? 'bg-dach-orange' : 'bg-white/25'"
              type="button"
              :aria-label="`Show ${slide.title}`"
              @click="currentHeroSlide = index"
            />
          </div>
        </Motion>

        <Motion
          id="quote"
          as="form"
          class="relative z-10 rounded-3xl border border-white/25 bg-dach-black/35 p-7 text-white shadow-2xl shadow-black/35 backdrop-blur-md"
          :initial="{ opacity: 0, y: 26 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.55, ease: 'easeOut', delay: 0.08 }"
          @submit.prevent="submitQuote"
        >
          <h2 class="text-2xl font-bold tracking-tight">Get Your Price</h2>
          <p class="mt-2 text-sm text-white/75">Enter your details. We confirm the rest.</p>

          <label class="mt-6 block text-sm font-semibold">Moving From</label>
          <div class="mt-2 flex items-center gap-3 rounded-2xl border border-white/15 bg-white/15 px-4 py-4 text-white/75">
            <FontAwesomeIcon icon="location-dot" />
            <input v-model="quoteForm.pickupPostcode" class="w-full bg-transparent text-white outline-none placeholder:text-white/55" placeholder="Enter postcode" />
          </div>

          <label class="mt-4 block text-sm font-semibold">Moving To</label>
          <div class="mt-2 flex items-center gap-3 rounded-2xl border border-white/15 bg-white/15 px-4 py-4 text-white/75">
            <FontAwesomeIcon icon="route" />
            <input v-model="quoteForm.deliveryPostcode" class="w-full bg-transparent text-white outline-none placeholder:text-white/55" placeholder="Enter postcode" />
          </div>

          <div class="mt-4 grid gap-3 md:grid-cols-2">
            <select v-model="quoteForm.serviceType" class="rounded-2xl border border-white/15 bg-white/15 px-4 py-4 text-white outline-none">
              <option>House Move</option>
              <option>Man and Van</option>
              <option>Office Relocation</option>
              <option>Packing and Storage</option>
            </select>
            <input v-model="quoteForm.preferredDate" class="rounded-2xl border border-white/15 bg-white/15 px-4 py-4 text-white outline-none placeholder:text-white/55" placeholder="dd/mm/yyyy" />
          </div>

          <button class="mt-6 w-full rounded-full bg-dach-orange px-6 py-4 font-semibold text-white transition hover:bg-white hover:text-dach-black">
            See Price Instantly <FontAwesomeIcon icon="arrow-right" class="ml-2" />
          </button>
          <p class="mt-4 text-center text-sm text-white/75"><FontAwesomeIcon icon="lock" class="mr-1" /> Secure enquiry. Rated 4.8/5.</p>
        </Motion>
      </Motion>
    </section>

    <section id="how-it-works" class="section-wrap relative overflow-hidden py-24">
      <span class="route-lines right-0 top-10 opacity-50" />
      <span class="corner-mark bottom-12 left-0 opacity-60" />
      <Motion
        as="div"
        class="grid gap-6 lg:grid-cols-[0.75fr_1fr] lg:items-end"
        :initial="{ opacity: 0, y: 28 }"
        :whileInView="{ opacity: 1, y: 0 }"
        :inViewOptions="{ once: true, margin: '-80px' }"
        :transition="{ duration: 0.5, ease: 'easeOut' }"
      >
        <h2 class="text-4xl font-bold tracking-tight md:text-5xl">How it works.</h2>
        <p class="text-lg leading-8 text-dach-muted">Five quick steps from quote to moving day.</p>
      </Motion>
      <div class="mt-12 grid gap-5 md:grid-cols-5">
        <Motion
          v-for="(step, index) in quoteSteps"
          :key="step.title"
          as="article"
          class="rounded-3xl border border-dach-line bg-white p-6 shadow-sm"
          :initial="{ opacity: 0, y: 26 }"
          :whileInView="{ opacity: 1, y: 0 }"
          :inViewOptions="{ once: true, margin: '-70px' }"
          :whileHover="{ y: -5 }"
          :transition="{ duration: 0.45, ease: 'easeOut', delay: index * 0.05 }"
        >
          <span class="mb-7 flex h-11 w-11 items-center justify-center rounded-2xl bg-dach-orange text-white"><FontAwesomeIcon :icon="step.icon" /></span>
          <span class="text-sm font-semibold text-dach-muted">0{{ index + 1 }}</span>
          <h3 class="mt-3 text-lg font-bold">{{ step.title }}</h3>
          <p class="mt-3 text-sm leading-6 text-dach-muted">{{ step.copy }}</p>
        </Motion>
      </div>
    </section>

    <section id="services" class="relative overflow-hidden border-y border-dach-line bg-white py-20">
      <div class="abstract-grid absolute inset-y-0 right-0 w-1/2 opacity-50" />
      <div class="section-wrap relative">
        <Motion
          as="div"
          class="grid gap-4 lg:grid-cols-[0.7fr_1fr] lg:items-end"
          :initial="{ opacity: 0, y: 28 }"
          :whileInView="{ opacity: 1, y: 0 }"
          :inViewOptions="{ once: true, margin: '-80px' }"
          :transition="{ duration: 0.5, ease: 'easeOut' }"
        >
          <h2 class="text-4xl font-bold tracking-tight md:text-5xl">Services.</h2>
          <p class="max-w-xl text-lg leading-8 text-dach-muted">Home, office, student, furniture, packing, and storage moves.</p>
        </Motion>

        <div class="mt-10 grid gap-5 md:grid-cols-2 xl:grid-cols-3">
          <Motion
            v-for="(service, index) in services"
            :key="service.title"
            as="article"
            class="relative overflow-hidden rounded-3xl border border-dach-line bg-white shadow-sm transition hover:shadow-xl hover:shadow-orange-100"
            :initial="{ opacity: 0, y: 30 }"
            :whileInView="{ opacity: 1, y: 0 }"
            :inViewOptions="{ once: true, margin: '-70px' }"
            :whileHover="{ y: -6 }"
            :transition="{ duration: 0.45, ease: 'easeOut', delay: index * 0.04 }"
          >
            <img v-if="index < 2" :src="service.image" :alt="service.title" class="h-44 w-full object-cover" />
            <div class="p-6">
              <span class="mb-5 flex h-11 w-11 items-center justify-center rounded-2xl bg-dach-orange text-white"><FontAwesomeIcon :icon="service.icon" /></span>
              <h3 class="text-xl font-bold">{{ service.title }}</h3>
              <p class="mt-3 min-h-14 leading-7 text-dach-muted">{{ service.description }}</p>
              <ul v-if="index < 2" class="mt-4 grid gap-2 text-sm text-dach-muted">
                <li v-for="item in service.idealFor.slice(0, 3)" :key="item"><FontAwesomeIcon icon="check" class="mr-2 text-dach-orange" />{{ item }}</li>
              </ul>
              <a href="#quote" class="mt-6 inline-flex items-center gap-2 font-semibold text-dach-orange">Get Quote <FontAwesomeIcon icon="arrow-right" /></a>
            </div>
          </Motion>
        </div>
      </div>
    </section>

    <section id="areas" class="relative overflow-hidden border-y border-dach-line bg-gray-50 py-24">
      <span class="route-lines left-10 top-24 opacity-40" />
      <div class="section-wrap grid gap-10 lg:grid-cols-[0.75fr_1.25fr]">
        <Motion
          as="div"
          :initial="{ opacity: 0, x: -28 }"
          :whileInView="{ opacity: 1, x: 0 }"
          :inViewOptions="{ once: true, margin: '-80px' }"
          :transition="{ duration: 0.5, ease: 'easeOut' }"
        >
          <h2 class="text-4xl font-bold tracking-tight md:text-5xl">Areas we cover.</h2>
          <p class="mt-5 leading-8 text-dach-muted">Local and long-distance removals across the UK.</p>
          <div class="mt-8 flex flex-wrap gap-3">
            <span v-for="routePair in routePairs" :key="routePair" class="rounded-full bg-white px-4 py-2 text-sm font-semibold shadow-sm">{{ routePair }}</span>
          </div>
        </Motion>
        <div class="grid gap-5 md:grid-cols-2">
          <Motion
            v-for="([area, cities], index) in areas"
            :key="area"
            as="article"
            class="rounded-3xl bg-white p-6 shadow-sm"
            :initial="{ opacity: 0, y: 24 }"
            :whileInView="{ opacity: 1, y: 0 }"
            :inViewOptions="{ once: true, margin: '-70px' }"
            :whileHover="{ y: -4 }"
            :transition="{ duration: 0.42, ease: 'easeOut', delay: index * 0.04 }"
          >
            <h3 class="text-lg font-bold"><FontAwesomeIcon icon="route" class="mr-2 text-dach-orange" />{{ area }}</h3>
            <ul class="mt-4 grid gap-2 text-sm text-dach-muted">
              <li v-for="city in cities" :key="city">-> {{ city }}</li>
            </ul>
          </Motion>
        </div>
      </div>
    </section>

    <Motion
      id="faq"
      as="section"
      class="section-wrap grid gap-10 py-24 lg:grid-cols-[0.7fr_1.3fr]"
      :initial="{ opacity: 0, y: 28 }"
      :whileInView="{ opacity: 1, y: 0 }"
      :inViewOptions="{ once: true, margin: '-80px' }"
      :transition="{ duration: 0.5, ease: 'easeOut' }"
    >
      <div>
        <h2 class="text-4xl font-bold tracking-tight md:text-5xl">FAQ.</h2>
        <p class="mt-5 leading-8 text-dach-muted">Quick answers before you book.</p>
        <p class="mt-8 font-semibold">Need help? <a :href="`mailto:${brand.email}`" class="text-dach-orange">{{ brand.email }}</a></p>
      </div>
      <div class="space-y-3">
        <Motion
          v-for="(faq, index) in faqs"
          :key="faq"
          as="article"
          class="overflow-hidden rounded-2xl border border-dach-line"
          :initial="{ opacity: 0, x: 18 }"
          :whileInView="{ opacity: 1, x: 0 }"
          :inViewOptions="{ once: true, margin: '-70px' }"
          :transition="{ duration: 0.35, ease: 'easeOut', delay: index * 0.04 }"
        >
          <button class="flex w-full items-center justify-between p-5 text-left font-semibold" @click="openFaq = openFaq === index ? -1 : index">
            {{ faq }}
            <FontAwesomeIcon :icon="openFaq === index ? 'chevron-down' : 'arrow-right'" class="text-dach-orange" />
          </button>
          <p v-if="openFaq === index" class="px-5 pb-5 leading-7 text-dach-muted">
            Send the quote form or call us. We will confirm the right service and availability.
          </p>
        </Motion>
      </div>
    </Motion>

    <section class="relative overflow-hidden bg-dach-black py-20 text-white">
      <span class="corner-mark right-24 top-8 border-white/15 opacity-60" />
      <Motion
        as="div"
        class="section-wrap grid gap-8 lg:grid-cols-[1fr_auto] lg:items-center"
        :initial="{ opacity: 0, y: 28 }"
        :whileInView="{ opacity: 1, y: 0 }"
        :inViewOptions="{ once: true, margin: '-80px' }"
        :transition="{ duration: 0.5, ease: 'easeOut' }"
      >
        <div>
          <h2 class="text-4xl font-bold tracking-tight md:text-5xl">Ready to move?</h2>
          <p class="mt-4 max-w-2xl leading-8 text-white/70">Get a quote or call the team.</p>
        </div>
        <div class="flex flex-wrap gap-3">
          <a href="#quote" class="rounded-full bg-dach-orange px-6 py-4 font-semibold text-white">Get a Quote</a>
          <a :href="`tel:${brand.phone}`" class="rounded-full border border-white/20 px-6 py-4 font-semibold text-white"><FontAwesomeIcon icon="phone" class="mr-2" />{{ brand.phone }}</a>
        </div>
      </Motion>
    </section>

    <footer class="relative overflow-hidden bg-[#101010] py-16 text-white">
      <img src="/images/hero-removals.png" alt="" class="absolute inset-0 h-full w-full object-cover opacity-20" />
      <div class="absolute inset-0 bg-gradient-to-r from-[#101010] via-[#101010]/92 to-[#101010]/78" />
      <div class="absolute inset-0 bg-dach-black/55" />
      <div class="section-wrap relative grid gap-10 md:grid-cols-4">
        <div>
          <img src="/images/logo.jpg" alt="Dach Removals" class="mb-5 h-12 w-auto bg-white object-contain" />
          <p class="text-white/60">UK removals. Clear pricing. Careful crews.</p>
        </div>
        <div>
          <h3 class="font-semibold">Services</h3>
          <p v-for="service in services.slice(0, 5)" :key="service.title" class="mt-3 text-white/55">{{ service.title }}</p>
        </div>
        <div>
          <h3 class="font-semibold">Company</h3>
          <p class="mt-3 text-white/55">About Us</p>
          <p class="mt-3 text-white/55">Reviews</p>
          <p class="mt-3 text-white/55">Get a Quote</p>
          <p class="mt-3 text-white/55">Help Centre</p>
        </div>
        <div>
          <h3 class="font-semibold">Contact</h3>
          <p class="mt-3 text-white/55">{{ brand.phone }}</p>
          <p class="mt-3 text-white/55">{{ brand.email }}</p>
          <p class="mt-8 text-white/35">© 2026 Dach Removals.</p>
        </div>
      </div>
    </footer>
  </main>

  <div v-else-if="activeAdminView === 'login'" class="relative isolate min-h-screen overflow-hidden bg-dach-black text-white">
    <img src="/images/hero-removals-2.png" alt="" class="absolute inset-0 -z-20 h-full w-full object-cover opacity-70" />
    <div class="absolute inset-0 -z-10 bg-gradient-to-r from-dach-black via-dach-black/80 to-dach-black/45" />
    <div class="abstract-grid absolute inset-0 -z-10 opacity-20" />
    <span class="corner-mark left-16 top-16 border-white/20" />

    <div class="mx-auto grid min-h-screen w-full max-w-6xl items-center gap-12 px-6 py-12 lg:grid-cols-[1fr_440px]">
      <Motion
        as="section"
        class="max-w-xl"
        :initial="{ opacity: 0, x: -28 }"
        :animate="{ opacity: 1, x: 0 }"
        :transition="{ duration: 0.55, ease: 'easeOut' }"
      >
        <img src="/images/logo.jpg" alt="Dach Removals" class="h-14 w-auto rounded-xl bg-white object-contain" />
        <p class="mt-8 text-sm font-semibold uppercase tracking-[0.18em] text-dach-orange">Operations Portal</p>
        <h1 class="mt-4 max-w-lg text-5xl font-bold leading-tight">Manage quotes, bookings, and move-day work.</h1>
        <p class="mt-5 max-w-lg text-lg leading-8 text-white/70">For the Dach Removals team only. Keep customer requests, schedules, and dispatch notes organised.</p>
      </Motion>

      <Motion
        as="form"
        class="rounded-3xl border border-white/18 bg-white/12 p-8 shadow-2xl shadow-black/40 backdrop-blur-xl"
        :initial="{ opacity: 0, y: 28 }"
        :animate="{ opacity: 1, y: 0 }"
        :transition="{ duration: 0.55, ease: 'easeOut', delay: 0.08 }"
        @submit.prevent="loginAdmin"
      >
        <div class="mb-8">
          <p class="text-sm font-semibold text-white/60">Secure sign in</p>
          <h2 class="mt-2 font-google-sans text-3xl font-bold">Admin access</h2>
        </div>
        <label class="text-xs font-bold uppercase tracking-[0.18em] text-white/55">Username</label>
        <input v-model="adminLogin.username" class="mt-2 w-full rounded-2xl border border-white/14 bg-white/14 px-5 py-4 text-white outline-none placeholder:text-white/45 focus:border-dach-orange" />
        <label class="mt-5 block text-xs font-bold uppercase tracking-[0.18em] text-white/55">Password</label>
        <input v-model="adminLogin.password" type="password" class="mt-2 w-full rounded-2xl border border-white/14 bg-white/14 px-5 py-4 text-white outline-none placeholder:text-white/45 focus:border-dach-orange" />
        <button class="mt-6 block w-full rounded-full bg-dach-orange px-6 py-4 text-center font-bold text-white transition hover:bg-white hover:text-dach-black" type="submit">
          Sign In <FontAwesomeIcon icon="arrow-right" class="ml-2" />
        </button>
        <p v-if="loginError" class="mt-4 rounded-2xl bg-dach-orange/15 px-4 py-3 text-sm font-semibold text-white">{{ loginError }}</p>
        <a href="/" class="mt-8 inline-flex items-center gap-2 text-sm font-semibold text-white/60 transition hover:text-white"><FontAwesomeIcon icon="arrow-left" /> Back to website</a>
      </Motion>
    </div>
  </div>

  <div v-else class="grid min-h-screen grid-cols-[288px_1fr] bg-[#f7f3ef] text-dach-black">
    <aside class="sticky top-0 flex h-screen flex-col bg-[#111111] text-white">
      <div class="border-b border-white/10 px-6 py-7">
        <img src="/images/logo.jpg" alt="Dach Removals" class="h-12 w-auto bg-white object-contain" />
      </div>

      <nav class="space-y-1 px-4 py-5">
        <NuxtLink
          v-for="item in adminNavItems"
          :key="item.label"
          :to="item.path"
          class="group flex items-center gap-4 px-4 py-4 text-sm font-semibold text-white/55 transition hover:bg-white/5 hover:text-white"
          :class="route.path === item.path ? 'bg-dach-orange text-white shadow-lg shadow-dach-orange/20 hover:bg-dach-orange' : ''"
        >
          <span class="grid h-9 w-9 place-items-center bg-white/8 text-white/60 group-hover:text-white" :class="route.path === item.path ? 'bg-white/15 text-white' : ''">
            <FontAwesomeIcon :icon="item.icon" />
          </span>
          <span>{{ item.label }}</span>
          <span
            v-if="(item.label === 'Quote Requests' && quotes.length > 0) || (item.label === 'Messages' && unreadMessages.length > 0)"
            class="ml-auto rounded-full px-2.5 py-1 text-xs font-bold"
            :class="route.path === item.path ? 'bg-white text-dach-orange' : 'bg-dach-orange text-white'"
          >
            {{ item.label === 'Quote Requests' ? quotes.length : unreadMessages.length }}
          </span>
        </NuxtLink>
      </nav>

      <div class="mx-4 mt-auto border-t border-white/10 py-5 text-sm font-semibold text-white/45">
        <a href="/" class="flex items-center gap-3 px-4 py-3 transition hover:text-white"><FontAwesomeIcon icon="globe" />View Website</a>
        <NuxtLink to="/admin/login" class="flex items-center gap-3 px-4 py-3 transition hover:text-white"><FontAwesomeIcon icon="right-from-bracket" />Sign Out</NuxtLink>
      </div>
    </aside>

    <main class="min-w-0">
      <header class="sticky top-0 z-30 border-b border-dach-line bg-white/95 px-9 py-5 backdrop-blur">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="font-google-sans text-3xl font-bold">
              {{ activeAdminView === 'dashboard' ? 'Dashboard' : activeAdminView === 'quotes' ? 'Quote Requests' : activeAdminView === 'bookings' ? 'Bookings Calendar' : activeAdminView === 'messages' ? 'Messages' : 'Business Setup' }}
            </h1>
          </div>
          <div class="flex items-center gap-4">
            <button class="border border-dach-line bg-white px-4 py-3 text-sm font-semibold text-dach-muted" type="button" @click="loadAdminData">
              <FontAwesomeIcon icon="rotate-right" class="mr-2 text-dach-orange" />Refresh
            </button>
            <p class="text-sm text-dach-muted">Mon, 7 September 2026</p>
          </div>
        </div>
      </header>

      <section v-if="activeAdminView === 'dashboard'" class="p-9">
        <div class="grid gap-5 xl:grid-cols-4">
          <article v-for="card in dashboardStats" :key="card.label" class="border border-dach-line bg-white p-6 shadow-sm">
            <div class="flex items-start justify-between">
              <p class="text-xs font-bold uppercase tracking-[0.18em] text-dach-muted">{{ card.label }}</p>
              <span class="grid h-10 w-10 place-items-center bg-dach-orange/10 text-dach-orange"><FontAwesomeIcon :icon="card.icon" /></span>
            </div>
            <p class="mt-5 font-google-sans text-4xl font-bold">{{ card.value }}</p>
            <p :class="card.tone" class="mt-2 text-sm font-medium">{{ card.note }}</p>
          </article>
        </div>

        <div class="mt-8 grid gap-6 xl:grid-cols-[1.1fr_0.9fr]">
          <article class="border border-dach-line bg-white shadow-sm">
            <div class="flex items-center justify-between border-b border-dach-line px-6 py-5">
              <h2 class="font-google-sans text-xl font-bold">Recent Quotes</h2>
              <NuxtLink to="/admin/quotes" class="border border-dach-line px-4 py-2 text-sm font-semibold">View All</NuxtLink>
            </div>
            <div v-if="quotes.length" class="divide-y divide-dach-line">
              <div v-for="quote in quotes.slice(0, 5)" :key="quote.id" class="grid grid-cols-[1fr_auto] gap-4 px-6 py-4">
                <div>
                  <p class="font-semibold">{{ quote.fullName || 'Website Visitor' }}</p>
                  <p class="mt-1 text-sm text-dach-muted">{{ quote.pickupPostcode || 'Pickup' }} -> {{ quote.deliveryPostcode || 'Delivery' }} · {{ quote.serviceType }}</p>
                </div>
                <span class="self-start bg-dach-orange/10 px-3 py-1 text-xs font-bold uppercase text-dach-orange">{{ quote.status }}</span>
              </div>
            </div>
            <div v-else class="grid h-72 place-items-center p-8 text-center text-dach-muted">
              <div>
                <FontAwesomeIcon icon="clipboard" class="mb-4 text-4xl text-dach-orange/70" />
                <p class="font-semibold text-dach-black">No quote requests yet</p>
                <p class="mt-2 max-w-sm text-sm leading-6">New website enquiries will appear here with route, service, date, and contact details.</p>
              </div>
            </div>
          </article>

          <article class="border border-dach-line bg-white shadow-sm">
            <div class="flex items-center justify-between border-b border-dach-line px-6 py-5">
              <h2 class="font-google-sans text-xl font-bold">Upcoming Bookings</h2>
              <NuxtLink to="/admin/bookings" class="border border-dach-line px-4 py-2 text-sm font-semibold">Calendar</NuxtLink>
            </div>
            <div v-if="bookings.length" class="divide-y divide-dach-line">
              <div v-for="booking in bookings.slice(0, 5)" :key="booking.id" class="px-6 py-4">
                <p class="font-semibold">{{ booking.customerName }}</p>
                <p class="mt-1 text-sm text-dach-muted">{{ booking.bookingDate || 'Date pending' }} · {{ booking.pickupPostcode }} -> {{ booking.deliveryPostcode }}</p>
              </div>
            </div>
            <div v-else class="grid h-72 place-items-center p-8 text-center text-dach-muted">
              <div>
                <FontAwesomeIcon icon="calendar-days" class="mb-4 text-4xl text-dach-orange/70" />
                <p class="font-semibold text-dach-black">No booked moves yet</p>
                <p class="mt-2 max-w-sm text-sm leading-6">Accepted quotes become scheduled jobs for the office and moving crew.</p>
              </div>
            </div>
          </article>
        </div>

        <div class="mt-6 grid gap-6 xl:grid-cols-[0.9fr_1.1fr]">
          <article class="border border-dach-line bg-white p-6 shadow-sm">
            <h2 class="font-google-sans text-xl font-bold">Move Pipeline</h2>
            <div class="mt-6 grid gap-3">
              <div v-for="stage in ['New inquiry', 'Qualified', 'Quote sent', 'Booked', 'Dispatch ready']" :key="stage" class="flex items-center justify-between border border-dach-line px-4 py-3">
                <span class="font-semibold">{{ stage }}</span>
                <span class="text-sm text-dach-muted">{{ stage === 'New inquiry' ? quotes.length : stage === 'Booked' ? bookings.length : 0 }}</span>
              </div>
            </div>
          </article>
          <article class="border border-dach-line bg-[#151515] p-6 text-white shadow-sm">
            <h2 class="font-google-sans text-xl font-bold">Dispatch Checklist</h2>
            <div class="mt-6 grid gap-4 md:grid-cols-2">
              <p class="border border-white/10 bg-white/5 p-4 text-sm text-white/70"><FontAwesomeIcon icon="route" class="mr-2 text-dach-orange" />Confirm pickup, delivery, access, and parking.</p>
              <p class="border border-white/10 bg-white/5 p-4 text-sm text-white/70"><FontAwesomeIcon icon="user-group" class="mr-2 text-dach-orange" />Assign crew size and van type.</p>
              <p class="border border-white/10 bg-white/5 p-4 text-sm text-white/70"><FontAwesomeIcon icon="shield-halved" class="mr-2 text-dach-orange" />Check insurance and fragile item notes.</p>
              <p class="border border-white/10 bg-white/5 p-4 text-sm text-white/70"><FontAwesomeIcon icon="phone" class="mr-2 text-dach-orange" />Send final confirmation to the customer.</p>
            </div>
          </article>
        </div>
      </section>

      <section v-if="activeAdminView === 'quotes'" class="p-9">
        <div class="mb-6 flex items-center justify-between gap-4">
          <div>
            <h2 class="font-google-sans text-2xl font-bold">All Quote Requests</h2>
            <p class="mt-1 text-dach-muted">{{ filteredQuotes.length }} requests found</p>
          </div>
          <input v-model="quoteSearch" class="w-80 border border-dach-line bg-white px-5 py-4 outline-none focus:border-dach-orange" placeholder="Search name or postcode..." />
        </div>
        <div class="overflow-hidden border border-dach-line bg-white shadow-sm">
          <table v-if="filteredQuotes.length" class="w-full text-left text-sm">
            <thead class="bg-[#151515] text-white">
              <tr>
                <th class="px-5 py-4 font-semibold">Customer</th>
                <th class="px-5 py-4 font-semibold">Route</th>
                <th class="px-5 py-4 font-semibold">Service</th>
                <th class="px-5 py-4 font-semibold">Date</th>
                <th class="px-5 py-4 font-semibold">Status</th>
                <th class="px-5 py-4 font-semibold">Action</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-dach-line">
              <tr v-for="quote in filteredQuotes" :key="quote.id" class="hover:bg-dach-cream/60">
                <td class="px-5 py-4">
                  <p class="font-semibold">{{ quote.fullName || 'Website Visitor' }}</p>
                  <p class="text-dach-muted">{{ quote.phoneNumber || quote.emailAddress || 'No contact supplied' }}</p>
                </td>
                <td class="px-5 py-4">{{ quote.pickupPostcode || '-' }} -> {{ quote.deliveryPostcode || '-' }}</td>
                <td class="px-5 py-4">{{ quote.serviceType || '-' }}</td>
                <td class="px-5 py-4">{{ quote.preferredDate || 'Flexible' }}</td>
                <td class="px-5 py-4"><span class="bg-dach-orange/10 px-3 py-1 text-xs font-bold uppercase text-dach-orange">{{ quote.status }}</span></td>
                <td class="px-5 py-4">
                  <button class="rounded-full bg-dach-orange px-4 py-2 text-xs font-bold text-white" type="button" @click="openBookingForm(quote)">Book</button>
                </td>
              </tr>
            </tbody>
          </table>
          <div v-else class="grid h-[440px] place-items-center text-center text-dach-muted">
            <p><FontAwesomeIcon icon="clipboard" class="mb-4 text-4xl text-dach-orange/70" /><br />No quote requests yet.<br />They will appear here when customers submit the form.</p>
          </div>
        </div>
      </section>

      <section v-if="activeAdminView === 'bookings'" class="p-9">
        <div class="mb-6 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <button class="border border-dach-line bg-white px-4 py-3" type="button">&lt;</button>
            <h2 class="font-google-sans text-2xl font-bold">September 2026</h2>
            <button class="border border-dach-line bg-white px-4 py-3" type="button">&gt;</button>
          </div>
          <button class="rounded-full bg-dach-orange px-5 py-4 font-semibold text-white" type="button" @click="openBookingForm()"><FontAwesomeIcon icon="plus" class="mr-2" />Add Booking</button>
        </div>
        <div class="mb-6 grid gap-5 xl:grid-cols-3">
          <article class="border border-dach-line bg-white p-5 shadow-sm">
            <p class="text-sm font-semibold text-dach-muted">Today's jobs</p>
            <p class="mt-2 text-3xl font-bold">{{ todaysJobs.length }}</p>
          </article>
          <article class="border border-dach-line bg-white p-5 shadow-sm">
            <p class="text-sm font-semibold text-dach-muted">Confirmed bookings</p>
            <p class="mt-2 text-3xl font-bold">{{ confirmedBookings.length }}</p>
          </article>
          <article class="border border-dach-line bg-white p-5 shadow-sm">
            <p class="text-sm font-semibold text-dach-muted">Next move</p>
            <p class="mt-2 text-lg font-bold">{{ nextBooking?.customerName || 'No job scheduled' }}</p>
          </article>
        </div>
        <div class="overflow-hidden border border-dach-line bg-white shadow-sm">
          <div class="grid grid-cols-7 border-b border-dach-line bg-[#151515] text-center text-xs font-bold uppercase tracking-[0.14em] text-white">
            <span v-for="day in calendarDays" :key="day" class="py-4">{{ day }}</span>
          </div>
          <div class="grid grid-cols-7">
            <div v-for="(cell, index) in calendarCells" :key="`${cell}-${index}`" class="min-h-32 border-r border-t border-dach-line p-3 text-sm" :class="cell === '7' && index === 7 ? 'bg-dach-orange/5 ring-1 ring-inset ring-dach-orange text-dach-orange' : 'bg-white'">
              <span class="font-semibold">{{ cell }}</span>
              <div v-if="calendarCellBookings(cell, index).length" class="mt-3 grid gap-2">
                <button
                  v-for="booking in calendarCellBookings(cell, index).slice(0, 2)"
                  :key="booking.id"
                  class="rounded-xl bg-green-100 px-3 py-2 text-left text-xs font-semibold text-green-800"
                  type="button"
                >
                  {{ booking.customerName }}<br />
                  <span class="font-normal">{{ booking.pickupPostcode }} -> {{ booking.deliveryPostcode }}</span>
                </button>
              </div>
            </div>
          </div>
        </div>
        <p class="mt-5 text-sm text-dach-muted"><span class="text-dach-orange">■</span> New <span class="ml-6 text-green-600">■</span> Confirmed</p>
      </section>

      <div v-if="showBookingModal" class="fixed inset-0 z-50 grid place-items-center bg-dach-black/65 p-6 backdrop-blur-sm">
        <form class="w-full max-w-3xl rounded-3xl bg-white p-7 shadow-2xl" @submit.prevent="createBooking">
          <div class="mb-6 flex items-start justify-between gap-6 border-b border-dach-line pb-5">
            <div>
              <p class="text-sm font-semibold uppercase tracking-[0.14em] text-dach-orange">New Move Booking</p>
              <h2 class="mt-2 font-google-sans text-3xl font-bold">Schedule a removal job</h2>
              <p class="mt-2 text-dach-muted">Capture the customer, route, service, date, and crew notes.</p>
            </div>
            <button class="rounded-full border border-dach-line px-4 py-2 text-sm font-semibold text-dach-muted" type="button" @click="showBookingModal = false">Close</button>
          </div>

          <div class="grid gap-4 md:grid-cols-2">
            <label class="block">
              <span class="mb-2 block text-sm font-semibold">Customer name</span>
              <input v-model="bookingForm.customerName" required class="w-full rounded-2xl border border-dach-line bg-dach-cream px-4 py-4 outline-none focus:border-dach-orange" placeholder="e.g. Sarah Moore" />
            </label>
            <label class="block">
              <span class="mb-2 block text-sm font-semibold">Phone number</span>
              <input v-model="bookingForm.phoneNumber" class="w-full rounded-2xl border border-dach-line bg-dach-cream px-4 py-4 outline-none focus:border-dach-orange" placeholder="+44..." />
            </label>
            <label class="block">
              <span class="mb-2 block text-sm font-semibold">Pickup postcode</span>
              <input v-model="bookingForm.pickupPostcode" required class="w-full rounded-2xl border border-dach-line bg-dach-cream px-4 py-4 outline-none focus:border-dach-orange" placeholder="SW4 0EX" />
            </label>
            <label class="block">
              <span class="mb-2 block text-sm font-semibold">Delivery postcode</span>
              <input v-model="bookingForm.deliveryPostcode" required class="w-full rounded-2xl border border-dach-line bg-dach-cream px-4 py-4 outline-none focus:border-dach-orange" placeholder="E8 1ND" />
            </label>
            <label class="block">
              <span class="mb-2 block text-sm font-semibold">Service type</span>
              <select v-model="bookingForm.serviceType" class="w-full rounded-2xl border border-dach-line bg-dach-cream px-4 py-4 outline-none focus:border-dach-orange">
                <option>House Move</option>
                <option>Man and Van</option>
                <option>Office Relocation</option>
                <option>Student Move</option>
                <option>Furniture Delivery</option>
                <option>Packing and Storage</option>
              </select>
            </label>
            <label class="block">
              <span class="mb-2 block text-sm font-semibold">Booking date</span>
              <input v-model="bookingForm.bookingDate" required type="date" class="w-full rounded-2xl border border-dach-line bg-dach-cream px-4 py-4 outline-none focus:border-dach-orange" />
            </label>
          </div>

          <label class="mt-4 block">
            <span class="mb-2 block text-sm font-semibold">Crew and access notes</span>
            <textarea v-model="bookingForm.notes" class="min-h-28 w-full rounded-2xl border border-dach-line bg-dach-cream px-4 py-4 outline-none focus:border-dach-orange" placeholder="Stairs, lift, parking, fragile items, van size, crew size..." />
          </label>

          <p v-if="bookingError" class="mt-4 rounded-2xl bg-dach-orange/10 px-4 py-3 text-sm font-semibold text-dach-orange">{{ bookingError }}</p>

          <div class="mt-6 flex items-center justify-end gap-3">
            <button class="rounded-full border border-dach-line px-5 py-3 font-semibold text-dach-muted" type="button" @click="showBookingModal = false">Cancel</button>
            <button class="rounded-full bg-dach-orange px-6 py-3 font-semibold text-white disabled:opacity-60" type="submit" :disabled="bookingSaving">
              {{ bookingSaving ? 'Saving...' : 'Save Booking' }}
            </button>
          </div>
        </form>
      </div>

      <section v-if="activeAdminView === 'messages'" class="p-9">
        <div class="mb-6 flex items-end justify-between">
          <div>
            <h2 class="font-google-sans text-2xl font-bold">Customer Messages</h2>
            <p class="mt-1 text-dach-muted">{{ unreadMessages.length }} unread</p>
          </div>
        </div>
        <div class="grid min-h-[620px] grid-cols-[380px_1fr] overflow-hidden border border-dach-line bg-white shadow-sm">
          <div class="border-r border-dach-line">
            <div v-if="messages.length" class="divide-y divide-dach-line">
              <button v-for="message in messages" :key="message.id" class="block w-full p-5 text-left hover:bg-dach-cream/60" type="button">
                <p class="font-semibold">{{ message.senderName }}</p>
                <p class="mt-1 text-sm text-dach-muted">{{ message.subject }}</p>
              </button>
            </div>
            <div v-else class="grid h-full place-items-center p-12 text-center text-dach-muted">
              <p><FontAwesomeIcon icon="envelope" class="mb-4 text-4xl text-dach-orange/70" /><br />No messages yet</p>
            </div>
          </div>
          <div class="grid place-items-center p-10 text-center text-dach-muted">
            <div>
              <FontAwesomeIcon icon="eye" class="mb-4 text-4xl text-dach-orange/70" />
              <p>Select a message to read</p>
            </div>
          </div>
        </div>
      </section>

      <section v-if="activeAdminView === 'content'" class="p-9">
        <div class="mb-6">
          <h2 class="font-google-sans text-2xl font-bold">Business Setup</h2>
          <p class="mt-1 text-dach-muted">Keep quote handling, contact details, and operations rules in one place.</p>
        </div>
        <div class="grid grid-cols-[300px_1fr] gap-6">
          <div class="overflow-hidden border border-dach-line bg-white shadow-sm">
            <button
              v-for="item in availableContentSections"
              :key="item.label"
              class="block w-full border-b border-dach-line p-5 text-left transition hover:bg-dach-cream/60"
              :class="selectedContent === item.id ? 'bg-dach-orange text-white hover:bg-dach-orange' : ''"
              type="button"
              @click="selectedContent = item.id"
            >
              <strong>{{ item.label }}</strong>
              <span class="mt-1 block text-sm" :class="selectedContent === item.id ? 'text-white/70' : 'text-dach-muted'">{{ Object.keys(item.fields || {}).length || 0 }} settings</span>
            </button>
          </div>
          <div class="min-h-72 border border-dach-line bg-white p-7 shadow-sm">
            <template v-if="selectedContentSection">
              <div class="mb-6 flex items-center justify-between border-b border-dach-line pb-5">
                <h3 class="font-google-sans text-xl font-bold">{{ selectedContentSection.label }}</h3>
                <button class="bg-dach-orange px-4 py-3 text-sm font-semibold text-white" type="button">Save Changes</button>
              </div>
              <label v-for="(value, key) in selectedContentSection.fields" :key="key" class="mb-5 block">
                <span class="mb-2 block text-xs font-bold uppercase tracking-[0.14em] text-dach-muted">{{ key }}</span>
                <textarea class="min-h-24 w-full border border-dach-line bg-dach-cream px-4 py-3 outline-none focus:border-dach-orange" :value="value" />
              </label>
            </template>
            <p v-else class="text-dach-muted">Select a section on the left to edit.</p>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>
