export const adminNavItems = [
  { label: 'Dashboard', icon: 'gauge-high', path: '/admin/dashboard' },
  { label: 'Quote Requests', icon: 'clipboard', path: '/admin/quotes', badge: 0 },
  { label: 'Bookings', icon: 'calendar-days', path: '/admin/bookings' },
  { label: 'Messages', icon: 'envelope', path: '/admin/messages', badge: 0 },
  { label: 'Business Setup', icon: 'building', path: '/admin/content' },
]

export const calendarDays = ['MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT', 'SUN']

export const calendarCells = [
  '31',
  ...Array.from({ length: 30 }, (_, index) => String(index + 1)),
  '1',
  '2',
  '3',
  '4',
]
