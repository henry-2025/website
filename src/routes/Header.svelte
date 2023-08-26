<script>
    import DropdownMenu from './DropdownMenu.svelte';
</script>
<header class="bg-gradient-to-l from-white">
  <nav class="mx-auto flex w-4/5 items-center justify-between p-6 lg:px-8" aria-label="Global">
    <div class="flex lg:flex-1">
      <a href="/" class="-m-1.5 p-1.5 logo mix-blend-color-dodge">
        JHP
      </a>
    </div>
    <div class="flex hidden">
            <DropdownMenu />
    </div>
    <div class="lg:flex lg:flex-1 lg:gap-x-12 lg:justify-end">
      <a href="/about" class="text-m font-mono font-light leading-6 text-gray-900">About</a>
      <a href="/projects" class="text-m font-mono font-light leading-6 text-gray-900">Projects</a>
      <a href="/contact" class="text-m font-mono font-light leading-6 text-gray-900">Contact</a>
    </div>
  </nav>
</header>

<style>
  .logo {
    font-size: 20pt;
  }
</style>
