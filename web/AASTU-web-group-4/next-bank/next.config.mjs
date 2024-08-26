/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    domains: ['cdn.freelogovectors.net'], // Add the external image domain here
  },
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
