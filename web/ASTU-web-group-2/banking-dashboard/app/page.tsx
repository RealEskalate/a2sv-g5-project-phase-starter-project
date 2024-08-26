"use client";
import React, { useState, useEffect, useRef } from "react";
import LandingNav from "./components/landing/LandingNav";
import LandingHome from "./components/landing/LandingHome";
import Services from "./components/landing/services";
import About from "./components/landing/about";
import Footer from "./components/landing/footer";

const Page = () => {
  const [bgWhite, setBgWhite] = useState(false);
  const homeRef = useRef<HTMLDivElement | null>(null);
  const servicesRef = useRef<HTMLDivElement | null>(null);
  const aboutRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const handleScroll = () => {
      if (homeRef.current) {
        const { bottom } = homeRef.current.getBoundingClientRect();
        setBgWhite(bottom <= 0); // Change background if the user has scrolled past LandingHome
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  return (
    <div>
      <div className={`${bgWhite ? "fixed" : ""} w-[100%]`}>
        <LandingNav
          bgWhite={bgWhite}
          homeRef={homeRef}
          servicesRef={servicesRef}
          aboutRef={aboutRef}
        />
      </div>
      <div ref={homeRef}>
        <LandingHome />
      </div>
      <div ref={servicesRef}>
        <Services />
      </div>
      <div ref={aboutRef}>
        <About />
      </div>
      <Footer />
    </div>
  );
};

export default Page;
