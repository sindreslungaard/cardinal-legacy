<script>
import { Router, Route, Link } from "svelte-navigator";
import Index from "./pages/Index.svelte";
import Containers from "./pages/Containers.svelte";
import Servers from "./pages/Servers.svelte";
import Login from "./pages/Login.svelte";
import AuthGuard from "./components/AuthGuard.svelte";
import Logout from "./pages/Logout.svelte";
</script>

<Router>
  <Route path="/login" component={Login} />
  <Route path="/logout" component={Logout} />

  <AuthGuard>
    <Route path="/" component={Index} />
    <Route path="/containers" component={Containers} />
    <Route path="/servers" component={Servers} />
  </AuthGuard>
</Router>