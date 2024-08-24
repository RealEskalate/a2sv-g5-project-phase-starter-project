import { useUser } from "@/contexts/UserContext";
import React from "react";

export const Loading = () => {
  const { isDarkMode } = useUser();

  return (
    <div
      className={`flex space-x-2 justify-center items-center h-screen pb-32 pr-10 transition-colors duration-300 `}
    >
      <span className="sr-only">Loading...</span>
      <div
        className={`h-5 w-5 rounded-full animate-bounce ${
          isDarkMode ? "bg-white" : "bg-gray-400"
        } [animation-delay:-0.3s]`}
      ></div>
      <div
        className={`h-5 w-5 rounded-full animate-bounce ${
          isDarkMode ? "bg-white" : "bg-gray-400"
        } [animation-delay:-0.15s]`}
      ></div>
      <div
        className={`h-5 w-5 rounded-full animate-bounce ${
          isDarkMode ? "bg-white" : "bg-gray-400"
        }`}
      ></div>
    </div>
  );
};
