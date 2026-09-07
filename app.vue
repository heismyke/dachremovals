<script setup lang="ts">
import { Motion } from 'motion-v'
import { adminNavItems, calendarCells, calendarDays, dashboardCards, editableSections } from './data/admin'
import { areas, brand, comparison, faqs, navigation, quoteSteps, resources, routePairs, services, testimonials, trustBadges } from './data/content'

const route = useRoute()
const featuredServices = services.slice(0, 2)
const openFaq = ref(0)
const isAdminRoute = computed(() => route.path.startsWith('/admin'))
const activeAdminView = computed(() => {
  if (route.path.includes('/quotes')) return 'quotes'
  if (route.path.includes('/bookings')) return 'bookings'
  if (route.path.includes('/messages')) return 'messages'
  if (route.path.includes('/content')) return 'content'
  if (route.path.includes('/login')) return 'login'
  return 'dashboard'
})
</script>

<template>
  <div class="hidden">
    <NuxtPage />
  </div>

  <main v-if="!isAdminRoute" class="min-h-screen bg-white text-dach-black">
    <header class="sticky top-0 z-50 border-b border-dach-line bg-white/95 backdrop-blur">
      <div class="section-wrap flex h-20 items-center justify-between">
        <a href="#" class="flex items-center gap-3">
          <span class="grid h-11 w-11 place-items-center bg-dach-orange text-xl font-black text-dach-black">D</span>
          <span class="leading-none">
            <span class="display-title block text-xl font-extrabold">ACH</span>
            <span class="block text-[10px] font-bold uppercase tracking-[0.42em] text-dach-orange">Removals</span>
          </span>
        </a>

        <nav class="hidden items-center gap-8 text-sm font-semibold lg:flex">
          <a v-for="item in navigation" :key="item.href" :href="item.href" class="transition hover:text-dach-orange">
            {{ item.label }}
          </a>
        </nav>

        <div class="flex items-center gap-3">
          <a :href="`tel:${brand.phone}`" class="hidden items-center gap-3 bg-dach-cream px-5 py-3 text-sm font-semibold md:flex">
            <FontAwesomeIcon icon="phone" />
            {{ brand.phone }}
          </a>
          <a href="/admin" class="bg-dach-orange px-5 py-3 text-sm font-semibold text-white transition hover:bg-dach-black">Log In</a>
        </div>
      </div>
    </header>

    <section class="relative overflow-hidden bg-dach-black text-white">
      <img src="/images/hero-removals.png" alt="Dach Removals crew loading a van" class="absolute inset-0 h-full w-full object-cover opacity-35" />
      <div class="absolute inset-0 bg-gradient-to-r from-dach-black via-dach-black/80 to-dach-black/25" />

      <div class="section-wrap relative grid min-h-[720px] items-center gap-12 py-20 lg:grid-cols-[1fr_470px]">
        <Motion
          as="div"
          :initial="{ opacity: 0, y: 24 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.55, ease: 'easeOut' }"
        >
          <div class="mb-8 flex flex-wrap gap-3">
            <span v-for="badge in trustBadges" :key="badge" class="bg-white/12 px-4 py-2 text-sm font-medium backdrop-blur">{{ badge }}</span>
          </div>
          <h1 class="max-w-3xl text-5xl font-extrabold leading-[1.02] tracking-tight md:text-7xl">
            Professional UK removals, <span class="text-dach-orange">made simple.</span>
          </h1>
          <p class="mt-6 max-w-2xl text-lg leading-8 text-white/82">
            Fast quotes, careful movers, and reliable man and van support for homes, offices, student moves, furniture deliveries, packing, and storage.
          </p>
          <div class="mt-9 flex flex-wrap gap-4">
            <a href="#quote" class="bg-dach-orange px-6 py-4 font-semibold text-white transition hover:bg-white hover:text-dach-black">
              Get a Quote <FontAwesomeIcon icon="arrow-right" class="ml-2" />
            </a>
            <a :href="`tel:${brand.phone}`" class="border border-white/30 px-6 py-4 font-semibold text-white transition hover:bg-white hover:text-dach-black">
              <FontAwesomeIcon icon="phone" class="mr-2" />{{ brand.phone }}
            </a>
          </div>
        </Motion>

        <Motion
          id="quote"
          as="form"
          class="bg-white p-7 text-dach-black shadow-2xl shadow-black/30"
          :initial="{ opacity: 0, y: 26 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.55, ease: 'easeOut', delay: 0.08 }"
        >
          <h2 class="text-2xl font-bold tracking-tight">Get Your Price Instantly</h2>
          <p class="mt-2 text-sm text-dach-muted">No waiting around. Send the details and we confirm clearly.</p>

          <label class="mt-6 block text-sm font-semibold">Moving From</label>
          <div class="mt-2 flex items-center gap-3 bg-dach-cream px-4 py-4 text-dach-muted">
            <FontAwesomeIcon icon="location-dot" />
            <input class="w-full bg-transparent outline-none" placeholder="Enter postcode" />
          </div>

          <label class="mt-4 block text-sm font-semibold">Moving To</label>
          <div class="mt-2 flex items-center gap-3 bg-dach-cream px-4 py-4 text-dach-muted">
            <FontAwesomeIcon icon="route" />
            <input class="w-full bg-transparent outline-none" placeholder="Enter postcode" />
          </div>

          <div class="mt-4 grid gap-3 md:grid-cols-2">
            <select class="bg-dach-cream px-4 py-4 text-dach-muted outline-none">
              <option>House Move</option>
              <option>Man and Van</option>
              <option>Office Relocation</option>
              <option>Packing and Storage</option>
            </select>
            <input class="bg-dach-cream px-4 py-4 text-dach-muted outline-none" placeholder="dd/mm/yyyy" />
          </div>

          <button class="mt-6 w-full bg-dach-orange px-6 py-4 font-semibold text-white transition hover:bg-dach-black">
            See Price Instantly <FontAwesomeIcon icon="arrow-right" class="ml-2" />
          </button>
          <p class="mt-4 text-center text-sm text-dach-muted"><FontAwesomeIcon icon="lock" class="mr-1" /> Secure enquiry. Rated 4.8/5.</p>
        </Motion>
      </div>
    </section>

    <section id="how-it-works" class="section-wrap py-24">
      <p class="text-sm font-semibold uppercase tracking-[0.2em] text-dach-orange">Fast quote process</p>
      <div class="mt-4 grid gap-6 lg:grid-cols-[0.75fr_1fr] lg:items-end">
        <h2 class="text-4xl font-bold tracking-tight md:text-5xl">Get a clear removals quote without the back and forth.</h2>
        <p class="text-lg leading-8 text-dach-muted">The quote flow keeps the important details simple: locations, move type, date, access, and the level of help required.</p>
      </div>
      <div class="mt-12 grid gap-5 md:grid-cols-5">
        <Motion
          v-for="(step, index) in quoteSteps"
          :key="step.title"
          as="article"
          class="border border-dach-line bg-white p-6 shadow-sm"
          :whileHover="{ y: -5 }"
          :transition="{ duration: 0.2, ease: 'easeOut' }"
        >
          <span class="mb-7 flex h-11 w-11 items-center justify-center bg-dach-orange text-white"><FontAwesomeIcon :icon="step.icon" /></span>
          <span class="text-sm font-semibold text-dach-muted">0{{ index + 1 }}</span>
          <h3 class="mt-3 text-lg font-bold">{{ step.title }}</h3>
          <p class="mt-3 text-sm leading-6 text-dach-muted">{{ step.copy }}</p>
        </Motion>
      </div>
    </section>

    <section id="services" class="border-y border-dach-line bg-dach-cream py-24">
      <div class="section-wrap">
        <p class="text-sm font-semibold uppercase tracking-[0.2em] text-dach-orange">Services</p>
        <div class="mt-4 grid gap-6 lg:grid-cols-[0.7fr_1fr] lg:items-end">
          <h2 class="text-4xl font-bold tracking-tight md:text-5xl">Removal services for every type of move.</h2>
          <p class="text-lg leading-8 text-dach-muted">From one item to a full home, Dach Removals gives you the right crew, vehicle, and support for the job.</p>
        </div>

        <div class="mt-12 grid gap-6 lg:grid-cols-2">
          <article v-for="service in featuredServices" :key="service.title" class="bg-white shadow-sm">
            <img :src="service.image" :alt="service.title" class="h-64 w-full object-cover" />
            <div class="p-7">
              <h3 class="text-2xl font-bold">{{ service.title }}</h3>
              <p class="mt-3 leading-7 text-dach-muted">{{ service.description }}</p>
              <ul class="mt-5 grid gap-2 text-sm text-dach-muted">
                <li v-for="item in service.idealFor" :key="item"><FontAwesomeIcon icon="check" class="mr-2 text-dach-orange" />{{ item }}</li>
              </ul>
              <p class="mt-5 leading-7">{{ service.detail }}</p>
              <div class="mt-6 flex gap-3">
                <a href="#quote" class="bg-dach-orange px-5 py-3 font-semibold text-white">Get Quote</a>
                <a href="#faq" class="border border-dach-line px-5 py-3 font-semibold">Learn More</a>
              </div>
            </div>
          </article>
        </div>

        <div class="mt-8 grid gap-5 md:grid-cols-3">
          <article v-for="service in services.slice(2)" :key="service.title" class="bg-white p-7 shadow-sm">
            <span class="mb-6 flex h-11 w-11 items-center justify-center bg-dach-orange text-white"><FontAwesomeIcon :icon="service.icon" /></span>
            <h3 class="text-xl font-bold">{{ service.title }}</h3>
            <p class="mt-3 leading-7 text-dach-muted">{{ service.description }}</p>
            <a href="#quote" class="mt-6 inline-flex items-center gap-2 font-semibold text-dach-orange">Get Quote <FontAwesomeIcon icon="arrow-right" /></a>
          </article>
        </div>
      </div>
    </section>

    <section class="section-wrap py-24">
      <p class="text-sm font-semibold uppercase tracking-[0.2em] text-dach-orange">Why choose Dach</p>
      <div class="mt-4 grid gap-10 lg:grid-cols-[0.8fr_1.2fr] lg:items-start">
        <div>
          <h2 class="text-4xl font-bold tracking-tight md:text-5xl">Reliable movers, clear communication, prepared crews.</h2>
          <p class="mt-5 leading-8 text-dach-muted">A professional removals page should make the offer obvious. This one focuses on practical proof, direct contact, and simple booking actions.</p>
        </div>
        <div class="grid gap-4 md:grid-cols-2">
          <article v-for="[label, ours, traditional] in comparison" :key="label" class="border border-dach-line p-6">
            <p class="text-sm font-semibold uppercase tracking-[0.16em] text-dach-orange">{{ label }}</p>
            <p class="mt-4 font-semibold"><FontAwesomeIcon icon="check" class="mr-2 text-dach-orange" />{{ ours }}</p>
            <p class="mt-2 text-sm text-dach-muted">Typical issue: {{ traditional }}</p>
          </article>
        </div>
      </div>
    </section>

    <section id="areas" class="border-y border-dach-line bg-gray-50 py-24">
      <div class="section-wrap grid gap-10 lg:grid-cols-[0.75fr_1.25fr]">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.2em] text-dach-orange">UK coverage</p>
          <h2 class="mt-4 text-4xl font-bold tracking-tight md:text-5xl">Areas we cover across the UK.</h2>
          <p class="mt-5 leading-8 text-dach-muted">Dach Removals supports local and long-distance moves across major UK regions.</p>
          <div class="mt-8 flex flex-wrap gap-3">
            <span v-for="routePair in routePairs" :key="routePair" class="bg-white px-4 py-2 text-sm font-semibold shadow-sm">{{ routePair }}</span>
          </div>
        </div>
        <div class="grid gap-5 md:grid-cols-2">
          <article v-for="[area, cities] in areas" :key="area" class="bg-white p-6 shadow-sm">
            <h3 class="text-lg font-bold"><FontAwesomeIcon icon="route" class="mr-2 text-dach-orange" />{{ area }}</h3>
            <ul class="mt-4 grid gap-2 text-sm text-dach-muted">
              <li v-for="city in cities" :key="city">-> {{ city }}</li>
            </ul>
          </article>
        </div>
      </div>
    </section>

    <section id="reviews" class="section-wrap py-24">
      <div class="flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.2em] text-dach-orange">Reviews</p>
          <h2 class="mt-4 text-4xl font-bold tracking-tight md:text-5xl">What customers say.</h2>
        </div>
        <p class="text-dach-muted"><FontAwesomeIcon icon="star" class="text-dach-orange" /> 4.8/5 from 2,500+ reviews</p>
      </div>
      <div class="mt-10 grid gap-5 md:grid-cols-4">
        <article v-for="[name, city, quote] in testimonials" :key="name" class="border border-dach-line p-6">
          <p class="text-dach-orange">*****</p>
          <p class="mt-5 leading-7">"{{ quote }}"</p>
          <p class="mt-8 font-bold">{{ name }}</p>
          <p class="text-sm text-dach-muted">{{ city }}</p>
        </article>
      </div>
    </section>

    <section class="bg-dach-cream py-24">
      <div class="section-wrap">
        <p class="text-sm font-semibold uppercase tracking-[0.2em] text-dach-orange">Guides</p>
        <h2 class="mt-4 text-4xl font-bold tracking-tight md:text-5xl">Helpful moving resources.</h2>
        <div class="mt-10 grid gap-5 md:grid-cols-3">
          <article v-for="[icon, title, copy] in resources" :key="title" class="bg-white p-6 shadow-sm">
            <div class="mb-6 flex items-center justify-between border-b border-dach-line pb-6">
              <span class="flex h-12 w-12 items-center justify-center bg-dach-orange text-white">
                <FontAwesomeIcon :icon="icon" />
              </span>
              <span class="text-sm font-semibold uppercase tracking-[0.16em] text-dach-muted">Guide</span>
            </div>
            <h3 class="text-xl font-bold">{{ title }}</h3>
            <p class="mt-3 leading-7 text-dach-muted">{{ copy }}</p>
            <a href="#quote" class="mt-6 inline-flex items-center gap-2 font-semibold text-dach-orange">
              Get advice <FontAwesomeIcon icon="arrow-right" />
            </a>
          </article>
        </div>
      </div>
    </section>

    <section id="faq" class="section-wrap grid gap-10 py-24 lg:grid-cols-[0.7fr_1.3fr]">
      <div>
        <p class="text-sm font-semibold uppercase tracking-[0.2em] text-dach-orange">FAQ</p>
        <h2 class="mt-4 text-4xl font-bold tracking-tight md:text-5xl">Frequently asked questions.</h2>
        <p class="mt-5 leading-8 text-dach-muted">Answers to the common questions people ask before booking a removals crew.</p>
        <p class="mt-8 font-semibold">Need help? <a :href="`mailto:${brand.email}`" class="text-dach-orange">{{ brand.email }}</a></p>
      </div>
      <div class="space-y-3">
        <article v-for="(faq, index) in faqs" :key="faq" class="border border-dach-line">
          <button class="flex w-full items-center justify-between p-5 text-left font-semibold" @click="openFaq = openFaq === index ? -1 : index">
            {{ faq }}
            <FontAwesomeIcon :icon="openFaq === index ? 'chevron-down' : 'arrow-right'" class="text-dach-orange" />
          </button>
          <p v-if="openFaq === index" class="px-5 pb-5 leading-7 text-dach-muted">
            Send your details through the quote form or call us directly. We will confirm the right service, availability, and next steps for your move.
          </p>
        </article>
      </div>
    </section>

    <section class="bg-dach-black py-20 text-white">
      <div class="section-wrap grid gap-8 lg:grid-cols-[1fr_auto] lg:items-center">
        <div>
          <p class="text-sm font-semibold uppercase tracking-[0.2em] text-dach-orange">Get started</p>
          <h2 class="mt-4 text-4xl font-bold tracking-tight md:text-5xl">Ready to plan your move?</h2>
          <p class="mt-4 max-w-2xl leading-8 text-white/70">Send the details once. We will come back with clear pricing, availability, and the best crew for the job.</p>
        </div>
        <div class="flex flex-wrap gap-3">
          <a href="#quote" class="bg-dach-orange px-6 py-4 font-semibold text-white">Get a Quote</a>
          <a :href="`tel:${brand.phone}`" class="border border-white/20 px-6 py-4 font-semibold text-white"><FontAwesomeIcon icon="phone" class="mr-2" />{{ brand.phone }}</a>
        </div>
      </div>
    </section>

    <footer class="bg-[#101010] py-14 text-white">
      <div class="section-wrap grid gap-10 md:grid-cols-4">
        <div>
          <div class="mb-5 flex items-center gap-3">
            <span class="grid h-10 w-10 place-items-center bg-dach-orange font-black text-dach-black">D</span>
            <span>
              <span class="block text-lg font-extrabold">ACH</span>
              <span class="block text-[10px] font-bold uppercase tracking-[0.35em] text-dach-orange">Removals</span>
            </span>
          </div>
          <p class="text-white/60">Professional removals across the UK. Reliable crews, direct support, and clear pricing.</p>
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

  <div v-else-if="activeAdminView === 'login'" class="grid min-h-screen place-items-center bg-dach-black">
    <form class="w-[420px] border-t-4 border-dach-orange bg-white p-12 shadow-2xl">
      <div class="mb-8 flex items-center gap-3">
        <span class="grid h-12 w-12 place-items-center bg-dach-orange text-2xl font-black">D</span>
        <span>
          <span class="display-title block text-2xl font-extrabold">ACH</span>
          <span class="block text-[10px] font-bold uppercase tracking-[0.45em] text-dach-orange">Removals</span>
        </span>
      </div>
      <p class="mb-8 text-sm font-bold uppercase tracking-[0.25em] text-dach-muted">Admin Portal</p>
      <label class="text-xs font-bold uppercase tracking-[0.2em] text-dach-muted">Username</label>
      <input value="admin" class="mt-2 w-full border border-dach-line bg-dach-cream px-5 py-4 outline-none" />
      <label class="mt-5 block text-xs font-bold uppercase tracking-[0.2em] text-dach-muted">Password</label>
      <input type="password" value="password" class="mt-2 w-full border border-dach-line bg-dach-cream px-5 py-4 outline-none" />
      <NuxtLink to="/admin/dashboard" class="mt-6 block bg-dach-orange px-6 py-4 text-center font-black uppercase text-white">
        Sign In <FontAwesomeIcon icon="arrow-right" />
      </NuxtLink>
      <a href="/" class="mt-8 block text-center text-sm text-dach-muted"><FontAwesomeIcon icon="arrow-left" /> Back to website</a>
    </form>
  </div>

  <div v-else class="grid min-h-screen grid-cols-[300px_1fr] bg-dach-cream">
    <aside class="flex flex-col bg-dach-black text-white">
      <div class="flex h-28 items-center gap-3 border-b border-white/10 px-7">
        <span class="grid h-12 w-12 place-items-center bg-dach-orange text-2xl font-black text-dach-black">D</span>
        <span>
          <span class="display-title block text-2xl font-extrabold">ACH</span>
          <span class="block text-[10px] font-bold uppercase tracking-[0.45em] text-dach-orange">Removals</span>
          <span class="block text-[11px] uppercase tracking-[0.25em] text-white/40">Admin Panel</span>
        </span>
      </div>

      <nav class="mt-4">
        <NuxtLink
          v-for="item in adminNavItems"
          :key="item.label"
          :to="item.path"
          class="flex items-center gap-4 px-8 py-5 text-white/45"
          :class="route.path === item.path ? 'border-l-4 border-dach-orange bg-dach-orange/20 text-dach-orange' : 'hover:text-white'"
        >
          <FontAwesomeIcon :icon="item.icon" />
          <span>{{ item.label }}</span>
          <span v-if="item.badge !== undefined" class="ml-auto rounded-full bg-dach-orange px-3 py-1 text-xs font-black text-white">{{ item.badge }}</span>
        </NuxtLink>
      </nav>

      <div class="mt-auto border-t border-white/10 p-7 text-white/45">
        <a href="/" class="mb-4 block"><FontAwesomeIcon icon="globe" class="mr-2" />View Website</a>
        <NuxtLink to="/admin/login"><FontAwesomeIcon icon="right-from-bracket" class="mr-2" />Sign Out</NuxtLink>
      </div>
    </aside>

    <main>
      <header class="flex h-20 items-center justify-between bg-white px-9">
        <h1 class="display-title text-2xl font-extrabold">
          {{ activeAdminView === 'dashboard' ? 'Dashboard' : activeAdminView === 'quotes' ? 'Quote Requests' : activeAdminView === 'bookings' ? 'Bookings Calendar' : activeAdminView === 'messages' ? 'Messages' : 'Edit Content' }}
        </h1>
        <p class="text-dach-muted">Mon, 7 September 2026 <span class="ml-3 text-green-600">Live</span></p>
      </header>

      <section v-if="activeAdminView === 'dashboard'" class="p-9">
        <div class="grid gap-6 xl:grid-cols-4">
          <article v-for="card in dashboardCards" :key="card.label" class="border border-dach-line bg-white p-8">
            <p class="text-sm font-bold uppercase tracking-[0.2em] text-dach-muted">{{ card.label }}</p>
            <p class="display-title mt-3 text-5xl font-extrabold">{{ card.value }}</p>
            <p :class="card.tone" class="mt-3">{{ card.note }}</p>
          </article>
        </div>
        <div class="mt-9 grid gap-8 xl:grid-cols-2">
          <article>
            <div class="mb-6 flex items-center justify-between">
              <h2 class="display-title text-2xl font-extrabold">Recent Quotes</h2>
              <NuxtLink to="/admin/quotes" class="border border-dach-line bg-white px-5 py-3 font-bold">View All</NuxtLink>
            </div>
            <div class="grid h-80 place-items-center text-dach-muted">
              <p class="text-center text-lg"><FontAwesomeIcon icon="clipboard" class="mb-4 text-4xl" /><br />No quotes yet</p>
            </div>
          </article>
          <article>
            <div class="mb-6 flex items-center justify-between">
              <h2 class="display-title text-2xl font-extrabold">Upcoming Bookings</h2>
              <NuxtLink to="/admin/bookings" class="border border-dach-line bg-white px-5 py-3 font-bold">Calendar</NuxtLink>
            </div>
            <div class="grid h-80 place-items-center text-dach-muted">
              <p class="text-center text-lg"><FontAwesomeIcon icon="calendar-days" class="mb-4 text-4xl" /><br />No upcoming bookings</p>
            </div>
          </article>
        </div>
      </section>

      <section v-if="activeAdminView === 'quotes'" class="p-9">
        <div class="flex items-center justify-between">
          <h2 class="display-title text-2xl font-extrabold">All Quote Requests</h2>
          <input class="border border-dach-line bg-white px-5 py-4" placeholder="Search name or postcode..." />
        </div>
        <div class="grid h-[520px] place-items-center text-center text-dach-muted">
          <p><FontAwesomeIcon icon="clipboard" class="mb-4 text-4xl" /><br />No quote requests yet.<br />They'll appear here when customers fill in the form on your website.</p>
        </div>
      </section>

      <section v-if="activeAdminView === 'bookings'" class="p-9">
        <div class="mb-8 flex items-center justify-between">
          <div class="flex items-center gap-4">
            <button class="border bg-white px-4 py-3">&lt;</button>
            <h2 class="display-title text-2xl font-extrabold">September 2026</h2>
            <button class="border bg-white px-4 py-3">&gt;</button>
          </div>
          <button class="bg-dach-orange px-6 py-4 font-black text-white"><FontAwesomeIcon icon="plus" /> Add Booking</button>
        </div>
        <div class="grid grid-cols-7 text-center text-sm font-bold text-dach-muted">
          <span v-for="day in calendarDays" :key="day">{{ day }}</span>
        </div>
        <div class="mt-4 grid grid-cols-7">
          <div v-for="cell in calendarCells" :key="cell" class="h-28 border border-dach-line bg-white p-3" :class="cell === '7' ? 'border-dach-orange text-dach-orange' : ''">
            {{ cell }}
          </div>
        </div>
        <p class="mt-5 text-dach-muted"><span class="text-dach-orange">■</span> New <span class="ml-6 text-green-600">■</span> Confirmed</p>
      </section>

      <section v-if="activeAdminView === 'messages'" class="p-9">
        <h2 class="display-title mb-6 text-2xl font-extrabold">Customer Messages</h2>
        <div class="grid min-h-[620px] grid-cols-[360px_1fr] border border-dach-line bg-white">
          <div class="grid place-items-start border-r border-dach-line p-20 text-center text-dach-muted">
            <FontAwesomeIcon icon="envelope" class="mb-4 text-4xl" />No messages yet
          </div>
          <div class="grid place-items-center text-dach-muted">Select a message to read</div>
        </div>
      </section>

      <section v-if="activeAdminView === 'content'" class="p-9">
        <h2 class="display-title mb-6 text-2xl font-extrabold">Edit Website Content</h2>
        <div class="grid grid-cols-[280px_1fr] gap-8">
          <div class="border border-dach-line bg-white">
            <button v-for="item in editableSections" :key="item.label" class="block w-full border-b border-dach-line p-6 text-left">
              <strong>{{ item.label }}</strong>
              <span class="block text-sm text-dach-muted">{{ item.fields }} fields</span>
            </button>
          </div>
          <div class="min-h-56 border border-dach-line bg-white p-8 text-dach-muted">Select a section on the left to edit</div>
        </div>
      </section>
    </main>
  </div>
</template>
