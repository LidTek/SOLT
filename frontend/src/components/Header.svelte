<script>
  export let y;
  import { signIn, signOut } from "@auth/sveltekit/client";
  import { page } from "$app/stores";
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
    {#if $page.data.session?.user}
      <p class="mx-auto poppins font-semibold text-sm sm:text-md md:text-l">
        Signed in as {$page.data.session.user.email}
      </p>
      <button
        class="blueShadow mx-auto lg:mr-auto lg:ml-0 text-base sm:text-sm md:text-md poppins relative justify-center overflow-hidden px-6 py-3 group rounded-full bg-white text-slate-950"
        on:click={signOut}>Sign out</button
      >
    {:else}
      <p class="mx-auto poppins font-semibold text-sm sm:text-md md:text-l">
        Not signed in.
      </p>
      <button
        class="blueShadow mx-auto lg:mr-auto lg:ml-0 text-base sm:text-sm md:text-md poppins relative justify-center overflow-hidden px-6 py-3 group rounded-full bg-white text-slate-950"
        on:click={() => signIn("google")}
        >Sign In With Google <i class="fa-brands fa-google"></i></button
      >
      <button
        class="blueShadow mx-auto lg:mr-auto lg:ml-0 text-base sm:text-sm md:text-md poppins relative justify-center overflow-hidden px-6 py-3 group rounded-full bg-white text-slate-950"
        on:click={() => signIn("github")}
        >Sign In With Github <i class="fa-brands fa-github"></i></button
      >
    {/if}
  </div>
</header>
