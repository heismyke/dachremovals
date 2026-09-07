export const adminNavItems = [
  { label: 'Dashboard', icon: 'bolt', path: '/admin/dashboard' },
  { label: 'Quote Requests', icon: 'clipboard', path: '/admin/quotes', badge: 0 },
  { label: 'Bookings', icon: 'calendar-days', path: '/admin/bookings' },
  { label: 'Messages', icon: 'envelope', path: '/admin/messages', badge: 0 },
  { label: 'Edit Content', icon: 'pen', path: '/admin/content' },
]

export const dashboardCards = [
  { label: 'New Quotes', value: '0', note: 'Awaiting response', tone: 'text-dach-orange' },
  { label: 'Bookings This Month', value: '0', note: 'Confirmed jobs', tone: 'text-green-600' },
  { label: 'Unread Messages', value: '0', note: 'Needs attention', tone: 'text-dach-orange' },
  { label: 'Next Job', value: '-', note: '', tone: 'text-dach-muted' },
]

export const editableSections = [
  { label: 'Hero Section', fields: 3 },
  { label: 'About Us', fields: 3 },
  { label: 'Contact Info', fields: 2 },
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
