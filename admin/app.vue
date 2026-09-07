<script setup lang="ts">
import { dashboardCards, editableSections, navItems } from './data/admin'

const route = useRoute()
const activeView = computed(() => {
  const path = route.path
  if (path.includes('/quotes')) return 'quotes'
  if (path.includes('/bookings')) return 'bookings'
  if (path.includes('/messages')) return 'messages'
  if (path.includes('/content')) return 'content'
  if (path.includes('/login')) return 'login'
  return 'dashboard'
})

const days = ['MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT', 'SUN']
const calendarCells = ['31', ...Array.from({ length: 30 }, (_, i) => String(i + 1)), '1', '2', '3', '4']
</script>

<template>
  <div class="hidden">
    <NuxtPage />
  </div>

  <div v-if="activeView === 'login'" class="grid min-h-screen place-items-center bg-dach-black">
    <form class="w-[420px] border-t-4 border-dach-orange bg-white p-12 shadow-2xl">
      <div class="mb-8 flex items-center gap-3">
        <span class="grid h-12 w-12 place-items-center bg-dach-orange text-2xl font-black">D</span>
        <span>
          <span class="display-title block text-2xl">ACH</span>
          <span class="block text-[10px] font-bold uppercase tracking-[0.45em] text-dach-orange">Removals</span>
        </span>
      </div>
      <p class="mb-8 text-sm font-bold uppercase tracking-[0.25em] text-dach-muted">Admin Portal</p>
      <label class="text-xs font-bold uppercase tracking-[0.2em] text-dach-muted">Username</label>
      <input value="admin" class="mt-2 w-full border border-dach-line bg-dach-cream px-5 py-4 outline-none" />
      <label class="mt-5 block text-xs font-bold uppercase tracking-[0.2em] text-dach-muted">Password</label>
      <input type="password" value="password" class="mt-2 w-full border border-dach-line bg-dach-cream px-5 py-4 outline-none" />
      <NuxtLink to="/dashboard" class="mt-6 block bg-dach-orange px-6 py-4 text-center font-black uppercase text-white">
        Sign In <FontAwesomeIcon icon="arrow-left" class="rotate-180" />
      </NuxtLink>
      <a href="/" class="mt-8 block text-center text-sm text-dach-muted"><FontAwesomeIcon icon="arrow-left" /> Back to website</a>
    </form>
  </div>

  <div v-else class="grid min-h-screen grid-cols-[300px_1fr] bg-dach-cream">
    <aside class="flex flex-col bg-dach-black text-white">
      <div class="flex h-28 items-center gap-3 border-b border-white/10 px-7">
        <span class="grid h-12 w-12 place-items-center bg-dach-orange text-2xl font-black text-dach-black">D</span>
        <span>
          <span class="display-title block text-2xl">ACH</span>
          <span class="block text-[10px] font-bold uppercase tracking-[0.45em] text-dach-orange">Removals</span>
          <span class="block text-[11px] uppercase tracking-[0.25em] text-white/40">Admin Panel</span>
        </span>
      </div>

      <nav class="mt-4">
        <NuxtLink
          v-for="item in navItems"
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
        <NuxtLink to="/login"><FontAwesomeIcon icon="right-from-bracket" class="mr-2" />Sign Out</NuxtLink>
      </div>
    </aside>

    <main>
      <header class="flex h-20 items-center justify-between bg-white px-9">
        <h1 class="display-title text-3xl">
          {{ activeView === 'dashboard' ? 'Dashboard' : activeView === 'quotes' ? 'Quote Requests' : activeView === 'bookings' ? 'Bookings Calendar' : activeView === 'messages' ? 'Messages' : 'Edit Content' }}
        </h1>
        <p class="text-dach-muted">Mon, 7 September 2026 <span class="ml-3 text-green-600">● Live</span></p>
      </header>

      <section v-if="activeView === 'dashboard'" class="p-9">
        <div class="grid gap-6 xl:grid-cols-4">
          <article v-for="card in dashboardCards" :key="card.label" class="border border-dach-line bg-white p-8">
            <p class="text-sm font-bold uppercase tracking-[0.2em] text-dach-muted">{{ card.label }}</p>
            <p class="display-title mt-3 text-5xl">{{ card.value }}</p>
            <p :class="card.tone" class="mt-3">{{ card.note }}</p>
          </article>
        </div>
        <div class="mt-9 grid gap-8 xl:grid-cols-2">
          <article class="min-h-[360px]">
            <div class="mb-6 flex items-center justify-between"><h2 class="display-title text-2xl">Recent Quotes</h2><button class="border border-dach-line bg-white px-5 py-3 font-bold">View All</button></div>
            <div class="grid h-80 place-items-center text-dach-muted"><p class="text-center text-lg">📋<br />No quotes yet</p></div>
          </article>
          <article class="min-h-[360px]">
            <div class="mb-6 flex items-center justify-between"><h2 class="display-title text-2xl">Upcoming Bookings</h2><button class="border border-dach-line bg-white px-5 py-3 font-bold">Calendar</button></div>
            <div class="grid h-80 place-items-center text-dach-muted"><p class="text-center text-lg">📅<br />No upcoming bookings</p></div>
          </article>
        </div>
      </section>

      <section v-if="activeView === 'quotes'" class="p-9">
        <div class="flex items-center justify-between"><h2 class="display-title text-2xl">All Quote Requests</h2><input class="border border-dach-line bg-white px-5 py-4" placeholder="Search name or postcode..." /></div>
        <div class="grid h-[520px] place-items-center text-center text-dach-muted">📋<br />No quote requests yet.<br />They'll appear here when customers fill in the form on your website.</div>
      </section>

      <section v-if="activeView === 'bookings'" class="p-9">
        <div class="mb-8 flex items-center justify-between">
          <div class="flex items-center gap-4"><button class="border bg-white px-4 py-3">&lt;</button><h2 class="display-title text-2xl">September 2026</h2><button class="border bg-white px-4 py-3">&gt;</button></div>
          <button class="bg-dach-orange px-6 py-4 font-black text-white"><FontAwesomeIcon icon="plus" /> Add Booking</button>
        </div>
        <div class="grid grid-cols-7 text-center text-sm font-bold text-dach-muted"><span v-for="day in days" :key="day">{{ day }}</span></div>
        <div class="mt-4 grid grid-cols-7">
          <div v-for="cell in calendarCells" :key="cell" class="h-28 border border-dach-line bg-white p-3" :class="cell === '7' ? 'border-dach-orange text-dach-orange' : ''">{{ cell }}</div>
        </div>
        <p class="mt-5 text-dach-muted"><span class="text-dach-orange">■</span> New <span class="ml-6 text-green-600">■</span> Confirmed</p>
      </section>

      <section v-if="activeView === 'messages'" class="p-9">
        <h2 class="display-title mb-6 text-2xl">Customer Messages</h2>
        <div class="grid min-h-[620px] grid-cols-[360px_1fr] border border-dach-line bg-white">
          <div class="grid place-items-start border-r border-dach-line p-20 text-center text-dach-muted">✉️<br />No messages yet</div>
          <div class="grid place-items-center text-dach-muted">Select a message to read</div>
        </div>
      </section>

      <section v-if="activeView === 'content'" class="p-9">
        <h2 class="display-title mb-6 text-2xl">Edit Website Content</h2>
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
