import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      colors: {
        customBlue: '#718EBF', // Example custom color
        customGreen: '#17BF63', // Another custom color
        customGray: {
          light: '#F7FAFC',
          DEFAULT: '#E2E8F0',
          dark: '#4A5568',
        },}
    },
  },
  plugins: [],
};
export default config;
