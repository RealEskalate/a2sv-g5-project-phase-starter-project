/** @type {import('next').NextConfig} */
const nextConfig = {
  // Your existing Next.js config
  webpack: (config, { isServer }) => {
    if (!isServer) {
      // SVG loader configuration for client-side
      config.module.rules.push({
        test: /\.svg$/,
        use: ['@svgr/webpack'],
      });
    }
    return config;
  },
};

export default nextConfig;
