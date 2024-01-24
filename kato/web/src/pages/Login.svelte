<script>
  import { navigate } from "svelte-routing";
  import { LOGIN_TOKEN_KEY } from "../const";
  import { userToken } from "../store";
  import { onMount } from "svelte";

  function login() {
    window.localStorage.setItem(LOGIN_TOKEN_KEY, "user");
    $userToken = "user";
    navigate("/", { replace: true });
  }

  onMount(() => {
    if (Boolean(window.localStorage.getItem(LOGIN_TOKEN_KEY))) {
      navigate("/", { replace: true });
    }
  });
</script>

<button on:click={login}>Login</button>
