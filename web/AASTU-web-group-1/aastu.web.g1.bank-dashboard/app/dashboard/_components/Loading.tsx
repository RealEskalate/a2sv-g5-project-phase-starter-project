import { useUser } from "@/contexts/UserContext";
import React from "react";

export const Loading = () => {
  const { isDarkMode } = useUser();

  return (
    <div
      className={`flex space-x-2 justify-center items-center h-screen transition-colors duration-300 ${
        isDarkMode ? "bg-gray-900" : "bg-white"
      }`}
    >
      <span className="sr-only">Loading...</span>
      <div
        className={`h-8 w-8 rounded-full animate-bounce ${
          isDarkMode ? "bg-white" : "bg-black"
        } [animation-delay:-0.3s]`}
      ></div>
      <div
        className={`h-8 w-8 rounded-full animate-bounce ${
          isDarkMode ? "bg-white" : "bg-black"
        } [animation-delay:-0.15s]`}
      ></div>
      <div
        className={`h-8 w-8 rounded-full animate-bounce ${
          isDarkMode ? "bg-white" : "bg-black"
        }`}
      ></div>
    </div>
  );
};
