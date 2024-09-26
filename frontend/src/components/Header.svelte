<script>
  export let y;
</script>

<header
  class="absolute sticky z-[10] top-0 duration-200 px-6 py-6 flex items-center justify-between"
>
  <h1 class="font-medium">
    <b class="font-bold poppins cursor-pointer"
      ><span class="font-bold poppins text-blue-600">Draw</span>niverse</b
    >
  </h1>
  <div class="sm:flex flex wrap items-center content-end space-x-4">
    <a
      target="_blank"
      href="/sign_up"
      class="blueShadow relative overflow-hidden px-5 py-2 group rounded-full bg-white text-slate-950"
    >
      <div
        class="absolute top-0 right-full w-full h-full bg-violet-400 opacity-20 group-hover:translate-x-full z-0 duration-200"
      ></div>
      <h4 class="relative z-9">
        <p>Sign Up</p>
      </h4>
    </a>
    <a
      target="_blank"
      href="/login"
      class="blueShadow relative overflow-hidden px-5 py-2 group rounded-full bg-white text-slate-950"
    >
      <div
        class="absolute top-0 right-full w-full h-full bg-violet-400 opacity-20 group-hover:translate-x-full z-0 duration-200"
      ></div>
      <h4 class="relative z-9"><p>Login</p></h4>
    </a>
  </div>
</header>
