<script>
    import { useNavigate, useLocation } from "svelte-navigator";
    import { user } from "../store.js";

    const navigate = useNavigate();
    const location = useLocation();

    $user = null
    localStorage.removeItem("user")

    navigate("/login", {
      state: { from: $location.pathname },
      replace: true,
    });
</script>
