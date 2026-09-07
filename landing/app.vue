<script setup lang="ts">
import { Motion } from 'motion-v'
import {
  appFeatures,
  areas,
  brand,
  comparison,
  faqs,
  navigation,
  quoteSteps,
  resources,
  routePairs,
  services,
  testimonials,
  trustBadges,
} from './data/content'

const featuredServices = services.slice(0, 2)
const hubServices = services.slice(0, 5)
const openFaq = ref(1)
</script>

<template>
  <main class="mx-auto min-h-screen max-w-[1560px] overflow-hidden rounded-[24px] bg-white shadow-2xl shadow-orange-200/40">
    <nav class="sticky top-0 z-50 flex h-24 items-center justify-between border-b border-dach-line bg-white/95 px-8 backdrop-blur">
      <a href="#" class="flex items-center gap-3">
        <span class="grid h-12 w-12 place-items-center bg-dach-orange text-2xl font-black text-dach-black">D</span>
        <span class="leading-none">
          <span class="display-title block text-2xl">ACH</span>
          <span class="block text-[10px] font-bold uppercase tracking-[0.45em] text-dach-orange">Removals</span>
        </span>
      </a>

      <div class="hidden items-center gap-8 text-sm font-semibold lg:flex">
        <a v-for="item in navigation" :key="item.href" :href="item.href" class="hover:text-dach-orange">
          {{ item.label }}
        </a>
      </div>

      <div class="flex items-center gap-3">
        <a :href="`tel:${brand.phone}`" class="hidden items-center gap-2 rounded-xl bg-dach-cream px-5 py-4 text-sm font-bold md:flex">
          <FontAwesomeIcon icon="phone" />
          {{ brand.phone }}
        </a>
        <a href="/admin" class="rounded-xl bg-dach-orange px-6 py-4 text-sm font-bold text-white">Log In</a>
      </div>
    </nav>

    <section class="relative min-h-[760px] overflow-hidden">
      <img
        src="https://images.unsplash.com/photo-1600518464441-9306b65d6d3c?auto=format&fit=crop&w=2200&q=85"
        alt="Dach Removals crew loading a van"
        class="absolute inset-0 h-full w-full object-cover"
      />
      <div class="absolute inset-0 bg-gradient-to-t from-dach-black via-dach-black/35 to-transparent" />

      <div class="relative z-10 grid min-h-[760px] items-end gap-10 px-10 pb-20 lg:grid-cols-[1.1fr_0.9fr]">
        <Motion
          as="div"
          class="max-w-3xl text-white"
          :initial="{ opacity: 0, y: 28 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.65, ease: 'easeOut' }"
        >
          <div class="mb-7 flex flex-wrap gap-3">
            <span class="rounded-xl bg-dach-black/70 px-4 py-2 text-sm"><FontAwesomeIcon icon="star" class="mr-2 text-dach-orange" />4.8/5</span>
            <span class="rounded-xl bg-dach-black/70 px-4 py-2 text-sm">BAR Registered</span>
            <span class="rounded-xl bg-dach-black/70 px-4 py-2 text-sm"><FontAwesomeIcon icon="shield-halved" class="mr-2" />Fully Insured</span>
            <span class="rounded-xl bg-dach-black/70 px-4 py-2 text-sm">UK-Wide</span>
          </div>
          <h1 class="display-title text-6xl leading-[0.95] md:text-8xl">
            Giant Van - UK<br />Removals, <span class="text-dach-orange">Simplified!</span>
          </h1>
          <p class="mt-7 max-w-xl text-xl text-white/90">
            Moving home? Need a man with van? Looking for removal companies near me? Dach Removals makes it happen in 60 seconds.
          </p>
        </Motion>

        <Motion
          id="quote"
          as="form"
          class="rounded-[28px] bg-white p-7 shadow-2xl lg:mb-[-120px]"
          :initial="{ opacity: 0, y: 34, scale: 0.97 }"
          :animate="{ opacity: 1, y: 0, scale: 1 }"
          :transition="{ duration: 0.7, ease: 'easeOut', delay: 0.12 }"
        >
          <h2 class="text-2xl font-black">Get Your Price Instantly</h2>
          <p class="mt-2 text-sm text-dach-muted">No phone calls. No waiting. Instant fixed price.</p>
          <label class="mt-6 block text-sm font-bold">Moving From</label>
          <div class="mt-2 flex items-center gap-3 rounded-xl bg-dach-cream px-4 py-4 text-dach-muted">
            <FontAwesomeIcon icon="location-dot" />
            <input class="w-full bg-transparent outline-none" placeholder="Enter postcode" />
          </div>
          <label class="mt-4 block text-sm font-bold">Moving To</label>
          <div class="mt-2 flex items-center gap-3 rounded-xl bg-dach-cream px-4 py-4 text-dach-muted">
            <FontAwesomeIcon icon="route" />
            <input class="w-full bg-transparent outline-none" placeholder="Enter postcode" />
          </div>
          <div class="mt-4 grid gap-3 md:grid-cols-2">
            <select class="rounded-xl bg-dach-cream px-4 py-4 text-dach-muted outline-none">
              <option>House Move</option>
              <option>Man and Van</option>
              <option>Office Relocation</option>
            </select>
            <input class="rounded-xl bg-dach-cream px-4 py-4 text-dach-muted outline-none" placeholder="dd/mm/yyyy" />
          </div>
          <button class="mt-6 w-full rounded-xl bg-dach-orange px-6 py-5 font-black text-white transition hover:-translate-y-1">
            See Price Instantly <FontAwesomeIcon icon="arrow-right" class="ml-2" />
          </button>
          <div class="mt-4 grid grid-cols-2 gap-3">
            <button type="button" class="rounded-lg bg-black px-4 py-3 text-sm font-bold text-white"><FontAwesomeIcon :icon="['fab', 'google-play']" class="mr-2" />Google Play</button>
            <button type="button" class="rounded-lg bg-black px-4 py-3 text-sm font-bold text-white"><FontAwesomeIcon :icon="['fab', 'apple']" class="mr-2" />App Store</button>
          </div>
          <p class="mt-5 text-center text-sm text-dach-muted"><FontAwesomeIcon icon="lock" /> Secure checkout - Rated 4.8/5</p>
        </Motion>
      </div>
    </section>

    <section id="how-it-works" class="px-10 py-24">
      <span class="rounded-xl border border-dach-orange/30 bg-orange-50 px-4 py-2 text-sm font-semibold text-dach-orange">Mobile Exclusive</span>
      <h2 class="mt-6 text-5xl font-black tracking-tight text-dach-black">Get Your Quote in <em class="font-normal text-dach-orange">60 Seconds</em></h2>
      <p class="mt-5 max-w-2xl text-lg text-dach-muted">Traditional removal companies make you wait. Phone calls. Callbacks. Home visits. Days before you get a price.</p>
      <div class="mt-14 grid gap-6 md:grid-cols-5">
        <Motion
          v-for="(step, index) in quoteSteps"
          :key="step.title"
          as="article"
          class="rounded-3xl bg-white p-7 shadow-xl shadow-orange-100/60"
          :class="index % 2 ? 'bg-orange-50/70' : 'bg-gray-50'"
          :whileHover="{ y: -8 }"
          :transition="{ duration: 0.22, ease: 'easeOut' }"
        >
          <span class="grid h-14 w-14 place-items-center rounded-xl bg-dach-orange text-white shadow-lg shadow-orange-300"><FontAwesomeIcon :icon="step.icon" /></span>
          <h3 class="mt-8 text-xl font-black">{{ step.title }}</h3>
          <p class="mt-4 text-dach-muted">{{ step.copy }}</p>
        </Motion>
      </div>
    </section>

    <section class="mx-10 grid gap-10 rounded-[32px] bg-gray-50 p-12 lg:grid-cols-2">
      <div>
        <span class="rounded-xl border border-dach-orange/30 bg-white px-4 py-2 text-sm font-semibold text-dach-orange">Mobile Exclusive</span>
        <h2 class="mt-7 text-5xl font-black text-dach-black">Download<br />the <em class="font-normal text-dach-orange">Dach App</em></h2>
        <p class="mt-6 text-lg text-dach-muted">The easiest way to book removals. Our app puts everything at your fingertips.</p>
        <div class="mt-10 grid gap-7 md:grid-cols-2">
          <div v-for="[icon, title, copy] in appFeatures" :key="title" class="flex gap-4">
            <span class="grid h-11 w-11 shrink-0 place-items-center rounded-xl bg-white text-dach-orange shadow"><FontAwesomeIcon :icon="icon" /></span>
            <p><strong class="block text-lg">{{ title }}</strong><span class="text-dach-muted">{{ copy }}</span></p>
          </div>
        </div>
      </div>
      <div class="relative min-h-[520px]">
        <div class="absolute right-10 top-0 h-[520px] w-[300px] rounded-[42px] border-[14px] border-dach-black bg-white p-6 shadow-2xl">
          <div class="mx-auto mb-6 h-7 w-28 rounded-full bg-dach-black" />
          <div class="rounded-2xl border border-orange-100 p-5 shadow">
            <p class="text-sm text-dach-muted">TOTAL</p>
            <p class="text-2xl font-black">£52.00</p>
            <div class="mt-3 h-2 rounded-full bg-dach-orange" />
          </div>
          <div class="mt-5 rounded-2xl bg-orange-50 p-5">
            <strong>£10 Off Your First Booking</strong>
            <p class="mt-2 text-sm text-dach-muted">Available on iOS and Android.</p>
          </div>
        </div>
      </div>
    </section>

    <section id="services" class="px-10 py-24">
      <div class="text-center">
        <span class="rounded-xl border border-dach-orange/30 bg-orange-50 px-4 py-2 text-sm font-semibold text-dach-orange">Services</span>
        <h2 class="mt-7 text-5xl font-black">Our Removal <em class="font-normal text-dach-orange">Services</em></h2>
        <p class="mx-auto mt-5 max-w-4xl text-lg text-dach-muted">Dach Removals offers a complete range of removal and delivery services across the United Kingdom.</p>
      </div>
      <div class="mt-14 grid gap-8 lg:grid-cols-2">
        <article v-for="service in featuredServices" :key="service.title" class="overflow-hidden rounded-3xl bg-white shadow-xl shadow-orange-100/60">
          <img :src="service.image" :alt="service.title" class="h-72 w-full object-cover" />
          <div class="p-7">
            <h3 class="text-2xl font-black">{{ service.title }}</h3>
            <p class="mt-3 text-dach-muted">{{ service.description }}</p>
            <p class="mt-5 font-black">Ideal for:</p>
            <ul class="mt-3 space-y-3 text-dach-muted">
              <li v-for="item in service.idealFor" :key="item"><FontAwesomeIcon icon="check" class="mr-2 text-dach-orange" />{{ item }}</li>
            </ul>
            <p class="mt-6">{{ service.detail }}</p>
            <div class="mt-7 grid gap-3 md:grid-cols-2">
              <a href="#quote" class="rounded-xl bg-dach-orange px-5 py-4 text-center font-black text-white">Get Man and Van Quote</a>
              <a href="#faq" class="rounded-xl border border-dach-line px-5 py-4 text-center font-black">Learn More</a>
            </div>
          </div>
        </article>
      </div>
    </section>

    <section class="bg-gray-50 px-10 py-24">
      <span class="rounded-xl border border-dach-orange/30 bg-white px-4 py-2 text-sm font-semibold text-dach-orange">Our Capabilities</span>
      <div class="mt-6 grid gap-8 lg:grid-cols-2">
        <h2 class="text-5xl font-black">Removals Hub - <em class="font-normal text-dach-orange">All Services</em></h2>
        <p class="text-lg text-dach-muted">Comprehensive moving solutions designed for every need. Select a service to get your instant fixed-price quote.</p>
      </div>
      <div class="mt-12 grid gap-6 md:grid-cols-3">
        <Motion
          v-for="service in hubServices"
          :key="service.title"
          as="article"
          class="rounded-3xl bg-white p-8 shadow-lg shadow-orange-100/50"
          :whileHover="{ y: -6 }"
          :transition="{ duration: 0.22, ease: 'easeOut' }"
        >
          <span class="grid h-14 w-14 place-items-center rounded-xl bg-dach-orange text-white"><FontAwesomeIcon :icon="service.icon" /></span>
          <h3 class="mt-8 text-2xl font-black">{{ service.title }}</h3>
          <p class="mt-3 text-dach-muted">{{ service.description }}</p>
          <a href="#quote" class="mt-8 inline-flex items-center gap-3 rounded-xl border border-dach-line px-5 py-3 font-bold">Get Quote <FontAwesomeIcon icon="arrow-right" /></a>
        </Motion>
      </div>
    </section>

    <section class="mx-10 my-24 rounded-[32px] bg-gray-50 p-10 text-center">
      <span class="rounded-xl border border-dach-orange/30 bg-white px-4 py-2 text-sm font-semibold text-dach-orange">Why Choose Us</span>
      <h2 class="mt-6 text-4xl font-black">Why Choose <em class="font-normal text-dach-orange">Dach Removals</em></h2>
      <div class="mx-auto mt-12 grid max-w-5xl overflow-hidden rounded-3xl bg-white shadow-xl md:grid-cols-3">
        <div class="space-y-7 p-8 text-left font-bold text-dach-muted">
          <p v-for="[label] in comparison" :key="label">{{ label }}</p>
        </div>
        <div class="rounded-3xl border-t-4 border-dach-orange bg-orange-50 p-8 text-left shadow-xl">
          <div class="mb-6 h-14 w-32 bg-dach-orange" />
          <p v-for="[, ours] in comparison" :key="ours" class="mb-6"><FontAwesomeIcon icon="check" class="mr-2 text-dach-orange" />{{ ours }}</p>
          <a href="#quote" class="block rounded-xl bg-dach-orange px-5 py-4 text-center font-black text-white">Get a Quote <FontAwesomeIcon icon="arrow-right" /></a>
        </div>
        <div class="space-y-6 p-8 text-left text-dach-muted">
          <p v-for="[, , traditional] in comparison" :key="traditional">x {{ traditional }}</p>
        </div>
      </div>
      <div class="mt-10 grid gap-5 md:grid-cols-4">
        <div v-for="badge in trustBadges" :key="badge" class="rounded-2xl bg-white p-7 font-bold text-dach-orange shadow">{{ badge }}</div>
      </div>
    </section>

    <section id="areas" class="grid gap-10 px-10 py-24 lg:grid-cols-[0.8fr_1.2fr]">
      <div>
        <span class="rounded-xl border border-dach-orange/30 bg-orange-50 px-4 py-2 text-sm font-semibold text-dach-orange">UK Coverage</span>
        <h2 class="mt-6 text-4xl font-black">Areas We Cover <em class="font-normal text-dach-orange">In The UK</em></h2>
        <p class="mt-5 text-dach-muted">Dach Removals operates across the entire United Kingdom. From the Scottish Highlands to Cornwall, we have crews ready.</p>
      </div>
      <div class="grid gap-5 md:grid-cols-2">
        <article v-for="[area, cities] in areas" :key="area" class="rounded-3xl bg-gray-50 p-7">
          <h3 class="text-xl font-black"><FontAwesomeIcon icon="route" class="mr-3 text-dach-orange" />{{ area }}</h3>
          <ul class="mt-5 space-y-2 text-dach-muted">
            <li v-for="city in cities" :key="city">-> {{ city }}</li>
          </ul>
        </article>
      </div>
    </section>

    <section class="px-10 py-24 text-center">
      <span class="rounded-xl border border-dach-orange/30 bg-orange-50 px-4 py-2 text-sm font-semibold text-dach-orange">Our Routes</span>
      <h2 class="mt-6 text-4xl font-black">Most Popular <em class="font-normal text-dach-orange">Routes of Dach Removals</em></h2>
      <div class="mx-auto mt-12 grid max-w-4xl gap-4 md:grid-cols-5">
        <span v-for="route in routePairs" :key="route" class="rounded-full bg-orange-50 px-5 py-3 font-bold text-dach-orange">{{ route }}</span>
      </div>
    </section>

    <section id="reviews" class="bg-gray-50 px-10 py-24">
      <h2 class="text-4xl font-black">What Our <em class="font-normal text-dach-orange">Customers Say</em></h2>
      <p class="mt-4 text-dach-muted"><FontAwesomeIcon icon="star" class="text-dach-orange" /> 4.8/5 on Trustpilot from 2,500+ reviews</p>
      <div class="mt-10 grid gap-6 md:grid-cols-4">
        <article v-for="[name, city, quote] in testimonials" :key="name" class="rounded-3xl bg-white p-7 shadow">
          <p class="text-dach-orange">*****</p>
          <p class="mt-6">"{{ quote }}"</p>
          <p class="mt-10 font-black">{{ name }}</p>
          <p class="text-sm text-dach-muted">{{ city }}</p>
        </article>
      </div>
    </section>

    <section class="px-10 py-24">
      <h2 class="text-center text-4xl font-black">Helpful Resources <em class="font-normal text-dach-orange">for You</em></h2>
      <div class="mt-10 grid gap-6 md:grid-cols-3">
        <article v-for="[title, copy] in resources" :key="title" class="rounded-3xl bg-white p-7 shadow-xl shadow-orange-100/60">
          <div class="mb-6 h-40 rounded-2xl bg-gradient-to-br from-orange-100 to-white" />
          <h3 class="text-xl font-black">{{ title }}</h3>
          <p class="mt-3 text-dach-muted">{{ copy }}</p>
        </article>
      </div>
    </section>

    <section id="faq" class="grid gap-10 px-10 py-24 lg:grid-cols-[0.7fr_1.3fr]">
      <div>
        <span class="rounded-xl border border-dach-orange/30 bg-orange-50 px-4 py-2 text-sm font-semibold text-dach-orange">FAQ</span>
        <h2 class="mt-6 text-4xl font-black">Frequently Ask <em class="font-normal text-dach-orange">Questions</em></h2>
        <p class="mt-5 text-dach-muted">If you have questions, we're here to help. Our FAQ section covers common moving topics.</p>
      </div>
      <div class="space-y-4">
        <article v-for="(faq, index) in faqs" :key="faq" class="rounded-2xl bg-gray-50">
          <button class="flex w-full items-center justify-between p-6 text-left font-black" @click="openFaq = openFaq === index ? -1 : index">
            {{ faq }}
            <FontAwesomeIcon :icon="openFaq === index ? 'chevron-down' : 'arrow-right'" class="text-dach-orange" />
          </button>
          <p v-if="openFaq === index" class="px-6 pb-6 text-dach-muted">
            Contact the Dach Removals team for an instant fixed-price quote and service guidance for your move.
          </p>
        </article>
      </div>
    </section>

    <section class="mx-10 mb-20 grid gap-8 rounded-[32px] bg-gradient-to-br from-orange-50 to-white p-10 lg:grid-cols-2">
      <div>
        <div class="flex flex-wrap gap-3 text-sm font-bold text-dach-orange">
          <span class="rounded-full bg-white px-4 py-2">Trustpilot 4.8/5</span>
          <span class="rounded-full bg-white px-4 py-2">2,500+ Reviews</span>
          <span class="rounded-full bg-white px-4 py-2">BAR Registered</span>
        </div>
        <h2 class="mt-8 text-5xl font-black">Book Your <em class="font-normal text-dach-orange">Move Today</em></h2>
        <p class="mt-4 text-dach-muted">Ready to get started? Get your instant quote now and enjoy a stress-free move.</p>
        <div class="mt-8 flex flex-wrap gap-4">
          <a href="#quote" class="rounded-xl bg-dach-orange px-6 py-4 font-black text-white">Get Your Instant Quote <FontAwesomeIcon icon="arrow-right" /></a>
          <a :href="`tel:${brand.phone}`" class="rounded-xl bg-white px-6 py-4 font-black"><FontAwesomeIcon icon="phone" /> {{ brand.phone }}</a>
        </div>
      </div>
    </section>

    <footer class="grid gap-10 bg-dach-black px-10 py-16 text-white md:grid-cols-4">
      <div>
        <div class="mb-5 h-12 w-32 bg-dach-orange" />
        <p class="text-white/60">Professional removal services across the entire UK. Technology-driven, customer-focused.</p>
        <div class="mt-6 flex">
          <input class="w-full rounded-l-xl bg-white/10 px-4 py-3 outline-none" placeholder="Enter your email" />
          <button class="rounded-r-xl bg-dach-orange px-5 font-black">Subscribe</button>
        </div>
      </div>
      <div>
        <h3 class="font-black">Services</h3>
        <p v-for="service in services.slice(0, 5)" :key="service.title" class="mt-4 text-white/60">{{ service.title }}</p>
      </div>
      <div>
        <h3 class="font-black">Resources</h3>
        <p class="mt-4 text-white/60">Van Size Guide</p>
        <p class="mt-4 text-white/60">Packing Guide</p>
        <p class="mt-4 text-white/60">Moving Day Guide</p>
        <p class="mt-4 text-white/60">Help Centre</p>
      </div>
      <div>
        <h3 class="font-black">Company</h3>
        <p class="mt-4 text-white/60">About Us</p>
        <p class="mt-4 text-white/60">Careers</p>
        <p class="mt-4 text-white/60">Trustpilot Reviews</p>
        <p class="mt-4 text-white/60">BAR Accreditation</p>
      </div>
    </footer>
  </main>
</template>
