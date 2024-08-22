'use client';
import React, { useState, useEffect } from "react";
import LandingNav from "./components/landing/LandingNav";
import LandingHome from "./components/landing/LandingHome";
import Services from "./components/landing/services";
import About from "./components/landing/about";
import Footer from "./components/landing/footer";

const Page = () => {
  const [bgWhite, setBgWhite] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      const homeSection = document.getElementById('home-section');
      if (homeSection) {
        const { bottom } = homeSection.getBoundingClientRect();
        setBgWhite(bottom <= 0); // Change background if the user has scrolled past LandingHome
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  return (
    <div>
      <div className={`${bgWhite&&"fixed"}  w-[100%]`}>
      <LandingNav bgWhite={bgWhite} />
      </div>
      <div id="home-section">
        <LandingHome />
      </div>
      <div id="services-section">
        <Services />
      </div>
      <div id="about-section">
        <About />
      </div>
      <Footer />
    </div>
  );
};

export default Page;
