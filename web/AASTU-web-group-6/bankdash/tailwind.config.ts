import type { Config } from "tailwindcss";
import { Lato } from "next/font/google";

const config: Config = {
  darkMode: ["class"],
  content: [
    "./pages/**/*.{ts,tsx,js,jsx,mdx}",
    "./components/**/*.{ts,tsx,js,jsx,mdx}",
    "./app/**/*.{ts,tsx,js,jsx,mdx}",
    "./src/**/*.{ts,tsx,js,jsx,mdx}",
  ],
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    // container: {
    //   center: true,
    //   padding: "2rem",
    //   screens: {
    //     "2xl": "1400px",
    //   },
    // },
    extend: {
      fontFamily: {
        Lato: "Lato",
      },
      colors: {
        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsl(var(--background))",
        foreground: "hsl(var(--foreground))",
        primary: {
          DEFAULT: "hsl(var(--primary))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsl(var(--secondary))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsl(var(--destructive))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsl(var(--muted))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsl(var(--accent))",
          foreground: "hsl(var(--accent-foreground))",
        },
        popover: {
          DEFAULT: "hsl(var(--popover))",
          foreground: "hsl(var(--popover-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card))",
          foreground: "hsl(var(--card-foreground))",
        },
        colorBody: {
          1: "#343C6A",
          2: "#9199AF",
        },
        colorIcon: {
          1: "#FFF5D9",
          2: "#E7EDFF",
        },
        colorWhite: {
          1: "rgba(255, 255, 255, 0.1)",
          2: "rgba(255, 255, 255, 0.2)",
          3: "rgba(255, 255, 255, 0.3)",
          4: "rgba(255, 255, 255, 0.4)",
          5: "rgba(255, 255, 255, 0.5)",
          6: "rgba(255, 255, 255, 0.6)",
          7: "rgba(255, 255, 255, 0.7)",
          8: "rgba(255, 255, 255, 0.8)",
          9: "rgba(255, 255, 255, 0.9)",
          10: "rgba(255, 255, 255, 1)",
        },
      },
      borderRadius: {
        lg: "var(--radius)",
        md: "calc(var(--radius) - 2px)",
        sm: "calc(var(--radius) - 4px)",
      },
      keyframes: {
        "accordion-down": {
          from: { height: "0" },
          to: { height: "var(--radix-accordion-content-height)" },
        },
        "accordion-up": {
          from: { height: "var(--radix-accordion-content-height)" },
          to: { height: "0" },
        },
      },
      animation: {
        "accordion-down": "accordion-down 0.2s ease-out",
        "accordion-up": "accordion-up 0.2s ease-out",
      },
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        page: "#F5F7FA",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
        "card-gradient":
          "linear-gradient(45deg, rgba(76,73,237,1) 41%, rgba(10,6,244,1) 100%)",
        "card-box-light":
          "linear-gradient(180deg, rgba(255,255,255,0.15) 0%, rgba(255,255,255,0) 100%)",
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
};

export default config;
