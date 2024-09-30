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
    <!-- <a
      target="_blank"
      href=""
      class="blueShadow relative overflow-hidden px-5 py-2 group rounded-full bg-white text-slate-950"
    >
      <div
        class="absolute top-0 right-full w-full h-full bg-violet-400 opacity-20 group-hover:translate-x-full z-0 duration-200"
      ></div>
      <h4 class="relative z-9">
        <p>Sign Up</p>
      </h4>
    </a> -->
    {#if $page.data.session?.user}
      <p>Signed in as {$page.data.session.user.email}</p>
      <button on:click={signOut}>Sign out</button>
    {:else}
      <p>Not signed in.</p>
      <button class="bg-black rounded-full" on:click={() => signIn("google")}
        >Sign In With Google <i class="fa-brands fa-google"></i></button
      >
      <button class="bg-black rounded-full" on:click={() => signIn("github")}
        >Sign In With Github <i class="fa-brands fa-github"></i></button
      >
    {/if}
  </div>
</header>
