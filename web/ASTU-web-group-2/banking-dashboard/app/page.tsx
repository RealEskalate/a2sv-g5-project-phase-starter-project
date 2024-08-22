import React from "react";
import LandingNav from "./components/landing/LandingNav";
import LandingHome from "./components/landing/LandingHome";
import Services from "./components/landing/services";
import About from "./components/landing/about";
import Footer from "./components/landing/footer";

const page = () => {
  return (
    <div>
      <LandingNav />
      <LandingHome />
      <Services />
      <About />
      <Footer />
    </div>
  );
};

export default page;
