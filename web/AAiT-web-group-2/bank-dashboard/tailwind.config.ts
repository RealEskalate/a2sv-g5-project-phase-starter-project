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
      },
      colors: {
        "custom-light-purple": "#718EBF",
        "custom-purple": "#123288",
        "custom-bright-purple": "#1814F3",
        "custom-pink-red": "#FE5C73",
        "custom-greenish": "#16DBAA",
        "custom-light-grey": "#DFEAF2",
        "custom-light-dark": "#232323",
        "custom-faint-white": "#F4F5F7",
        "background": "#F5F7FA",
      },
      fontFamily: {
        body: ["Inter", "sans-serif"]

      }
    },
  },
  plugins: [],
};
export default config;
