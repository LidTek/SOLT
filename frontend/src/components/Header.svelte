<script>
  export let y;

  //   let tabs = [
  //     { name: "Projects", link: "#projects" },
  //     { name: "About Me", link: "#about" },
  //   ];
</script>

<header
  class="sticky z-[10] top-0 duration-200 px-6 py-6 flex items-center justify-between"
>
  <h1 class="font-medium">
    <b class="font-bold poppins cursor-pointer"
      ><span class="font-bold poppins text-blue-600">Draw</span>niverse</b
    >
  </h1>
  <div class="sm:flex flex wrap items-center gap-4">
    <!-- {#each tabs as tab, index}
      <a href={tab.link} class="duration-200 hover:text-violet-400">
        <p>{tab.name}</p>
      </a>
    {/each} -->
    <a
      target="_blank"
      href="https://x.com/temidaradev"
      class="blueShadow relative overflow-hidden px-5 py-2 group rounded-full bg-white text-slate-950"
    >
      <div
        class="absolute top-0 right-full w-full h-full bg-violet-400 opacity-20 group-hover:translate-x-full z-0 duration-200"
      ></div>
      <h4 class="relative z-9">Get In Touch</h4>
    </a>
  </div>
</header>
