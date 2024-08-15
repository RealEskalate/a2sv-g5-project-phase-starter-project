import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
        "credit-card-gradient":
          "linear-gradient(107.38deg, #4C49ED 2.61%, #0A06F4 101.2%)",
        "secondary-credit-card":
          "linear-gradient(107.38deg, #2D60FF 2.61%, #539BFF 101.2%)",
      },
      boxShadow: {
        "custom-shadow": "4px 4px 18px -2px rgba(231, 228, 232, 0.8)",
      },
    },
  },
  plugins: [],
};
export default config;
