import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      boxShadow: {
        'custom-shadow': '4px 4px 18px -2px #E7E4E8CC',
      },
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
          "blue-gradient": "linear-gradient(107.38deg, #4C49ED 2.61%, #0A06F4 101.2%)",
        
      },

      screens: {
        // Override default breakpoints with custom values
        mobile: '440px',    // Small screens: 480px and up
        tablet: '1024px',    // Medium screens: 768px and up (default)
        desktop: '1440px',   // Large screens: 1024px and up (default)
        xl: '1280px',   // Extra Large screens: 1280px and up (default)
        '2xl': '1440px' // 2x Extra Large screens: 1440px and up
      },

      colors: {
        "Very-Pale-Blue" : "#E6EFF5",
        "Very-Light-Grey" : "#B1B1B1",
        "Dark-Slate-Blue" : "#343C6A",
        "Very-Light-White" : "#F5F7FA",
      }
    },
  },
  plugins: [],
};
export default config;
