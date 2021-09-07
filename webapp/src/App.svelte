<script>
import { Router, Route, Link } from "svelte-navigator";
import Index from "./pages/Index.svelte";
import Containers from "./pages/Containers.svelte";
import Servers from "./pages/Servers.svelte";
</script>

<Router>
  <Route path="/" component={Index} />
  <Route path="/containers" component={Containers} />
  <Route path="/servers" component={Servers} />
</Router>