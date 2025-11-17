---
layout: home
pageClass: home-page

hero:
  name: Cruise 
  text: A Container Management TUI 
  actions:
    - theme: brand 
      text: Getting Started 
      link: /docs/
    - theme: alt
      text: Documentation 
      link: /docs/config/

features:
  - title: Comprehensive Container Management
    details: Manage containers, images, volumes, and networks from a single TUI, covering the full lifecycle of artifacts.
  - title: Real-Time Monitoring & Logs 
    details: Get instant visibility into container health, system metrics, and live log streams, all in one dashboard.
  - title: Intuitive & Customizable UI 
    details: A clean terminal interface with keyboard shortcuts, themes, and filters for a smooth developer experience.
  - title: Powerful Search & Filtering
    details: Quickly locate containers, images, or volumes with flexible search, sorting, and context-aware actions.
---

<style>

.home-page .VPFeatures .VPFeature {
  border: 1px solid #b4befe; /* subtle bluish-gray border */
  border-radius: 12px;
  background-color: rgba(30, 32, 48, 0.4); /* semi-transparent background */
  box-shadow: 0 4px 14px rgba(125, 207, 255, 0.06); /* soft blue glow */
  transition: border-color 0.3s ease, box-shadow 0.3s ease;

}

.home-page .VPFeatures .VPFeature:hover {
  box-shadow: 0 6px 20px rgba(125, 207, 255, 0.1);
}
:root {
  /* === Hero Title Gradient (Indigo to Cyan) === */
  --vp-home-hero-name-color: transparent;
  --vp-home-hero-name-background: -webkit-linear-gradient(
    135deg,
    #b4befe,   /* Soft Purple */
    #89b4fa  /* Light Blue */
  );

  /* === Hero Image Glow Background === */
  --vp-home-hero-image-background-image: linear-gradient(
    -45deg,
    #1a1b26, /* Base background */
    #b4befe  /* Glow */
  );

  --vp-home-hero-image-filter: blur(56px);
  --overlay-gradient: color-mix(in srgb, #b4befe, transparent 80%);
}

.home-page {
  background:
    linear-gradient(225deg, var(--overlay-gradient), transparent 40%),
    radial-gradient(var(--overlay-gradient), transparent 60%) no-repeat -40vw -20vh / 120vw 180vh,
    radial-gradient(var(--overlay-gradient), transparent 70%) no-repeat 50% calc(100% + 20rem) / 60rem 30rem;

  .VPFeature code {
    background-color: #1e2030; /* Slightly lighter than main bg */
    color: #7dcfff;
    padding: 3px 6px;
    border-radius: 4px;
  }

  /* === Transparent Footer === */
  .VPFooter {
    background-color: transparent !important;
    border: none;
  }

  /* === Frosted Glass NavBar === */
  .VPNavBar:not(.top) {
    background-color: rgba(30, 32, 48, 0.6) !important;
    -webkit-backdrop-filter: blur(14px);
    backdrop-filter: blur(14px);
  }
}

/* === Responsive Hero Blur Tweaks === */
@media (min-width: 640px) {
  :root {
    --vp-home-hero-image-filter: blur(64px);
  }
}
@media (min-width: 960px) {
  :root {
    --vp-home-hero-image-filter: blur(72px);
  }
}
</style>
