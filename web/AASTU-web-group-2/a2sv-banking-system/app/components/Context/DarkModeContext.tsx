// DarkModeContext.tsx
"use client";

import React, {
  createContext,
  useState,
  useContext,
  ReactNode,
  useEffect,
} from "react";

interface DarkModeContextType {
  darkMode: boolean | undefined;
  toggleDarkMode: () => void;
}

const DarkModeContext = createContext<DarkModeContextType | undefined>(
  undefined
);

export const DarkModeProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [darkMode, setDarkMode] = useState<boolean | undefined>(undefined);

  useEffect(() => {
    if (typeof window !== "undefined") {
      const storedMode = localStorage.getItem("darkMode");
      setDarkMode(storedMode === "true");
    }
  }, []);

  useEffect(() => {
    if (darkMode !== undefined) {
      document.body.classList.toggle("dark", darkMode);
      if (typeof window !== "undefined") {
        localStorage.setItem("darkMode", String(darkMode));
      }
    }
  }, [darkMode]);

  const toggleDarkMode = () => {
    setDarkMode((prevMode) => !prevMode);
  };

  return (
    <DarkModeContext.Provider value={{ darkMode: darkMode ?? false, toggleDarkMode }}>
      {children}
    </DarkModeContext.Provider>
  );
};

export const useDarkMode = () => {
  const context = useContext(DarkModeContext);
  if (context === undefined) {
    throw new Error("useDarkMode must be used within a DarkModeProvider");
  }
  return context;
};
