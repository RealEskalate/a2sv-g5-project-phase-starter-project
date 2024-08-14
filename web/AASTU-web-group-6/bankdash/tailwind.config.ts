import { Lato } from "next/font/google";
import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        Lato: "Lato",
      },
      colors: {
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
  plugins: [],
};
export default config;
