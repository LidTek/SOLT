<script>
  import { signIn, signOut } from "@auth/sveltekit/client";
  import { page } from "$app/stores";
</script>

<login class="flex flex-col justify-center items-center gap-5">
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
</login>
