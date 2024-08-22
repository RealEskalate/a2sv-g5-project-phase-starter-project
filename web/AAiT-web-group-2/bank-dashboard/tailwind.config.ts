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
        "my-card-bg-1": "linear-gradient(107.38deg, #4C49ED 2.61%, #0A06F4 101.2%)",
        "my-card-bg-2": "linear-gradient(180deg, rgba(255, 255, 255, 0.15) 0%, rgba(255, 255, 255, 0) 100%)",
        "my-bg-1": "#F5F7FA",
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
        "custom-light-orange": "#FFF5D9",
        "custom-light-blue": "#E7EDFF",
        "custom-light-teal": "#DCFAF8",
        "custom-red": "#FF4B4A",
        "custom-font-color": "#343C6A",
        "my-color-8": "#41D4A8",
        "my-color-10": "#EDF0F7",
      },
      fontFamily: {
        body: ["Inter", "sans-serif"]
      },
      width: {
        "chip-size": "34.77px",
        "my-card-width": "350px",
        "my-card-width-2": "350px",
      },
      height: {
        "my-card-height": "235px",
        "my-card-height-2": "70px",
        "my-card-height-3": "235px",
      },
      borderRadius: {
        "my-card-radius": "25px",
      },
      screens: {
        'max-sm': {'max': '380px'},  
        'max-md' :{'max': '1025px'}
      },
    },
  },
  plugins: [],
};
export default config;
