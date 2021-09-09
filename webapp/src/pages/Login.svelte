<script>
import { useNavigate, useLocation } from "svelte-navigator";
import { user } from "../store.js";

const navigate = useNavigate();
const location = useLocation();

function auth() {
    $user = { user: "Test", token: "token" }
    localStorage.setItem("user", JSON.stringify($user))

    navigate("/", {
      state: { from: $location.pathname },
      replace: true,
    });
}
</script>

<main>
    <button on:click={auth}>Login</button>
</main>