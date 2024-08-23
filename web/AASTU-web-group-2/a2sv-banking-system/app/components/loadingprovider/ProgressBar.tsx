"use client";
import React, { useEffect } from "react";
import { Next13ProgressBar } from "next13-progressbar";

const ProgressBar = ({ children }: { children: React.ReactNode }) => {
  useEffect(() => {
    console.log("ProgressBar rendered");
  }, []);

  return (
    <>
      {children}
      <Next13ProgressBar
        height="4px"
        color="#0A2FFF"
        options={{ showSpinner: false }}
        showOnShallow
      />
    </>
  );
};

export default ProgressBar;
