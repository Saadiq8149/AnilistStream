import { render } from "preact";
import { LocationProvider, Router, Route } from "preact-iso";
import NotFound from "./pages/_404.jsx";
import Home from "./pages/Home/index.jsx";
import "./index.css";
import Configure from "./pages/Configure/index.jsx";

export function App() {
  return (
    <LocationProvider>
      <main>
        <Router>
          <Route path="/" component={Home} />
          <Route path="/configure" component={Configure} />
          <Route default component={NotFound} />
        </Router>
      </main>
    </LocationProvider>
  );
}

render(<App />, document.getElementById("app"));
