"use client";
import React from "react";
import Image from "next/image";
import DescriptionCard from "./components/Shimmer/SimmerCard";

const NotFound = () => {
  const li = [1, 2, 3, 4];
  return (
    <div className="w-full flex flex-col px-6 gap-2">
      {li.map((item, key) => (
        <DescriptionCard key={key} />
      ))}
    </div>
  );
};

export default NotFound;
