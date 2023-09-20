import { Route, Routes } from "@solidjs/router";
import Navbar from "./components/Navbar";
import Home from "./pages/Home";
import Info from "./pages/Info";

function App() {
  return (
    <>
        <Navbar />
      <div class="bg-base-100 h-screen w-screen flex justify-center items-center">
        <Routes>
          <Route path={"/"} component={Home} />
          <Route path={"/info"} component={Info} />
        </Routes>
      </div>
    </>
  );
}

export default App;
