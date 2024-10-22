const cacheName = 'travelbangla-cache-v1';
const cacheAssets = [
  '/',
  '/resources/css/home.css',
  '/resources/js/main.js',
  '/resources/package_img/', // Cache package images if needed
  '/admin-resources/css/style.min.css',
  '/admin-resources/plugins/jquery/dist/jquery.min.js',
  '/admin-resources/plugins//bootstrap/dist/js/bootstrap.bundle.min.js',
  '/admin-resources/js/sidebarmenu.js',
  '/admin-resources/js/pages/dashboards/dashboard1.js',
  '/admin-resources/js/custom.js',
  // Add more paths to assets you want to cache
];

// Install event
self.addEventListener('install', (e) => {
  e.waitUntil(
    caches.open(cacheName).then((cache) => {
      return cache.addAll(cacheAssets);
    })
  );
});

// Fetch event
self.addEventListener('fetch', (e) => {
  e.respondWith(
    fetch(e.request)
      .then((res) => {
        // Make a clone of the response
        const resClone = res.clone();
        // Open cache
        caches.open(cacheName).then((cache) => {
          // Add the response to the cache
          cache.put(e.request, resClone);
        });
        return res;
      })
      .catch(() => caches.match(e.request).then((res) => res))
  );
});